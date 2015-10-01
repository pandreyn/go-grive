package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

  "github.com/gorilla/mux"
  "strconv"
)

var (
	hostname     string
	port         int
	topStaticDir string
)

func init() {
	// Flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [default_static_dir]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&hostname, "h", "localhost", "hostname")
	flag.IntVar(&port, "p", 4001, "port")
	flag.StringVar(&topStaticDir, "static_dir", "", "static directory in addition to default static directory")
}

func appendStaticRoute(sr StaticRoutes, dir string) StaticRoutes {
	if _, err := os.Stat(dir); err != nil {
		log.Fatal(err)
	}
	return append(sr, http.Dir(dir))
}

type StaticRoutes []http.FileSystem

func (sr StaticRoutes) Open(name string) (f http.File, err error) {
	for _, s := range sr {
		if f, err = s.Open(name); err == nil {
			f = disabledDirListing{f}
			return
		}
	}
	return
}

type disabledDirListing struct {
	http.File
}

func (f disabledDirListing) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func Search(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Gorilla!\n"))
}

func main() {
	// Parse flags
	flag.Parse()
	staticDir := flag.Arg(0)

	//getFilesFromGDrive()

	// Setup static routes
	staticRoutes := make(StaticRoutes, 0)
	if topStaticDir != "" {
		staticRoutes = appendStaticRoute(staticRoutes, topStaticDir)
	}
	if staticDir == "" {
		staticDir = "./"
	}
	staticRoutes = appendStaticRoute(staticRoutes, staticDir)

  err := openBrowser(hostname, port)
  if err != nil {
    log.Fatalf("Unable to open browser.", err)
  }


  r := mux.NewRouter()
  r.HandleFunc("/search/{searchTerm}", Search)

  r.PathPrefix("/").Handler(http.FileServer(staticRoutes))
  http.Handle("/", r)
  address := ":"+strconv.Itoa(port)
  log.Fatal(http.ListenAndServe(address, r))

}
