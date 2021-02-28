// This file is part of atreugo/session which is released under Apache 2.0 license.
// Go to https://github.com/atreugo/session/blob/master/LICENSE for full license details.

package redis

import (
	"github.com/fasthttp/session/v2/providers/redis"
)

// Config provider settings.
type Config redis.Config

// New returns a new configured redis provider.
func New(cfg Config) (*redis.Provider, error) {
	provider, err := redis.New(redis.Config(cfg))

	return provider, err
}
