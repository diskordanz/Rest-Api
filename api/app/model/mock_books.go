package model

func (db *MockDB) GetBook(book *Book) error {
	rets := db.Called(book)
	return rets.Error(0)
}
func (db *MockDB) CreateBook(book *Book) error {
	rets := db.Called(book)
	return rets.Error(0)
}

func (db *MockDB) UpdateBook(book *Book) error {
	rets := db.Called(book)
	return rets.Error(0)
}

func (db *MockDB) DeleteBook(book *Book) error {
	rets := db.Called(book)
	return rets.Error(0)
}

func (db *MockDB) GetBooks() ([]Book, error) {
	rets := db.Called()
	return rets.Get(0).([]Book), rets.Error(1)
}

func (db *MockDB) GetFilterBooks(filterString string) ([]Book, error) {
	rets := db.Called()
	return rets.Get(0).([]Book), rets.Error(1)
}
func (db *MockDB) GetBooksByAuthor(author *Author) ([]Book, error) {
	rets := db.Called()
	return rets.Get(0).([]Book), rets.Error(1)
}

func (db *MockDB) GetFilterBooksByAuthor(idAuthor int, filterString string) ([]Book, error) {
	rets := db.Called()
	return rets.Get(0).([]Book), rets.Error(1)
}
// func (db *MockDB) GetBookByAuthor(book *Book) error {
// 	rets := db.Called(book)
// 	return rets.Error(0)
// }


