package main

import (
    "fmt"
    "github.com/CarsonSlovoka/painter/app"
    "github.com/CarsonSlovoka/painter/app/log"
    "github.com/CarsonSlovoka/painter/app/server"
    "github.com/CarsonSlovoka/painter/app/urls"
    "os/exec"
)

func main() {
    file := log.InitLog("painter.temp.log")
    if file != nil {
        defer func() {
            log.Trace.Printf("Exit App.")
            _ = file.Close()
        }()
    }
    app.LoadConfig()

    quit := make(chan bool)
    port := app.Config.Server.Port
    go func() {
        urls.InitURLs()
        if err := server.ListenAndServe(port); err != nil {
            log.Trace.Println(err)
        }
        quit <- true
    }()

    rootURL := fmt.Sprintf("http://localhost:%d", port)
    go func() {
        if err := exec.Command("rundll32", "url.dll,FileProtocolHandler",
            rootURL,
        ).Start(); err != nil {
            panic(err)
        }
    }()

    for {
        select {
        case <-quit:
            // log.Println("Close App.")
            return
        }
    }
}
