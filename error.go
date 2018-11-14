package loteria

type (
	// Error defines a concrete Loteria error.
	Error string
)

// Error returns the concrete error message.
func (e Error) Error() string {
	return string(e)
}
