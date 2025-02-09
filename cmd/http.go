package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/thrifting-app-service/helpers"
	productAPI "github.com/hilmiikhsan/thrifting-app-service/internal/api/product"
	userAPI "github.com/hilmiikhsan/thrifting-app-service/internal/api/user"
	"github.com/hilmiikhsan/thrifting-app-service/internal/interfaces"
	productRepository "github.com/hilmiikhsan/thrifting-app-service/internal/repository/product"
	userRepository "github.com/hilmiikhsan/thrifting-app-service/internal/repository/user"
	productServices "github.com/hilmiikhsan/thrifting-app-service/internal/services/product"
	userServices "github.com/hilmiikhsan/thrifting-app-service/internal/services/user"
	"github.com/hilmiikhsan/thrifting-app-service/internal/validator"
	"github.com/sirupsen/logrus"
)

func ServeHTTP() {
	dependency := dependencyInject()

	router := gin.Default()

	routerV1 := router.Group("/api/v1")
	routerV1.POST("/login", dependency.UserAPI.Login)

	routerV1WithAuth := routerV1.Use()
	routerV1WithAuth.POST("/logout", dependency.MiddlewareValidateAuth, dependency.UserAPI.Logout)
	routerV1WithAuth.POST("/product", dependency.MiddlewareValidateAuth, dependency.ProductAPI.CreateProduct)
	routerV1WithAuth.GET("/product", dependency.MiddlewareValidateAuth, dependency.ProductAPI.GetAllProduct)
	routerV1WithAuth.GET("/product/:id", dependency.MiddlewareValidateAuth, dependency.ProductAPI.GetDetailProduct)
	routerV1WithAuth.GET("/profile", dependency.MiddlewareValidateAuth, dependency.UserAPI.GetUserProfile)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		helpers.Logger.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	Logger *logrus.Logger

	UserAPI    interfaces.IUserHandler
	ProductAPI interfaces.IProductHandler
}

func dependencyInject() Dependency {
	helpers.SetupLogger()

	userRepo := &userRepository.UserRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	productRepo := &productRepository.ProductRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	userSvc := &userServices.UserService{
		UserRepo: userRepo,
		Logger:   helpers.Logger,
		Redis:    helpers.RedisClient,
	}
	userAPI := &userAPI.UserHandler{
		UserService: userSvc,
		Validator:   validator,
	}

	productSvc := &productServices.ProductService{
		ProductRepo: productRepo,
		Logger:      helpers.Logger,
	}
	productAPI := &productAPI.ProductHandler{
		ProductService: productSvc,
		Validator:      validator,
	}

	return Dependency{
		Logger:     helpers.Logger,
		UserAPI:    userAPI,
		ProductAPI: productAPI,
	}
}
