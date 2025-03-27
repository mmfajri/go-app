package routes

import (
	"fmt"
	"go-app/controllers"
	"go-app/docs"
	"go-app/middlewares"
	"go-app/repositories"
	"log"

	_ "go-app/docs"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

/*
Run this command
swag init -g routes/route.go (<~/ {folder_name_that_has_root} / {route_file_name}.go>)
*/
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetupRoutes(db *gorm.DB) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost:8080"
	//[Line 3]
	httpRouter := gin.Default()

	httpRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Initialize casbin adapter
	//[Line 6]
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	//Load model configuration file and policy store adapter
	//[Line 12]
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	//add policy for req
	//one time only
	//[Line 17-26]
	if hasPolicy, _ := enforcer.HasPolicy("doctor", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("doctor", "report", "read")
	}
	if hasPolicy, _ := enforcer.HasPolicy("doctor", "report", "write"); !hasPolicy {
		enforcer.AddPolicy("doctor", "report", "write")
	}
	if hasPolicy, _ := enforcer.HasPolicy("patient", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("patient", "report", "read")
	}

	//Registry All Repositories
	userRepository := repositories.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User Migrate err", err)
	}
	reportRepository := repositories.NewReportRepository(db)
	if err := reportRepository.Migrate(); err != nil {
		log.Fatal("Report Migrate err", err)
	}

	//Registry All Controllers
	userController := controllers.NewUserController(userRepository)
	reportController := controllers.NewReportController(reportRepository)

	//Setup Route
	apiRoutes := httpRouter.Group("/api")
	{
		apiRoutes.POST("/registry", userController.AddUser(enforcer))
		apiRoutes.POST("/signin", userController.SignInUser)
		//Testing API
		apiRoutes.GET("/report/get_all", reportController.GetAll)
	}

	reportProtectedRoutes := apiRoutes.Group("/report", middlewares.AuthorizeJWT())
	{
		reportProtectedRoutes.GET("/", reportController.Get)
		reportProtectedRoutes.POST("/add", reportController.Add)
		reportProtectedRoutes.DELETE("/delete", reportController.Delete)
		reportProtectedRoutes.PUT("/update", reportController.Update)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middlewares.AuthorizeJWT())
	{
		//[Line 44-49]
		userProtectedRoutes.GET("/", middlewares.Authorize("report", "read", enforcer), userController.GetAllUser)
		userProtectedRoutes.GET("/:user", middlewares.Authorize("report", "read", enforcer), userController.GetUser)
		userProtectedRoutes.PUT("/:user", middlewares.Authorize("report", "write", enforcer), userController.UpdateUser)
		userProtectedRoutes.DELETE("/:user", middlewares.Authorize("report", "write", enforcer), userController.DeleteUser)
	}

	roleProtectedRoutes := apiRoutes.Group("/roles", middlewares.AuthorizeJWT())
	{
		roleProtectedRoutes.GET("/", middlewares.Authorize("role", "read", enforcer))
		roleProtectedRoutes.GET("/:role", middlewares.Authorize("role", "read", enforcer))
		roleProtectedRoutes.PUT("/:role", middlewares.Authorize("role", "write", enforcer))
		roleProtectedRoutes.POST("/add", middlewares.Authorize("role", "write", enforcer))
		roleProtectedRoutes.DELETE("/:user", middlewares.Authorize("role", "write", enforcer))
	}

	httpRouter.Run(":" + "8080")
}
