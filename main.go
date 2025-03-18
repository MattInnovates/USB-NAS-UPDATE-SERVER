package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// handler serves files and provides directory browsing.
func handler(w http.ResponseWriter, r *http.Request) {
	publicDir := "./public"
	// Check that the public directory exists.
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		http.Error(w, "Error: public directory does not exist", http.StatusNotFound)
		return
	}

	// Clean the request path and join with publicDir.
	reqPath := filepath.Clean(r.URL.Path)
	fullPath := filepath.Join(publicDir, reqPath)

	// Check if the path exists.
	fi, err := os.Stat(fullPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fi.IsDir() {
		// If directory, list files.
		files, err := os.ReadDir(fullPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<html><body><h1>Index of %s</h1><ul>", r.URL.Path)

		// Link to parent directory if not at root.
		if r.URL.Path != "/" {
			parent := filepath.Dir(r.URL.Path)
			// Ensure parent always starts with "/"
			if !strings.HasPrefix(parent, "/") {
				parent = "/" + parent
			}
			fmt.Fprintf(w, `<li><a href="%s">..</a></li>`, parent)
		}

		// List each file or directory.
		for _, file := range files {
			name := file.Name()
			displayName := name
			linkPath := filepath.Join(r.URL.Path, name)
			// Make sure URL paths use forward slashes.
			linkPath = filepath.ToSlash(linkPath)
			if file.IsDir() {
				displayName = name + "/"
				linkPath = linkPath + "/"
			}
			fmt.Fprintf(w, `<li><a href="%s">%s</a></li>`, linkPath, displayName)
		}
		fmt.Fprint(w, "</ul></body></html>")
	} else {
		// If file, serve it.
		http.ServeFile(w, r, fullPath)
	}
}

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)

	fmt.Printf("Server is running on port %d...\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
