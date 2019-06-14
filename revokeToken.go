// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func revokeToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		accessToken := GetAccessToken(c)
		if accessToken != "" {
			result := iden.RevokeToken(accessToken)
			if result {
				c.JSON(http.StatusOK, gin.H{
					"success": true,
				})
			} else {
				rushLogger.Warn("revoke token failed,token may not exist")
				c.JSON(http.StatusBadRequest, gin.H{"message": "revoke token failed, token may not exist"})
			}
		}
		c.Abort()
	}
}
