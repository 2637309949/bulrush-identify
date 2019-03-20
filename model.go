/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush identify plugin]
 */

package identify

import (
	"time"

	"github.com/2637309949/bulrush-addition/redis"
)

// RedisTokensGroup adapter for redis
type RedisTokensGroup struct {
	TokensGroup
	Redis *redis.Redis
}

// Save save a token
func (group *RedisTokensGroup) Save(token map[string]interface{}) {
	accessToken, _ := token["accessToken"].(string)
	refreshToken, _ := token["refreshToken"].(string)
	group.Redis.Hooks.SetJSON("TOKEN:"+accessToken, token, 36*time.Hour)
	group.Redis.Hooks.SetJSON("TOKEN:"+refreshToken, token, 36*time.Hour)
}

// Revoke revoke a token
func (group *RedisTokensGroup) Revoke(accessToken string) bool {
	imapGet := group.Redis.Hooks.GetJSON("TOKEN:" + accessToken)
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
func (group *RedisTokensGroup) Find(accessToken string, refreshToken string) map[string]interface{} {
	if accessToken != "" {
		imapGet := group.Redis.Hooks.GetJSON("TOKEN:" + accessToken)
		accessTokenRaw, _ := imapGet["accessToken"].(string)
		if accessTokenRaw == accessToken {
			return imapGet
		}
	} else if refreshToken != "" {
		imapGet := group.Redis.Hooks.GetJSON("TOKEN:" + refreshToken)
		refreshTokenRaw, _ := imapGet["refreshToken"].(string)
		if refreshTokenRaw == refreshToken {
			return imapGet
		}
	}
	return nil
}
