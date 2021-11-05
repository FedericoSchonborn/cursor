package cursor_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/fdschonborn/go-cursor"
)

// TODO(fdschonborn): Port missing tests and examples.

func Example() {
	writeTenBytesAtEnd := func(ws io.WriteSeeker) error {
		if _, err := ws.Seek(-10, io.SeekEnd); err != nil {
			return err
		}

		for i := byte(0); i < 10; i++ {
			if _, err := ws.Write([]byte{i}); err != nil {
				return err
			}
		}

		return nil
	}

	buf := cursor.New(make([]byte, 15))
	if err := writeTenBytesAtEnd(buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.Bytes()[5:15])
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleCursor_Position() {
	buf := cursor.New([]byte{1, 2, 3, 4, 5})
	fmt.Println(buf.Position())

	buf.Seek(2, io.SeekCurrent)
	fmt.Println(buf.Position())

	buf.Seek(-1, io.SeekCurrent)
	fmt.Println(buf.Position())

	// Output:
	// 0
	// 2
	// 1
}

func ExampleCursor_SetPosition() {
	buf := cursor.New([]byte{1, 2, 3, 4, 5})
	fmt.Println(buf.Position())

	buf.SetPosition(2)
	fmt.Println(buf.Position())

	buf.SetPosition(4)
	fmt.Println(buf.Position())

	// Output:
	// 0
	// 2
	// 4
}

func ExampleCursor_Remaining() {
	buf := cursor.New([]byte{1, 2, 3, 4, 5})
	fmt.Println(buf.Remaining())

	buf.SetPosition(2)
	fmt.Println(buf.Remaining())

	buf.SetPosition(4)
	fmt.Println(buf.Remaining())

	buf.SetPosition(6)
	fmt.Println(buf.Remaining())

	// Output:
	// [1 2 3 4 5]
	// [3 4 5]
	// [5]
	// []
}

func ExampleCursor_IsEmpty() {
	buf := cursor.New([]byte{1, 2, 3, 4, 5})

	buf.SetPosition(2)
	fmt.Println(!buf.IsEmpty())

	buf.SetPosition(5)
	fmt.Println(buf.IsEmpty())

	buf.SetPosition(10)
	fmt.Println(buf.IsEmpty())

	// Output:
	// true
	// true
	// true
}

func TestWriter(t *testing.T) {
	w := cursor.New([]byte{})
	var (
		n   int
		err error
	)

	n, err = w.Write([]byte{0})
	if err != nil {
		t.Fatal(n)
	}

	if n != 1 {
		t.Errorf("Expected n to be 1, got %d", n)
	}

	n, err = w.Write([]byte{1, 2, 3})
	if err != nil {
		t.Fatal(n)
	}

	if n != 3 {
		t.Errorf("Expected n to be 3, got %d", n)
	}

	n, err = w.Write([]byte{4, 5, 6, 7})
	if err != nil {
		t.Fatal(n)
	}

	if n != 4 {
		t.Errorf("Expected n to be 4, got %d", n)
	}

	expected := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	actual := w.Bytes()

	result := bytes.Compare(expected, actual)
	if result != 0 {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
