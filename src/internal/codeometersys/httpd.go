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
            if r.Method == "POST" {
                fmt.Fprintf(w, "codeometer post")
            } else {
                fmt.Fprintf(w, "codeometer get")
            }
            break

        default:
            fmt.Fprintf(w, "404 error")
            break
    }
}
