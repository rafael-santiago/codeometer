// package loader concentrates all code loading stuff, however all you should call is LoadCode().
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package loader

import (
    "internal/ruler"
    "internal/options"
    "os"
    "fmt"
    "strings"
    "path/filepath"
    "io/ioutil"
    "archive/zip"
    "bytes"
    "io"
    "os/exec"
    "strconv"
    "time"
    "net/url"
    "net/http"
)

var gWorkingLoadersPerRecursionNr int = 20

var gAsyncLoadDir bool

var gSubtaskTimeout string = "10m"

var gHasGit bool

// Initializes internal stuff.
func init() {
    gAsyncLoadDir = options.GetBoolOption("async", false)
    nr, err := strconv.Atoi(options.GetOption("loaders-nr",
                                              fmt.Sprintf("%d", gWorkingLoadersPerRecursionNr)))
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: --loaders-nr option has invalid data : %s.\n", err)
        os.Exit(1)
    }
    if nr <= 0 {
        fmt.Fprintf(os.Stderr, "error: --loaders-nr option must be a positive number.\n")
        os.Exit(1)
    }
    gWorkingLoadersPerRecursionNr = nr
    gSubtaskTimeout = options.GetOption("subtask-timeout", gSubtaskTimeout)
    _, err = time.ParseDuration(gSubtaskTimeout)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: --subtask-timeout has invalid value : %s.\n", err)
        os.Exit(1)
    }
    cmd := exec.Command("git", "--version")
    if cmd.Run() == nil {
        gHasGit = true
    }
}

// This function expects as srcpath a file, directory or git-repo uri. At the end
// of a successful execution codestat will gather all relevant files info. The
// relevance a code file is given by exts... where you pass all relevant file
// extensions, including the dot symbol or not.
func LoadCode(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    codestat.Lock()
    if len(codestat.ProjectName) == 0 {
        codestat.ProjectName = filepath.Base(srcpath)
    }
    codestat.Unlock()
    var loader func(*ruler.CodeStat, string, ...string) error

    if strings.HasPrefix(srcpath, "https://") ||
       strings.HasPrefix(srcpath, "git://")   ||
       strings.HasPrefix(srcpath, "http://") {
        loader = loadGitRepoCode
    } else {
        st, err := os.Stat(srcpath)
        if err != nil {
            return err
        }
        mode := st.Mode()
        if mode.IsDir() {
            if !gAsyncLoadDir {
                loader = loadCodeDirSync
            } else {
                loader = loadCodeDirAsync
            }
        } else if mode.IsRegular() && strings.HasSuffix(strings.ToLower(srcpath), ".zip") {
            loader = loadZippedCode
        } else if mode.IsRegular() {
            loader = loadCodeFile
        } else {
            loader = func(*ruler.CodeStat, string, ...string) error {
                        return fmt.Errorf("ERROR: Unsupported file type.")
                     }
        }
    }
    return loader(codestat, srcpath, exts...)
}

// Checks if filename ends with some relevant file extension.
func isRelevantFile(filename string, exts...string) bool {
    if len(exts) == 0 {
        return true
    }
    for _, ext := range exts {
        if len(ext) > 0 && !strings.HasPrefix(ext, ".") {
            ext = "." + ext
        }
        if strings.HasSuffix(strings.ToLower(filename), strings.ToLower(ext)) {
            return true
        }
    }
    return false
}

// Gets information from a relevant code file.
func loadCodeFile(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    if isRelevantFile(srcpath, exts...) {
        st, err := os.Stat(srcpath)
        if err != nil {
            return err
        }
        codestat.Lock()
        if codestat.Files == nil {
            codestat.Files = make(map[string]ruler.CodeFileInfo)
        }
        codestat.Files[getCodeKey(srcpath)] = ruler.CodeFileInfo{st.Size()}
        defer codestat.Unlock()
    }
    return nil
}

// Scans synchronously a directory searching for relevant files to get information.
func loadCodeDirSync(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    files, err := ioutil.ReadDir(srcpath)
    if err != nil {
        return err
    }
    for _, file := range files {
        fullpath := filepath.Join(srcpath, file.Name())
        err := LoadCode(codestat, fullpath, exts...)
        if err != nil {
            return err
        }
    }
    return nil
}

