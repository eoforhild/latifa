package router

import (
	"context"
	"latifa/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func postLogin(c *gin.Context) {
	var r LoginRequest
	if err := c.BindJSON(&r); err != nil {
		return
	}

	filter := bson.M{
		"email": r.Login,
	}

	var user entity.User
	mongodb := ExtractMongoClient(c)
	collection := mongodb.Database("latifa_info").Collection("users")
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		NewError(err).Abort(c)
	}

	if err := verifyEncryptedPassword(user.Password, r.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "These credentials do not match our records.",
		})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func postLogout(c *gin.Context) {

}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func postRegister(c *gin.Context) {
	var r RegisterRequest
	if err := c.BindJSON(&r); err != nil {
		return
	}

	newPassword, err := encryptPassword(r.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "An unexpected error was encountered while processing this request.",
		})
	}

	user := &entity.User{
		Username: r.Username,
		Email:    r.Email,
		Password: newPassword,
		Device:   c.GetHeader("User-Agent"),
		Ip:       c.ClientIP(),
	}

	mongodb := ExtractMongoClient(c)
	collection := mongodb.Database("latifa_info").Collection("users")

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func encryptPassword(plainPassword string) (string, error) {
	cost := 12
	newPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), cost)
	if err != nil {
		return "", err
	}
	return string(newPassword), nil
}

func verifyEncryptedPassword(hashedPassword, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err
}
