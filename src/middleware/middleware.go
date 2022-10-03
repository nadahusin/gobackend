package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
type Adapter func(http.Handler) http.Handler

func Handle(result http.HandlerFunc, middle ...Middleware) http.HandlerFunc {
	for i := len(middle); i > 0; i-- {
		result = middle[i-1](result)
	}
	return result
}
func Adapts(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}
