package bootstraps

import (
	_authUc "keep-remind-app/businesses/auth"
	_labelUc "keep-remind-app/businesses/label"
	_noteUc "keep-remind-app/businesses/note"
	_ocrUc "keep-remind-app/businesses/ocr"
	_userUc "keep-remind-app/businesses/user"
	"keep-remind-app/configs"
	_labelRepo "keep-remind-app/repositories/label"
	_noteRepo "keep-remind-app/repositories/note"
	_userRepo "keep-remind-app/repositories/user"
	_authHandler "keep-remind-app/server/handlers/auth"
	_labelHandler "keep-remind-app/server/handlers/label"
	_noteHandler "keep-remind-app/server/handlers/note"
	_userHandler "keep-remind-app/server/handlers/user"

	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	App          *echo.Echo
	Configs      configs.Configs
	AuthHandler  *_authHandler.AuthHandler
	UserHandler  *_userHandler.UserHandler
	LabelHandler *_labelHandler.LabelHandler
	NoteHandler  *_noteHandler.NoteHandler
}

func Init(app *echo.Echo, configs configs.Configs) *Bootstrap {
	// factory repository
	userRepo := _userRepo.NewUserRepository(configs.DB)
	noteRepo := _noteRepo.NewNoteRepository(configs.DB)
	labelRepo := _labelRepo.NewLabelRepository(configs.DB)
	// factory usecase
	userUc := _userUc.NewUserUsecase(userRepo)
	ocrUc := _ocrUc.NewOCRUsecase()
	authUc := _authUc.NewAuthUsecase(userUc, &configs.JWT)
	labelUc := _labelUc.NewLabelUsecase(labelRepo)
	noteUc := _noteUc.NewNoteUsecase(noteRepo, ocrUc)
	// boot
	boot := Bootstrap{
		App:          app,
		Configs:      configs,
		AuthHandler:  _authHandler.NewAuthHandler(&configs, authUc),
		UserHandler:  _userHandler.NewUserHandler(userUc),
		LabelHandler: _labelHandler.NewLabelHandler(labelUc),
		NoteHandler:  _noteHandler.NewNoteHandler(&configs, noteUc),
	}
	return &boot
}
