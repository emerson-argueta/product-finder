package productfinder

// Achiever moodel errors.
const (
	ErrSearchIncompleteDetails = Error("incomplete details for search")
)

// Error represents a general domain error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
