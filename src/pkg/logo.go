package pkg

import (
	"image"
	"image/color"
	"image/draw"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Colors структура для хранения цветов
type Colors struct {
	Red   color.RGBA
	Aqua  color.RGBA
	White color.RGBA
}

var Colors_1 = Colors{
	Red:   color.RGBA{255, 0, 0, 255},     // Красный цвет
	Aqua:  color.RGBA{0, 255, 255, 255},   // Аквамариновый цвет
	White: color.RGBA{255, 255, 255, 255}, // Белый цвет
}

func DrawRoundedRect(img *image.NRGBA, x0, y0, x1, y1, radius int, col color.Color) {
	// Рисуем прямоугольник
	draw.Draw(img, image.Rect(x0+radius, y0, x1-radius, y1), &image.Uniform{col}, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(x0, y0+radius, x1, y1-radius), &image.Uniform{col}, image.Point{}, draw.Src)

	// Рисуем закруглённые углы
	drawCircle(img, x0+radius, y0+radius, radius, col) // Верхний левый
	drawCircle(img, x1-radius, y0+radius, radius, col) // Верхний правый
	drawCircle(img, x0+radius, y1-radius, radius, col) // Нижний левый
	drawCircle(img, x1-radius, y1-radius, radius, col) // Нижний правый
}

func drawCircle(img *image.NRGBA, cx, cy, radius int, col color.Color) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				img.Set(cx+x, cy+y, col)
			}
		}
	}
}

// Функция для добавления текста на изображение
func addLabel(img *image.NRGBA, label string, x, y int, col color.Color) {
	point := fixed.P(x, y)

	// Используем базовый шрифт
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	d.DrawString(label)
}

// Функция для добавления обрамления к цифрам
func AddBorder(img *image.NRGBA, label string, x, y int) {
	// Добавляем обрамление
	addLabel(img, label, x-1, y-1, Colors_1.Aqua) // Верхний левый
	addLabel(img, label, x+1, y-1, Colors_1.Aqua) // Верхний правый
	addLabel(img, label, x-1, y+1, Colors_1.Aqua) // Нижний левый
	addLabel(img, label, x+1, y+1, Colors_1.Aqua) // Нижний правый

	// Добавляем основной текст
	addLabel(img, label, x, y, Colors_1.Red) // Основной текст
}

func drawLine(img *image.NRGBA, x1, x2, y int, col color.Color) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	for x := x1; x <= x2; x++ {
		img.Set(x, y, col)
	}
}

