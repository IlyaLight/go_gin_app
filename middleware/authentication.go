package middleware

import (
	"go_gin_app/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var code int
		//var data interface{}

		token := c.Query("token")
		if token == "" {
			//code = http.StatusBadRequest
		} else {
			_, err := util.CheckToken(&token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					//code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					//code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		//if code != e.SUCCESS {
		//	c.JSON(http.StatusUnauthorized, gin.H{
		//		"code": code,
		//		"msg":  e.GetMsg(code),
		//		"data": data,
		//	})
		//
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}
