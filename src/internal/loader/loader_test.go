//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package loader

import (
    "testing"
    "runtime"
    "os"
    "internal/ruler"
    "archive/zip"
    "io/ioutil"
    "strings"
)

func TestLoadCode(t *testing.T) {
    files := []string {
        "../../codeometer.go",
        "../../internal/measurer/mi.go",
        "../../internal/measurer/km.go",
        "../../internal/measurer/m.go",
        "../../internal/measurer/mm.go",
        "../../internal/estimator/arcdetriomphe.go",
        "../../internal/estimator/christtheredeemer.go",
        "../../internal/estimator/empirestatebuilding.go",
        "../../internal/estimator/iguazufalls.go",
        "../../internal/estimator/pantheon.go",
        "../../internal/estimator/wallstreet.go",
        "../../internal/estimator/bigbang.go",
        "../../internal/estimator/coliseum.go",
        "../../internal/estimator/estimator.go",
        "../../internal/estimator/libertystatue.go",
        "../../internal/estimator/paulistaavenue.go",
        "../../internal/estimator/washingtonmonument.go",
        "../../internal/estimator/chinesegreatwall.go",
        "../../internal/estimator/eiffeltower.go",
        "../../internal/estimator/frogtraveler.go",
        "../../internal/estimator/niagarafalls.go",
        "../../internal/estimator/sistinechapel.go",
        "../../internal/loader/loader.go",
        "../../internal/loader/loader_test.go",
        "../../internal/loader/getcodekey_linux.go",
        "../../internal/loader/getcodekey_freebsd.go",
        "../../internal/loader/getcodekey_windows.go",
        "../../internal/magnitudes/magnitudes.go",
        "../../internal/ruler/ruler.go",
    }
    codestat := &ruler.CodeStat{}
    err := LoadCode(codestat, "../../", ".go")
    if err != nil {
        t.Errorf(`err != nil: %v`, err)
    }
    for _, file := range files {
        st, err := os.Stat(file)
        if err != nil {
            t.Errorf(`err != nil: %v`, err)
        }
        if runtime.GOOS == "windows" {
            file = strings.Replace(file, "/", "\\", -1)
        }
        info, fileFound := codestat.Files[file]
        if !fileFound {
            t.Errorf(`!fileFound : %v`, file)
        }
        if st.Size() != info.BytesTotal {
            t.Errorf(`st.Size() != info.BytesTotal: %v != %v`, st.Size(), info.BytesTotal)
        }
    }
}

func Test_isRelevantFile(t *testing.T) {
    type TestVectorCtx struct {
        Filename string
        Exts []string
        Expected bool
    }
    testVector := []TestVectorCtx {
        { "main.c", []string{".c", ".h"}, true },
        { "main.c", []string{".C", ".h"}, true },
        { "main.c", []string{"c", "h"}, true },
        { "mainc", []string{"c", "h"}, false },
        { "main.c", []string{"cc", "h"}, false },
        { "main.cc", []string{".c", ".h"}, false },
        { "main.cpp", []string{".c", ".h"}, false },
        { "main.go", []string{".c", ".cpp", ".go", ".pas", ".h", ".hpp"}, true },
        { "main.HpP", []string{".c", ".HPP"}, true },
    }
    for _, test := range testVector {
        if isRelevantFile(test.Filename, test.Exts...) != test.Expected {
            t.Errorf(`isRelevantFile(test.Filename, test.Exts...) != test.Expected`)
            t.Errorf(`test.Filename: %v / test.Exts: %v / test.Expected: %v`, test.Filename, test.Exts, test.Expected)
        }
    }
}

func Test_getCodeKey(t *testing.T) {
    type TestVectorCtx struct {
        SrcPath, Expected string
    }
    var testVector []TestVectorCtx
    if runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
        testVector = []TestVectorCtx {
            { "/tmp/codeometer-temp0120392/src/abc.d", "src/abc.d" },
            { "/tmp/src/abc.d", "/tmp/src/abc.d" },
        }
    } else if runtime.GOOS == "windows" {
        testVector = []TestVectorCtx {
            { "C:\\temp\\codeometer-temp9389127312\\src\\abc.d", "src\\abc.d" },
            { "C:\\temp\\src\\abc.d", "C:\\temp\\src\\abc.d" },
        }
    } else {
        t.Errorf("Unsupported platform: %s.\n", runtime.GOOS)
        t.Fail()
    }
    for _, test := range testVector {
        key := getCodeKey(test.SrcPath)
        if key != test.Expected {
            t.Errorf(`key != test.Expected : %v != %v`, key, test.Expected)
        }
    }
}

func Test_loadCodeFile(t *testing.T) {
    files := []string {
        "loader_test.go",
        "loader.go",
        "getcodekey_linux.go",
        "getcodekey_freebsd.go",
        "getcodekey_windows.go",
    }
    for _, file := range files {
        st, errStat := os.Stat(file)
        if errStat != nil {
            t.Errorf(`errStat != nil: %v`, errStat)
        }
        codestat := &ruler.CodeStat{}
        err := loadCodeFile(codestat, file, ".go")
        if err != nil {
            t.Errorf(`err != nil: %v`, err)
        }
        info, fileFound := codestat.Files[getCodeKey(file)]
        if !fileFound {
            t.Error(`!fileFound`)
        }
        if info.BytesTotal != st.Size() {
            t.Errorf(`info.BytesTotal != st.Size(): %v != %v`, info.BytesTotal, st.Size())
        }
    }
}

