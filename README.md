# Bookslib

Simple RestAPI app allowing to manage library of books. Supports get, post, delete queries.

## Features

* **/POST/books** - add book to the storage. Example of json book format see below

```json
 {
    "book-id": "3",
    "name": "Harry Potter and the Philosopher's Stone",
    "author": "J.K. Rowling",
    "genre": "fantasy",
    "year": 1996
 }
```

* **/GET/books** - return all books instances from library storage as slice of books

```json
[
    {
        "book-id": "3",
        "name": "Harry Potter and the Philosopher's Stone",
        "author": "J.K. Rowling",
        "genre": "fantasy",
        "year": 1997
    },
    {
        "book-id": "1",
        "name": "The Hitchhiker's Guide to the Galaxy",
        "author": "Douglas Adams",
        "genre": "Comic science fiction",
        "year": 1979
    }
]
```

* **/GET/books/{id}** - return same json book structure as posted from storage by ID.
* **/DELETE/books/{id}** - delete book from the storage by ID.

## Used in project

* router based on [go-restful](https://github.com/emicklei/go-restful) library.
* [sqlx](https://github.com/jmoiron/sqlx) library for connect to database.
* logger - [zerolog](https://github.com/rs/zerolog).
* database - PostgreSQL.