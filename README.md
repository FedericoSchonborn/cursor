# Cursor

[![Go Reference](https://pkg.go.dev/badge/github.com/FedericoSchonborn/go-cursor.svg)](https://pkg.go.dev/github.com/FedericoSchonborn/go-cursor)

Package `cursor` provides an implementation of Rust's [`std::io::Cursor`][std-io-cursor] for Go.

## Example

```go
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
```

## Status

| Method         | Implemented       | Tests | Examples |
| -------------- | ----------------- | ----- | -------- |
| `new`          | Yes (`From`)      | No    | No       |
| `into_inner`   | Yes (`Unwrap`)    | No    | No       |
| `position`     | Yes (`Offset`)    | No    | Yes      |
| `set_position` | Yes (`SetOffset`) | No    | Yes      |

### Traits

| Trait     | Interface | Implemented                   | Tests | Examples |
| --------- | --------- | ----------------------------- | ----- | -------- |
| `Read`    | `Reader`  | Yes (`Read`)                  | No    | No       |
| `Write`   | `Writer`  | Yes (`Write`)                 | No    | No       |
| `Seek`    | `Seeker`  | Yes (`Seek`)                  | No    | No       |
| `Clone`   | N/A       | Yes (`Clone` and `CloneFrom`) | No    | No       |
| `Eq`      | N/A       | Yes (`Equal`)                 | No    | No       |
| `Default` | N/A       | Yes (`New`)                   | No    | No       |

### Unstable Features

| Feature/Build Tag  | Method            | Implemented       | Tests | Examples |
| ------------------ | ----------------- | ----------------- | ----- | -------- |
| `cursor_remaining` | `remaining_slice` | Yes (`Remaining`) | No    | Yes      |
| `cursor_remaining` | `is_empty`        | Yes (`IsEmpty`)   | No    | Yes      |

[std-io-cursor]: https://doc.rust-lang.org/stable/std/io/struct.Cursor.html
