package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spikeKill/pkg/e"
	"spikeKill/pkg/util"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		authHeader := c.Request.Header.Get("Authorization")
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			code = e.INVALID_PARAMS
		} else {
			token := parts[1]
			if token == "" {
				code = e.INVALID_PARAMS
			} else {
				claims, err := util.ParseToken(token)
				if err != nil {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
