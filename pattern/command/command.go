package main

import "fmt"

type Command interface {
	Execute()
}

type CutCommand struct {
	Receiver *Editor
}

func (c *CutCommand) Execute() {
	c.Receiver.Cut()
}

type CopyCommand struct {
	Receiver *Editor
}

func (c *CopyCommand) Execute() {
	c.Receiver.Copy()
}

type PasteCommand struct {
	Receiver *Editor
}

func (c *PasteCommand) Execute() {
	c.Receiver.Paste()
}

type Editor struct{}

func (e *Editor) Cut() {
	fmt.Println("Текст вырезан")
}

func (e *Editor) Copy() {
	fmt.Println("Текст скопирован")
}

func (e *Editor) Paste() {
	fmt.Println("Текст вставлен")
}

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

func main() {
	editor := &Editor{}

	cutCommand := &CutCommand{Receiver: editor}
	copyCommand := &CopyCommand{Receiver: editor}
	pasteCommand := &PasteCommand{Receiver: editor}

	button1 := &Button{Command: cutCommand}
	button2 := &Button{Command: copyCommand}
	button3 := &Button{Command: pasteCommand}

	button1.Press() // Текст вырезан
	button2.Press() // Текст скопирован
	button3.Press() // Текст вставлен
}
