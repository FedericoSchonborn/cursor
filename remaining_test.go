//go:build cursor_remaining

package cursor_test

import (
	"fmt"

	"github.com/FedericoSchonborn/cursor"
)

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
