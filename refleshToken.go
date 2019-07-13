// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func refleshToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		token := iden.GetToken(c)
		if token != nil {
			token, err := iden.Model.Find(token)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error()})
				return
			}
			if err := iden.Model.Revoke(token); err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error()})
				return
			}
			token.CreatedAt = time.Now().Unix()
			token.UpdatedAt = time.Now().Unix()
			token.AccessToken = RandString(32)
			token, err = iden.Model.Save(token)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, token)
		}
		c.Abort()
	}
}
