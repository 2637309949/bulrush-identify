// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func verifyToken(iden *Identify) func(*gin.Context) {
	FakeURLs := iden.FakeURLs
	FakeTokens := iden.FakeTokens
	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		accessToken := GetAccessToken(c)
		fakeURL := Find(FakeURLs, func(regex interface{}) bool {
			r, _ := regexp.Compile(regex.(string))
			return r.MatchString(reqPath)
		})
		fakeToken := Find(FakeTokens, func(token interface{}) bool {
			return accessToken == token.(string)
		})
		if fakeURL != nil {
			c.Next()
		} else if fakeToken != nil {
			c.Next()
		} else {
			ret := iden.VerifyToken(accessToken)
			if ret {
				rawToken := iden.Model.Find(accessToken, "")
				setAccessData(c, rawToken["extra"])
				c.Next()
			} else {
				rushLogger.Warn("invalid token")
				c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
				c.Abort()
			}
		}
	}
}
