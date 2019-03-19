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

func obtainToken(iden *Identify) func(c *gin.Context) {
	return func(c *gin.Context) {
		authData, err := iden.Auth(c)
		if authData != nil {
			data := iden.ObtainToken(authData)
			c.JSON(http.StatusOK, data)
			c.Abort()
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}
	}
}
