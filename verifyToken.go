// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func verifyToken(iden *Identify) func(*gin.Context) {
	FakeURLs := iden.FakeURLs
	FakeTokens := iden.FakeTokens
	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path

		token := iden.GetToken(c)
		if token == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"stack":   "invalid token",
			})
			return
		}

		fakeURL := funk.Find(FakeURLs, func(regex string) bool {
			r, _ := regexp.Compile(regex)
			return r.MatchString(reqPath)
		})
		fakeToken := funk.Find(FakeTokens, func(fake string) bool {
			return token.AccessToken == fake
		})

		if fakeURL != nil {
			c.Next()
		} else if fakeToken != nil {
			c.Next()
		} else {
			token, err := iden.Model.Find(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
				c.Abort()
				return
			}
			if token.ISValid() {
				iden.setToken(c, token)
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
				c.Abort()
			}
		}
	}
}
