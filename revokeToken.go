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
		token := iden.GetToken(c)
		if token != nil {
			if err := iden.Model.Revoke(token); err != nil {
				rushLogger.Warn("revoke token failed,token may not exist")
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	}
}
