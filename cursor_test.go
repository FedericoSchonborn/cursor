package cursor_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/FedericoSchonborn/cursor"
)

// TODO: Port missing tests and examples (Reader, Seeker).

func Example() {
	writeTenBytesAtEnd := func(ws io.WriteSeeker) error {
		if _, err := ws.Seek(-10, io.SeekEnd); err != nil {
			return err
		}

		if _, err := ws.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}); err != nil {
			return err
		}

		return nil
	}

	buf := cursor.From(make([]byte, 15))
	if err := writeTenBytesAtEnd(buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.Bytes()[5:15])
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleCursor_Offset() {
	buf := cursor.From([]byte{1, 2, 3, 4, 5})
	fmt.Println(buf.Offset())

	if _, err := buf.Seek(2, io.SeekCurrent); err != nil {
		panic(err)
	}
	fmt.Println(buf.Offset())

	if _, err := buf.Seek(-1, io.SeekCurrent); err != nil {
		panic(err)
	}
	fmt.Println(buf.Offset())

	// Output:
	// 0
	// 2
	// 1
}

func ExampleCursor_SetOffset() {
	buf := cursor.From([]byte{1, 2, 3, 4, 5})
	fmt.Println(buf.Offset())

	buf.SetOffset(2)
	fmt.Println(buf.Offset())

	buf.SetOffset(4)
	fmt.Println(buf.Offset())

	// Output:
	// 0
	// 2
	// 4
}

func TestWriter(t *testing.T) {
	w := cursor.New()
	var (
		n   int
		err error
	)

	n, err = w.Write([]byte{0})
	if err != nil {
		t.Fatal(err)
	}

	if n != 1 {
		t.Errorf("Expected n to be 1, got %d", n)
	}

	n, err = w.Write([]byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}

	if n != 3 {
		t.Errorf("Expected n to be 3, got %d", n)
	}

	n, err = w.Write([]byte{4, 5, 6, 7})
	if err != nil {
		t.Fatal(err)
	}

	if n != 4 {
		t.Errorf("Expected n to be 4, got %d", n)
	}

	expected := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	result := w.Bytes()
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReader(t *testing.T) {
	r := cursor.From([]byte{0, 1, 2, 3, 4, 5, 6, 7})
	var (
		buf    []byte
		b      []byte
		n      int
		offset int
		err    error
	)

	buf = make([]byte, 0)
	n, err = r.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 0 {
		t.Errorf("expected n to be 0, got %d", n)
	}

	offset = r.Offset()
	if offset != 0 {
		t.Errorf("expected offset to be 0, got %d", offset)
	}

	buf = make([]byte, 1)
	n, err = r.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 1 {
		t.Errorf("expected n to be 1, got %d", n)
	}

	offset = r.Offset()
	if offset != 1 {
		t.Errorf("expected offset to be 1, got %d", offset)
	}

	buf = make([]byte, 4)
	n, err = r.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 4 {
		t.Errorf("expected n to be 4, got %d", n)
	}

	offset = r.Offset()
	if offset != 5 {
		t.Errorf("expected offset to be 5, got %d", offset)
	}

	b = []byte{1, 2, 3, 4}
	if !bytes.Equal(buf, b) {
		t.Errorf("expected buffer to be %v, got %v", b, buf)
	}

	n, err = r.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 3 {
		t.Errorf("expected n to be 3, got %d", n)
	}

	b = []byte{5, 6, 7}
	buf3 := buf[:3]
	if !bytes.Equal(buf3, b) {
		t.Errorf("expected buffer to be %v, got %v", b, buf3)
	}

	n, err = r.Read(buf)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			t.Fatal(err)
		}
	}

	if n != 0 {
		t.Errorf("expected n to be 0, got %d", n)
	}
}

func TestReadAll(t *testing.T) {
	bs := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	r := cursor.From(bs)
	v, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(bs, v) {
		t.Errorf("Expected result to be %v, got %v", bs, v)
	}
}
