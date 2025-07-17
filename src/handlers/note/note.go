package noteHandler

import (
	"go-fiber-server/database"
	"go-fiber-server/src/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	note.ID = uuid.New()

	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Note created", "data": note})
}

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Notes not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Note found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"text"`
	}

	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": nil})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Note updated", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": nil})
	}

	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Note deleted"})
}
