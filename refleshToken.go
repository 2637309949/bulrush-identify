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

func refleshToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		refreshToken := GetAccessToken(c)
		if refreshToken != "" {
			originToken := iden.RefleshToken(refreshToken)
			// reflesh and save
			if originToken != nil {
				c.JSON(http.StatusOK, originToken)
			} else {
				rushLogger.Warn("reflesh token failed,token may not exist")
				c.JSON(http.StatusBadRequest, gin.H{"message": "reflesh token failed, token may not exist"})
			}
		}
		c.Abort()
	}
}
