// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"github.com/gin-gonic/gin"
)

const key = "accessToken"

func accessToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		var accessToken string
		if queryToken := c.Query(key); queryToken != "" {
			accessToken = queryToken
		} else if formToken := c.PostForm(key); formToken != "" {
			accessToken = formToken
		} else if headerToken := c.Request.Header.Get(key); headerToken != "" {
			accessToken = headerToken
		} else if cookieToken, err := c.Cookie(key); err == nil {
			accessToken = cookieToken
		}
		setAccessToken(c, accessToken)
		c.Next()
	}
}
