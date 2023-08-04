package controller

import (
	"desa-sragen/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cx *Controller) Register(c *gin.Context) {

	c.HTML(200, "register.html", gin.H{
		"url": cx.Env(),
	})
}

func (cx *Controller) RegisterHandler(c *gin.Context)  {
	var req domain.RegisterParam
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, "bad request")
		return
	}


	checkUser, err := cx.Repo.GetUser(req.Username)
	if err != nil || checkUser != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse(404, "username sudah digunakan", domain.Empty{}))
		return
	}

	data := domain.Auth{
		Username: req.Username,
		Email: req.Email,
		Password: domain.HashAndSalt([]byte(req.Password)),
		RoleId: 3,
		IsActive: true,
		CreatedAt: domain.TimeNow(),
		UpdatedAt: domain.TimeNow(),
	}

	if err := cx.Repo.CreateUser(data); err != nil {
		c.JSON(http.StatusInternalServerError, domain.JsonResponse(500, err.Error(), domain.Empty{}))
		return
	}

	c.JSON(201, domain.JsonResponse(201, "Success", domain.Empty{}))
}