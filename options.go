// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import "github.com/gin-gonic/gin"

// Option defined implement of option
type (
	Option func(*Identify) interface{}
)

// option defined implement of option
func (o Option) apply(r *Identify) *Identify { return o(r).(*Identify) }

// option defined implement of option
func (o Option) check(r *Identify) interface{} { return o(r) }

// AuthOption defined auth
func AuthOption(auth func(ctx *gin.Context) (interface{}, error)) Option {
	return Option(func(r *Identify) interface{} {
		r.auth = auth
		return r
	})
}

// ModelOption defined model
func ModelOption(model Model) Option {
	return Option(func(r *Identify) interface{} {
		r.model = model
		return r
	})
}

// IdenOption defined iden
func IdenOption(iden func(c *gin.Context)) Option {
	return Option(func(r *Identify) interface{} {
		r.iden = iden
		return r
	})
}

// FakeTokensOption defined model
func FakeTokensOption(f []string) Option {
	return Option(func(r *Identify) interface{} {
		r.fakeTokens = f
		return r
	})
}

// FakeURLsOption defined model
func FakeURLsOption(f []string) Option {
	return Option(func(r *Identify) interface{} {
		r.fakeURLs = f
		return r
	})
}
