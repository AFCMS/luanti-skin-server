package main

import (
	"context"
	"image"
	"image/png"
	"log"
	"luanti-skin-server/utils"
	"os"

	"github.com/urfave/cli/v3"
)

func LoadImageFromFile(src string) (*image.RGBA, error) {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	rgbImage, err := utils.ImageToRGBA(img)
	if err != nil {
		return nil, err
	}

	return &rgbImage, nil
}

func main() {
	cmd := &cli.Command{
		Name:  "luanti-skin-converter",
		Usage: "A CLI tool for converting and optimising Minecraft/Luanti skins",
		Commands: []*cli.Command{
			{
				Name:  "luanti",
				Usage: "Luanti skin commands",
				Commands: []*cli.Command{
					{
						Name:  "head",
						Usage: "Extract the head from a skin",
						Arguments: []cli.Argument{
							&cli.StringArg{
								Name: "src",
							},
							&cli.StringArg{
								Name: "out",
							},
						},
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "no-oxipng",
								Usage: "Do not use oxipng to optimise the final image",
							},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							img, err := LoadImageFromFile(cmd.StringArg("src"))
							if err != nil {
								log.Fatal(err)
							}

							head := utils.SkinExtractHead(img)

							outFile, err := os.Create(cmd.StringArg("out"))
							if err != nil {
								log.Fatal(err)
							}
							defer outFile.Close()

							err = png.Encode(outFile, head)
							if err != nil {
								log.Fatal(err)
							}

							log.Printf("Head extracted to %s\n", cmd.StringArg("out"))

							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
