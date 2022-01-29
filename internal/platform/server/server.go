package server

import (
	"log"
	"net/http"

	"github.com/Trepka/bookslib/internal/config"

	"github.com/Trepka/bookslib/internal/platform/database"
	st "github.com/Trepka/bookslib/internal/structs"
	"github.com/emicklei/go-restful"
)

type Storage struct {
	BookRepository database.PostgressBooksStorage
}

func WebService(s Storage) *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/").To(s.findAllBooks).
		Writes([]st.Book{}).
		Returns(200, "OK", []st.Book{}))

	ws.Route(ws.GET("/{book-id}").To(s.findBook).
		Param(ws.PathParameter("book-id", "identifier of the book").DataType("integer").DefaultValue("1")).
		Writes(st.Book{}).
		Returns(200, "OK", st.Book{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{book-id}").To(s.updateBook).
		Param(ws.PathParameter("book-id", "identifier of the book").DataType("string")).
		Reads(st.Book{}))

	ws.Route(ws.POST("").To(s.createBook).
		Reads(st.Book{}))

	ws.Route(ws.DELETE("/{book-id}").To(s.removeBook).
		Param(ws.PathParameter("book-id", "identifier of the book").DataType("string")))

	return ws
}

func (s Storage) findAllBooks(request *restful.Request, response *restful.Response) {
	list, err := s.BookRepository.GetAllBooks()
	if err != nil {
		log.Fatal(err)
	}
	response.WriteEntity(list)
}

func (s Storage) findBook(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("book-id")
	b, err := s.BookRepository.GetBook(id)
	if err != nil {
		log.Fatal(err)
	}
	if len(string(b.ID)) == 0 {
		response.WriteErrorString(http.StatusNotFound, "Book could not be found.")
	} else {
		response.WriteEntity(b)
	}
}

func (s *Storage) updateBook(request *restful.Request, response *restful.Response) {
	b := new(st.Book)
	err := request.ReadEntity(&b)
	if err == nil {
		s.BookRepository.UpdateBook(b.ID, b.Name)
		response.WriteEntity(b)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (s *Storage) createBook(request *restful.Request, response *restful.Response) {
	b := st.Book{ID: request.PathParameter("book-id")}
	err := request.ReadEntity(&b)
	if err == nil {
		s.BookRepository.PutBook(database.BookStructure(b))
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (s *Storage) removeBook(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("book-id")
	s.BookRepository.DeleteBook(id)
}

func RunServer(cfg config.DBConfig) {
	webStorage := Storage{BookRepository: database.ConnectDB(cfg)}
	restful.DefaultContainer.Add(WebService(webStorage))

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
