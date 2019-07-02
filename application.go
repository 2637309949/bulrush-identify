// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package identify

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
)

// RoutesGroup iden routes
type RoutesGroup struct {
	ObtainTokenRoute  string
	RevokeTokenRoute  string
	RefleshTokenRoute string
}

// Model token store
type Model interface {
	Save(map[string]interface{})
	Find(string, string) map[string]interface{}
	Revoke(string) bool
}

// Identify authentication interface
type Identify struct {
	Auth       func(c *gin.Context) (interface{}, error)
	ExpiresIn  int
	Routes     RoutesGroup
	Model      Model
	FakeURLs   []interface{}
	FakeTokens []interface{}
}

// ObtainToken accessToken
func (iden *Identify) ObtainToken(authData interface{}) interface{} {
	if authData != nil {
		data := map[string]interface{}{
			"accessToken":  RandString(32),
			"refreshToken": RandString(32),
			"expiresIn":    Some(iden.ExpiresIn, 86400),
			"created":      now.New(time.Now()).Unix(),
			"updated":      now.New(time.Now()).Unix(),
			"extra":        authData,
		}
		iden.Model.Save(data)
		return data
	}
	return nil
}

// RevokeToken accessToken
func (iden *Identify) RevokeToken(token string) bool {
	return iden.Model.Revoke(token)
}

// RefleshToken accessToken
func (iden *Identify) RefleshToken(refreshToken string) interface{} {
	originToken := iden.Model.Find("", refreshToken)
	if originToken != nil {
		accessToken, _ := originToken["accessToken"]
		iden.Model.Revoke(accessToken.(string))
		originToken["created"] = now.New(time.Now()).Unix()
		originToken["updated"] = now.New(time.Now()).Unix()
		originToken["accessToken"] = RandString(32)
		iden.Model.Save(originToken)
		return originToken
	}
	return nil
}

// VerifyToken accessToken
func (iden *Identify) VerifyToken(token string) bool {
	verifyToken := iden.Model.Find(token, "")
	now := time.Now().Unix()
	if verifyToken == nil {
		return false
	}
	expiresIn, _ := verifyToken["expiresIn"]
	created, _ := verifyToken["created"]
	if (expiresIn.(float64) + created.(float64)) < float64(now) {
		return false
	}
	return true
}

// Plugin for bulrush
func (iden *Identify) Plugin(router *gin.RouterGroup) {
	obtainTokenRoute := Some(iden.Routes.ObtainTokenRoute, "/obtainToken").(string)
	revokeTokenRoute := Some(iden.Routes.RevokeTokenRoute, "/revokeToken").(string)
	refleshTokenRoute := Some(iden.Routes.RefleshTokenRoute, "/refleshToken").(string)
	router.Use(accessToken(iden))
	router.POST(obtainTokenRoute, obtainToken(iden))
	router.POST(revokeTokenRoute, revokeToken(iden))
	router.POST(refleshTokenRoute, refleshToken(iden))
	router.Use(verifyToken(iden))
}
