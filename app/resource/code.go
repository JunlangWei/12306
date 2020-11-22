package resource

import "net/http"

var (
	Success   = http.StatusOK           // 成功
	AuthError = http.StatusUnauthorized // 权限错误
	Parameter = http.StatusBadRequest   // 参数错误
)
