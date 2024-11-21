package keyauth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/keyauth"
)

func MyKeyauth() app.HandlerFunc {
	return keyauth.New(
		keyauth.WithValidator(func(ctx context.Context, requestContext *app.RequestContext, s string) (bool, error) {
			if s == "test_admin" {
				return true, nil
			}
			return false, nil
		}),
	)
}
