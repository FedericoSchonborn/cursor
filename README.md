# Cursor

[![Go Reference](https://pkg.go.dev/badge/github.com/FedericoSchonborn/cursor.svg)](https://pkg.go.dev/github.com/FedericoSchonborn/cursor)

Package `cursor` provides an implementation of Rust's [`std::io::Cursor`] for Go.

## Example

```go
package main

import (
    "fmt"
    "io"

    "github.com/FedericoSchonborn/cursor"
)

func main() {
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

### Associated Items

|                | Implemented       | Tests | Examples |
| -------------- | ----------------- | ----- | -------- |
| `new`          | Yes (`New`)       | No    | No       |
| `into_inner`   | Yes (`Unwrap`)    | No    | No       |
| `position`     | Yes (`Offset`)    | No    | Yes      |
| `set_position` | Yes (`SetOffset`) | No    | Yes      |

### Traits

|         | Interface | Implemented   | Tests | Examples |
| ------- | --------- | ------------- | ----- | -------- |
| `Read`  | `Reader`  | Yes (`Read`)  | No    | No       |
| `Write` | `Writer`  | Yes (`Write`) | No    | No       |
| `Seek`  | `Seeker`  | Yes (`Seek`)  | No    | No       |
| `Clone` | N/A       | Yes (`Clone`) | No    | No       |

### Unstable Features

|                   | Implemented       | Feature/Build Tag  | Tests | Examples |
| ----------------- | ----------------- | ------------------ | ----- | -------- |
| `remaining_slice` | Yes (`Remaining`) | `cursor_remaining` | No    | Yes      |
| `is_empty`        | Yes (`IsEmpty`)   | `cursor_remaining` | No    | Yes      |

[`std::io::Cursor`]: https://doc.rust-lang.org/stable/std/io/struct.Cursor.html
