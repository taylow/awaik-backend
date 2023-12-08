package interceptor

import (
	"context"

	"connectrpc.com/connect"
)

var _ connect.Interceptor = (*validationInterceptor)(nil)

// validator is an interface for validating request parameters
type validator interface {
	ValidateAll() error
}

// NewValidateInterceptor creates an interceptor that validates the request parameters
func NewValidateInterceptor() *validationInterceptor {
	return &validationInterceptor{}
}

type validationInterceptor struct{}

// WrapUnary implements validation for unary connect handlers
func (v *validationInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		if v, ok := req.Any().(validator); ok {
			err := v.ValidateAll()
			if err != nil {
				return nil, connect.NewError(
					connect.CodeInvalidArgument,
					err,
				)
			}
		}

		return next(ctx, req)
	}
}

// WrapStreamingClient validation for streaming connect clients
func (v *validationInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		// TODO: implement validation for streaming clients
		return next(ctx, spec)
	}
}

// WrapStreamingHandler validation for streaming connect handlers
func (v *validationInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		// TODO: implement validation for streaming handlers
		return next(ctx, conn)
	}
}
