//go:build cursor_remaining

package cursor_test

import (
	"fmt"

	"github.com/fdschonborn/go-cursor"
)

func ExampleCursor_Remaining() {
	buf := cursor.New([]byte{1, 2, 3, 4, 5})
	fmt.Println(buf.Remaining())

	buf.SetOffset(2)
	fmt.Println(buf.Remaining())

	buf.SetOffset(4)
	fmt.Println(buf.Remaining())

	buf.SetOffset(6)
	fmt.Println(buf.Remaining())

	// Output:
	// [1 2 3 4 5]
	// [3 4 5]
	// [5]
	// []
}

func ExampleCursor_IsEmpty() {
	buf := cursor.New([]byte{1, 2, 3, 4, 5})

	buf.SetOffset(2)
	fmt.Println(!buf.IsEmpty())

	buf.SetOffset(5)
	fmt.Println(buf.IsEmpty())

	buf.SetOffset(10)
	fmt.Println(buf.IsEmpty())

	// Output:
	// true
	// true
	// true
}
