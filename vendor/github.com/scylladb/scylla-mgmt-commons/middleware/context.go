// Copyright (C) 2017 ScyllaDB

package middleware

// ctxt is a context key type.
type ctxt byte

// ctxt enumeration.
const (
	ctxInteractive ctxt = iota
	ctxHost
	ctxNoRetry
	ctxCustomTimeout
)
