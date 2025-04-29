package oxipng

import (
	"bytes"
	"os/exec"
)

// OxipngPresent Return true if Oxipng is present on the system
//
// Try to execute `oxipng --version`
func OxipngPresent() bool {
	c := exec.Command("oxipng", "--version")
	c.Stdout = nil // /dev/null

	return c.Run() == nil
}

// OxipngBytes Run Oxipng on a byte array
func OxipngBytes(input []byte) (output []byte, err error) {
	cmd := exec.Command("oxipng", "-", "--strip", "all", "--opt", "max", "--zopfli", "--stdout")

	cmd.Stdin = bytes.NewReader(input)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
