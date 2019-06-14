// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

// GetAccessToken get from ctx
func GetAccessToken(ctx *gin.Context) string {
	accessToken, _ := ctx.Get("accessToken")
	return accessToken.(string)
}

func setAccessToken(ctx *gin.Context, accessToken string) {
	ctx.Set("accessToken", accessToken)
}

// GetAccessData get from ctx
func GetAccessData(ctx *gin.Context) map[string]interface{} {
	accessData, _ := ctx.Get("accessData")
	if accessData != nil {
		return accessData.(map[string]interface{})
	}
	return nil
}

func setAccessData(ctx *gin.Context, data interface{}) {
	ctx.Set("accessData", data)
}

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

// Find target from array
func Find(arrs []interface{}, matcher func(interface{}) bool) interface{} {
	var target interface{}
	for _, item := range arrs {
		match := matcher(item)
		if match {
			target = item
			break
		}
	}
	return target
}
