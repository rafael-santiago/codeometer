// package codeometersys - gathers default values and useful application system functions.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "fmt"
    "os"
    "net/http"
    "internal/options"
    "os/signal"
    "syscall"
    "net/url"
    "strings"
    "path/filepath"
    "io/ioutil"
    "encoding/base64"
    "regexp"
    "runtime"
)

// The 'httpd' command handler.
func httpd() int {
    addr := options.GetOption("peer-addr", "")
    if len(addr) == 0 {
        fmt.Fprintf(os.Stderr, "error: --peer-addr option is missing.\n")
        return 1
    }
    go listen(addr)
    sigintWatchdog := make(chan os.Signal, 1)
    signal.Notify(sigintWatchdog, os.Interrupt)
    signal.Notify(sigintWatchdog, syscall.SIGINT|syscall.SIGTERM)
    <-sigintWatchdog
    fmt.Fprintf(os.Stdout, "\ninfo: codeometer finished.\n")
    return 0
}

// The 'httpd' command helper.
func httpdHelp() int {
    fmt.Fprintf(os.Stdout, "use: codeometer httpd --peer-addr=<host:port>\n"+
                           "                       [--cert=<cert file path> --key=<key file path>\n" +
                           "                        --async --loaders-nr=<number> --subtask-timeout=<duration string>]\n")
    return 0
}

// Initializes the HTTPd and listens to connections.
func listen(peerAddr string) {
    cert := options.GetOption("cert", "")
    key := options.GetOption("key", "")
    if len(key) > 0 && len(cert) == 0 {
        fmt.Fprintf(os.Stderr, "error: --cert option is missing.\n")
        os.Exit(1)
    } else if len(cert) > 0 && len(key) == 0 {
        fmt.Fprintf(os.Stderr, "error: --key option is missing.\n")
        os.Exit(1)
    }

    http.HandleFunc("/", handle)
    var err error
    if len(key) == 0 {
        err = http.ListenAndServe(peerAddr, nil)
    } else {
        err = http.ListenAndServeTLS(peerAddr, cert, key, nil)
    }
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s.\n", err)
        os.Exit(1)
    }
}

