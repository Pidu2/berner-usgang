package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func proxyHandler(target string) http.Handler {
    url, _ := url.Parse(target)
    return httputil.NewSingleHostReverseProxy(url)
}

func main() {
    // Serve static files for your frontend
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    // Proxy API requests to the backend
    http.Handle("/api/", http.StripPrefix("/api", proxyHandler("http://localhost:8080")))

    log.Println("Starting proxy server on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

