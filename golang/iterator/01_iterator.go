"implement iterator pattern"

package main

import "fmt"

type Iterator interface {
	hasNext() bool
	next() interface{}
}

func NewBookShelfIterator(this *BookShelf) *BookShelfIterator {
	FIRST_ITEM := 0
	return &BookShelfIterator{this, FIRST_ITEM}
}

type BookShelfIterator struct {
	bookshelf *BookShelf
	index     int
}

func (bsi *BookShelfIterator) hasNext() bool {
	return bsi.index < bsi.bookshelf.last
}

func (bsi *BookShelfIterator) next() interface{} {
	if bsi.index >= bsi.bookshelf.last {
		return nil
	}
	book := bsi.bookshelf.BookAt(bsi.index)
	bsi.index += 1
	return book

}

type Aggregate interface {
	iterator() *Iterator
}

// represent iteratable bookshelf
type BookShelf struct {
	books []Book
	last  int
}

func (bs *BookShelf) BookAt(n int) Book {
	return bs.books[n]
}

func (bs *BookShelf) AppendBook(b Book) {
	bs.books = append(bs.books, b)
	bs.last += 1
}

func (bs *BookShelf) Volumes() int {
	return len(bs.books)
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}

// represent book
type Book struct {
	name string
}

func (book *Book) Name() string {
	return book.name
}

func main() {
	books := []Book{
		Book{"foo"},
		Book{"bar"},
		Book{"hoge"},
		Book{"fuga"},
		Book{"weeeei"},
	}
	bs := &BookShelf{books, len(books)}
	bs.AppendBook(Book{"appended book"})
	bsi := bs.Iterator()
	// pattern 1
	var book Book
	for bsi.hasNext() {
		book = bsi.next().(Book)
		fmt.Printf("%#v\n", book)
	}
	// pattern 2
	//for book = bsi.next(); book != nil; book = bsi.next() {
	//	fmt.Printf("%#v\n", book)
	//}
	fmt.Printf("Volumes: %#d\n", bs.Volumes())
}
