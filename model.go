// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"encoding/json"
	"errors"
	"time"

	redisext "github.com/2637309949/bulrush-addition/redis"
	goRedis "github.com/go-redis/redis"
)

// RedisModel adapter for redis
type RedisModel struct {
	Model
	Redis *redisext.Redis
}

// Save save a token
func (model *RedisModel) Save(token *Token) (*Token, error) {
	dataByte, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}
	if err = model.Redis.Client.Set("TOKEN:"+token.AccessToken, string(dataByte), 24*time.Hour).Err(); err != nil {
		return nil, err
	}
	if err = model.Redis.Client.Set("TOKEN:"+token.RefreshToken, string(dataByte), 24*time.Hour).Err(); err != nil {
		return nil, err
	}
	return token, nil
}

// Revoke revoke a token
func (model *RedisModel) Revoke(token *Token) error {
	if _, err := model.Redis.Client.Del("TOKEN:" + token.AccessToken).Result(); err != nil {
		return err
	}
	if _, err := model.Redis.Client.Del("TOKEN:" + token.RefreshToken).Result(); err != nil {
		return err
	}
	return nil
}

// Find find a token
func (model *RedisModel) Find(token *Token) (*Token, error) {
	if token.AccessToken != "" {
		data, err := model.Redis.Client.Get("TOKEN:" + token.AccessToken).Result()
		if err == goRedis.Nil {
			return nil, errors.New("invalid token")
		} else if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(data), token); err != nil {
			return nil, err
		}
		return token, nil
	} else if token.RefreshToken != "" {
		data, err := model.Redis.Client.Get("TOKEN:" + token.RefreshToken).Result()
		if err == goRedis.Nil {
			return nil, errors.New("invalid token")
		} else if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(data), token); err != nil {
			return nil, err
		}
		return token, nil
	}
	return nil, errors.New("no token found")
}
