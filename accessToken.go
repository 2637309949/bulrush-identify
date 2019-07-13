// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	utils "github.com/2637309949/bulrush-utils"
	"github.com/gin-gonic/gin"
)

const tokenKey = "accessToken"

func accessToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		token := utils.Until(
			c.Query(tokenKey),
			c.PostForm(tokenKey),
			c.Request.Header.Get(tokenKey),
			func() interface{} {
				value, _ := c.Cookie(tokenKey)
				return value
			},
		).(string)
		iden.setToken(c, &Token{AccessToken: token})
		c.Next()
	}
}
