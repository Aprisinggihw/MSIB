package builder

import (
	"tugas-6/configs"
	"tugas-6/internal/http/handler"
	"tugas-6/internal/http/router"
	"tugas-6/internal/repository"
	"tugas-6/internal/service"
	"tugas-6/pkg/cache"
	"tugas-6/pkg/route"
	"tugas-6/pkg/token"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *configs.Config, db *gorm.DB, rdb *redis.Client) []route.Route {
	cacheable := cache.NewCacheable(rdb)
	userRepository := repository.NewUserRepository(db)
	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)
	userService := service.NewUserService(userRepository, tokenUseCase, cacheable)
	userHandler := handler.NewUserHandler(userService)
	return router.PublicRoutes(userHandler)
}

func BuildPrivateRoutes(cfg *configs.Config, db *gorm.DB, rdb *redis.Client) []route.Route {
	cacheable := cache.NewCacheable(rdb)
	userRepository := repository.NewUserRepository(db)
	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)
	userService := service.NewUserService(userRepository, tokenUseCase, cacheable)
	userHandler := handler.NewUserHandler(userService)
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository, tokenUseCase, cacheable)
	todoHandler := handler.NewTodoHandler(todoService)
	return router.PrivateRoutes(userHandler,*todoHandler)
}
