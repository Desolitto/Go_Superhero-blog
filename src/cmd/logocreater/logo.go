package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"superhero-blog/pkg"
)

func main() {
	width := 300
	height := 300

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	// Прозрачны фон
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)
	// Рисуем красный прямоугольник
	pkg.DrawRoundedRect(img, 27, 27, 273, 273, 50, pkg.Colors_1.Aqua) // 50 - радиус закругления
	pkg.DrawRoundedRect(img, 30, 30, 270, 270, 50, pkg.Colors_1.Red)  // 50 - радиус закругления

	pkg.DrawLines(img)
	// Добавление обрамления по пиеселю к цифрам
	pkg.AddBorder(img, "01", 270, 30)

	// Создание файла в корне
	err := saveImage("amazing_logo.png", img)
	if err != nil {
		panic(err)
	}

	// Создание файла в папке images
	err = saveImage("images/amazing_logo.png", img)
	if err != nil {
		panic(err)
	}
}

// saveImage сохраняет изображение в указанный файл
func saveImage(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
