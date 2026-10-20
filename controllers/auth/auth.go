package auth

import (
	"scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	userInfo := UserInfo{}
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		r.JSON(401, config.NewResponse(401, err.Error(), nil))
		return
	}
	logs.Debug(map[string]any{"用户名": userInfo.Username, "密码": userInfo.Password}, "用户登录请求")

	// 这里可以添加实际的登录逻辑，例如验证用户名和密码
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		token, err := jwtutil.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]any{"用户名": userInfo.Username}, "生成token失败")
			r.JSON(500, config.NewResponse(500, "生成token失败", nil))
			return
		}
		logs.Debug(map[string]any{"用户名": userInfo.Username, "token": token}, "用户登录成功，生成token")
		r.JSON(200, config.NewResponse(200, "登录成功", map[string]any{"token": token}))

		// r.Set("Authorization", token)
	} else {
		logs.Warning(map[string]any{"用户名": userInfo.Username}, "用户名或密码错误")
		r.JSON(401, config.NewResponse(401, "用户名或密码错误", nil))
	}
}

func Logout(r *gin.Context) {
	logs.Debug(nil, "用户登出请求")
	// 这里可以添加实际的登出逻辑，例如清除会话或令牌
	r.JSON(200, config.NewResponse(200, "登出成功", nil))
}
