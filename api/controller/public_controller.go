package controller

import (
	"bytes"
	"desa-sragen/bootrstrap"
	"desa-sragen/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{
	Config *bootrstrap.Env
	Repo domain.Repo
}

func (cx *Controller) Env()  string {
	return cx.Config.ServerUrl
}

func (cx *Controller) Dashboard(c *gin.Context) {


	c.HTML(200, "index.html", gin.H{
		"url":cx.Env(),
	})
}

func (cx *Controller) ValidateToken(c *gin.Context)  {
	c.JSON(200, domain.JsonResponse(200, "Validated", domain.Empty{}))
}

func (cx *Controller) ValidateTokenAdmin(c *gin.Context)  {
	c.JSON(200, domain.JsonResponse(200, "Validated", domain.Empty{}))
}

func (cx *Controller) ValidateTokenSuper(c *gin.Context)  {
	c.JSON(200, domain.JsonResponse(200, "Validated", domain.Empty{}))
}

func (cx *Controller) SendEmail(c *gin.Context)  {
	var request domain.EmailParam
	if c.ShouldBindJSON(&request)!= nil {
		c.JSON(400, domain.JsonResponse(400, "bad request", domain.Empty{}))
		return
	}
	
	// Replace with your SendGrid API key
	apiKey := "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"

	// Compose the email message
	toEmail := "desabendo1703@gmail.com"
	fromEmail := request.Email
	subject := request.Subject
	message := request.Message
	name := request.Name

	url := "https://api.brevo.com/v3/smtp/email"
	method := "POST"

	payload :=fmt.Sprintf(`{
		"sender": {
			"name": "%s",
			"email": "%s"
		},
		"to": [
			{
				"email": "%s",
				"name": "%s"
			}
		],
		"subject": "%s",
		"htmlContent": "<html><head></head><body>%s.</p></body></html>"
	}`, name, fromEmail, toEmail, name, subject, message)


	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set request headers
	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", apiKey)
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		c.JSON(500, domain.JsonResponse(500, "failed to send email", domain.Empty{}))
		return 
	}
	defer resp.Body.Close()

	// Read the response body
	var responseBody bytes.Buffer
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		c.JSON(500, domain.JsonResponse(500, "failed to send email", domain.Empty{}))
		return
	}

	if resp.StatusCode != 201 {
		c.JSON(400, domain.JsonResponse(400,responseBody.String(), domain.Empty{}))
		return
	}

	c.JSON(200, domain.JsonResponse(200,"success", domain.Empty{}))

}