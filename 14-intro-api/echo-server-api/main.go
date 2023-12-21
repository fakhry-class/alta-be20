package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func main() {
	// create new instance echo
	e := echo.New()
	// define endpoint
	e.POST("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST Hello, World!")
	})

	e.GET("/hello", HelloController)
	e.GET("/users", GetUsersController)

	// url param
	e.GET("/users/:user_id", GetUserByIdController)
	// query param
	e.GET("/articles", GetArticlesController)
	//bind
	e.POST("/articles", CreateArticleController)
	// form value
	e.POST("/users", CreateUserController)

	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

func HelloController(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status
	// return c.String(http.StatusOK, "Hello World")
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "hello world",
	})
}

func GetUsersController(c echo.Context) error {
	user := User{Name: "Ismail", Email: "ismail@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func GetArticlesController(c echo.Context) error {
	qPage := c.QueryParam("page")
	qSearch := c.QueryParam("search")

	// dummy data
	var data = []article{
		{1, "lorem", "lorem"},
		{2, "ipsum", "ipsum"},
		{3, "alterra", "academy"},
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"data":    data,
		"page":    qPage,
		"search":  qSearch,
	})
}

// URL PARAM
func GetUserByIdController(c echo.Context) error {
	idParam := c.Param("user_id")
	idInt, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "error",
			"message": "error convert param",
		})
	}
	// logic
	// ....
	return c.JSON(http.StatusOK, map[string]any{
		"status":    "success",
		"message":   "success read data",
		"param":     idParam,
		"param_int": idInt,
	})
}

// request body via BINDING
func CreateArticleController(c echo.Context) error {
	newArticle := article{}
	// newArticle := new(article)
	errBind := c.Bind(&newArticle)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "error",
			"message": "error bind data" + errBind.Error(),
		})
	}

	// simpan ke db
	//
	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add data",
		"result":  newArticle,
	})
}

func CreateUserController(c echo.Context) error {
	newUser := User{}
	name := c.FormValue("name")
	email := c.FormValue("email")
	newUser.Name = name
	newUser.Email = email
	// c.Request().FormFile("image")

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add data",
		"result":  newUser,
	})
}
