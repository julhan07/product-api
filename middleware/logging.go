package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lg := logrus.New()
		lg.SetFormatter(&logrus.JSONFormatter{})
		lg.Printf("host :%s, path :%s, query :%v, body :%v", r.Host, r.URL.Path, r.URL.Query(), r.Body)
		next.ServeHTTP(w, r)
	})
}
