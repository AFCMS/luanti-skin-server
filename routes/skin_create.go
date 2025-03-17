package routes

import (
	"bytes"
	"image/png"
	"luanti-skin-server/database"
	"luanti-skin-server/models"
	"luanti-skin-server/utils"
	"mime/multipart"
	"sync"
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

	headImg := utils.SkinExtractHead(img)
	err = png.Encode(&headBuffer, headImg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Server error", Data: "Cannot extract head from image"})
	}
	headB := headBuffer.Bytes()

	var skinBOpti = skinB
	var headBOpti = headB

	// Optionally run Oxipng
	// TODO: Run them async to get them done faster
	// https://stackoverflow.com/questions/27792389/golang-functions-parallel-execution-with-return

	var out1 []byte
	var err1 error

	var out2 []byte
	var err2 error

	var sg sync.WaitGroup

	sg.Add(2)

	go func(out *[]byte, err *error, sg *sync.WaitGroup) {
		*out, *err = utils.OxipngBytes(skinB)
		sg.Done()
	}(&out1, &err1, &sg)

	go func(out *[]byte, err *error, sg *sync.WaitGroup) {
		*out, *err = utils.OxipngBytes(headB)
		sg.Done()
	}(&out2, &err2, &sg)

	sg.Wait()

	if err1 != nil || err2 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorOutput{Message: "Server error", Data: "Cannot obtimize image"})
	}

	skinBOpti = out1
	headBOpti = out2

	// Create entry in database
	var l = models.Skin{
		Description: form.Value["description"][0],
		Public:      form.Value["public"][0] == "true",
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
