package book

type Book struct {
	name string
}

func NewBook(name string) Book {
	return Book{
		name: name,
	}
}

func (b Book) GetName() string {
	return b.name
}

type Bookshelf struct {
	books []Book
}

func NewBookshelf(bs []Book) Bookshelf {
	return Bookshelf{
		books: bs,
	}
}

func (bs *Bookshelf) GetBookAt(i int) Book {
	return bs.books[i]
}

func (bs *Bookshelf) AppendBook(b Book) {
	bs.books = append(bs.books, b)
}

func (bs *Bookshelf) GetLength() int {
	return len(bs.books)
}

func (bs *Bookshelf) Iterator() *BookshelfIterator {
	return NewBookshelfIterator(*bs)
}

type BookshelfIterator struct {
	bookshelf Bookshelf
	index     int
}

func NewBookshelfIterator(bs Bookshelf) *BookshelfIterator {
	return &BookshelfIterator{
		bookshelf: bs,
	}
}

func (it *BookshelfIterator) HasNext() bool {
	return it.index < it.bookshelf.GetLength()
}

func (it *BookshelfIterator) Next() Book {
	b := it.bookshelf.GetBookAt(it.index)
	it.index++
	return b
}
