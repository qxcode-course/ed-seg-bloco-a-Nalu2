package main

import (
	"fmt"
)

func pentatree(pen *Pen, x, y, comprimento, reducao float64, nivel int) {
	if nivel == 0 {
		return
	}

	for i := 0; i < 6; i++ {
		angulo := float64(i)*72.0 - 90.0
		
		pen.Up()
		pen.SetPosition(x, y)
		pen.SetHeading(angulo)
		pen.Down()

		pen.Walk(comprimento)

		novoX := pen.x
		novoY := pen.y

		pentatree(pen, novoX, novoY, comprimento*reducao, reducao, nivel-1)
	}
}

func main() {
	largura, altura := 700, 700
	pen := NewPen(largura, altura)

	pen.dc.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)
	pen.SetLineWidth(0.6)

	centroX := float64(largura) / 2.0
	centroY := float64(altura) / 2.0

	pentatree(pen, centroX, centroY, 165.0, 0.38, 5)
	pen.SavePNG("pentatree.png")
	fmt.Println("Pentatree'pentatree.png'!")
}