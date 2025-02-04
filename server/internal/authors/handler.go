package authors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthorHandler struct {
	authorService AuthorService
}

func NewAuthorHandler(authorRoute fiber.Router, authorService AuthorService) *AuthorHandler {
	handler := &AuthorHandler{authorService: authorService}
	
	authorRoute.Get("", handler.getAuthors)
	authorRoute.Get("/:id", handler.getAuthor)
	authorRoute.Post("", handler.createAuthor)
	authorRoute.Patch("/:id", handler.updateAuthor)
	authorRoute.Delete("/:id", handler.deleteAuthor)

	return &AuthorHandler{authorService: authorService}
}

func (h *AuthorHandler) getAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uuid.Validate(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid author id",
			"status": "fail",
		})
	}

	author, err := h.authorService.GetAuthor(c.Context(), id)
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Author not found",
			"status": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data": author,
	})
}

func (h *AuthorHandler) getAuthors(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 10)

	authors, err := h.authorService.GetAuthors(c.Context(), offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}

	if authors == nil {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"data": []*GetManyAuthorsDTO{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data": authors,
	})
}

func (h *AuthorHandler) createAuthor(c *fiber.Ctx) error {
	author := &CreateAuthorDTO{}

	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
			"status": "fail",
		})
	}

	err := h.authorService.CreateAuthor(c.Context(), author)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"status": "fail",
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"data": author,
	})
}

func (h *AuthorHandler) updateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uuid.Validate(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid author id",
			"status": "fail",
		})
	}
	
	author := &UpdateAuthorDTO{
		ID: id,
	}

	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
			"status": "fail",
		})
	}

	err := h.authorService.UpdateAuthor(c.Context(), author)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Author not found",
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
		"data": author,
	})
}

func (h *AuthorHandler) deleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uuid.Validate(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid author id",
			"status": "fail",
		})
	}
	
	err := h.authorService.DeleteAuthor(c.Context(), id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Author not found",
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
		"message": "Author deleted successfully",
	})
}