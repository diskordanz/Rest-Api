package model

func (db *MockDB) GetAuthor(author *Author) error {
	rets := db.Called(author)
	return rets.Error(0)
}

func (db *MockDB) UpdateAuthor(author *Author) error {
	rets := db.Called(author)
	return rets.Error(0)
}

func (db *MockDB) DeleteAuthor(author *Author) error {
	rets := db.Called(author)
	return rets.Error(0)
}

func (db *MockDB) CreateAuthor(author *Author) error {
	rets := db.Called(author)
	return rets.Error(0)
}

func (db *MockDB) GetAuthors() ([]Author, error) {
	rets := db.Called()
	return rets.Get(0).([]Author), rets.Error(1)
}

func (db *MockDB) GetFilterAuthors(filterString string) ([]Author, error) {
	rets := db.Called()
	return rets.Get(0).([]Author), rets.Error(1)
}