func DrawLines(img *image.NRGBA) {
	drawLine(img, 138, 165, 60, Colors_1.Aqua)
	drawLine(img, 130, 172, 61, Colors_1.Aqua)
	drawLine(img, 122, 179, 62, Colors_1.Aqua)
	drawLine(img, 120, 182, 63, Colors_1.Aqua)
	drawLine(img, 116, 186, 64, Colors_1.Aqua)
	drawLine(img, 113, 190, 65, Colors_1.Aqua)
	drawLine(img, 110, 193, 66, Colors_1.Aqua)
	drawLine(img, 107, 196, 67, Colors_1.Aqua)
	drawLine(img, 104, 199, 68, Colors_1.Aqua)
	drawLine(img, 102, 201, 69, Colors_1.Aqua)
	drawLine(img, 100, 203, 70, Colors_1.Aqua)
	drawLine(img, 98, 205, 71, Colors_1.Aqua)
	drawLine(img, 96, 207, 72, Colors_1.Aqua)
	drawLine(img, 95, 209, 73, Colors_1.Aqua)
	drawLine(img, 93, 211, 74, Colors_1.Aqua)

	drawLine(img, 91, 123, 75, Colors_1.Aqua)
	drawLine(img, 180, 212, 75, Colors_1.Aqua)
	drawLine(img, 125, 178, 75, Colors_1.Aqua)

	drawLine(img, 90, 119, 76, Colors_1.Aqua)
	drawLine(img, 184, 213, 76, Colors_1.Aqua)
	drawLine(img, 126, 177, 76, Colors_1.Aqua)

	drawLine(img, 89, 116, 77, Colors_1.Aqua)
	drawLine(img, 187, 214, 77, Colors_1.Aqua)
	drawLine(img, 126, 177, 77, Colors_1.Aqua)

	drawLine(img, 87, 113, 78, Colors_1.Aqua)
	drawLine(img, 190, 216, 78, Colors_1.Aqua)
	drawLine(img, 126, 177, 78, Colors_1.Aqua)

	drawLine(img, 86, 110, 79, Colors_1.Aqua)
	drawLine(img, 193, 217, 79, Colors_1.Aqua)
	drawLine(img, 126, 177, 79, Colors_1.Aqua)

	drawLine(img, 85, 107, 80, Colors_1.Aqua)
	drawLine(img, 196, 218, 80, Colors_1.Aqua)
	drawLine(img, 127, 176, 80, Colors_1.Aqua)

	drawLine(img, 84, 105, 81, Colors_1.Aqua)
	drawLine(img, 198, 219, 81, Colors_1.Aqua)
	drawLine(img, 127, 176, 81, Colors_1.Aqua)

	drawLine(img, 83, 103, 82, Colors_1.Aqua)
	drawLine(img, 200, 220, 82, Colors_1.Aqua)
	drawLine(img, 127, 176, 82, Colors_1.Aqua)

	drawLine(img, 83, 101, 83, Colors_1.Aqua)
	drawLine(img, 202, 220, 83, Colors_1.Aqua)
	drawLine(img, 127, 176, 83, Colors_1.Aqua)

	drawLine(img, 82, 99, 84, Colors_1.Aqua)
	drawLine(img, 204, 221, 84, Colors_1.Aqua)
	drawLine(img, 128, 175, 84, Colors_1.Aqua)

	drawLine(img, 81, 99, 85, Colors_1.Aqua)
	drawLine(img, 204, 222, 85, Colors_1.Aqua)
	drawLine(img, 128, 175, 85, Colors_1.Aqua)

	drawLine(img, 81, 97, 86, Colors_1.Aqua)
	drawLine(img, 206, 222, 86, Colors_1.Aqua)
	drawLine(img, 128, 175, 86, Colors_1.Aqua)

	drawLine(img, 80, 95, 87, Colors_1.Aqua)
	drawLine(img, 208, 223, 87, Colors_1.Aqua)
	drawLine(img, 129, 174, 87, Colors_1.Aqua)

	drawLine(img, 80, 94, 88, Colors_1.Aqua)
	drawLine(img, 209, 223, 88, Colors_1.Aqua)
	drawLine(img, 129, 174, 88, Colors_1.Aqua)

	drawLine(img, 79, 93, 89, Colors_1.Aqua)
	drawLine(img, 210, 224, 89, Colors_1.Aqua)
	drawLine(img, 129, 174, 89, Colors_1.Aqua)

	drawLine(img, 79, 92, 90, Colors_1.Aqua)
	drawLine(img, 211, 224, 90, Colors_1.Aqua)
	drawLine(img, 130, 173, 90, Colors_1.Aqua)

	drawLine(img, 79, 91, 91, Colors_1.Aqua)
	drawLine(img, 212, 224, 91, Colors_1.Aqua)
	drawLine(img, 130, 173, 91, Colors_1.Aqua)

	drawLine(img, 79, 91, 92, Colors_1.Aqua)
	drawLine(img, 212, 224, 92, Colors_1.Aqua)
	drawLine(img, 130, 173, 92, Colors_1.Aqua)

	drawLine(img, 79, 91, 93, Colors_1.Aqua)
	drawLine(img, 212, 224, 93, Colors_1.Aqua)
	drawLine(img, 131, 172, 93, Colors_1.Aqua)

	drawLine(img, 79, 91, 94, Colors_1.Aqua)
	drawLine(img, 212, 224, 94, Colors_1.Aqua)
	drawLine(img, 131, 172, 94, Colors_1.Aqua)

	drawLine(img, 79, 91, 95, Colors_1.Aqua)
	drawLine(img, 212, 224, 95, Colors_1.Aqua)
	drawLine(img, 131, 172, 95, Colors_1.Aqua)

	drawLine(img, 79, 91, 96, Colors_1.Aqua)
	drawLine(img, 212, 224, 96, Colors_1.Aqua)
	drawLine(img, 132, 171, 96, Colors_1.Aqua)

	drawLine(img, 79, 91, 97, Colors_1.Aqua)
	drawLine(img, 212, 224, 97, Colors_1.Aqua)
	drawLine(img, 132, 171, 97, Colors_1.Aqua)

	drawLine(img, 79, 91, 98, Colors_1.Aqua)
	drawLine(img, 212, 224, 98, Colors_1.Aqua)
	drawLine(img, 132, 171, 98, Colors_1.Aqua)

	drawLine(img, 79, 91, 99, Colors_1.Aqua)
	drawLine(img, 212, 224, 99, Colors_1.Aqua)
	drawLine(img, 133, 170, 99, Colors_1.Aqua)

	drawLine(img, 79, 91, 100, Colors_1.Aqua)
	drawLine(img, 212, 224, 100, Colors_1.Aqua)
	drawLine(img, 133, 170, 100, Colors_1.Aqua)

	drawLine(img, 79, 91, 101, Colors_1.Aqua)
	drawLine(img, 212, 224, 101, Colors_1.Aqua)
	drawLine(img, 133, 170, 101, Colors_1.Aqua)

	drawLine(img, 79, 91, 102, Colors_1.Aqua)
	drawLine(img, 212, 224, 102, Colors_1.Aqua)
	drawLine(img, 134, 169, 102, Colors_1.Aqua)

	drawLine(img, 79, 91, 103, Colors_1.Aqua)
	drawLine(img, 212, 224, 103, Colors_1.Aqua)
	drawLine(img, 134, 169, 103, Colors_1.Aqua)

	drawLine(img, 79, 91, 104, Colors_1.Aqua)
	drawLine(img, 212, 224, 104, Colors_1.Aqua)
	drawLine(img, 135, 168, 104, Colors_1.Aqua)

	drawLine(img, 79, 91, 105, Colors_1.Aqua)
	drawLine(img, 212, 224, 105, Colors_1.Aqua)
	drawLine(img, 135, 168, 105, Colors_1.Aqua)

	drawLine(img, 79, 91, 106, Colors_1.Aqua)
	drawLine(img, 212, 224, 106, Colors_1.Aqua)
	drawLine(img, 136, 167, 106, Colors_1.Aqua)

	drawLine(img, 79, 91, 107, Colors_1.Aqua)
	drawLine(img, 212, 224, 107, Colors_1.Aqua)
	drawLine(img, 136, 167, 107, Colors_1.Aqua)

	drawLine(img, 79, 92, 108, Colors_1.Aqua)
	drawLine(img, 211, 224, 108, Colors_1.Aqua)
	drawLine(img, 137, 166, 108, Colors_1.Aqua)

	drawLine(img, 79, 92, 109, Colors_1.Aqua)
	drawLine(img, 211, 224, 109, Colors_1.Aqua)
	drawLine(img, 138, 165, 109, Colors_1.Aqua)

	drawLine(img, 79, 92, 110, Colors_1.Aqua)
	drawLine(img, 211, 224, 110, Colors_1.Aqua)

	drawLine(img, 79, 92, 111, Colors_1.Aqua)
	drawLine(img, 211, 224, 111, Colors_1.Aqua)

	drawLine(img, 79, 92, 112, Colors_1.Aqua)
	drawLine(img, 211, 224, 112, Colors_1.Aqua)

	drawLine(img, 79, 92, 113, Colors_1.Aqua)
	drawLine(img, 211, 224, 113, Colors_1.Aqua)

	for i := 0; i < 7; i++ {
		drawLine(img, 79, 93, 113+i, Colors_1.Aqua)
		drawLine(img, 210, 224, 113+i, Colors_1.Aqua)
	}
	for i := 0; i < 6; i++ {
		drawLine(img, 79, 94, 120+i, Colors_1.Aqua)
		drawLine(img, 209, 224, 120+i, Colors_1.Aqua)
	}
	for i := 0; i < 5; i++ {
		drawLine(img, 78, 95, 126+i, Colors_1.Aqua)
		drawLine(img, 208, 225, 126+i, Colors_1.Aqua)
	}
	for i := 0; i < 11; i++ {
		drawLine(img, 78, 96, 131+i, Colors_1.Aqua)
		drawLine(img, 207, 225, 131+i, Colors_1.Aqua)
	}

	drawLine(img, 78, 97, 137, Colors_1.Aqua)
	drawLine(img, 206, 225, 137, Colors_1.Aqua)
	drawLine(img, 78, 97, 138, Colors_1.Aqua)
	drawLine(img, 206, 225, 138, Colors_1.Aqua)

	for i := 0; i < 3; i++ {
		drawLine(img, 78, 95, 142+i, Colors_1.Aqua)
		drawLine(img, 208, 225, 142+i, Colors_1.Aqua)
	}

	for i := 0; i < 3; i++ {
		drawLine(img, 78, 94, 145+i, Colors_1.Aqua)
		drawLine(img, 209, 225, 145+i, Colors_1.Aqua)
	}

	drawLine(img, 101, 106, 147, Colors_1.Aqua)
	drawLine(img, 197, 202, 147, Colors_1.Aqua)

	for i := 0; i < 3; i++ {
		drawLine(img, 78, 93, 148+i, Colors_1.Aqua)
		drawLine(img, 210, 225, 148+i, Colors_1.Aqua)
	}
	drawLine(img, 101, 112, 148, Colors_1.Aqua)
	drawLine(img, 192, 202, 148, Colors_1.Aqua)

	drawLine(img, 101, 117, 149, Colors_1.Aqua)
	drawLine(img, 186, 202, 149, Colors_1.Aqua)

	drawLine(img, 100, 122, 150, Colors_1.Aqua)
	drawLine(img, 181, 203, 150, Colors_1.Aqua)

	for i := 0; i < 3; i++ {
		drawLine(img, 78, 92, 151+i, Colors_1.Aqua)
		drawLine(img, 211, 225, 151+i, Colors_1.Aqua)
	}

	drawLine(img, 100, 127, 151, Colors_1.Aqua)
	drawLine(img, 176, 203, 151, Colors_1.Aqua)

	drawLine(img, 100, 132, 152, Colors_1.Aqua)
	drawLine(img, 171, 203, 152, Colors_1.Aqua)

	drawLine(img, 99, 138, 153, Colors_1.Aqua)
	drawLine(img, 165, 204, 153, Colors_1.Aqua)

	for i := 0; i < 3; i++ {
		drawLine(img, 78, 91, 154+i, Colors_1.Aqua)
		drawLine(img, 212, 225, 154+i, Colors_1.Aqua)
	}
	drawLine(img, 99, 141, 154, Colors_1.Aqua)
	drawLine(img, 162, 204, 154, Colors_1.Aqua)

	for i := 0; i < 3; i++ {
		drawLine(img, 78, 90, 157+i, Colors_1.Aqua)
		drawLine(img, 213, 225, 157+i, Colors_1.Aqua)
	}
	for i := 0; i < 5; i++ {
		drawLine(img, 99+i, 141, 154+i, Colors_1.Aqua)
		drawLine(img, 162, 204-i, 154+i, Colors_1.Aqua)
	}

	for i := 0; i < 5; i++ {
		drawLine(img, 103+i, 142-i, 158+i, Colors_1.Aqua)
		drawLine(img, 161+i, 200-i, 158+i, Colors_1.Aqua)
	}

	drawLine(img, 110, 137, 163, Colors_1.Aqua)
	drawLine(img, 166, 193, 163, Colors_1.Aqua)

	drawLine(img, 115, 136, 164, Colors_1.Aqua)
	drawLine(img, 167, 189, 164, Colors_1.Aqua)

	drawLine(img, 118, 135, 165, Colors_1.Aqua)
	drawLine(img, 168, 185, 165, Colors_1.Aqua)

	// Низ маски
	for i := 0; i < 15; i++ {
		drawLine(img, 78, 89, 159+i, Colors_1.Aqua)
		drawLine(img, 214, 225, 159+i, Colors_1.Aqua)
	}

	for i := 0; i < 36; i++ {
		drawLine(img, 77, 88, 174+i, Colors_1.Aqua)
		drawLine(img, 215, 226, 174+i, Colors_1.Aqua)
	}
	for i := 0; i < 11; i++ {
		drawLine(img, 89, 89, 174+i, Colors_1.Aqua)
		drawLine(img, 214, 214, 174+i, Colors_1.Aqua)
	}

	// Вертикально
	for i := 0; i < 10; i++ {
		if i == 0 {
			drawLine(img, 90, 90, 176+i, Colors_1.Aqua)
			drawLine(img, 213, 213, 176+i, Colors_1.Aqua)
		} else {
			drawLine(img, 90, 91, 176+i, Colors_1.Aqua)
			drawLine(img, 212, 213, 176+i, Colors_1.Aqua)
		}

		drawLine(img, 91, 91, 177+i, Colors_1.Aqua)
		drawLine(img, 212, 212, 177+i, Colors_1.Aqua)
		drawLine(img, 92, 92, 178+i, Colors_1.Aqua)
		drawLine(img, 211, 211, 178+i, Colors_1.Aqua)

		if i < 9 {
			drawLine(img, 92, 92, 178+i, Colors_1.Aqua)
			drawLine(img, 211, 211, 178+i, Colors_1.Aqua)
			drawLine(img, 93, 93, 180+i, Colors_1.Aqua)
			drawLine(img, 210, 210, 180+i, Colors_1.Aqua)
			drawLine(img, 94, 94, 181+i, Colors_1.Aqua)
			drawLine(img, 209, 209, 181+i, Colors_1.Aqua)
			drawLine(img, 95, 95, 182+i, Colors_1.Aqua)
			drawLine(img, 208, 208, 182+i, Colors_1.Aqua)
			drawLine(img, 94, 94, 181+i, Colors_1.Aqua)
			drawLine(img, 209, 209, 181+i, Colors_1.Aqua)
			drawLine(img, 95, 95, 182+i, Colors_1.Aqua)
			drawLine(img, 208, 208, 182+i, Colors_1.Aqua)
			drawLine(img, 96, 96, 183+i, Colors_1.Aqua)
			drawLine(img, 207, 207, 183+i, Colors_1.Aqua)
		}
	}

	for i := 0; i < 8; i++ {
		drawLine(img, 97, 97, 185+i, Colors_1.Aqua)
		drawLine(img, 206, 206, 185+i, Colors_1.Aqua)
		drawLine(img, 98, 98, 186+i, Colors_1.Aqua)
		drawLine(img, 205, 205, 186+i, Colors_1.Aqua)
		drawLine(img, 99, 99, 187+i, Colors_1.Aqua)
		drawLine(img, 204, 204, 187+i, Colors_1.Aqua)
		drawLine(img, 100, 100, 188+i, Colors_1.Aqua)
		drawLine(img, 203, 203, 188+i, Colors_1.Aqua)

	}
	for i := 0; i < 17; i++ {
		drawLine(img, 101, 101, 190+i, Colors_1.Aqua)
		drawLine(img, 202, 202, 190+i, Colors_1.Aqua)
	}

	for i := 0; i < 16; i++ {
		drawLine(img, 102, 102, 191+i, Colors_1.Aqua)
		drawLine(img, 201, 201, 191+i, Colors_1.Aqua)
		drawLine(img, 103, 103, 192+i, Colors_1.Aqua)
		drawLine(img, 200, 200, 192+i, Colors_1.Aqua)
		drawLine(img, 104, 104, 193+i, Colors_1.Aqua)
		drawLine(img, 199, 199, 193+i, Colors_1.Aqua)
	}
	for i := 0; i < 14; i++ {
		drawLine(img, 105, 105, 195+i, Colors_1.Aqua)
		drawLine(img, 198, 198, 195+i, Colors_1.Aqua)
		drawLine(img, 106, 106, 196+i, Colors_1.Aqua)
		drawLine(img, 197, 197, 196+i, Colors_1.Aqua)
		drawLine(img, 107, 107, 197+i, Colors_1.Aqua)
		drawLine(img, 196, 196, 197+i, Colors_1.Aqua)
	}

	for i := 0; i < 6; i++ {
		drawLine(img, 108, 108, 205+i, Colors_1.Aqua)
		drawLine(img, 195, 195, 205+i, Colors_1.Aqua)
		drawLine(img, 109, 109, 206+i, Colors_1.Aqua)
		drawLine(img, 194, 194, 206+i, Colors_1.Aqua)
		drawLine(img, 110, 110, 206+i, Colors_1.Aqua)
		drawLine(img, 193, 193, 206+i, Colors_1.Aqua)
		drawLine(img, 111, 111, 207+i, Colors_1.Aqua)
		drawLine(img, 192, 192, 207+i, Colors_1.Aqua)
		drawLine(img, 112, 112, 207+i, Colors_1.Aqua)
		drawLine(img, 191, 191, 207+i, Colors_1.Aqua)
		drawLine(img, 113, 113, 208+i, Colors_1.Aqua)
		drawLine(img, 190, 190, 208+i, Colors_1.Aqua)
		drawLine(img, 114, 114, 208+i, Colors_1.Aqua)
		drawLine(img, 189, 189, 208+i, Colors_1.Aqua)
		drawLine(img, 115, 115, 209+i, Colors_1.Aqua)
		drawLine(img, 188, 188, 209+i, Colors_1.Aqua)
		drawLine(img, 116, 116, 209+i, Colors_1.Aqua)
		drawLine(img, 188, 188, 209+i, Colors_1.Aqua)
	}
	for i := 0; i < 5; i++ {
		drawLine(img, 117, 117, 210+i, Colors_1.Aqua)
		drawLine(img, 187, 187, 210+i, Colors_1.Aqua)
		drawLine(img, 118, 118, 211+i, Colors_1.Aqua)
		drawLine(img, 186, 186, 211+i, Colors_1.Aqua)
		drawLine(img, 119, 119, 211+i, Colors_1.Aqua)
		drawLine(img, 185, 185, 211+i, Colors_1.Aqua)
		drawLine(img, 120, 120, 211+i, Colors_1.Aqua)
		drawLine(img, 184, 184, 211+i, Colors_1.Aqua)
	}
	for i := 0; i < 5; i++ {
		drawLine(img, 121, 121, 210+i, Colors_1.Aqua)
		drawLine(img, 183, 183, 211+i, Colors_1.Aqua)
		drawLine(img, 122, 122, 209+i, Colors_1.Aqua)
		drawLine(img, 182, 182, 210+i, Colors_1.Aqua)
		drawLine(img, 123, 123, 208+i, Colors_1.Aqua)
		drawLine(img, 181, 181, 209+i, Colors_1.Aqua)
		drawLine(img, 124, 124, 207+i, Colors_1.Aqua)
		drawLine(img, 180, 180, 208+i, Colors_1.Aqua)
		drawLine(img, 125, 125, 206+i, Colors_1.Aqua)
		drawLine(img, 179, 179, 207+i, Colors_1.Aqua)
		drawLine(img, 126, 126, 205+i, Colors_1.Aqua)
		drawLine(img, 178, 178, 206+i, Colors_1.Aqua)
		drawLine(img, 177, 177, 205+i, Colors_1.Aqua)
	}
	for i := 0; i < 5; i++ {
		drawLine(img, 127, 176, 204+i, Colors_1.Aqua)
	}

	drawLine(img, 77, 89, 209, Colors_1.Aqua)
	drawLine(img, 214, 226, 209, Colors_1.Aqua)
	drawLine(img, 77, 90, 210, Colors_1.Aqua)
	drawLine(img, 213, 226, 210, Colors_1.Aqua)
	drawLine(img, 78, 92, 211, Colors_1.Aqua)
	drawLine(img, 211, 225, 211, Colors_1.Aqua)
	drawLine(img, 78, 94, 212, Colors_1.Aqua)
	drawLine(img, 209, 225, 212, Colors_1.Aqua)
	drawLine(img, 79, 96, 213, Colors_1.Aqua)
	drawLine(img, 207, 224, 213, Colors_1.Aqua)
	drawLine(img, 79, 98, 214, Colors_1.Aqua)
	drawLine(img, 205, 224, 214, Colors_1.Aqua)
	drawLine(img, 80, 100, 215, Colors_1.Aqua)
	drawLine(img, 203, 223, 215, Colors_1.Aqua)
	drawLine(img, 81, 102, 216, Colors_1.Aqua)
	drawLine(img, 201, 222, 216, Colors_1.Aqua)
	drawLine(img, 81, 104, 217, Colors_1.Aqua)
	drawLine(img, 199, 222, 217, Colors_1.Aqua)
	drawLine(img, 82, 106, 218, Colors_1.Aqua)
	drawLine(img, 197, 221, 218, Colors_1.Aqua)
	for i := 0; i < 8; i++ {
		drawLine(img, 84+i+i, 108+i+i, 219+i, Colors_1.Aqua)
		drawLine(img, 195-i-i, 219-i-i, 219+i, Colors_1.Aqua)
	}
	for i := 0; i < 10; i++ {
		drawLine(img, 100+i+i, 203-i-i, 227+i, Colors_1.Aqua)
	}
	drawLine(img, 121, 182, 237, Colors_1.Aqua)
	drawLine(img, 126, 177, 238, Colors_1.Aqua)
	//бородка
	drawLine(img, 137, 166, 220, Colors_1.Aqua)
	drawLine(img, 135, 168, 221, Colors_1.Aqua)
	drawLine(img, 133, 170, 222, Colors_1.Aqua)
	for i := 0; i < 4; i++ {
		drawLine(img, 132-i, 171+i, 223+i, Colors_1.Aqua)
	}
}

func LoadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}
