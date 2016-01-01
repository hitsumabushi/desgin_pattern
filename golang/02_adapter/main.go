/* adapter pattern */

package main

import "fmt"

type Print interface {
	PrintWeak()
	PrintStrong()
}

type Banner struct {
	str string
}

func (b *Banner) showWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) showWithAster() {
	fmt.Printf("*%s*\n", b.str)
}

type PrintBanner struct {
	*Banner
}

func (this *PrintBanner) PrintWeak()   { this.showWithParen() }
func (this *PrintBanner) PrintStrong() { this.showWithAster() }

func main() {
	println("--- start adapter pattern ---")
	p := PrintBanner{&Banner{str: "hello"}}
	p.PrintWeak()
	p.PrintStrong()
}