func Test_loadCodeDirSync(t *testing.T) {
    files := []string {
        "../../codeometer.go",
        "../../internal/measurer/mi.go",
        "../../internal/measurer/km.go",
        "../../internal/measurer/m.go",
        "../../internal/measurer/mm.go",
        "../../internal/estimator/arcdetriomphe.go",
        "../../internal/estimator/christtheredeemer.go",
        "../../internal/estimator/empirestatebuilding.go",
        "../../internal/estimator/iguazufalls.go",
        "../../internal/estimator/pantheon.go",
        "../../internal/estimator/wallstreet.go",
        "../../internal/estimator/bigbang.go",
        "../../internal/estimator/coliseum.go",
        "../../internal/estimator/estimator.go",
        "../../internal/estimator/libertystatue.go",
        "../../internal/estimator/paulistaavenue.go",
        "../../internal/estimator/washingtonmonument.go",
        "../../internal/estimator/chinesegreatwall.go",
        "../../internal/estimator/eiffeltower.go",
        "../../internal/estimator/frogtraveler.go",
        "../../internal/estimator/niagarafalls.go",
        "../../internal/estimator/sistinechapel.go",
        "../../internal/loader/loader.go",
        "../../internal/loader/loader_test.go",
        "../../internal/loader/getcodekey_linux.go",
        "../../internal/loader/getcodekey_freebsd.go",
        "../../internal/loader/getcodekey_windows.go",
        "../../internal/magnitudes/magnitudes.go",
        "../../internal/ruler/ruler.go",
    }
    codestat := &ruler.CodeStat{}
    err := loadCodeDirSync(codestat, "../../", ".go")
    if err != nil {
        t.Errorf(`err != nil: %v`, err)
    }
    for _, file := range files {
        st, err := os.Stat(file)
        if err != nil {
            t.Errorf(`err != nil: %v`, err)
        }
        if runtime.GOOS == "windows" {
            file = strings.Replace(file, "/", "\\", -1)
        }
        info, fileFound := codestat.Files[file]
        if !fileFound {
            t.Errorf(`!fileFound : %v`, file)
        }
        if st.Size() != info.BytesTotal {
            t.Errorf(`st.Size() != info.BytesTotal: %v != %v`, st.Size(), info.BytesTotal)
        }
    }
}

func Test_loadCodeDirAsync(t *testing.T) {
    files := []string {
        "../../codeometer.go",
        "../../internal/measurer/mi.go",
        "../../internal/measurer/km.go",
        "../../internal/measurer/m.go",
        "../../internal/measurer/mm.go",
        "../../internal/estimator/arcdetriomphe.go",
        "../../internal/estimator/christtheredeemer.go",
        "../../internal/estimator/empirestatebuilding.go",
        "../../internal/estimator/iguazufalls.go",
        "../../internal/estimator/pantheon.go",
        "../../internal/estimator/wallstreet.go",
        "../../internal/estimator/bigbang.go",
        "../../internal/estimator/coliseum.go",
        "../../internal/estimator/estimator.go",
        "../../internal/estimator/libertystatue.go",
        "../../internal/estimator/paulistaavenue.go",
        "../../internal/estimator/washingtonmonument.go",
        "../../internal/estimator/chinesegreatwall.go",
        "../../internal/estimator/eiffeltower.go",
        "../../internal/estimator/frogtraveler.go",
        "../../internal/estimator/niagarafalls.go",
        "../../internal/estimator/sistinechapel.go",
        "../../internal/loader/loader.go",
        "../../internal/loader/loader_test.go",
        "../../internal/loader/getcodekey_linux.go",
        "../../internal/loader/getcodekey_freebsd.go",
        "../../internal/loader/getcodekey_windows.go",
        "../../internal/magnitudes/magnitudes.go",
        "../../internal/ruler/ruler.go",
    }
    codestat := &ruler.CodeStat{}
    err := loadCodeDirAsync(codestat, "../../", ".go")
    if err != nil {
        t.Errorf(`err != nil: %v`, err)
    }
    for _, file := range files {
        st, err := os.Stat(file)
        if err != nil {
            t.Errorf(`err != nil: %v`, err)
        }
        if runtime.GOOS == "windows" {
            file = strings.Replace(file, "/", "\\", -1)
        }
        info, fileFound := codestat.Files[file]
        if !fileFound {
            t.Errorf(`!fileFound : %v`, file)
        }
        if st.Size() != info.BytesTotal {
            t.Errorf(`st.Size() != info.BytesTotal: %v != %v`, st.Size(), info.BytesTotal)
        }
    }
}

