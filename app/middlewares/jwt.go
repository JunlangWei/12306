package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"time"
)

var jwtSecret = []byte(getSignalKey())

func getSignalKey() string {
	cfg, _ := ini.Load("conf/config.ini")
	return cfg.Section("server").Key("sign_key").String()
}


type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "12306",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		claims, err := ParseToken(token)
		if err == nil {
			context.Set("claims", claims)
		}
		context.Next()
	}
}

//func jwtAbort(c *gin.Context, msg string) {
//	c.JSON(http.StatusUnauthorized, gin.H{
//		"status":  "error",
//		"message": msg,
//	})
//	c.Abort()
//}
