// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import "github.com/gin-gonic/gin"

func accessToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		var accessToken string
		queryToken := c.Query("accessToken")
		formToken := c.PostForm("accessToken")
		headerToken := c.Request.Header.Get("Authorization")
		if queryToken != "" {
			accessToken = queryToken
		} else if formToken != "" {
			accessToken = formToken
		} else if headerToken != "" {
			accessToken = headerToken
		}
		setAccessToken(c, accessToken)
		c.Next()
	}
}
