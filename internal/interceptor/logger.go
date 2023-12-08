package interceptor

import (
	"context"
	"log"
	"time"

	"connectrpc.com/connect"
)

// LogRequest logs the request parameters.
// It logs all kinds of requests.
func NewLogInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			begin := time.Now()
			defer func() {
				log.Printf("%s took %v\n", req.Spec().Procedure, time.Since(begin))
			}()

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
