package middleware

import (
	"desa-sragen/domain"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJwt(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			token := c.Query("token")
			if token == "" {
				c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
				c.Abort()
				return
			}
			authHeader = "Bearer "+token
		}

		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := IsAuthorized(authToken, secret)
			if authorized {
				claimedToken, err := ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
			c.Abort()
			return
				}
				c.Set("username", claimedToken.Username)
				c.Set("role", claimedToken.Role)
				c.Set("userId", claimedToken.UserId)
				c.Next()
				return
			}
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
			c.Abort()
			return
		}
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
			c.Abort()

	}
}

func ValidateSuper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		privilege := ctx.GetInt("role")

		if privilege == 1 {
			ctx.Next()
			return
		}

		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		ctx.Abort()
	}
}

func ValidateAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		privilege := ctx.GetInt("role")
		fmt.Println(privilege)
		if privilege == 2 || privilege == 1 {
			ctx.Next()
			return
		}

		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		ctx.Abort()
	}
}

func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		privilege := ctx.GetInt("role")
		if privilege == 3 {
			ctx.Next()
			return
		}

		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		ctx.Abort()
	}
}

func ValidateToken(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Query("token")
		if authHeader == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
			c.Abort()
			return
		}

		authorized, _ := IsAuthorized(authHeader, secret)
		if authorized {
			claimedToken, err := ExtractIDFromToken(authHeader, secret)
			if err != nil {
				c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		c.Abort()
		return
			}
			c.Set("username", claimedToken.Username)
			c.Set("role", claimedToken.Role)
			c.Set("userId", claimedToken.UserId)
			c.Next()
			return
		}
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		c.Abort()

	}
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (*domain.ClaimedToken, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid Token")
	}

	var claim domain.ClaimedToken

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token userId")
	}
	roleId, ok := claims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token roleId")
	}
	username, ok := claims["user_name"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token username")
	}


	claim.UserId = int(userId)
	claim.Username = username
	claim.Role = domain.RoleTo(roleId)

	return &claim, nil
}