func Test_loadZippedCode(t *testing.T) {
    files := []string{
        "loader_test.go",
        "loader.go",
        "getcodekey_linux.go",
        "getcodekey_freebsd.go",
        "getcodekey_windows.go",
    }
    buf, _ := os.Create("src.zip")
    defer buf.Close()
    defer os.Remove("src.zip")
    wzip := zip.NewWriter(buf)
    for _, file := range files {
        f, _ := wzip.Create(file)
        data, _ := ioutil.ReadFile(file)
        f.Write(data)
    }
    wzip.Close()
    codestat := &ruler.CodeStat{}
    err := loadZippedCode(codestat, "src.zip", ".go")
    if err != nil {
        t.Errorf(`err != nil: %v`, err)
    }
    for _, file := range files {
        st, err := os.Stat(file)
        if err != nil {
            t.Error(`err != nil: %v`, err)
        }
        if runtime.GOOS == "windows" {
            file = strings.Replace(file, "/", "\\", -1)
        }
        info, fileFound := codestat.Files[file]
        if !fileFound {
            t.Errorf(`!fileFound: %v`, file)
            t.Error(codestat.Files)
        }
        if st.Size() != info.BytesTotal {
            t.Errorf(`st.Size() != info.BytesTotal: %v != %v`, st.Size(), info.BytesTotal)
        }
    }
}

func OffTest_loadGitRepoCode(t *testing.T) {
    files := []struct {
        Filepath, Key string
    }{
        {"../../codeometer.go", "src/codeometer.go"},
        {"../../internal/measurer/mi.go", "src/internal/measurer/mi.go"},
        {"../../internal/measurer/km.go", "src/internal/measurer/km.go"},
        {"../../internal/measurer/m.go", "src/internal/measurer/m.go"},
        {"../../internal/measurer/mm.go", "src/internal/measurer/mm.go"},
        {"../../internal/estimator/arcdetriomphe.go", "src/internal/estimator/arcdetriomphe.go"},
        {"../../internal/estimator/christtheredeemer.go", "src/internal/estimator/christtheredeemer.go"},
        {"../../internal/estimator/empirestatebuilding.go", "src/internal/estimator/empirestatebuilding.go"},
        {"../../internal/estimator/iguazufalls.go", "src/internal/estimator/iguazufalls.go"},
        {"../../internal/estimator/pantheon.go", "src/internal/estimator/pantheon.go"},
        {"../../internal/estimator/wallstreet.go", "src/internal/estimator/wallstreet.go"},
        {"../../internal/estimator/bigbang.go", "src/internal/estimator/bigbang.go"},
        {"../../internal/estimator/coliseum.go", "src/internal/estimator/coliseum.go"},
        {"../../internal/estimator/estimator.go", "src/internal/estimator/estimator.go"},
        {"../../internal/estimator/libertystatue.go", "src/internal/estimator/libertystatue.go"},
        {"../../internal/estimator/paulistaavenue.go", "src/internal/estimator/paulistaavenue.go"},
        {"../../internal/estimator/washingtonmonument.go", "src/internal/estimator/washingtonmonument.go"},
        {"../../internal/estimator/chinesegreatwall.go", "src/internal/estimator/chinesegreatwall.go"},
        {"../../internal/estimator/eiffeltower.go", "src/internal/estimator/eiffeltower.go"},
        {"../../internal/estimator/frogtraveler.go", "src/internal/estimator/frogtraveler.go"},
        {"../../internal/estimator/niagarafalls.go", "src/internal/estimator/niagarafalls.go"},
        {"../../internal/estimator/sistinechapel.go", "src/internal/estimator/sistinechapel.go"},
        {"../../internal/loader/loader.go", "src/internal/loader/loader.go"},
        {"../../internal/loader/loader_test.go", "src/internal/loader/loader_test.go"},
        {"../../internal/loader/getcodekey_linux.go", "src/internal/loader/getcodekey_linux.go"},
        {"../../internal/loader/getcodekey_freebsd.go", "src/internal/loader/getcodekey_freebsd.go"},
        {"../../internal/loader/getcodekey_windows.go", "src/internal/loader/getcodekey_windows.go"},
        {"../../internal/magnitudes/magnitudes.go", "src/internal/magnitudes/magnitudes.go"},
        {"../../internal/ruler/ruler.go", "src/internal/ruler/ruler.go"},
    }
    codestat := &ruler.CodeStat{}
    err := loadGitRepoCode(codestat, "https://github.com/rafael-santiago/codeometer", ".go")
    if err != nil {
        t.Errorf(`err != nil: %v`, err)
    }
    for _, file := range files {
        st, err := os.Stat(file.Filepath)
        if err != nil {
            t.Errorf(`err != nil: %v`, err)
        }
        if runtime.GOOS == "windows" {
            file.Key = strings.Replace(file.Key, "/", "\\", -1)
        }
        info, fileFound := codestat.Files[file.Key]
        if !fileFound {
            t.Errorf(`!fileFound: %v`, file.Key)
        }
        if info.BytesTotal != st.Size() {
            t.Errorf(`info.BytesTotal != st.Size(): %v != %v`, info.BytesTotal, st.Size())
        }
    }
}

