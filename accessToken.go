// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"github.com/gin-gonic/gin"
)

const tokenKey = "accessToken"

func token(c *gin.Context) string {
	var accessToken string
	if queryToken := c.Query(tokenKey); queryToken != "" {
		accessToken = queryToken
	} else if formToken := c.PostForm(tokenKey); formToken != "" {
		accessToken = formToken
	} else if headerToken := c.Request.Header.Get(tokenKey); headerToken != "" {
		accessToken = headerToken
	} else if cookieToken, err := c.Cookie(tokenKey); err == nil {
		accessToken = cookieToken
	}
	return accessToken
}

func accessToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		iden.setToken(c, &Token{AccessToken: token(c)})
		c.Next()
	}
}
