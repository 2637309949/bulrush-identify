// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"

	utils "github.com/2637309949/bulrush-utils"
	"github.com/gin-gonic/gin"
)

func obtainToken(iden *Identify) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := utils.Chain(
			func(ret interface{}) (interface{}, error) {
				return iden.Auth(c)
			},
			func(ret interface{}) (interface{}, error) {
				return iden.ObtainToken(ret)
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		token := data.(*Token)
		c.SetCookie(tokenKey, token.AccessToken, 60*60*24, "/", "", false, true)
		c.JSON(http.StatusOK, map[string]interface{}{
			"AccessToken":  token.AccessToken,
			"RefreshToken": token.RefreshToken,
			"ExpiresIn":    token.ExpiresIn,
			"CreatedAt":    token.CreatedAt,
			"UpdatedAt":    token.UpdatedAt,
		})
	}
}
