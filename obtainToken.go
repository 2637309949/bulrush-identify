// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func obtainToken(iden *Identify) func(c *gin.Context) {
	return func(c *gin.Context) {
		authData, err := iden.Auth(c)
		if authData != nil {
			data := iden.ObtainToken(authData)
			c.JSON(http.StatusOK, data)
			c.Abort()
		} else {
			rushLogger.Warn("auth error,%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}
	}
}
