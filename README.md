negroni-pprof
=============

[![GoDoc](https://godoc.org/github.com/mistifyio/negroni-pprof?status.svg)](https://godoc.org/github.com/mistifyio/negroni-pprof)

pprof middleware for Negroni

# pprof

[pprof](http://golang.org/pkg/net/http/pprof/) middleware for [Negroni](https://github.com/codegangsta/negroni).

A simple middleware for handling pprof.  This handles "/debug/pprof" paths

Note: this implementation is sub-optimal, but works.

## Usage

~~~ go
import (
    "github.com/codegangsta/negroni"
    "github.com/mistifyio/negroni-pprof"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
      fmt.Fprintf(w, "Welcome to the home page!")
    }
    n := negroni.Classic()
    n.Use(pprof.Pprof())
    n.UseHandler(mux)
    n.Run(":3000")
}
