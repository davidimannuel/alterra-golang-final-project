package bootstraps

import (
	_authUc "keep-remind-app/businesses/auth"
	_noteUc "keep-remind-app/businesses/note"
	_userUc "keep-remind-app/businesses/user"
	"keep-remind-app/configs"
	_noteRepo "keep-remind-app/drivers/repositories/note"
	_userRepo "keep-remind-app/drivers/repositories/user"
	_authHandler "keep-remind-app/server/handlers/auth"
	_noteHandler "keep-remind-app/server/handlers/note"
	_userHandler "keep-remind-app/server/handlers/user"

	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	App         *echo.Echo
	Configs     *configs.Configs
	AuthHandler *_authHandler.Handler
	UserHandler *_userHandler.Handler
	NoteHandler *_noteHandler.Handler
}

func Init(app *echo.Echo, configs *configs.Configs) *Bootstrap {
	//init middleware

	// factory repository
	userRepo := _userRepo.NewPostgreSQLRepository(configs.DB)
	noteRepo := _noteRepo.NewPostgreSQLRepository(configs.DB)
	// factory usecase
	userUc := _userUc.NewUsecase(userRepo)
	authUc := _authUc.NewUsecase(userUc, configs.JWT)
	noteUc := _noteUc.NewUsecase(noteRepo)
	// boot
	boot := Bootstrap{
		App:         app,
		Configs:     configs,
		AuthHandler: _authHandler.NewHandler(configs, authUc),
		UserHandler: _userHandler.NewHandler(configs, userUc),
		NoteHandler: _noteHandler.NewHandler(configs, noteUc),
	}
	return &boot
}
