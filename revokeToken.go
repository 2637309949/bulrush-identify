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

	"github.com/gin-gonic/gin"
)

func revokeToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		accessToken := GetAccessToken(c)
		if accessToken != "" {
			result := iden.RevokeToken(accessToken)
			if result {
				c.JSON(http.StatusOK, gin.H{
					"success": true,
				})
			} else {
				rushLogger.Warn("revoke token failed,token may not exist")
				c.JSON(http.StatusBadRequest, gin.H{"message": "revoke token failed, token may not exist"})
			}
		}
		c.Abort()
	}
}
