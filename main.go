package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/rafael-abuawad/refactored-stack/app/controllers"
	"github.com/rafael-abuawad/refactored-stack/app/models"
	"github.com/rafael-abuawad/refactored-stack/config"
	"github.com/rafael-abuawad/refactored-stack/config/database"
	"github.com/rafael-abuawad/refactored-stack/utils"
	"github.com/rafael-abuawad/refactored-stack/utils/middleware"
)

func main() {
	database.Connect()
	database.Migrate(&models.User{})

	engine := html.New("./app/views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		ErrorHandler:      utils.ErrorHandler,
		Views:             engine,
		ViewsLayout:       "layouts/_base",
		PassLocalsToViews: true,
	})

	app.Use(logger.New())
	app.Use(middleware.FlashMiddleware)
	app.Use(middleware.AuthMiddleware)

	controllers.AuthController(app)
	controllers.TodosController(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.PORT)))
}
