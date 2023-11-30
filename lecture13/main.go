package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/static/"):]
	filePath := filepath.Join("files", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", http.DetectContentType(make([]byte, 512)))
	w.Header().Set("Content-Length", string(fileInfo.Size()))

	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}

func main() {
	http.HandleFunc("/static/", staticHandler)

	http.ListenAndServe(":8080", nil)
}
