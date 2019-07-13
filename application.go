// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ctxIndenKey = "identify"
)

type (
	// Identify authentication interface
	Identify struct {
		Auth       func(c *gin.Context) (interface{}, error)
		Iden       func(c *gin.Context)
		ExpiresIn  int64
		Routes     RoutesGroup
		Model      Model
		FakeURLs   []string
		FakeTokens []string
	}
	// Model token store
	Model interface {
		Save(*Token) (*Token, error)
		Find(*Token) (*Token, error)
		Revoke(*Token) error
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

// New defined return a new struct
func New() *Identify {
	iden := &Identify{
		ExpiresIn: 86400,
		Routes:    RoutesGroup{},
		Auth: func(c *gin.Context) (interface{}, error) {
			return nil, errors.New("user authentication failed")
		},
	}
	iden.Iden = func(c *gin.Context) {
		c.JSON(http.StatusOK, iden.GetToken(c).Extra)
	}
	return iden
}

// Init Identify
func (iden *Identify) routesGroup() *Identify {
	if iden.Routes.ObtainTokenRoute == "" {
		iden.Routes.ObtainTokenRoute = "/obtainToken"
	}
	if iden.Routes.RevokeTokenRoute == "" {
		iden.Routes.RevokeTokenRoute = "/revokeToken"
	}
	if iden.Routes.RefleshTokenRoute == "" {
		iden.Routes.RefleshTokenRoute = "/refleshToken"
	}
	if iden.Routes.IdenTokenRoute == "" {
		iden.Routes.IdenTokenRoute = "/idenToken"
	}
	return iden
}

// Init Identify
func (iden *Identify) Init(init func(*Identify)) *Identify {
	init(iden)
	iden.routesGroup()
	return iden
}

// ObtainToken accessToken
func (iden *Identify) ObtainToken(extra interface{}) (*Token, error) {
	return iden.Model.Save(&Token{
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
	if token, exists := ctx.Get(ctxIndenKey); exists {
		return token.(*Token)
	}
	return nil
}

// setToken set to ctx
func (iden *Identify) setToken(ctx *gin.Context, token *Token) {
	ctx.Set(ctxIndenKey, token)
}

// Plugin for bulrush
func (iden *Identify) Plugin(router *gin.RouterGroup, httpProxy *gin.Engine) {
	router.Use(accessToken(iden))
	router.POST(iden.Routes.ObtainTokenRoute, obtainToken(iden))
	router.POST(iden.Routes.RevokeTokenRoute, revokeToken(iden))
	router.POST(iden.Routes.RefleshTokenRoute, refleshToken(iden))
	router.Use(verifyToken(iden))
	router.GET(iden.Routes.IdenTokenRoute, iden.Iden)
}
