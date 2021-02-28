// This file is part of atreugo/session which is released under Apache 2.0 license.
// Go to https://github.com/atreugo/session/blob/master/LICENSE for full license details.

package memcache

import (
	"github.com/fasthttp/session/v2/providers/memcache"
)

// Config provider settings.
type Config memcache.Config

// New returns a new memcache provider configured.
func New(cfg Config) (*memcache.Provider, error) {
	provider, err := memcache.New(memcache.Config(cfg))

	return provider, err
}
