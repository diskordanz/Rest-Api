package model

func (db *GormDB) GetAuthors() ([]Author, error) {
	var authors []Author
	if err := db.DB.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (db *GormDB) GetAuthor(author *Author) error {
	return db.DB.Find(&author).Error
}

func (db *GormDB) UpdateAuthor(author *Author) error {
	return db.DB.Save(&author).Error
}

func (db *GormDB) DeleteAuthor(author *Author) error {
	return db.DB.Delete(&author).Error
}

func (db *GormDB) CreateAuthor(author *Author) error {
	if err := db.DB.Create(&author).Error; err != nil {
		return err
	}
	return nil
}
func (db *GormDB) GetFilterAuthors(filterString string) error{
	var authors []Author
	if err := db.Where("name LIKE ?", filterString).Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}