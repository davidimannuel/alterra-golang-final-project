mocks: 
	mockery --dir "./businesses/label" --name "LabelRepository" --output "./businesses/label/mocks"
	mockery --dir "./businesses/note" --name "NoteRepository" --output "./businesses/note/mocks"
	mockery --dir "./businesses/user" --name "UserRepository" --output "./businesses/user/mocks"
	mockery --dir "./businesses/redis" --name "RedisRepository" --output "./businesses/redis/mocks"
	mockery --dir "./businesses/telegramUser" --name "TelegramUserRepository" --output "./businesses/telegramUser/mocks"