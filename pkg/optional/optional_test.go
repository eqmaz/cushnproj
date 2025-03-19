package optional

import (
	"testing"
)

func TestOptional(t *testing.T) {
	t.Run("TestSomeAndNone", func(t *testing.T) {
		// Testing with int type.
		optInt := Some(42)
		if !optInt.IsSome() {
			t.Error("Some value should return true for IsSome")
		}
		if optInt.IsNone() {
			t.Error("Some value should return false for IsNone")
		}
		if val := optInt.Unwrap(); val != 42 {
			t.Errorf("Expected Unwrap to return 42, got %d", val)
		}
		if val := optInt.UnwrapOr(100); val != 42 {
			t.Errorf("Expected UnwrapOr to return 42, got %d", val)
		}

		optIntNone := None[int]()
		if optIntNone.IsSome() {
			t.Error("None should return false for IsSome")
		}
		if !optIntNone.IsNone() {
			t.Error("None should return true for IsNone")
		}
		if val := optIntNone.UnwrapOr(100); val != 100 {
			t.Errorf("Expected UnwrapOr to return default 100, got %d", val)
		}
	})

	t.Run("TestUnwrapPanics", func(t *testing.T) {
		// Ensure that Unwrap panics when called on a None.
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Unwrap to panic when called on None")
			}
		}()
		_ = None[string]().Unwrap()
	})

	t.Run("TestValueAlias", func(t *testing.T) {
		// Value() should behave as an alias to Unwrap.
		optStr := Some("test")
		if val := optStr.Value(); val != "test" {
			t.Errorf("Expected Value() to return 'test', got %s", val)
		}
	})

	t.Run("TestOptionalWithStruct", func(t *testing.T) {
		// Testing Optional with a custom struct type.
		type User struct {
			Name string
			Age  int
		}
		user := User{Name: "Alice", Age: 30}
		optUser := Some(user)
		if !optUser.IsSome() {
			t.Error("Some value should return true for IsSome")
		}
		if got := optUser.Unwrap(); got != user {
			t.Errorf("Expected Unwrap to return %+v, got %+v", user, got)
		}

		defaultUser := User{Name: "Default", Age: 0}
		if got := None[User]().UnwrapOr(defaultUser); got != defaultUser {
			t.Errorf("Expected UnwrapOr to return default %+v, got %+v", defaultUser, got)
		}
	})
}
