# Disallow Robot (Negroni middleware)

Disallows indexation by search engines (with the robots.txt file) whenever you're not in production.

## Usage

~~~ go
package main

import (
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/rabeesh/negroni-disallowrobots"
    "net/http"
)


func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(rw, "Welcome to the home page!")
    })

    isProduction := false

    n := negroni.New()
    n.Use(disallowrobots.New(isProduction))
    n.UseHandler(mux)
    n.Run(":5000")
}
~~~

`X-Robots-Tag` info is available [here](https://developers.google.com/webmasters/control-crawl-index/docs/robots_meta_tag).


## Authors

- Rab [@rabeesh](http://stackoverflow.com/users/1722625/rab)
- Inspired by ruby [Cylon](https://github.com/dmathieu/cylon).