package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

const blockWidth = 4
const maxImageWidth = 320

func main() {
	// Define two colors.
	colors := []color.RGBA{
		{128, 128, 128, 255},  // 0: Grey
		{211, 211, 211, 255},  // 1: Light Grey	
	}

	// Define the flags for encoding and decoding.
	encode := flag.Bool("e", false, "To encode")
	decode := flag.Bool("d", false, "To decode")
	flag.Parse()

	// The file name is the first unchecked argument.
	filename := flag.Arg(0)

	if *encode {
		// Read the number string from the standard input.
		reader := bufio.NewReader(os.Stdin)
		numbers, _ := reader.ReadString('\n')
		numbers = strings.TrimSpace(numbers) // Remove the newline character.

		// Calculate the number of rows needed.
		numRows := (len(numbers)*blockWidth + maxImageWidth - 1) / maxImageWidth

		// Create a new image.
		img := image.NewRGBA(image.Rect(0, 0, maxImageWidth, numRows*blockWidth))

		// Fill the image with the corresponding colors.
		for i := 0; i < len(numbers); i++ {
			// Convert the number to a color.
			n, _ := strconv.Atoi(string(numbers[i]))
			for x := (i % (maxImageWidth / blockWidth)) * blockWidth; x < (i%(maxImageWidth/blockWidth)+1)*blockWidth; x++ {
				for y := (i / (maxImageWidth / blockWidth)) * blockWidth; y < ((i/(maxImageWidth/blockWidth))+1)*blockWidth; y++ {
					img.Set(x, y, colors[n])
				}
			}
		}

		// Save the image as a .png file.
		f, _ := os.Create(filename)
		defer f.Close()
		png.Encode(f, img)
	} else if *decode {
		// Open the .png file.
		in, _ := os.Open(filename)
		defer in.Close()
		imgIn, _, _ := image.Decode(in)

		// Decode the colors back into numbers and print in a single line.
		for y := 0; y < imgIn.Bounds().Dy(); y += blockWidth {
			for x := 0; x < imgIn.Bounds().Dx(); x += blockWidth {
				c := color.RGBAModel.Convert(imgIn.At(x, y)).(color.RGBA)
				for j, col := range colors {
					if c.R == col.R && c.G == col.G && c.B == col.B {
						fmt.Print(j)
						break
					}
				}
			}
		}

		// Add a newline at the end to separate the prompt.
		fmt.Println()
	} else {
		fmt.Println("Please enter either -e for encoding or -d for decoding.")
	}
}

