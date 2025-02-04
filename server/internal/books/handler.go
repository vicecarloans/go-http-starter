package books

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type BookHandler struct {
	bookService BookService
}

func NewBookHandler(bookRoute fiber.Router, bookService BookService) *BookHandler {
	handler := &BookHandler{bookService: bookService}

	bookRoute.Get("", handler.getBooks)
	bookRoute.Get("/:id", handler.getBook)
	bookRoute.Post("", handler.createBook)
	bookRoute.Patch("/:id", handler.updateBook)	
	bookRoute.Delete("/:id", handler.deleteBook)
	return &BookHandler{bookService: bookService}
}


func (h *BookHandler) getBooks(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 10)

	books, err := h.bookService.GetBooks(c.Context(), offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data": books,
	})
}

func (h *BookHandler) getBook(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uuid.Validate(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid book id",
			"status": "fail",
		})
	}

	book, err := h.bookService.GetBook(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}

	if book == nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Book not found",
			"status": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data": book,
	})
}

func (h *BookHandler) createBook(c *fiber.Ctx) error {
	book := &CreateBookDTO{}

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
			"status": "fail",
		})
	}

	book.ID = uuid.New().String()

	err := h.bookService.CreateBook(c.Context(), book)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"data": book,
	})
}


func (h *BookHandler) updateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uuid.Validate(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid book id",
			"status": "fail",
		})
	}

	book := &UpdateBookDTO{
		ID: id,
	}

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
			"status": "fail",
		})
	}

	err := h.bookService.UpdateBook(c.Context(), book)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Book not found",
			"status": "fail",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data": book,
	})
}

func (h *BookHandler) deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uuid.Validate(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid book id",
			"status": "fail",
		})
	}

	err := h.bookService.DeleteBook(c.Context(), id)
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Book not found",
			"status": "fail",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "Book deleted successfully",
	})
}