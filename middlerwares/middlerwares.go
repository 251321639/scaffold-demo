package middlerwares

import (
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	requestUrl := c.FullPath()
	logs.Debug(map[string]any{"请求路径": requestUrl}, "")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		c.Next()
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		logs.Warning(nil, "请求头中缺少Authorization字段")
		c.JSON(401, gin.H{
			"message": "请求头中缺少Authorization字段",
		})
		c.Abort()
		return
	}
	claims, err := jwtutil.ParseToken(token)
	if err != nil {
		logs.Warning(map[string]any{"token": token}, "解析token失败")
		c.JSON(401, gin.H{
			"message": "解析token失败",
		})
		c.Abort()
		return
	}
	logs.Debug(map[string]any{"用户名": claims.Username}, "解析token成功")
	c.Set("claims", claims)
	c.Next()	
}
