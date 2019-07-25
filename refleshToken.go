// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"
	"time"

	utils "github.com/2637309949/bulrush-utils"
	"github.com/gin-gonic/gin"
)

func refleshToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		ret, err := utils.Chain(
			func(ret interface{}) (interface{}, error) {
				return iden.GetToken(c), nil
			},
			func(ret interface{}) (interface{}, error) {
				token := ret.(*Token)
				return token, iden.Model.Revoke(token)
			},
			func(ret interface{}) (interface{}, error) {
				token := ret.(*Token)
				token.CreatedAt = time.Now().Unix()
				token.UpdatedAt = time.Now().Unix()
				token.AccessToken = RandString(32)
				return iden.Model.Save(token)
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"stack":   err.Error(),
			})
			return
		}
		token := ret.(*Token)
		c.JSON(http.StatusInternalServerError, token)
	}
}
