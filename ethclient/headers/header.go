// Package headers is used to provide ethclient with the ability to inject http headers
package headers

import (
	"context"
	"net/http"
)

// defines a set of wellknown header keys
const (
	KeyConsistentHash = "x-blockchain-gateway-consistent-hash"
)

type headerKey struct{}

// WithHTTPHeader is used to attach header information to the context
func WithHTTPHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, headerKey{}, header)
}

// FromContext is used to extract header information from context
func FromContext(ctx context.Context) (http.Header, bool) {
	value := ctx.Value(headerKey{})
	if value == nil {
		return nil, false
	}
	return value.(http.Header), true
}

// Inject is used to extract headers information from the context and inject it into the provided headers
func Inject(ctx context.Context, headers http.Header) {
	if got, ok := FromContext(ctx); ok {
		override(headers, got)
	}
}

func override(low http.Header, high http.Header) {
	for key, values := range high {
		low.Del(key)
		for _, val := range values {
			low.Add(key, val)
		}
	}
}
