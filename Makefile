mocks: 
	mockery --dir "./businesses/label" --name "LabelRepository" --output "./businesses/label/mocks"
	mockery --dir "./businesses/note" --name "NoteRepository" --output "./businesses/note/mocks"