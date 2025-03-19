// Optional type implementation in Go:
//  - inspired by Rust's Option type.
//  - supports generics.
//  - allows for safe handling of nil values as function parameters and return values
// If you're coming from Rust, you'll feel right at home.

package optional

type Optional[T any] struct {
	value *T
}

// Some creates an Optional with a value.
func Some[T any](v T) Optional[T] {
	return Optional[T]{value: &v}
}

// None creates an empty Optional.
func None[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

// IsSome checks if Optional has a value.
func (o Optional[T]) IsSome() bool {
	return o.value != nil
}

// IsNone checks if Optional is empty.
func (o Optional[T]) IsNone() bool {
	return o.value == nil
}

// Unwrap returns the underlying value or panics if None.
func (o Optional[T]) Unwrap() T {
	if o.value == nil {
		panic("called Unwrap on None value")
	}
	return *o.value
}

// Value - alias for Unwrap
func (o Optional[T]) Value() T {
	return o.Unwrap()
}

// UnwrapOr returns the underlying value or a default value if None.
func (o Optional[T]) UnwrapOr(def T) T {
	if o.value == nil {
		return def
	}
	return *o.value
}
