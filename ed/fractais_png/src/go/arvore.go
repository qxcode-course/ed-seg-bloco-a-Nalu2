package main

import (
	"math"
)

func desenhaArvore(p *Pen, x, y, length, angle float64, depth int) {
	if depth == 0 || length < 2 {
		return
	}

	p.Up()
	p.SetPosition(x, y)
	p.SetHeading(angle)
	p.Down()
	p.Walk(length)

	rad := angle * math.Pi / 180
	fimX := x + length*math.Cos(rad)
	fimY := y - length*math.Sin(rad) 

	anguloRamo := 30.0
	novoCompr := length * 0.7

	desenhaArvore(p, fimX, fimY, novoCompr, angle+anguloRamo, depth-1)

	desenhaArvore(p, fimX, fimY, novoCompr, angle-anguloRamo, depth-1)
}

func main() {
	largura, altura := 800, 800
	pen := NewPen(largura, altura)
	pen.SetRGB(0, 0, 0) 
	pen.SetLineWidth(2)
	baseX := float64(largura) / 2
	baseY := float64(altura) * 0.9 
	desenhaArvore(pen, baseX, baseY, 150, 90, 8)
	pen.SavePNG("arvore.png")
}