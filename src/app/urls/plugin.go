package urls

import (
    "fmt"
    "github.com/CarsonSlovoka/painter/app/server"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
)

type PluginViewHandler struct{}

func (vh *PluginViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:] // 0 is "/"
    log.Println(path)

    data, err := ioutil.ReadFile(path)

    if err != nil {
        w.WriteHeader(404)
        _, _ = w.Write([]byte("404 - " + http.StatusText(404)))
        return
    }

    contentTypeMap := map[string]string{
        "html":  "text/html",
        "css":   "text/css",
        "js":    "application/javascript",
        "png":   "image/png",
        "image": "image/jpeg",
        "svg":   "image/svg+xml",
        "mp4":   "video/mp4",
        "ico":   "image/x-icon",
    }

    contentType := filepath.Ext(path)[1:]
    if contentType, exists := contentTypeMap[contentType]; exists {
        w.Header().Add("Content-Type", contentType)
        _, _ = w.Write(data)
        return
    }

    w.WriteHeader(http.StatusNotAcceptable)
    _, _ = w.Write([]byte(fmt.Sprintf("406 - %s - %s", http.StatusText(406), contentType)))

}

func initPlugin() {
    pluginViewHandler := &PluginViewHandler{}
    /*
       pluginRouter := server.Mux.PathPrefix("/plugin/").Subrouter()
       pluginRouter.Handle("/{all}", http.StripPrefix("/plugin/", pluginViewHandler)) // 這個只能單一層
    */

    server.Mux.PathPrefix("/plugin/").Handler(
        http.StripPrefix("/plugin", pluginViewHandler),
    )
}
