package domain

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Empty struct{}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type jwtCustomClaim struct {
	UserID        int `json:"user_id"`
	Username      string `json:"user_name"`
	Role           string    `json:"role"`
	jwt.StandardClaims
}

func TimeNow() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc)
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}

func  GenerateToken(UserID int, secret, Username, role string) string {
	claims := &jwtCustomClaim{
		UserID,
		Username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			// Issuer:    j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return t
}

func RoleH(role int8) string {
	switch role{
	case 1:
		return "super"
	case 2:
		return "admin"
	case 3:
		return "user"
	default:
		return "user"
	}
}

func RoleTo(role string) int {
	switch role{
	case "super":
		return 1
	case "admin":
		return 2
	case "user":
		return 3
	default:
		return 3 
	}
}

func IntToString(i int) string {
	str := strconv.Itoa(i)

	return str
}

func StringToInt(str string) int {
	i , _ := strconv.Atoi(str)
	return i
}

func AuthProfile(c *gin.Context) ProfileAuth{
	return ProfileAuth{
		Username: c.GetString("username"),
		Role: int8(c.GetInt("role")),
		UserId: c.GetInt("userId"),
	}
}

func StatusActive(b bool) string {
	if b {
		return "Aktif"
	}

	return "Tidak Aktif"
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GenerateUuid() string {
	return uuid.New().String()
}