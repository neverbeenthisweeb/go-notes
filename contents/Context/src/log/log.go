package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

// NOTE: To avoid key naming collision we use an
// unexported type.
type key string

const requestIdKey = key("request-id")

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIdKey).(int64)
	if !ok {
		log.Println("invalid data type")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		// NOTE: Giving value to context should been through a solid
		// thinking process. This is not that recommended.
		// If you still want to do this, consider giving a request-specific value
		// (i.e. request id).
		ctx = context.WithValue(ctx, requestIdKey, id)
		f(w, r.WithContext(ctx))
	}
}
