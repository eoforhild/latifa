package router

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"latifa/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	database := mongodb.Database("latifa_info")
	usercoll := database.Collection("users")
	err := usercoll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		NewError(err).Abort(c)
	}

	if err := verifyEncryptedPassword(user.Password, r.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "These credentials do not match our records.",
		})
	}

	randomString, _ := generateRandomString(32)

	token := &entity.Token{
		UserId: user.ID,
		Token:  randomString,
	}

	tokencoll := database.Collection("auth_tokens")
	_, err = tokencoll.InsertOne(context.TODO(), token)
	if err != nil {
		NewError(err).Abort(c)
	}

	c.JSON(http.StatusOK, gin.H{"token": token.Token})
}

func postLogout(c *gin.Context) {
	user := ExtractUser(c)
	mongodb := ExtractMongoClient(c)
	auth_token, _ := c.Get("Authorization")

	filter := bson.M{
		"token":   auth_token,
		"user_id": user.ID,
	}
	database := mongodb.Database("latifa_info")
	tokencoll := database.Collection("auth_tokens")

	_, err := tokencoll.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Status(http.StatusNoContent)
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`

	// Keys
	Ik       string     `json:"ik" binding:"required"`
	Spk      string     `json:"spk" binding:"required"`
	SpkSig   string     `json:"spk_sig" binding:"required"`
	PqSpk    string     `json:"pqspk" binding:"required"`
	PqSpkSig string     `json:"pqspk_sig" binding:"required"`
	Opk      [32]string `json:"opk_arr" binding:"required"`
	PqOpk    [32]string `json:"pqopk_arr" binding:"required"`
	PqOpkSig [32]string `json:"pqopk_sig_arr" binding:"required"`
}

func postRegister(c *gin.Context) {
	var r RegisterRequest
	if err := c.Bind(&r); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mongodb := ExtractMongoClient(c)
	database := mongodb.Database("latifa_info")
	usercoll := database.Collection("users")

	filter := bson.M{
		"username": r.Username,
		"email":    r.Email,
	}
	var check bson.M
	if err := usercoll.FindOne(context.TODO(), filter).Decode(&check); err == nil {
		// document found
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"error": "A user with that username or email already exists.",
		})
		return
	}

	newPassword, err := encryptPassword(r.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "An unexpected error was encountered while processing this request.",
		})
	}

	newUser := bson.M{
		"username": r.Username,
		"email":    r.Email,
		"password": newPassword,
		"device":   c.GetHeader("User-Agent"),
		"ip":       c.ClientIP(),
	}

	result, err := usercoll.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
	}

	keycoll := database.Collection("public_keys")
	println("%v\n", result.InsertedID)
	key := &entity.Key{
		ID:       result.InsertedID.(primitive.ObjectID),
		Ik:       r.Ik,
		Spk:      r.Spk,
		SpkSig:   r.SpkSig,
		PqSpk:    r.PqSpk,
		PqSpkSig: r.PqSpkSig,
		Opk:      r.Opk,
		PqOpk:    r.PqOpk,
		PqOpkSig: r.PqOpkSig,
	}

	_, err = keycoll.InsertOne(context.TODO(), key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
	}

	c.Status(http.StatusNoContent)
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

func generateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := hex.EncodeToString(randomBytes)
	return randomString, nil
}
