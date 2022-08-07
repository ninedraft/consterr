# consterr

Small type helper for declaring constant errors.

```go
// marker type.
// It a compile type meta for constant errors.
// We need it to distinguish two errors with equal values.
type errPackage struct{}

// Err is a package level error.
type Err = consterr.Err[errPackage]

const (
    ErrBadArg Err = "bad argument"
    ErrUnexpectedValue Err = "unexpected value"
)
```