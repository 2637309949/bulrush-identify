/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush identify plugin]
 */

package identify

import (
	"encoding/json"
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
	accessToken, _ := token["accessToken"]
	refreshToken, _ := token["refreshToken"]
	value, _ := json.Marshal(token)
	group.Redis.Client.Set("TOKEN:"+accessToken.(string), value, 2*24*time.Hour)
	group.Redis.Client.Set("TOKEN:"+refreshToken.(string), value, 5*24*time.Hour)
}

// Revoke revoke a token
func (group *RedisTokensGroup) Revoke(accessToken string) bool {
	var imapGet map[string]interface{}
	if value, err := group.Redis.Client.Get("TOKEN:" + accessToken).Result(); err == nil {
		json.Unmarshal([]byte(value), &imapGet)
	}

	if status, err := group.Redis.Client.Del("TOKEN:" + accessToken).Result(); err != nil || status != 1 {
		return false
	}

	refreshToken, _ := imapGet["refreshToken"].(string)
	if status, err := group.Redis.Client.Del("TOKEN:" + refreshToken).Result(); err != nil || status != 1 {
		return false
	}
	return true
}

// Find find a token
func (group *RedisTokensGroup) Find(accessToken string, refreshToken string) map[string]interface{} {
	var imapGet map[string]interface{}
	if accessToken != "" {
		if value, err := group.Redis.Client.Get("TOKEN:" + accessToken).Result(); err == nil {
			json.Unmarshal([]byte(value), &imapGet)
			accessTokenRaw, _ := imapGet["accessToken"].(string)
			if accessTokenRaw == accessToken {
				return imapGet
			}
		}
	} else if refreshToken != "" {
		if value, err := group.Redis.Client.Get("TOKEN:" + refreshToken).Result(); err == nil {
			json.Unmarshal([]byte(value), &imapGet)
			refreshTokenRaw, _ := imapGet["refreshToken"].(string)
			if refreshTokenRaw == refreshToken {
				return imapGet
			}
		}
	}
	return nil
}
