/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-addition/redis"
	identify "github.com/2637309949/bulrush-identify"
	"github.com/gin-gonic/gin"
)

// Login login info bind
type Login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	Type     string `form:"type" json:"type" xml:"type"`
}

// GINMODE APP ENV
var GINMODE = bulrush.Some(os.Getenv("GIN_MODE"), "local")

// Redis store
var Redis = redis.New(bulrush.NewCfg(path.Join(".", fmt.Sprintf("%s.yaml", GINMODE))))
var auth = &identify.Identify{
	Auth: func(c *gin.Context) (interface{}, error) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			return nil, err
		}
		if login.Password == "xx" && login.UserName == "xx" {
			return map[string]interface{}{
				"id":       "3e4r56u80a55",
				"username": login.UserName,
			}, nil
		}
		return nil, errors.New("User authentication failed")
	},
	Tokens: &identify.RedisTokensGroup{
		Redis: Redis,
	},
	FakeTokens: []interface{}{"DEBUG"},
	FakeURLs:   []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`},
}

func main() {
	app := bulrush.Default()
	app.Config(path.Join(".", fmt.Sprintf("%s.yaml", GINMODE)))
	app.Inject("bulrushApp")
	app.Use(auth)
	app.Use(bulrush.PNQuick(func(testInject string, router *gin.RouterGroup) {
		router.GET("/identify", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": testInject,
			})
		})
		router.GET("/vaulting", func(c *gin.Context) {
			data := auth.ObtainToken(map[string]interface{}{
				"id":       "test",
				"username": "test",
			})
			c.JSON(http.StatusOK, data)
		})
		router.GET("/utils", func(c *gin.Context) {
			accessToken := identify.GetAccessToken(c)
			accessData := identify.GetAccessData(c)
			c.JSON(http.StatusOK, gin.H{
				"accessToken": accessToken,
				"accessData":  accessData,
			})
		})
	}))
	app.Run(func(err error, config *bulrush.Config) {
		if err != nil {
			panic(err)
		} else {
			name := config.GetString("name", "")
			port := config.GetString("port", "")
			fmt.Println("================================")
			fmt.Printf("App: %s\n", name)
			fmt.Printf("Listen on %s\n", port)
			fmt.Println("================================")
		}
	})
}
