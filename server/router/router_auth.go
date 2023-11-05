package router

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Device   string `json:"device"`
	Ip       string `json:"ip"`
}

func postLogin(c *gin.Context) {
	var r LoginRequest
	if err := c.BindJSON(&r); err != nil {
		return
	}
}

func postLogout(c *gin.Context) {

}

type LogoutRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Device   string `json:"device"`
	Ip       string `json:"ip"`
}

func postRegister(c *gin.Context) {
	var r LogoutRequest
	if err := c.BindJSON(&r); err != nil {
		return
	}
}
