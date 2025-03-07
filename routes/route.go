package routes

import (
	"fmt"
	"go-app/controllers"
	"go-app/middlewares"
	"go-app/repositories"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
	//[Line 3]
	httpRouter := gin.Default()

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
	if hasPolicy,_ := enforcer.HasPolicy("doctor", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("doctor", "report", "read")
	}	
	if hasPolicy,_ := enforcer.HasPolicy("doctor", "report", "write"); !hasPolicy {
		enforcer.AddPolicy("doctor", "report", "write")
	}	
	if hasPolicy,_ := enforcer.HasPolicy("patient", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("patient", "report", "read")
	}	

	//Registry All Repositories
	userRepository := repositories.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User Migrate err", err)
	}

	//Registry All Controllers
	userController := controllers.NewUserController(userRepository)	

	//Setup Route
	apiRoutes := httpRouter.Group("/api")
	{
		apiRoutes.POST("/registry", userController.AddUser(enforcer))
		apiRoutes.POST("/signin", userController.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middlewares.AuthorizeJWT())
	{
		//[Line 44-49]
		userProtectedRoutes.GET("/", middlewares.Authorize("report", "read", enforcer), userController.GetAllUser)
		userProtectedRoutes.GET("/:user", middlewares.Authorize("report", "read", enforcer), userController.GetUser)
		userProtectedRoutes.PUT("/:user", middlewares.Authorize("report", "write", enforcer), userController.UpdateUser)
		userProtectedRoutes.DELETE("/:user", middlewares.Authorize("report", "write", enforcer), userController.DeleteUser)
	}

	httpRouter.Run()
}
