package database

import (
	"fmt"

	"github.com/Trepka/bookslib/internal/config"
	"github.com/Trepka/bookslib/internal/logger"
	st "github.com/Trepka/bookslib/internal/structs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type BookStructure st.Book

type Library []BookStructure

type PostgressBooksStorage struct {
	db  *sqlx.DB
	log *logger.Logger
}

func ConnectDB(config config.DBConfig, logger *logger.Logger) PostgressBooksStorage {
	var storage PostgressBooksStorage
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", config.DbUser, config.DbName, config.DbPassword)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logger.Fatal().
			Str("package", "database").
			Str("func", "ConnectDB").
			Msg(err.Error())
	}
	storage.db = db
	storage.log = logger

	return storage
}

func (p *PostgressBooksStorage) GetAllBooks() (Library, error) {
	var l Library
	var b BookStructure
	rows, err := p.db.Queryx("SELECT * FROM books")
	if err != nil {
		p.log.Error().
			Str("package", "database").
			Str("func", "GetAllBooks").
			Msg(err.Error())
	}
	for rows.Next() {
		err = rows.StructScan(&b)
		if err != nil {
			p.log.Error().
				Str("package", "database").
				Str("func", "GetAllBooks next row").
				Msg(err.Error())
		}
		l = append(l, b)
	}
	p.log.Info().Timestamp().Msg("get all books from storage")
	return l, nil
}

func (p *PostgressBooksStorage) GetBook(id string) (BookStructure, error) {
	var b BookStructure
	row := p.db.QueryRowx("SELECT id, name, author, genre, year FROM books WHERE id = $1", id)
	err := row.Scan(&b.ID, &b.Name, &b.Author, &b.Genre, &b.Year)
	if err != nil {
		p.log.Error().
			Str("package", "database").
			Str("func", "GetBook").
			Msg(err.Error())
	}
	p.log.Info().Timestamp().Msg(b.String())
	return b, nil
}

func (p *PostgressBooksStorage) PutBook(b BookStructure) {
	tx := p.db.MustBegin()
	tx.MustExec("INSERT INTO books (id, name, author, genre, year) VALUES ($1, $2, $3, $4, $5)", b.ID, b.Name, b.Author, b.Genre, b.Year)
	p.log.Info().Timestamp().Msg(b.String())
	tx.Commit()
}

func (p *PostgressBooksStorage) DeleteBook(id string) {
	b, _ := p.GetBook(id)
	p.db.QueryRowx("DELETE FROM books WHERE id = $1", id)
	p.log.Info().Timestamp().Msg("delete book " + b.String())
}

func (b BookStructure) String() string {
	return fmt.Sprintf("ID:%s Name:%s Author:%s Genre:%s Year:%d", b.ID, b.Name, b.Author, b.Genre, b.Year)
}
