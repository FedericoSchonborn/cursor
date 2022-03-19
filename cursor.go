// Package cursor provides a Go implementation of Rust's Cursor.
package cursor

import (
	"bytes"
	"errors"
	"io"
	"math/bits"
)

// TODO: Improve documentation.

var _ io.ReadWriteSeeker = (*Cursor)(nil)

// Cursor wraps a byte slice and provides it with io.Reader, io.Writer and
// io.Seeker implementations.
type Cursor struct {
	bytes  []byte
	offset int
}

// New creates a new Cursor wrapping an uninitialized byte slice.
//
// New and Cursor{} are equivalent.
func New() *Cursor {
	return &Cursor{
		bytes: nil,
	}
}

// From creates a new Cursor wrapping the given byte slice.
func From(bytes []byte) *Cursor {
	return &Cursor{
		bytes: bytes,
	}
}

// Clone sets the wrapped byte slice to a copy of the byte slice contained
// in other.
func Clone(other *Cursor) *Cursor {
	bytes := make([]byte, len(other.bytes))
	copy(bytes, other.bytes)

	return &Cursor{
		bytes:  bytes,
		offset: other.offset,
	}
}

// Clone creates a new Cursor containing a copy of the wrapped byte slice.
func (c *Cursor) Clone() *Cursor {
	bytes := make([]byte, len(c.bytes))
	copy(bytes, c.bytes)

	return &Cursor{
		bytes:  bytes,
		offset: c.offset,
	}
}

// Bytes returns the wrapped byte slice.
func (c *Cursor) Bytes() []byte {
	return c.bytes
}

// Unwrap invalidates the Cursor and returns the wrapped byte slice.
func (c *Cursor) Unwrap() []byte {
	bytes := c.bytes
	c.bytes = nil
	c.offset = 0
	c = nil

	return bytes
}

// Offset returns the current offset.
func (c *Cursor) Offset() int {
	return c.offset
}

// SetOffset sets the current offset.
func (c *Cursor) SetOffset(pos int) {
	c.offset = pos
}

// Read implements io.Reader for Cursor.
func (c *Cursor) Read(p []byte) (n int, err error) {
	if c.offset >= len(c.bytes) {
		return 0, io.EOF
	}

	n = copy(p, c.bytes[c.offset:])
	c.offset += n
	return n, nil
}

// Write implements io.Writer for Cursor.
func (c *Cursor) Write(p []byte) (n int, err error) {
	if len(c.bytes) <= len(p) {
		bytes := make([]byte, len(c.bytes)+len(p))
		copy(bytes, c.bytes)
		c.bytes = bytes
	}

	pos := min(c.offset, len(c.bytes))
	count := copy(c.bytes[pos:], p)
	c.offset += count
	return count, nil
}

func (c *Cursor) Compare(other *Cursor) int {
	return bytes.Compare(c.bytes, other.bytes)
}

func (c *Cursor) Equal(other *Cursor) bool {
	return bytes.Equal(c.bytes, other.bytes)
}

// Seek implements io.Seeker for Cursor.
func (c *Cursor) Seek(offset int64, whence int) (int64, error) {
	var basePos int
	switch whence {
	case io.SeekStart:
		c.offset = int(offset)
		return offset, nil
	case io.SeekEnd:
		basePos = len(c.bytes)
	case io.SeekCurrent:
		basePos = c.offset
	}

	var (
		newPos uint64
		ok     bool
	)
	if offset >= 0 {
		newPos, ok = checkedAdd64(uint64(basePos), uint64(offset))
	} else {
		newPos, ok = checkedSub64(uint64(basePos), uint64(-offset))
	}

	if !ok {
		return -1, errors.New("invalid seek to negative or overflowing position")
	}

	c.offset = int(newPos)
	return int64(newPos), nil
}

func min(l, r int) int {
	if l < r {
		return l
	}

	return r
}

func checkedAdd64(x, y uint64) (_ uint64, ok bool) {
	result, carry := bits.Add64(x, y, 0)
	return result, carry == 0
}

func checkedSub64(x, y uint64) (_ uint64, ok bool) {
	result, borrow := bits.Sub64(x, y, 0)
	return result, borrow == 0
}
