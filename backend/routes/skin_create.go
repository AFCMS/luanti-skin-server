package routes

import (
	"bytes"
	"image/png"
	"luanti-skin-server/backend/database"
	"luanti-skin-server/backend/models"
	"luanti-skin-server/backend/utils"
	"luanti-skin-server/common/oxipng"
	"luanti-skin-server/common/skins"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v3"
)

// SkinCreate Handle Skin creation
//
// Use a multipart request
func SkinCreate(c fiber.Ctx) error {
	// Get User
	user := c.Locals("user").(models.Account)

	// input := new(types.InputSkinCreate)

	// Get the text fields
	var form *multipart.Form
	var err error
	if form, err = c.MultipartForm(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Invalid request body", Data: err.Error()})
	}

	// Get file part
	var skinB []byte
	if skinB, err = utils.LoadFormFile(c, "data"); err != nil {
		return err
	}

	// Decode image
	img, err := png.Decode(bytes.NewReader(skinB))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Cannot decode skin", Data: err.Error()})
	}

	// Validate image size
	bounds := img.Bounds()

	if bounds.Max.X != 64 || bounds.Max.Y != 32 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Invalid skin", Data: "Image have invalid size (64x32 expected)"})
	}

	// Extract head
	var headBuffer bytes.Buffer

	rgbaImg, err := skins.ImageToRGBA(img)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorOutput{Message: "Invalid image type", Data: "Image is not of type RGBA"})
	}

	headImg := skins.SkinExtractHead(&rgbaImg)
	err = png.Encode(&headBuffer, headImg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Server error", Data: "Cannot extract head from image"})
	}
	headB := headBuffer.Bytes()

	// Run Oxipng
	skinBOpti, err1 := oxipng.OxipngBytes(skinB)
	headBOpti, err2 := oxipng.OxipngBytes(headB)

	if err1 != nil || err2 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Server error", Data: "Cannot obtimize image"})
	}

	// Create entry in database
	var l = models.Skin{
		Description: form.Value["description"][0],
		Owner:       user,
		Data:        skinBOpti,
		DataHead:    headBOpti,
		CreatedAt:   time.Now(),
	}

	if err := database.DB.Create(&l).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Cannot interact with database", Data: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(l)
}
