package cursor

import (
	"bytes"
	"errors"
	"io"
)

var _ io.ReadWriteSeeker = (*Cursor)(nil)

type Cursor struct {
	buf []byte
	pos int
}

func New(buf []byte) *Cursor {
	return &Cursor{
		buf: buf,
		pos: 0,
	}
}

func Read(r io.Reader) (*Cursor, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return New(buf), nil
}

func (c *Cursor) Bytes() []byte {
	return c.buf
}

func (c *Cursor) Position() int {
	return c.pos
}

func (c *Cursor) SetPosition(pos int) {
	c.pos = pos
}

func (c *Cursor) Remaining() []byte {
	len := min(c.pos, len(c.buf))
	return c.buf[len:]
}

func (c *Cursor) IsEmpty() bool {
	return c.pos >= len(c.buf)
}

func (c *Cursor) Clone() *Cursor {
	buf := make([]byte, len(c.buf))
	copy(buf, c.buf)

	return &Cursor{
		buf: buf,
		pos: c.pos,
	}
}

func (c *Cursor) CloneFrom(other *Cursor) {
	c.buf = make([]byte, len(other.buf))
	copy(c.buf, other.buf)
	c.pos = other.pos
}

func (c *Cursor) Read(p []byte) (n int, err error) {
	r := bytes.NewReader(c.Remaining())
	if n, err := r.Read(p); err != nil {
		return n, err
	}

	c.pos += n
	return n, nil
}

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

func min(l, r int) int {
	if l < r {
		return l
	}

	return r
}
