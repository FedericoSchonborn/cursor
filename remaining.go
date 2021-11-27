//go:build cursor_remaining

package cursor

// Remaining returns the remaining slice.
func (c *Cursor) Remaining() []byte {
	len := min(c.offset, len(c.bytes))
	return c.bytes[len:]
}

// IsEmpty returns true if the remaining slice is empty.
func (c *Cursor) IsEmpty() bool {
	return c.offset >= len(c.bytes)
}
