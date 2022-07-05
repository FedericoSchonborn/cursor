//go:build cursor_remaining

package cursor

// Remaining returns the remaining slice.
func (c *Cursor) Remaining() []byte {
	len := min(c.off, len(c.buf))
	return c.buf[len:]
}

// IsEmpty returns true if the remaining slice is empty.
func (c *Cursor) IsEmpty() bool {
	return c.off >= len(c.buf)
}
