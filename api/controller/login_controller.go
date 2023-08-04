package controller

import (
	"desa-sragen/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cx *Controller) Login(c *gin.Context) {

	c.HTML(200, "login.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) LoginHandler(c *gin.Context)  {
	var req domain.LoginParam
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, "bad request")
		return
	}  

	if req.Password == "" || req.Username == "" {
		c.JSON(400, domain.JsonResponse(400, "Username atau Password harus diisi", domain.Empty{}))
		return
	}

	checkUser, err := cx.Repo.GetUser(req.Username)
	if err != nil || checkUser == nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse(404, "username tidak ditemukan", domain.Empty{}))
		return
	}

	if !checkUser.IsActive  {
		c.JSON(http.StatusBadRequest, domain.JsonResponse(404, "akun anda sudah tidak aktif", domain.Empty{}))
		return
	}
	
	if !domain.ComparePassword(checkUser.Password, []byte(req.Password)) {
		c.JSON(http.StatusBadRequest, domain.JsonResponse(404, "password tidak sesuai", domain.Empty{}))
		return
	}

	token := domain.GenerateToken(checkUser.AuthId, cx.Config.SecretKey, checkUser.Username, domain.RoleH(checkUser.RoleId))
	if token == "" {
		c.JSON(http.StatusBadRequest, domain.JsonResponse(404, "Login gagal", domain.Empty{}))
		return
	}

	c.JSON(200, domain.JsonResponse(200, "Success", domain.LoginResponse{
		Username: checkUser.Username,
		Email: checkUser.Email,
		Role: domain.RoleH(checkUser.RoleId),
		Token: token,
		IsActive: checkUser.IsActive,
	}))
}

func (cx *Controller) SuccessLogin(c *gin.Context)  {
	token := c.Query("token")

	// Redirect to a new URL
	newURL := cx.Env()+"dashboard"
	newToken := "Bearer "+ token

	targetReq, err := http.NewRequest(http.MethodGet, newURL, nil)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error creating request")
        return
    }

    // Set the Authorization header in the new request
    targetReq.Header.Set("Authorization", newToken)

    // Perform the request and get the response
    client := http.Client{}
    targetResp, err := client.Do(targetReq)
    if err != nil {
        c.String(http.StatusInternalServerError, "Error performing request")
        return
    }
    defer targetResp.Body.Close()
}