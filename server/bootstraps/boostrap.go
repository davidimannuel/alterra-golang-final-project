package bootstraps

import (
	_authUc "keep-remind-app/businesses/auth"
	_noteUc "keep-remind-app/businesses/note"
	_ocrUc "keep-remind-app/businesses/ocr"
	_userUc "keep-remind-app/businesses/user"
	"keep-remind-app/configs"
	_noteRepo "keep-remind-app/repositories/note"
	_userRepo "keep-remind-app/repositories/user"
	_authHandler "keep-remind-app/server/handlers/auth"
	_noteHandler "keep-remind-app/server/handlers/note"
	_userHandler "keep-remind-app/server/handlers/user"

	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	App         *echo.Echo
	Configs     configs.Configs
	AuthHandler *_authHandler.AuthHandler
	UserHandler *_userHandler.UserHandler
	NoteHandler *_noteHandler.NoteHandler
}

func Init(app *echo.Echo, configs configs.Configs) *Bootstrap {
	//init middleware

	// factory repository
	userRepo := _userRepo.NewUserRepository(configs.DB)
	noteRepo := _noteRepo.NewNoteRepository(configs.DB)
	// factory usecase
	userUc := _userUc.NewUserUsecase(userRepo)
	ocrUc := _ocrUc.NewOCRUsecase()
	// configJwt := configs.JWT
	authUc := _authUc.NewAuthUsecase(userUc, &configs.JWT)
	noteUc := _noteUc.NewNoteUsecase(noteRepo, ocrUc)
	// boot
	boot := Bootstrap{
		App:         app,
		Configs:     configs,
		AuthHandler: _authHandler.NewAuthHandler(&configs, authUc),
		UserHandler: _userHandler.NewUserHandler(&configs, userUc),
		NoteHandler: _noteHandler.NewNoteHandler(&configs, noteUc),
	}
	return &boot
}
