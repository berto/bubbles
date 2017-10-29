package routes

import (
	"io/ioutil"
	"log"
	"net/http"
)

type hookedResponseWriter struct {
	http.ResponseWriter
	ignore bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == 404 {
		hrw.ResponseWriter.Header().Set("Content-Type", "text/html")
		index, err := ioutil.ReadFile("./public/dist/index.html")
		hrw.ignore = true
		if err != nil {
			log.Fatal(err)
		}
		hrw.ResponseWriter.WriteHeader(200)
		hrw.ResponseWriter.Write(index)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.ignore {
		return len(p), nil
	}
	return hrw.ResponseWriter.Write(p)
}

type NotFoundHook struct {
	H http.Handler
}

func (nfh NotFoundHook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nfh.H.ServeHTTP(&hookedResponseWriter{ResponseWriter: w}, r)
}
