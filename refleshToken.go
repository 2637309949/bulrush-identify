// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func refleshToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		refreshToken := GetAccessToken(c)
		if refreshToken != "" {
			originToken := iden.RefleshToken(refreshToken)
			// reflesh and save
			if originToken != nil {
				c.JSON(http.StatusOK, originToken)
			} else {
				rushLogger.Warn("reflesh token failed,token may not exist")
				c.JSON(http.StatusBadRequest, gin.H{"message": "reflesh token failed, token may not exist"})
			}
		}
		c.Abort()
	}
}
