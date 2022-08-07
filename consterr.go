package consterr

// Err is an error type, which can be used in constant declaration.
//
// Example:
//	const ErrInvalidArg Err[marker] = "invalid argument"
//
// M parameter is used to distinguish two different errors with equal values.
//
// Example:
//
//	-- package a
//	type marker struct{}
//	type Err = consterr.Err[marker]
//	const ErrInvalidArg Err = "invalid argument"
//
//	-- package b
//	type marker struct{}
//	type Err = consterr.Err[marker]
//	const ErrInvalidArg Err = "invalid argument"
//
//	-- main.go
//	fmt.Println(a.ErrInvalidArg==b.ErrInvalidArg) // -> false
type Err[M ~struct{}] string

func (err Err[M]) Error() string {
	return string(err)
}
