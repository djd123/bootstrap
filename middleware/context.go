package middleware

import (
	"context"
	"net/http"

	"github.com/djd123/bootstrap"
)

func AddContext(fields map[string]interface{}) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			augmentedContext := r.Context()

			for k, v := range fields {
				augmentedContext = context.WithValue(augmentedContext, bootstrap.ContextKey(k), v)
			}

			next.ServeHTTP(w, r.WithContext(augmentedContext))
		}
		return http.HandlerFunc(fn)
	}
}
