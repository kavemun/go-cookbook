package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// A Book has an Author and Title
type Book struct {
	Author string
	Title  string
}

// Books is a named tupe for an array of books
type Books []Book

// ToCSV takes a set of Books and writes to an io.Writer
// it returns any errors
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	n.Flush()
	return n.Error()
}

// WriteCSVOutput initializes a set of books
// and writes the to os.Stdout
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "Simon Sinek",
			Title:  "The Infinite Game",
		},
	}

	return b.ToCSV(os.Stdout)

}

// WriteCSVBuffer returns a bugger vsv for
// a set of books

func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "Simon Sinek",
			Title:  "The Infinite Game",
		},
	}

	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err

}
