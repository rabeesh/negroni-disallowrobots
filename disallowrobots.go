package disallowrobots

import (
    "log"
    "os"
    "net/http"
    "github.com/codegangsta/negroni"
)

type DisallowRobots struct {
    Env bool
    Logger  *log.Logger
}

func New(env bool) *DisallowRobots {
    return &DisallowRobots{
        Env: env,
        Logger: log.New(os.Stdout, "[negroni disallow robots in production is enabled] ", 0),
    }
}

func (dr *DisallowRobots) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

    if dr.Env {
        next(rw, req)
        return
    }

    file := req.URL.Path
    wh := rw.Header()
    if file == "/robots.txt" {
        wh.Set("User-Agent", "*\nDisallow: /")
        wh.Set("Content-Type", "text/plain")
        rw.WriteHeader(http.StatusOK)
        return
    }

    newRw := negroni.NewResponseWriter(rw)
    newRw.Before(func (rw negroni.ResponseWriter){
        rw.Header().Set("X-Robots-Tag", "noindex, nofollow, noarchive")
    })

    next(newRw, req)
}
