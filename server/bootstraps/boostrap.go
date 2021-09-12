package bootstraps

import (
	_authUc "keep-remind-app/businesses/auth"
	_jwtUc "keep-remind-app/businesses/jwt"
	_userUc "keep-remind-app/businesses/user"
	"keep-remind-app/configs"
	_userRepo "keep-remind-app/drivers/repositories/user"
	_authHandler "keep-remind-app/server/handlers/auth"
	_userHandler "keep-remind-app/server/handlers/user"

	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	App         *echo.Echo
	Configs     *configs.Configs
	AuthHandler *_authHandler.Handler
	UserHandler *_userHandler.Handler
}

func (boot *Bootstrap) Init() {
	// factory repository
	userRepo := _userRepo.NewPostgreSQLRepository(boot.Configs.DB)
	// factory usecase
	userUc := _userUc.NewUsecase(userRepo)
	jwtUc := _jwtUc.NewUsecase(boot.Configs.JWT)
	authUc := _authUc.NewUsecase(userUc, jwtUc)
	// init handler
	boot.UserHandler = _userHandler.NewHandler(boot.Configs, userUc)
	boot.AuthHandler = _authHandler.NewHandler(boot.Configs, authUc)
}
