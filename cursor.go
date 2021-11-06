// Package cursor provides a Go implementation of Rust's Cursor.
package cursor

import (
	"bytes"
	"errors"
	"io"
)

// TODO(fdschonborn): Improve documentation.

var _ io.ReadWriteSeeker = (*Cursor)(nil)

// Cursor wraps an in-memory buffer and provides it with a Seek implementation.
type Cursor struct {
	buf []byte
	pos int
}

// New initializes a new Cursor wrapping buf.
func New(buf []byte) *Cursor {
	return &Cursor{
		buf: buf,
		pos: 0,
	}
}

// Empty initializes a new Cursor wrapping an empty slice.
func Empty() *Cursor {
	return New(nil)
}

// Read reads all of r and initializes a new Cursor wrapping the data.
func Read(r io.Reader) (*Cursor, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return New(buf), nil
}

// Clone initializes a new Cursor wrapping a copy of c's inner slice.
func (c *Cursor) Clone() *Cursor {
	buf := make([]byte, len(c.buf))
	copy(buf, c.buf)

	return &Cursor{
		buf: buf,
		pos: c.pos,
	}
}

// CloneFrom sets the inner slice of the current Cursor to a copy of other's.
func (c *Cursor) CloneFrom(other *Cursor) {
	c.buf = make([]byte, len(other.buf))
	copy(c.buf, other.buf)
	c.pos = other.pos
}

// Bytes returns the inner slice.
func (c *Cursor) Bytes() []byte {
	return c.buf
}

// Position returns the current position.
func (c *Cursor) Position() int {
	return c.pos
}

// SetPosition sets the current position.
func (c *Cursor) SetPosition(pos int) {
	c.pos = pos
}

// Remaining returns the remaining bytes.
func (c *Cursor) Remaining() []byte {
	len := min(c.pos, len(c.buf))
	return c.buf[len:]
}

// IsEmpty returns whether the Cursor is empty.
func (c *Cursor) IsEmpty() bool {
	return c.pos >= len(c.buf)
}

// Read implements io.Reader for Cursor.
func (c *Cursor) Read(p []byte) (n int, err error) {
	if c.pos >= len(c.buf) {
		return -1, io.EOF
	}

	n = copy(p, c.buf[c.pos:])
	c.pos += n
	return n, nil
}

// Write implements io.Writer for Cursor.
func (c *Cursor) Write(p []byte) (n int, err error) {
	if len(c.buf) <= len(p) {
		buf := make([]byte, len(c.buf)+len(p))
		copy(buf, c.buf)
		c.buf = buf
	}

	pos := min(c.pos, len(c.buf))
	count := copy(c.buf[pos:], p)
	c.pos += count
	return count, nil
}

// Seek implements io.Seeker for Cursor.
func (c *Cursor) Seek(offset int64, whence int) (int64, error) {
	var basePos int
	switch whence {
	case io.SeekStart:
		c.pos = int(offset)
		return offset, nil
	case io.SeekEnd:
		basePos = len(c.buf)
	case io.SeekCurrent:
		basePos = c.pos
	}

	newPos := basePos + int(offset)
	if newPos <= 0 {
		return -1, errors.New("invalid seek to negative position")
	}

	c.pos = newPos
	return int64(newPos), nil
}

func (c *Cursor) Compare(other *Cursor) int {
	return bytes.Compare(c.buf, other.buf)
}

func (c *Cursor) CompareBytes(buf []byte) int {
	return bytes.Compare(c.buf, buf)
}

func (c *Cursor) Equal(other *Cursor) bool {
	return bytes.Equal(c.buf, other.buf)
}

func (c *Cursor) EqualBytes(buf []byte) bool {
	return bytes.Equal(c.buf, buf)
}

func (c *Cursor) EqualFold(other *Cursor) bool {
	return bytes.EqualFold(c.buf, other.buf)
}

func (c *Cursor) EqualFoldBytes(buf []byte) bool {
	return bytes.EqualFold(c.buf, buf)
}

func min(l, r int) int {
	if l < r {
		return l
	}

	return r
}
