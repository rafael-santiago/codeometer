// package loader concentrates all code loading stuff, however all you should call is LoadCode().
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package loader

import (
    "internal/ruler"
    "os"
    "fmt"
    "strings"
    "path/filepath"
    "io/ioutil"
    "archive/zip"
    "bytes"
    "io"
    "os/exec"
)

// This function expects as srcpath a file, directory or git-repo uri. At the end
// of a successful execution codestat will gather all relevant files info. The
// relevance a code file is given by exts... where you pass all relevant file
// extensions, including the dot symbol or not.
func LoadCode(codestat *ruler.CodeStat, srcpath string, exts...string) error {
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
            loader = loadCodeDir
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

// Scans a directory searching for relevant files to get information.
func loadCodeDir(codestat *ruler.CodeStat, srcpath string, exts...string) error {
    files, err := ioutil.ReadDir(srcpath)
    if err != nil {
        return err
    }
    errors := make(chan error, len(files))
    for _, file := range files {
        fullpath := filepath.Join(srcpath, file.Name())
        go func() {
            err := LoadCode(codestat, fullpath, exts...)
            errors <- err
        }()
    }
    for range files {
        if err := <-errors; err != nil {
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
    tempdir, errTemp := ioutil.TempDir("", "codeometer-temp")
    if errTemp != nil {
        return errTemp
    }
    defer os.RemoveAll(tempdir)
    cmd := exec.Command("git", "clone", srcpath, "--recursive", tempdir)
    errCmd := cmd.Run()
    if errCmd != nil {
        return errCmd
    }
    return LoadCode(codestat, tempdir, exts...)
}
