// This file is part of atreugo/session which is released under Apache 2.0 license.
// Go to https://github.com/atreugo/session/blob/master/LICENSE for full license details.

package postgre

import (
	"github.com/fasthttp/session/v2/providers/postgre"
)

// Config provider settings.
type Config postgre.Config

// New returns a new configured postgres provider.
func New(cfg Config) (*postgre.Provider, error) {
	provider, err := postgre.New(postgre.Config(cfg))

	return provider, err
}

// NewDefaultConfig returns a default configuration.
func NewDefaultConfig() Config {
	return Config(postgre.NewDefaultConfig())
}

// NewConfigWith returns a new configuration with especific paremters.
func NewConfigWith(host string, port int64, username string, password string, dbName string, tableName string) Config {
	cfg := postgre.NewConfigWith(host, port, username, password, dbName, tableName)

	return Config(cfg)
}
