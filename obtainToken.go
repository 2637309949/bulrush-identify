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
		result, err := iden.Auth(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		data, err := iden.ObtainToken(result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.SetCookie(tokenKey, data.AccessToken, 60*60*24, "/", "", false, true)
		c.JSON(http.StatusOK, map[string]interface{}{
			"AccessToken":  data.AccessToken,
			"RefreshToken": data.RefreshToken,
			"ExpiresIn":    data.ExpiresIn,
			"CreatedAt":    data.CreatedAt,
			"UpdatedAt":    data.UpdatedAt,
		})
	}
}
