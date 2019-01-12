package middleware

import (
	"github.com/bomkvilt/tech-db-ap/utiles/walhalla"
	"net/http"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
	data   []byte
}

func (w *statusResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

func (w *statusResponseWriter) Write(data []byte) (int, error) {
	w.data = append(w.data, data...)
	return w.ResponseWriter.Write(data)
}

func Logger(next walhalla.GlobalMiddlewareFunction, ctx *walhalla.Context) walhalla.GlobalMiddlewareFunction {
	return func(rw http.ResponseWriter, r *http.Request) {
		srw := statusResponseWriter{
			ResponseWriter: rw,
		}

		// defer func(start time.Time) {
		// 	duration := time.Since(start).Nanoseconds() / int64(time.Microsecond)
		// 	ctx.Log.WithFields(walhalla.Fields{
		// 		"url":      r.URL.Path,
		// 		"type":     "handle",
		// 		"duration": duration,
		// 		"code":     srw.status,
		// 	}).Info(string(srw.data))
		// }(time.Now())

		next(&srw, r)
	}
}
