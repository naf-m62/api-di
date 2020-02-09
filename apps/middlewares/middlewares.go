package middlewares

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// PanicRecoveryMiddleware handles the panic in the handlers.
func PanicRecoveryMiddleware(h http.HandlerFunc, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				logger.Error(fmt.Sprint(rec))

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal error"))
			}
		}()

		h(w, r)
	}
}
