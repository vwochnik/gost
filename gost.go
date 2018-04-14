package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/vwochnik/gost/fileserver"
	"github.com/vwochnik/gost/template"
)

const Version = "0.1.2"

var args Arguments

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parseArguments(&args)

	if len(args.log) > 0 {
		file, err := os.OpenFile(args.log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		exitOnError(err)
		defer file.Close()
		log.SetOutput(file)
	} else if args.quiet {
		file, err := os.Open(os.DevNull)
		exitOnError(err)
		defer file.Close()
		log.SetOutput(file)
		log.SetFlags(0)
	}
}

func main() {
	listen := fmt.Sprintf("%s:%d", args.host, args.port)

	fileserver.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		template.ErrorTemplate(w, "File Not Found", 404)
	}

	fileserver.ErrorHandler = func(w http.ResponseWriter, r *http.Request, msg string, code int) {
		template.ErrorTemplate(w, msg, code)
	}

	fileserver.DirListHandler = func(w http.ResponseWriter, f http.File) {
		template.DirectoryTemplate(w, f)
	}

	http.Handle("/", fileserver.FileServer(http.Dir(args.directory)))
	handler := buildHttpHandler()

	log.Printf("Static file server running at %s. Ctrl+C to quit.\n", listen)
	err := http.ListenAndServe(listen, handler)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildHttpHandler() http.Handler {
	var handler http.Handler

	handler = http.DefaultServeMux

	if args.cors {
		handler = corsHandler(handler)
	}

	if args.noCache {
		handler = cacheHandler(handler)
	}

	if !args.quiet {
		handler = logHandler(handler)
	}

	return handler
}
