package model

func (db *GormDB) GetBooks() ([]Book, error) {
	var books []Book
	if err := db.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func (db *GormDB) GetBooksByAuthor(books *Book) ([]Book, error) {

	if err := db.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (db *GormDB) GetBook(book *Book) error {
	return db.DB.Find(&book).Error
}

func (db *GormDB) UpdateBook(book *Book) error {
	return db.DB.Save(&book).Error
}

func (db *GormDB) DeleteBook(book *Book) error {
	return db.DB.Delete(&book).Error
}

func (db *GormDB) CreateBook(book *Book) error {
	if err := db.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) GetFilterBooks(filterString string) error{
	var books []Book
	if err := db.Where("name LIKE ?", filterString).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (db *GormDB) GetFilterBooksByAuthor(idAuthor int, filterString string) ([]Book, error){
	var books []Book
	if err := db.Where("name LIKE ? AND author_id = ?", name, idAuthor).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}