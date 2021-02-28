// This file is part of atreugo/session which is released under Apache 2.0 license.
// Go to https://github.com/atreugo/session/blob/master/LICENSE for full license details.

package session

import (
	"github.com/emirmuminoglu/emir"
	"github.com/fasthttp/session/v2"
)

// New returns a configured manager.
func New(cfg Config) *Session {
	s := session.New(session.Config(cfg))

	return &Session{session: s}
}

// NewDefaultConfig returns a new default configuration.
func NewDefaultConfig() Config {
	return Config(session.NewDefaultConfig())
}

// SetProvider sets the session provider used by the sessions manager.
func (s *Session) SetProvider(provider Provider) error {
	return s.session.SetProvider(provider)
}

// Get returns the user session
// if it does not exist, it will be generated.
func (s *Session) Get(ctx emir.Context) (*session.Store, error) {
	return s.session.Get(ctx.FasthttpCtx())
}

// Save saves the user session
//
// Warning: Don't use the store after exec this function, because, you will lose the after data
// For avoid it, defer this function in your request handler.
func (s *Session) Save(ctx emir.Context, store *session.Store) error {
	return s.session.Save(ctx.FasthttpCtx(), store)
}

// Regenerate generates a new session id to the current user.
func (s *Session) Regenerate(ctx emir.Context) error {
	return s.session.Regenerate(ctx.FasthttpCtx())
}

// Destroy destroys the session of the current user.
func (s *Session) Destroy(ctx emir.Context) error {
	return s.session.Destroy(ctx.FasthttpCtx())
}
