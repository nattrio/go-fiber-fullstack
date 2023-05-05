package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-fiber-fullstack/database"
	"github.com/nattrio/go-fiber-fullstack/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Nattrio Trivia",
		"Subtitle": "List of facts",
		"Facts":    facts,
	})
}

// Create new Fact View handler
func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	result := database.DB.Db.First(&fact, c.Params("id"))
	if result.Error != nil {
		return c.Status(404).SendString("Fact not found")
	}

	return c.Render("show", fiber.Map{
		"Title":    "Fact",
		"Subtitle": "Fact details",
		"Fact":     fact,
	})
}
