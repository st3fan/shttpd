package main

import (
	"flag"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type server struct {
	root string
	mux  *http.ServeMux
}

func (s *server) walker(path string, info os.FileInfo, err error) error {
	if strings.HasPrefix(path, ".") || info.IsDir() {
		return nil
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	contentType := mime.TypeByExtension(filepath.Ext(path))
	if contentType == "" {
		contentType = "application/binary"
	}

	f := func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", contentType)
		res.Write(bytes)
	}

	s.mux.HandleFunc("/"+path, f)

	if strings.HasSuffix(path, "/index.html") {
		s.mux.HandleFunc("/"+strings.TrimSuffix(path, "index.html"), f)
	}

	log.Println("Registered", path)

	return nil
}

func NewServer(root string) (*server, error) {
	s := &server{
		root: root,
		mux:  http.NewServeMux(),
	}

	if err := os.Chdir(root); err != nil {
		return nil, err
	}

	if err := filepath.Walk(".", s.walker); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//log.Println(req.URL.String())
	s.mux.ServeHTTP(w, req)
}

func main() {
	root := flag.String("root", ".", "document root")
	addr := flag.String("addr", ":8080", "address to bind to")

	flag.Parse()

	s, err := NewServer(*root)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	if err := http.ListenAndServe(*addr, s); err != nil {
		log.Fatal("Failed to listen and server:", err)
	}
}
