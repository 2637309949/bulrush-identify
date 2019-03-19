/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush identify plugin]
 */

package identify

import "github.com/gin-gonic/gin"

func accessToken(iden *Identify) func(*gin.Context) {
	return func(c *gin.Context) {
		var accessToken string
		queryToken := c.Query("accessToken")
		formToken := c.PostForm("accessToken")
		headerToken := c.Request.Header.Get("Authorization")
		if queryToken != "" {
			accessToken = queryToken
		} else if formToken != "" {
			accessToken = formToken
		} else if headerToken != "" {
			accessToken = headerToken
		}
		setAccessToken(c, accessToken)
		c.Next()
	}
}
