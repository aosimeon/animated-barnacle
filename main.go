package main

// @title           AltSchool Capstone Project - URL Shortner
// @version         1.0
// @description     Shorten long urls.

// @contact.name   Olusola Amoo
// @contact.url    https://twitter.com/aosimeon
// @contact.email  aosimeon@outlook.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      url-shortner-3b6b.onrender.com
// @BasePath  /

import (
	"errors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"url-shortner/database"
	_ "url-shortner/docs"
	"url-shortner/models"
)

var db = database.InitDb()

type UrlPayload struct {
	LongUrl string
}

// Shorten Url godoc
// @Summary Shorten a url
// @Description Shorten a url
// @Accept json
// @Produce json
// @Param longUrl body UrlPayload true "Url to shorten e.g https://www.altschoolafrica.com/"
// @Success 200 {object} models.Url
// @Router /shorten [post]
func shortenUrl(c *gin.Context) {
	var url models.Url
	c.BindJSON(&url)

	id, _ := shortid.New(1, shortid.DefaultABC, 2342)
	result, _ := id.Generate()

	// Check of id already exists. If it does, regenerate id
	for err := models.GetUrl(db, &url, result); err == nil; {
		result, _ = id.Generate()
	}

	url.ShortUrl = result

	err := models.ShortenUrl(db, &url)
	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, url)
}

// Redirect to a Url godoc
// @Summary Redirect to a url
// @Description Redirect to a url
// @Accept json
// @Produce json
// @Param short-url path string true "Short url"
// @Success 301
// @Router /{short-url} [get]
func redirectUrl(c *gin.Context) {
	var url models.Url
	err := models.GetUrl(db, &url, c.Param("id"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.LongUrl)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	err := db.AutoMigrate(&models.Url{})
	if err != nil {
		return nil
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Shorten url
	r.POST("/shorten", shortenUrl)

	// Redirect to long url
	r.GET("/:id", redirectUrl)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
