package user

import "gorm.io/gorm"

type Repository interface {
	// punya method simpan ke DB dengan return User & err
	Save(user User) (User, error)
}

// perlu untuk meengakses koneksi dengan tipe private
type repository struct {
	db *gorm.DB
}

// buat fungsi newRepo untuk akses koneksi repo
func NewRepository(db *gorm.DB) *repository {
	// akhirnya kita akses private db
	return &repository{db}
}

// kita panggil fungsi Save()
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
