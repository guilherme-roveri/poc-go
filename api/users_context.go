package api

import "context"

type contextKey struct {
	name string
}

var ContextKeyAPIKey = &contextKey{"api-key"}

func APIKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(ContextKeyAPIKey).(string)
	return key, ok
}
