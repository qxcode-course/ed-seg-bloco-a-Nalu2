package main

import (
	"fmt"
	"math"
)

func circulo(pen *Pen, x, y, raio float64, nivel int) {
	if nivel == 0 {
		return
	}
    pen.Up()
    pen.SetPosition(x, y)
	pen.Down()
    pen.DrawCircle(raio)
	novoRaio := raio / 3.0
	distanciaCentros := raio

	for i := 0; i < 6; i++{
		angulo := float64(i) * 60.0
		rad := angulo * math.Pi / 180.0
		novoX := x + distanciaCentros*math.Cos(rad)
		novoY := y + distanciaCentros*math.Sin(rad)
		circulo(pen, novoX, novoY, novoRaio, nivel-1)
	}
}

func main() {
    largura, altura := 800, 800
	pen := NewPen(largura, altura)
	pen.dc.SetRGB(0, 0, 0) 
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255) 
	pen.SetLineWidth(1)

	centroX := float64(largura) / 2.0
	centroY := float64(altura) / 2.0
	raioInicial := 250.0 
	niveis := 6       
	circulo(pen, centroX, centroY, raioInicial, niveis)

    pen.SavePNG("circulos.png")
	fmt.Println("PNG gerado com sucesso!")
}