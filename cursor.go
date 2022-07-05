// Package cursor provides a Go implementation of Rust's Cursor.
package cursor

import (
	"errors"
	"io"
	"math/bits"
)

// TODO: Improve documentation.

var _ io.ReadWriteSeeker = (*Cursor)(nil)

// Cursor wraps a byte slice and provides it with io.Reader, io.Writer and
// io.Seeker implementations.
type Cursor struct {
	buf []byte
	off int
}

// New creates a new Cursor wrapping an uninitialized byte slice.
//
// New and Cursor{} are equivalent.
func New() *Cursor {
	return &Cursor{
		buf: nil,
	}
}

// From creates a new Cursor wrapping the given byte slice.
func From(slice []byte) *Cursor {
	return &Cursor{
		buf: slice,
	}
}

// Clone sets the wrapped byte slice to a copy of the byte slice contained
// in other.
func Clone(other *Cursor) *Cursor {
	bytes := make([]byte, len(other.buf))
	copy(bytes, other.buf)

	return &Cursor{
		buf: bytes,
		off: other.off,
	}
}

// Clone creates a new Cursor containing a copy of the wrapped byte slice.
func (c *Cursor) Clone() *Cursor {
	bytes := make([]byte, len(c.buf))
	copy(bytes, c.buf)

	return &Cursor{
		buf: bytes,
		off: c.off,
	}
}

// Bytes returns the wrapped byte slice.
func (c *Cursor) Bytes() []byte {
	return c.buf
}

// Unwrap invalidates the Cursor and returns the wrapped byte slice.
func (c *Cursor) Unwrap() []byte {
	bytes := c.buf
	c.buf = nil
	c.off = 0
	c = nil

	return bytes
}

// Offset returns the current offset.
func (c *Cursor) Offset() int {
	return c.off
}

// SetOffset sets the current offset.
func (c *Cursor) SetOffset(pos int) {
	c.off = pos
}

// Read implements io.Reader for Cursor.
func (c *Cursor) Read(p []byte) (n int, err error) {
	if c.off >= len(c.buf) {
		return 0, io.EOF
	}

	n = copy(p, c.buf[c.off:])
	c.off += n
	return n, nil
}

// Write implements io.Writer for Cursor.
func (c *Cursor) Write(p []byte) (n int, err error) {
	if len(c.buf) <= len(p) {
		bytes := make([]byte, len(c.buf)+len(p))
		copy(bytes, c.buf)
		c.buf = bytes
	}

	pos := min(c.off, len(c.buf))
	count := copy(c.buf[pos:], p)
	c.off += count
	return count, nil
}

// Seek implements io.Seeker for Cursor.
func (c *Cursor) Seek(offset int64, whence int) (int64, error) {
	var basePos int
	switch whence {
	case io.SeekStart:
		c.off = int(offset)
		return offset, nil
	case io.SeekEnd:
		basePos = len(c.buf)
	case io.SeekCurrent:
		basePos = c.off
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

	c.off = int(newPos)
	return int64(newPos), nil
}

func min(l, r int) int {
	if l < r {
		return l
	}

	return r
}

func checkedAdd64(x, y uint64) (_ uint64, ok bool) {
	value, carried := bits.Add64(x, y, 0)
	return value, carried == 0
}

func checkedSub64(x, y uint64) (_ uint64, ok bool) {
	value, borrowed := bits.Sub64(x, y, 0)
	return value, borrowed == 0
}
