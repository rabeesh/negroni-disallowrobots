package disallowrobots

import (
    "net/http"
    "net/http/httptest"
    "fmt"
    "testing"
)


func Test_DisallowRobots_RobotsUrl(t *testing.T) {

    isProduction := false
    dr := New(isProduction)

    rw := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "http://localhost/robots.txt", nil)

    if err != nil {
        t.Fatal(err)
    }

    dr.ServeHTTP(rw, req, func (rw http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(rw, "")
    })

    h := rw.Header()

    if rw.Code != http.StatusOK {
        t.Error("Expected a valid response code")
    }

    if "*\nDisallow: /" != h.Get("User-Agent") {
        t.Error("Expected a valid user agent header")
    }
}


func Test_DisallowRobots_OtherUrls(t *testing.T) {

    isProduction := false
    dr := New(isProduction)

    rw := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "http://localhost/blaa", nil)

    if err != nil {
        t.Fatal(err)
    }

    dr.ServeHTTP(rw, req, func (rw http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(rw, "")
    })

    h := rw.Header()

    if rw.Code != http.StatusOK {
        t.Error("Expected a valid response code")
    }

    if "noindex, nofollow, noarchive" != h.Get("X-Robots-Tag") {
        t.Error("Expected a valid response code")
    }
}

