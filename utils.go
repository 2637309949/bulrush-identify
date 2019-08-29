// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RandString gen random string
func RandString(n int) string {
	const seeds = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = seeds[rand.Intn(len(seeds))]
	}
	return string(bytes)
}

// Some get or a default value
func Some(target interface{}, initValue interface{}) interface{} {
	if target != nil && target != "" && target != 0 {
		return target
	}
	return initValue
}

func defaultIden(iden *Identify) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, iden.GetToken(c).Extra)
	}
}

func defaultAuth(c *gin.Context) (interface{}, error) {
	return nil, errors.New("user authentication failed")
}

var defaultRoutesGroup = RoutesGroup{
	ObtainTokenRoute:  "/obtainToken",
	RevokeTokenRoute:  "/revokeToken",
	RefleshTokenRoute: "/refleshToken",
	IdenTokenRoute:    "/idenToken",
}
