package oxipng

import (
	"os"
	"testing"
)

func TestOxipngWorks(t *testing.T) {
	if !OxipngPresent() {
		t.Error("Oxipng is not present on the system")
	}

	// Read png file in this directory
	data, err := os.ReadFile("oxipng_test_image.png")
	if err != nil {
		t.Error("Error on read file")
	}

	_, err = OxipngBytes(data)

	if err != nil {
		t.Error("Error on oxipng")
	}
}
