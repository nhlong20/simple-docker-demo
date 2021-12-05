package main

import (
	"btcn03-api/component/appctx"
	"btcn03-api/middleware"
	coursemodel "btcn03-api/modules/course/model"
	coursegin "btcn03-api/modules/course/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	//db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	errMigrate := appCtx.GetMainDBConnection().AutoMigrate(&coursemodel.Course{})
	if errMigrate != nil {
		panic("Error while creating DB model")
	}

	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	{
		restaurants := v1.Group("/courses")
		{
			restaurants.POST("", coursegin.CreateCourse(appCtx))
			restaurants.GET("", coursegin.ListCourse(appCtx))
		}
	}

	r.Run(":8080")
}