// Handle all HTTPd requests.
func handle(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "/codeometer":
            r.ParseForm()
            if r.Method == "POST" {
                htmlOut := webInterface
                //measureReport(src string, exts []string, fontSize string, wantedMeasures []string,
                //  statsPerFile bool, estimatives bool)
                src := r.Form.Get("src")
                rawMeasures := r.Form.Get("measures")
                rawMeasures = strings.Replace(rawMeasures, " ", "", -1)
                measures := strings.Split(rawMeasures, ",")
                rawExts := r.Form.Get("exts")
                rawExts = strings.Replace(rawExts, " ", "", -1)
                exts := strings.Split(rawExts, ",")
                data := r.Form.Get("data")
                statsPerFile := (r.Form.Get("statsPerFile") == "1")
                estimatives := (r.Form.Get("estimatives") == "1")
                fontSize := r.Form.Get("fontSize")
                r.Form.Set("waitImage", waitImage)
                // INFO(Rafael): Restoring user field values at web interface.
                r.Form.Set("edtQuery", src)
                r.Form.Set("edtExt", rawExts)
                if estimatives {
                    r.Form.Set("chkEstimatives", "checked")
                } else {
                    r.Form.Set("chkEstimatives", "")
                }
                if statsPerFile {
                    r.Form.Set("chkStatsPerFile", "checked")
                } else {
                    r.Form.Set("chkStatsPerFile", "")
                }
                if hasItem(measures, "km") {
                    r.Form.Set("chkKM", "checked")
                } else {
                    r.Form.Set("chkKM", "")
                }
                if hasItem(measures, "mi") {
                    r.Form.Set("chkMI", "checked")
                } else {
                    r.Form.Set("chkMI", "")
                }
                if hasItem(measures, "m") {
                    r.Form.Set("chkM", "checked")
                } else {
                    r.Form.Set("chkM", "")
                }
                if hasItem(measures, "mm") {
                    r.Form.Set("chkMM", "checked")
                } else {
                    r.Form.Set("chkMM", "")
                }
                if fontSize == "12px" {
                    r.Form.Set("fontSize12px", "selected")
                    r.Form.Set("fontSize10px", "")
                } else if fontSize == "10px" {
                    r.Form.Set("fontSize10px", "selected")
                    r.Form.Set("fontSize12px", "")
                }
                if len(src) == 0 {
                    r.Form.Add("div-type", "error")
                    r.Form.Add("info", "You need to specify a Git repo URL or upload a source code or zip file.")
                } else {
                    if len(data) == 0 {
                        // INFO(Rafael): A Git repo url was given.
                        makeHTMLReport(src, exts, fontSize, measures, statsPerFile, estimatives, &r.Form)
                    } else {
                        // INFO(Rafael): A file was uploaded.
                        tempDir, errTemp := ioutil.TempDir("", "codeometer-temp")
                        if errTemp != nil {
                            r.Form.Add("div-type", "error")
                            r.Form.Add("info", errTemp.Error())
                        } else {
                            defer os.RemoveAll(tempDir)
                            data, err := base64.StdEncoding.DecodeString(data)
                            if err != nil {
                                r.Form.Add("div-type", "error")
                                r.Form.Add("info", err.Error())
                            } else {
                                fileName := filepath.Base(src)
                                if runtime.GOOS != "windows" {
                                    pattern := regexp.MustCompile(`.*\\`)
                                    fileName = string(pattern.ReplaceAll([]byte(fileName), []byte("")))
                                }
                                src = filepath.Join(tempDir, fileName)
                                err := ioutil.WriteFile(src, data, os.ModePerm)
                                if err != nil {
                                    r.Form.Add("div-type", "error")
                                    r.Form.Add("info", err.Error())
                                } else {
                                    makeHTMLReport(src, exts, fontSize, measures, statsPerFile, estimatives, &r.Form,
                                                   tempDir, fileName)
                                    r.Form.Set("edtQuery", "")
                                }
                            }
                        }
                    }
                    htmlOut = expandTemplateActions(webInterface, r.Form)
                    fmt.Fprintf(w, "%s", htmlOut)
                }
            } else {
                r.Form.Set("edtQuery", "")
                r.Form.Set("edtExt", "")
                r.Form.Set("chkEstimatives", "")
                r.Form.Set("chkStatsPerFile", "")
                r.Form.Set("moreDiv", "none")
                r.Form.Add("div-type", "single-info")
                r.Form.Add("info", "")
                r.Form.Add("chkKM", "checked")
                r.Form.Add("chkMI", "")
                r.Form.Add("chkM", "")
                r.Form.Add("chkMM", "")
                r.Form.Add("fontSize12px", "selected")
                r.Form.Add("fontSize10px", "")
                r.Form.Set("waitImage", waitImage)
                fmt.Fprintf(w, "%s", expandTemplateActions(webInterface, r.Form))
            }
            break

        default:
            r.ParseForm()
            r.Form.Set("statusImage", webStatusImage)
            r.Form.Set("error", "404 Not Found")
            var home string
            if options.GetOption("cert", "") != "" {
                home = "https://"
            } else {
                home = "http://"
            }
            home += options.GetOption("peer-addr", "") + "/codeometer"
            r.Form.Set("home", home)
            fmt.Fprintf(w, "%s", expandTemplateActions(webErrorPage, r.Form))
            break
    }
}

// Expands all template actions from the passed template based on passed url.Values.
func expandTemplateActions(template string, userData url.Values) string {
    expandedData := template
    for k, _ := range userData {
        action := "{{." + k + "}}"
        data := userData.Get(k)
        expandedData = strings.Replace(expandedData, action, data, -1)
    }
    return expandedData
}

// Verifies the existence of a passed item inside the passed list.
func hasItem(list []string, item string) bool {
    for _, l := range list {
        if l == item {
            return true
        }
    }
    return false
}

// Gets plain text report and make it a HTML report.
func makeHTMLReport(src string, exts []string, fontSize string, wantedMeasures []string,
                    statsPerFile bool, estimatives bool, userData *url.Values, tempDirAndFileName...string) {
    var effectiveSrc string
    if len(tempDirAndFileName) < 2 {
        effectiveSrc = src
    } else {
        wd, err := os.Getwd()
        if err != nil {
            defer os.Chdir(wd)
        }
        os.Chdir(tempDirAndFileName[0])
        effectiveSrc = tempDirAndFileName[1]
    }
    info, err := measureReport(effectiveSrc, exts, fontSize, wantedMeasures, statsPerFile, estimatives)
    if err == nil {
        info := strings.Trim(info, "\n\n")
        if estimatives || statsPerFile {
            userData.Add("div-type", "info")
        } else {
            userData.Add("div-type", "single-info")
        }
        userData.Add("info", "&nbsp;" + filepath.Base(src) + " has " +
                             strings.Replace(info, "\n", "<br>&nbsp;", -1))
    } else {
        userData.Add("div-type", "error")
        errMsg := err.Error()
        if !strings.HasSuffix(errMsg, ".") {
            errMsg += "."
        }
        userData.Add("info", errMsg + " Unable to measure '" + src + "'.")
    }
}
