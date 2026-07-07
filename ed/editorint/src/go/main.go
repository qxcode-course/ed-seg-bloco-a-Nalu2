package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	text   *List[*List[rune]] // a lista de linhas
	itLine *Node[*List[rune]] // iterador para a linha corrente
	itChar *Node[rune]        // iterador para o caracter do cursor
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) getNodeAt(line *List[rune], col int) *Node[rune] {
	if col <= 0 {
		return line.Front()
	}
	size := line.Size()
	if col >= size {
		return line.End()
	}
	node := line.Front()
	for i := 0; i < col; i++ {
		node = node.Next()
	}
	return node
}


func (e *Editor) InsertChar(r rune) {
	e.itChar = e.itLine.Value.Insert(e.itChar, r) // insere antes do elemento apontado pelo cursor
	e.itChar = e.itChar.Next()                    // move o cursor para próxima posição
}

func (e *Editor) KeyLeft() {
	if e.itChar != e.itLine.Value.Front() { // Se o cursor não está no início da linha
		e.itChar = e.itChar.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Front() { // Se não está na primeira linha
		e.itLine = e.itLine.Prev()      // Atualiza iterador de linha para linha anterior
		e.itChar = e.itLine.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	newLine := NewList[rune]()
	it := e.itChar
	end := e.itLine.Value.End()

	for it != end {
		nextIt := it.Next()
		newLine.PushBack(it.Value)
		e.itLine.Value.Erase(it) 
		it = nextIt
	}

	e.text.Insert(e.itLine.Next(), newLine)
	e.itLine = e.itLine.Next()       
	e.itChar = e.itLine.Value.Front() 
}

func (e *Editor) KeyRight() {
	if e.itChar != e.itLine.Value.End() { 
		e.itChar = e.itChar.Next()
		return
	}
	if e.itLine.Next() != e.text.End() { 
		e.itLine = e.itLine.Next()    
		e.itChar = e.itLine.Value.Front() 
	}
}

func (e *Editor) KeyUp() {
	if e.itLine == e.text.Front() { 
		return 
	}
	col := e.itLine.Value.IndexOf(e.itChar) 
	e.itLine = e.itLine.Prev()
	e.itChar = e.getNodeAt(e.itLine.Value, col)
}

func (e *Editor) KeyDown() {
	if e.itLine.Next() == e.text.End() {
		return
	}
	col := e.itLine.Value.IndexOf(e.itChar)
	e.itLine = e.itLine.Next()
	e.itChar = e.getNodeAt(e.itLine.Value, col)
}

func (e *Editor) KeyBackspace() {
	if e.itChar != e.itLine.Value.Front() { 
		toErase := e.itChar.Prev()
		e.itLine.Value.Erase(toErase)        
		return
	}
	if e.itLine == e.text.Front() {
		return
	}
	prevLine := e.itLine.Prev()
	oldChar := e.itChar
	
	e.itChar = prevLine.Value.End()
	
	for it := oldChar; it != e.itLine.Value.End(); {
		nextIt := it.Next()
		prevLine.Value.Insert(prevLine.Value.End(), it.Value)
		it = nextIt
	}

	lineToRemove := e.itLine
	e.itLine = prevLine
	e.text.Erase(lineToRemove)
}

func (e *Editor) KeyDelete() {
	if e.itChar != e.itLine.Value.End() {
		e.itChar = e.itLine.Value.Erase(e.itChar)
		return
	}
	if e.itLine.Next() == e.text.End() {
		return
	}
	nextLine := e.itLine.Next()
	nextFront := nextLine.Value.Front()
	currentSize := e.itLine.Value.Size()
	
	for it := nextFront; it != nextLine.Value.End(); {
		nextIt := it.Next()
		e.itLine.Value.Insert(e.itLine.Value.End(), it.Value)
		it = nextIt
	}
	
	e.itChar = e.getNodeAt(e.itLine.Value, currentSize)
	e.text.Erase(nextLine)
}

func main() { // Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.text = NewList[*List[rune]]()
	e.text.PushBack(NewList[rune]())
	e.itLine = e.text.Front()
	e.itChar = e.itLine.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.text.Front(); line != e.text.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '↲'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.itChar {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
