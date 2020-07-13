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
)

var gFontSize string
var gWantedMeasures []string

func init() {
    gWantedMeasures = options.GetArrayOption("measures", "km")
    gFontSize = options.GetOption("font-size", "12px")
}

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
                           "                        --async --loaders-nr=<number>]\n")
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
                //measureReport(src string, exts []string, fontSize string, wantedMeasures []string,
                //  statsPerFile bool, estimatives bool)
                src := r.Form.Get("src")
                if len(src) == 0 {
                    r.Form.Add("div-type", "error")
                    r.Form.Add("info", "You need to specify a Git repo URL or upload a source code or zip file.")
                    fmt.Fprintf(w, "%s", expandTemplateActions(webInterface, r.Form))
                    return
                }
                rawExts := r.Form.Get("exts")
                rawExts = strings.Replace(rawExts, " ", "", -1)
                exts := strings.Split(rawExts, ",")
                data := r.Form.Get("data")
                statsPerFile := (r.Form.Get("statsPerFile") == "1")
                estimatives := (r.Form.Get("estimatives") == "1")
                if len(data) == 0 {
                    // INFO(Rafael): A Git repo url was given.
                    info, err := measureReport(src, exts, gFontSize, gWantedMeasures, statsPerFile, estimatives)
                    if err == nil {
                        info := strings.Trim(info, "\n\n")
                        if estimatives || statsPerFile {
                            r.Form.Add("div-type", "info")
                        } else {
                            r.Form.Add("div-type", "single-info")
                        }
                        r.Form.Add("info", "&nbsp;" + filepath.Base(src) + " has " + strings.Replace(info, "\n", "<br>&nbsp;", -1))
                    } else {
                        r.Form.Add("div-type", "error")
                        r.Form.Add("info", err.Error() + ". Unable to access '" + src + "'.")
                    }
                }
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
                fmt.Fprintf(w, "%s", expandTemplateActions(webInterface, r.Form))
            } else {
                r.Form.Set("edtQuery", "")
                r.Form.Set("edtExt", "")
                r.Form.Set("chkEstimatives", "")
                r.Form.Set("chkStatsPerFile", "")
                r.Form.Set("moreDiv", "none")
                r.Form.Add("div-type", "single-info")
                r.Form.Add("info", "")
                fmt.Fprintf(w, "%s", expandTemplateActions(webInterface, r.Form))
            }
            break

        default:
            fmt.Fprintf(w, "404 error")
            break
    }
}

func expandTemplateActions(template string, userData url.Values) string {
    expandedData := template
    for k, _ := range userData {
        action := "{{." + k + "}}"
        data := userData.Get(k)
        expandedData = strings.Replace(expandedData, action, data, -1)
    }
    return expandedData
}
