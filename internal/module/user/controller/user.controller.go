package controller

import (
	"github.com/HiBang15/sample-gorm.git/constant"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/dto"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/services"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/transformers"
	"github.com/HiBang15/sample-gorm.git/internal/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var createUserRequest *dto.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		util.SetResponse(c, http.StatusUnprocessableEntity, err, constant.InvalidRequestBody, nil)
		return
	}

	userTransformer := transformers.NewUserTransformer()
	errValid := userTransformer.VerifyCreateUserRequest(createUserRequest)
	if errValid != nil {
		util.SetResponse(c, http.StatusUnprocessableEntity, errValid, constant.InvalidRequestBody, nil)
		return
	}

	//create user via user service
	userClient := services.NewUserService()
	user, err := userClient.CreateUser(createUserRequest)
	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}

	util.SetResponse(c, http.StatusCreated, nil, constant.CreateUserSuccess, user)
	return
}

func GetUsers(c *gin.Context) {
	userClient := services.NewUserService()
	users, err := userClient.GetUsers()

	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}
	util.SetResponse(c, http.StatusOK, nil, constant.GetUserSuccess, users)
	return
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	userClient := services.NewUserService()
	user, err := userClient.GetUser(id)

	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}
	util.SetResponse(c, http.StatusOK, nil, constant.GetUserSuccess, user)
	return
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userClient := services.NewUserService()
	err := userClient.DeleteUser(id)
	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}
	util.SetResponse(c, http.StatusOK, nil, constant.DeleteUserSuccess, 0)
	return
}
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var createUserRequest *dto.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		util.SetResponse(c, http.StatusUnprocessableEntity, err, constant.InvalidRequestBody, nil)
		return
	}
	userTransformer := transformers.NewUserTransformer()
	errValid := userTransformer.VerifyCreateUserRequest(createUserRequest)
	if errValid != nil {
		util.SetResponse(c, http.StatusUnprocessableEntity, errValid, constant.InvalidRequestBody, nil)
		return
	}
	userClient := services.NewUserService()
	user, err := userClient.UpdateUser(id, createUserRequest)
	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}
	util.SetResponse(c, http.StatusOK, nil, constant.UpdateUserSuccess, user)
	return
}
func GetUsersWithPagination(c *gin.Context) {
	limit := c.Query("limit")
	pageNumber := c.Query("pageNumber")
	limitInt, _ := strconv.Atoi(limit)
	pageNumberInt, _ := strconv.Atoi(pageNumber)
	userClient := services.NewUserService()
	users, err := userClient.GetUsersWithPagination(limitInt, pageNumberInt)
	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}
	util.SetResponse(c, http.StatusOK, nil, constant.GetUserSuccess, users)
	return
}

func GetUserWithSearch(c *gin.Context) {
	query := c.Query("query")
	limit := c.Query("limit")
	pageNumber := c.Query("pageNumber")

	userClient := services.NewUserService()
	log.Println(query, "whitespace")
	users, limitInt, pageNumberInt, len, err := userClient.GetUserWithSearch(limit, pageNumber, query)

	if err != nil {
		util.SetResponse(c, http.StatusInternalServerError, err, err.Error(), 0)
		return
	}
	util.SetResponse(c, http.StatusOK, nil, constant.GetUserSuccess, map[string]interface{}{"users": users, "limit": limitInt, "page number": pageNumberInt, "total": len})
	return

}
