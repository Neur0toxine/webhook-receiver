package main

import (
    "fmt"
    "io"
    "log"
    "net/http"

    "github.com/fatih/color"
)

type Handler struct {
    body func(a ...interface{}) string
    method func(a ...interface{}) string
    keyPrint func(a ...interface{}) string
    prefixPrint func(a ...interface{}) string
}

func NewHandler() *Handler {
    return &Handler{
        body: color.New(color.Bold, color.FgHiGreen).SprintFunc(),
        method: color.New(color.Bold, color.FgBlack, color.BgGreen).SprintFunc(),
        keyPrint: color.New(color.FgYellow).SprintFunc(),
        prefixPrint: color.New(color.Faint, color.FgCyan).SprintFunc(),
    }
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s\n",
        color.HiMagentaString(r.Proto), h.method(r.Method), color.CyanString(r.URL.String()))
    log.Printf(" > %s: %s\n", h.keyPrint("RemoteAddr"), r.RemoteAddr)
    log.Printf(" > %s: %s\n", h.keyPrint("Host"), r.Host)
    log.Printf(" > %s: %d\n", h.keyPrint("Length"), r.ContentLength)

    for _, encoding := range r.TransferEncoding {
        log.Printf(" > %s: %s\n", h.keyPrint("Transfer-Encoding"), encoding)
    }

    for header, values := range r.Header {
        for _, value := range values {
            log.Printf("%s %s: %s\n", h.prefixPrint("[header]"), h.keyPrint(header), value)
        }
    }

    if r.Body == nil {
        log.Printf("%s\n", color.HiMagentaString("No body is present."))
    }

    data, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("%s\n", color.HiRedString("Cannot read body: " + err.Error()))
    }

    if len(data) == 0 {
        log.Printf("%s\n", color.HiMagentaString("Body is empty."))
    } else {
        log.Printf("%s\n%s\n", h.body("Body:"), string(data))
    }

    for header, values := range r.Trailer {
        for _, value := range values {
            log.Printf("%s %s: %s\n", h.prefixPrint("[trailer header]"), h.keyPrint(header), value)
        }
    }

    fmt.Println()

	w.WriteHeader(http.StatusOK)
}
