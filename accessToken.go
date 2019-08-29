// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"github.com/2637309949/bulrush-utils/funcs"
	"github.com/gin-gonic/gin"
)

const tokenKey = "accessToken"

func accessToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		token := funcs.Until(
			c.Query(tokenKey),
			c.PostForm(tokenKey),
			c.Request.Header.Get(tokenKey),
			func() interface{} {
				value, _ := c.Cookie(tokenKey)
				return value
			},
		).(string)
		iden.SetToken(c, &Token{AccessToken: token})
		c.Next()
	}
}
