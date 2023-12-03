package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"sync"
	"time"
)

//func staticHandler(w http.ResponseWriter, r *http.Request) {
//	fileName := chi.URLParam(r, "filename")
//	filePath := filepath.Join("files", fileName)
//
//	file, err := os.Open(filePath)
//	if err != nil {
//		http.NotFound(w, r)
//		return
//	}
//	defer file.Close()
//
//	fileInfo, err := file.Stat()
//	if err != nil {
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", http.DetectContentType(make([]byte, 512)))
//	w.Header().Set("Content-Length", string(fileInfo.Size()))
//
//	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
//}

func main() {
	// Основной роутер
	go func() {
		log.Printf("debug server is starting\n")
		log.Fatal(http.ListenAndServe("localhost:8081", profiler()))
	}()

	log.Printf("main server is starting\n")

	http.HandleFunc("/test", someTask())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func profiler() http.Handler {
	r := chi.NewRouter()
	r.Mount("/debug", middleware.Profiler())
	return r
}

func someTask() http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		wg := sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				log.Printf("goroutine %d", n)
				time.Sleep(5 * time.Minute)
			}(i)
		}
		wg.Wait()
	}
}
