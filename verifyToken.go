/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush identify plugin]
 */

package identify

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func verifyToken(iden *Identify) func(*gin.Context) {
	FakeURLs := iden.FakeURLs
	FakeTokens := iden.FakeTokens
	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		accessToken := GetAccessToken(c)
		fakeURL := Find(FakeURLs, func(regex interface{}) bool {
			r, _ := regexp.Compile(regex.(string))
			return r.MatchString(reqPath)
		})
		fakeToken := Find(FakeTokens, func(token interface{}) bool {
			return accessToken == token.(string)
		})
		if fakeURL != nil {
			c.Next()
		} else if fakeToken != nil {
			c.Next()
		} else {
			ret := iden.VerifyToken(accessToken)
			if ret {
				rawToken := iden.Tokens.Find(accessToken, "")
				setAccessData(c, rawToken["extra"])
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Invalid token",
				})
				c.Abort()
			}
		}
	}
}
