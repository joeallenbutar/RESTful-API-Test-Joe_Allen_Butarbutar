package controllers

import (
	services "RESTful-API-Test-Joe_Allen_Butarbutar/services"
	"RESTful-API-Test-Joe_Allen_Butarbutar/models"
	"fmt"
	_"log"
	"net/http"
	_"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/google/uuid"

)

type APIEnv struct {
	DB *gorm.DB
}

func (UserControllers *APIEnv) GetAllUsers(c *gin.Context) {
	requestPagination := models.Pagination{
		Page : c.Query("page"),
		Limit : c.Query("limit"),
		Sort : c.Query("sort"),
		Filter : c.Query("filter"),
	}
	generatePagination := GeneratePaginationFromRequest(c, requestPagination)
	Users, err := services.GetAllUsers(UserControllers.DB, &generatePagination)
	if err != nil {
		fmt.Println(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, Users)
}

func (UserControllers *APIEnv) CreateUser(c *gin.Context) {
	user := models.User{}
	uuid := uuid.New()
	user.Id = uuid.String()
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	if err := services.CreateUser(UserControllers.DB, &user); err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"data" : user,
	})
}

func (UserControllers *APIEnv) GetUserByID(c *gin.Context) {
	userId := c.Params.ByName("id")

	User, exists, err := services.GetUserByID(userId, UserControllers.DB)
	if err != nil {
		fmt.Println(c, http.StatusNotFound, err)
		return
	}

	if !exists {
		fmt.Println(c, http.StatusNotFound, err)
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, User)
}

func (UserControllers *APIEnv) UpdateUser(c *gin.Context) {
	userId := c.Params.ByName("id")

	dataUser, exists, err := services.GetUserByID(userId, UserControllers.DB)
	if err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		fmt.Println(c, http.StatusNotFound, err)
		c.JSON(http.StatusOK, gin.H{
			"message": "user not found",
		})
		return
	}

	updatedUser := models.User{}
	err = c.BindJSON(&updatedUser)
	if err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	updatedUser.Id = userId
	updatedUser.CreatedAt = dataUser.CreatedAt

	if err := services.UpdateUser(UserControllers.DB, &updatedUser); err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	UserControllers.GetUserByID(c)
}

func (UserControllers *APIEnv) DeleteUser(c *gin.Context) {
	userId := c.Params.ByName("id")

	_, exists, err := services.GetUserByID(userId, UserControllers.DB)
	if err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		fmt.Println(c, http.StatusNotFound, err)
		c.JSON(http.StatusOK, gin.H{
			"message": "user has been deleted",
		})
		return
	}

	err = services.DeleteUser(userId, UserControllers.DB)
	if err != nil {
		fmt.Println(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

func GeneratePaginationFromRequest(c *gin.Context, requestPagination models.Pagination) models.Pagination {
	limit := requestPagination.Limit
	page := requestPagination.Page
	sort := requestPagination.Sort
	filter := requestPagination.Filter
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit = queryValue
			break
		case "page":
			page = queryValue
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
		Filter:  filter,
	}

}