package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		//resp.Header().Add("header-x", "XXXXXX")
		fmt.Printf("URL %v called at the clock %v\n", req.URL.Path, time.Now())
		next.ServeHTTP(resp, req)
		fmt.Printf("URL %v finished calling at %v\n", req.URL.Path, time.Now())
	})
}