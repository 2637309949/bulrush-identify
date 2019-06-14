// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"time"

	"github.com/2637309949/bulrush-addition/redis"
)

// RedisModel adapter for redis
type RedisModel struct {
	Model
	Redis *redis.Redis
}

// Save save a token
func (group *RedisModel) Save(token map[string]interface{}) {
	accessToken, _ := token["accessToken"].(string)
	refreshToken, _ := token["refreshToken"].(string)
	group.Redis.API.SetJSON("TOKEN:"+accessToken, token, 36*time.Hour)
	group.Redis.API.SetJSON("TOKEN:"+refreshToken, token, 36*time.Hour)
}

// Revoke revoke a token
func (group *RedisModel) Revoke(accessToken string) bool {
	imapGet, _ := group.Redis.API.GetJSON("TOKEN:" + accessToken)
	refreshToken, _ := imapGet["refreshToken"].(string)
	if status, err := group.Redis.Client.Del("TOKEN:" + accessToken).Result(); err != nil || status != 1 {
		return false
	}
	if status, err := group.Redis.Client.Del("TOKEN:" + refreshToken).Result(); err != nil || status != 1 {
		return false
	}
	return true
}

// Find find a token
func (group *RedisModel) Find(accessToken string, refreshToken string) map[string]interface{} {
	if accessToken != "" {
		imapGet, _ := group.Redis.API.GetJSON("TOKEN:" + accessToken)
		accessTokenRaw, _ := imapGet["accessToken"].(string)
		if accessTokenRaw == accessToken {
			return imapGet
		}
	} else if refreshToken != "" {
		imapGet, _ := group.Redis.API.GetJSON("TOKEN:" + refreshToken)
		refreshTokenRaw, _ := imapGet["refreshToken"].(string)
		if refreshTokenRaw == refreshToken {
			return imapGet
		}
	}
	return nil
}
