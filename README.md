# Cursor

[![Go Reference](https://pkg.go.dev/badge/github.com/fdschonborn/go-cursor.svg)](https://pkg.go.dev/github.com/fdschonborn/go-cursor)

Package `cursor` provides an implementation of Rust's [`std::io::Cursor`][std-io-cursor] for Go.

## Progress

| Rust Method       | Implemented         | Tests | Examples |
|-------------------|---------------------|-------|----------|
| `new`             | Yes (`New`)         | No    | No       |
| `into_inner`      | Yes (`Bytes`)       | No    | No       |
| `position`        | Yes (`Position`)    | No    | Yes      |
| `set_position`    | Yes (`SetPosition`) | No    | Yes      |
| `remaining_slice` | Yes (`Remaining`)   | No    | Yes      |
| `is_empty`        | Yes (`IsEmpty`)     | No    | Yes      |

| Rust Trait | Go Interface | Implemented                | Tests | Examples |
|------------|--------------|----------------------------|-------|----------|
| `Read`     | `Reader`     | Yes (`Read`)               | No    | No       |
| `Write`    | `Writer`     | Yes (`Write`)              | No    | No       |
| `Seek`     | `Seeker`     | Yes (`Seek`)               | No    | No       |
| `Clone`    | N/A          | Yes (`Clone<From>`)        | No    | No       |
| `Eq`       | N/A          | Yes (`Equal<Fold><Bytes>`) | No    | No       |
| `Default`  | N/A          | Yes (`Empty`)              | No    | No       |

[std-io-cursor]: https://doc.rust-lang.org/stable/std/io/struct.Cursor.html
