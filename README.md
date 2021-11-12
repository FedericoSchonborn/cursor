# Cursor

[![Go Reference](https://pkg.go.dev/badge/github.com/fdschonborn/go-cursor.svg)](https://pkg.go.dev/github.com/fdschonborn/go-cursor)

Package `cursor` provides an implementation of Rust's [`std::io::Cursor`][std-io-cursor] for Go.

## Example

```go
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
```

## Status

| Method            | Implemented         | Tests | Examples |
| ----------------- | ------------------- | ----- | -------- |
| `new`             | Yes (`New`)         | No    | No       |
| `into_inner`      | Yes (`Bytes`)       | No    | No       |
| `position`        | Yes (`Position`)    | No    | Yes      |
| `set_position`    | Yes (`SetPosition`) | No    | Yes      |
| `remaining_slice` | Yes (`Remaining`)   | No    | Yes      |
| `is_empty`        | Yes (`IsEmpty`)     | No    | Yes      |

| Trait     | Interface | Implemented                | Tests | Examples |
| --------- | --------- | -------------------------- | ----- | -------- |
| `Read`    | `Reader`  | Yes (`Read`)               | No    | No       |
| `Write`   | `Writer`  | Yes (`Write`)              | No    | No       |
| `Seek`    | `Seeker`  | Yes (`Seek`)               | No    | No       |
| `Clone`   | N/A       | Yes (`Clone<From>`)        | No    | No       |
| `Eq`      | N/A       | Yes (`Equal<Fold><Bytes>`) | No    | No       |
| `Default` | N/A       | Yes (`Empty`)              | No    | No       |

[std-io-cursor]: https://doc.rust-lang.org/stable/std/io/struct.Cursor.html