// Scans asynchronously a directory searching for relevant files to get information.
func loadCodeDirAsync(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    files, err := ioutil.ReadDir(srcpath)
    if err != nil {
        return err
    }
    errors := make(chan error, gWorkingLoadersPerRecursionNr)
    load := func(codestat *ruler.CodeStat, srcpath string, exts...string) {
                err := LoadCode(codestat, srcpath, exts...)
                errors <- err
    }
    var chanNr int
    for _, file := range files {
        fullpath := filepath.Join(srcpath, file.Name())
        if chanNr < gWorkingLoadersPerRecursionNr {
            go load(codestat, fullpath, exts...)
            chanNr++
        } else {
            for n := 0; n < gWorkingLoadersPerRecursionNr; n++ {
                err = <-errors
                if err != nil {
                    close(errors)
                    return err
                }
                chanNr--
            }
            go load(codestat, fullpath, exts...)
            chanNr++
        }
    }
    max := chanNr
    for n := 0; n < max; n++ {
        err = <-errors
        if err != nil {
            close(errors)
            return err
        }
    }

    return nil
}

// Uncompresses a zip file to a temporary directory and scans this directory searching
// for relevant files to get information. The temporary directory is always removed.
func loadZippedCode(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    tempdir, errTemp := ioutil.TempDir("", "codeometer-temp")
    if errTemp != nil {
        return errTemp
    }
    defer os.RemoveAll(tempdir)
    zipfile, errZip := zip.OpenReader(srcpath)
    if errZip != nil {
        return errZip
    }
    defer zipfile.Close()
    for _, file := range zipfile.File {
        if file.Mode().IsRegular() {
            fullpath := filepath.Join(tempdir, file.Name)
            os.MkdirAll(filepath.Dir(fullpath), os.ModePerm)
            rc, errRc := file.Open()
            if errRc != nil {
                return errRc
            }
            fileData := bytes.NewBuffer([]byte(""))
            _, errCopy := io.CopyN(fileData, rc, int64(file.UncompressedSize))
            rc.Close()
            if errCopy != nil {
                return errCopy
            }
            errWrite := ioutil.WriteFile(fullpath, fileData.Bytes(), os.ModePerm)
            if errWrite != nil {
                return errWrite
            }
        }
    }
    return LoadCode(codestat, tempdir, exts...)
}

// Clones a git repository recursively into a temporary directory and scans this
// directory searching for relevant code files. The temporary directory is always
// removed.
func loadGitRepoCode(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    if !gHasGit {
        return giveHTTPGetZIPaTry(codestat, srcpath, exts...)
    }
    tempdir, errTemp := ioutil.TempDir("", "codeometer-temp")
    if errTemp != nil {
        return errTemp
    }
    defer os.RemoveAll(tempdir)
    cmd := exec.Command("git", "clone", srcpath, "--recursive", tempdir)
    errChan := make(chan error, 1)
    go func() {
        errChan <- cmd.Run()
    }()
    smt, _ := time.ParseDuration(gSubtaskTimeout)
    select {
        case errCmd := <-errChan:
            if errCmd != nil {
                return errCmd
            }
        case <-time.After(smt):
            cmd.Process.Kill()
            return fmt.Errorf("Git clone aborted due to processing timeout.")
    }
    return LoadCode(codestat, tempdir, exts...)
}

// If Git clone has failed try to download the zip file of this repository.
func giveHTTPGetZIPaTry(codestat *ruler.CodeStat, repoURL string, exts...string) error {
    var URL string
    ru, errURLParse := url.Parse(repoURL)
    if errURLParse != nil {
        return errURLParse
    }
    hostname := ru.Host
    if strings.Contains(hostname, "github.com") {
        URL = repoURL + "/archive/master.zip"
    } else if strings.Contains(hostname, "gitlab.com") {
        URL = repoURL + "/-/archive/master/" + filepath.Base(repoURL) + "-master.zip"
    } else {
        URL = repoURL
    }
    tempdir, errTemp := ioutil.TempDir("", "codeometer-temp")
    if errTemp != nil {
        return errTemp
    }
    defer os.RemoveAll(tempdir)
    content, errGet := http.Get(URL)
    if errGet != nil {
        return errGet
    }
    zipPath := filepath.Join(tempdir, filepath.Base(repoURL) + ".zip")
    file, errFile := os.Create(zipPath)
    if errFile != nil {
        return errFile
    }
    defer file.Close()

    _, errCopy := io.Copy(file, content.Body)
    if errCopy != nil {
        return errCopy
    }
    return loadZippedCode(codestat, zipPath, exts...)
}
