// golang でtemplate method をうまく実装できてない...
// 継承がないから仕方ないのかもしれない

package main

import "fmt"

// 単純に、override で実装しようと思っても、 golangでは、structを埋め込んだ時には、
// 親(埋め込まれたstruct)のメソッドから、子(親のstructのポインタを持つstruct)のメソッドや変数にアクセスできない。
/*
type AbstractDisplay struct {}

func (self *AbstractDisplay) open() {};
...
func (self *AbstractDisplay) display() {
	self.open()
	...
}

type CharDisplay struct {
	*AbstractDisplay
}

func (self *CharDisplay) open() {
	fmt.Print("<<<")
}
...
*/
// そういうわけで、引数に渡す方式でやることにする
// これは、単にデフォルト実装を与えたいだけ、という場合の実装パターンの一つとして使う
type printer interface {
	open()
	print()
	close()
}

type AbstractDisplay struct{}

func (self *AbstractDisplay) Display(printer printer) {
	printer.open()
	for i := 0; i < 5; i++ {
		printer.print()
	}
	printer.close()
}

type CharDisplay struct {
	*AbstractDisplay
	char rune
}

func (self *CharDisplay) open() {
	fmt.Print("<<<")
}
func (self *CharDisplay) print() {
	fmt.Print(string(self.char))
}
func (self *CharDisplay) close() {
	fmt.Print(">>>\n")
}

func main() {
	println("--- start 03. Template method. ---")
	d1 := &CharDisplay{&AbstractDisplay{}, 'A'}
	d1.Display(d1)
}
