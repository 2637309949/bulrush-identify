// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

// IdenKey defined ctx key
const IdenKey = "identify"

type (
	// Identify authentication interface
	Identify struct {
		auth       func(c *gin.Context) (interface{}, error)
		iden       func(c *gin.Context)
		ExpiresIn  int64
		Routes     RoutesGroup
		model      Model
		fakeURLs   []string
		fakeTokens []string
	}
	// RoutesGroup iden routes
	RoutesGroup struct {
		ObtainTokenRoute  string
		RevokeTokenRoute  string
		RefleshTokenRoute string
		IdenTokenRoute    string
	}
	// Token defined Token info
	Token struct {
		AccessToken  string
		RefreshToken string
		ExpiresIn    int64
		CreatedAt    int64
		UpdatedAt    int64
		Extra        interface{}
	}
)

// Unmarshal defined from json
func (t *Token) Unmarshal(data string) error {
	return json.Unmarshal([]byte(data), t)
}

// Marshal defined to json
func (t *Token) Marshal() (string, error) {
	dataByte, err := json.Marshal(t)
	return string(dataByte), err
}

// MarshalExtra defined Extra to spec type
func (t *Token) MarshalExtra(target interface{}) error {
	dataByte, err := json.Marshal(t.Extra)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, target)
}

// ExtraValue defined Extra to spec type
func (t *Token) ExtraValue(key string) interface{} {
	return t.Extra.(map[string]interface{})[key]
}

// ISValid defined token ISValid
func (t *Token) ISValid() bool {
	return (t.ExpiresIn + t.CreatedAt) > time.Now().Unix()
}

// New defined return a new struct
func New() *Identify {
	iden := &Identify{
		ExpiresIn: 86400,
		Routes:    defaultRoutesGroup,
		auth:      defaultAuth,
	}
	iden.iden = defaultIden(iden)
	return iden
}

// AddOptions defined add option
func (iden *Identify) AddOptions(opts ...Option) *Identify {
	for _, v := range opts {
		v.apply(iden)
	}
	return iden
}

// Init Identify
func (iden *Identify) Init(init func(*Identify)) *Identify {
	init(iden)
	return iden
}

// ObtainToken accessToken
func (iden *Identify) ObtainToken(extra interface{}) (*Token, error) {
	return iden.model.Save(&Token{
		AccessToken:  RandString(32),
		RefreshToken: RandString(32),
		ExpiresIn:    iden.ExpiresIn,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		Extra:        extra,
	})
}

// GetToken get from ctx
func (iden *Identify) GetToken(ctx *gin.Context) *Token {
	if token, exists := ctx.Get(IdenKey); exists {
		return token.(*Token)
	}
	return nil
}

// SetToken set to ctx
func (iden *Identify) SetToken(ctx *gin.Context, token *Token) {
	ctx.Set(IdenKey, token)
}

// Plugin for bulrush
func (iden *Identify) Plugin(router *gin.RouterGroup, httpProxy *gin.Engine) {
	router.Use(accessToken(iden))
	router.POST(iden.Routes.ObtainTokenRoute, obtainToken(iden))
	router.Use(verifyToken(iden))
	router.POST(iden.Routes.RefleshTokenRoute, refleshToken(iden))
	router.POST(iden.Routes.RevokeTokenRoute, revokeToken(iden))
	router.POST(iden.Routes.IdenTokenRoute, iden.iden)
}
