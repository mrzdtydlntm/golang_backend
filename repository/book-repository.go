package repository

import (
	"github.com/mrzdtydlntm/golang_backend/entity"
	"gorm.io/gorm"
)

// BookRepository
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	AllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

// NewBookRepository creates an instance BookRepository
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)                 //menyimpan data buku ke dalam db
	db.connection.Preload("User").Find(&b) //untuk mendapatkan object user dari book entity, harus di preload terlebih dahulu
	return b
}

func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books)
	return books
}
