package database

import (
	"fmt"
	"log"

	"github.com/Trepka/bookslib/internal/config"
	st "github.com/Trepka/bookslib/internal/structs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type BookStructure st.Book

type Library []BookStructure

type PostgressBooksStorage struct {
	db *sqlx.DB
	// log *logger.Logger
}

func ConnectDB(config config.DBConfig) PostgressBooksStorage {
	var storage PostgressBooksStorage
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", config.DbUser, config.DbName, config.DbPassword)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	storage.db = db
	// defer db.Close()

	return storage
}

func (p *PostgressBooksStorage) GetAllBooks() (Library, error) {
	// сделать запрос и каждую строку сохранить в структуру Book
	var l Library
	var b BookStructure
	rows, err := p.db.Queryx("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err = rows.StructScan(&b)
		if err != nil {
			log.Fatalln(err)
		}
		l = append(l, b)
	}
	return l, nil
}

func (p *PostgressBooksStorage) GetBook(id string) (BookStructure, error) {
	var b BookStructure
	row := p.db.QueryRowx("SELECT id, name, author, genre, year FROM books WHERE id = $1", id)
	err := row.Scan(&b.ID, &b.Name, &b.Author, &b.Genre, &b.Year)
	if err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (p *PostgressBooksStorage) UpdateBook(id string, name string) {
	p.db.QueryRowx("UPDATE books SET name=$1 WHERE id = $2", name, id)
}

func (p *PostgressBooksStorage) PutBook(b BookStructure) {
	tx := p.db.MustBegin()
	tx.MustExec("INSERT INTO books (id, name, author, genre, year) VALUES ($1, $2, $3, $4, $5)", b.ID, b.Name, b.Author, b.Genre, b.Year)
	tx.Commit()
}

func (p *PostgressBooksStorage) DeleteBook(id string) {
	p.db.QueryRowx("DELETE FROM books WHERE id = $1", id)
}
