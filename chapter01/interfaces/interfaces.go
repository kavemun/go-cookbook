package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy copies data from in to out first directly
// then using a bugger. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	//we write to out, but also to Stdout
	w := io.MultiWriter(out, os.Stdout)

	// a standard copy, this can be dangerous if theres a lot of data in in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	//buffered write using 64 byte chunks
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// lets pring a new line
	fmt.Println()

	return nil
}
