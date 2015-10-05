package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
  "time"
  "fmt"
  "runtime"
  "os/exec"
)

//flags
var (
	webDir  string
  webHost string
	webPort int
)

func init() {
	log.Println("Initializing...")
	//web flags
	flag.StringVar(&webDir, "webdir", "web/dist", "Change the web directory")
  flag.StringVar(&webHost, "webHost", "localhost", "Change the host")
	flag.IntVar(&webPort, "wport", 4444, "Change port for the web server to listen on")

	flag.Parse()

	webDir = filepath.Clean(webDir)

	log.Println("All systems ready!")
}

func Logger(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    inner.ServeHTTP(w, r)

    log.Printf(
      "%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
    )
  })
}

func openBrowser(hostname string, port int) error {

  var err error
  address := "http://"+hostname+":"+strconv.Itoa(port)+"/"

  switch runtime.GOOS {
  case "linux":
    err = exec.Command("xdg-open", address).Start()
  case "windows", "darwin":
    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", address).Start()
  default:
    err = fmt.Errorf("unsupported platform")
  }

  return err
}

func main() {
	//setup all of our routes
	r := mux.NewRouter()

	// Declare any api handlers here before the asset dirs
	// or you can use a separate subrouter to put your api on a subdomain
	// we like to do this in production but it's much simpler to not use it
  r.HandleFunc("/task", TodoIndex)

	// We use RequestPathHandler instead of http.FileServer because it allows us to have clean urls
	r.PathPrefix("/app/").HandlerFunc(RequestPathHandler)
	r.PathPrefix("/lib/").HandlerFunc(RequestPathHandler)
	r.PathPrefix("/css/").HandlerFunc(RequestPathHandler)
  r.PathPrefix("/js/").HandlerFunc(RequestPathHandler)
	r.PathPrefix("/img/").HandlerFunc(RequestPathHandler)
	r.PathPrefix("/fonts/").HandlerFunc(RequestPathHandler)

	// General
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath.Join(webDir, "/favicon.ico"))
	})

	// For all other routes we serve up our apps index file
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
		http.ServeFile(w, req, filepath.Join(webDir, "/index.html"))
	})

	err := openBrowser(webHost, webPort)
	if err != nil {
		log.Fatalf("Unable to open browser: %v", err)
	}

	log.Println("Starting web server")
	if err := http.ListenAndServe(":"+strconv.Itoa(webPort), r); err != nil {
		log.Fatalln("Server error:", err)
	}
}

// Serves files relative to the webDir
// Only safe if you use with PathPrefix() or similar functions
func RequestPathHandler(w http.ResponseWriter, req *http.Request) {

	path := filepath.Join(webDir, req.URL.Path)

	log.Println("Serving file:", path)

	//do not show directories
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Println("[FileHandler] error path does not exist:", err)
		http.NotFound(w, req)
		return
	} else if err != nil {
		http.NotFound(w, req)
		log.Println("[FileHandler] error checking if file is dir:", err)
		http.NotFound(w, req)
		return
	}
	if fi.IsDir() {
		http.NotFound(w, req)
		return
	}

	http.ServeFile(w, req, path)
}

//func openBrowser(hostname string, port int) error {
//
//  var err error
//  address := "http://"+hostname+":"+strconv.Itoa(port)+"/"
//
//  switch runtime.GOOS {
//  case "linux":
//    err = exec.Command("xdg-open", address).Start()
//  case "windows", "darwin":
//    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", address).Start()
//  default:
//    err = fmt.Errorf("unsupported platform")
//  }
//
//  return err
//}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Todo Index!")
}
