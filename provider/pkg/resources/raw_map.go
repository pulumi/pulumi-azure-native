package resources

import (
	"fmt"
	"unsafe"

	"github.com/segmentio/encoding/json"
)

type PartialMap[T any] struct {
	partialMap map[string]json.RawMessage
}

func NewPartialMap[T any]() PartialMap[T] {
	return PartialMap[T]{
		partialMap: make(map[string]json.RawMessage),
	}
}

func (m *PartialMap[T]) UnmarshalJSON(b []byte) error {
	if err := unmarshal(b, &m.partialMap); err != nil {
		return err
	}
	return nil
}

func (m PartialMap[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.partialMap)
}

func (m *PartialMap[T]) Get(key string) (T, bool, error) {
	rawMessage, ok := m.partialMap[key]
	if !ok {
		return *new(T), true, nil
	}
	var value T
	if _, err := json.Parse([]byte(rawMessage), &value, json.ZeroCopy); err != nil {
		return *new(T), false, err
	}

	return value, true, nil
}

//// via github.com/segmentio, MIT licensed:

func unmarshal(b []byte, x interface{}) error {
	r, err := json.Parse(b, x, json.ZeroCopy)
	if len(r) != 0 {
		if _, ok := err.(*json.SyntaxError); !ok {
			// The encoding/json package prioritizes reporting errors caused by
			// unexpected trailing bytes over other issues; here we emulate this
			// behavior by overriding the error.
			err = syntaxError(r, "invalid character '%c' after top-level value", r[0])
		}
	}
	return err
}

func prefix(b []byte) string {
	if len(b) < 32 {
		return string(b)
	}
	return string(b[:32]) + "..."
}

var syntaxErrorMsgOffset = ^uintptr(0)

func syntaxError(b []byte, msg string, args ...interface{}) error {
	e := new(json.SyntaxError)
	i := syntaxErrorMsgOffset
	if i != ^uintptr(0) {
		s := "json: " + fmt.Sprintf(msg, args...) + ": " + prefix(b)
		p := unsafe.Pointer(e)
		// Hack to set the unexported `msg` field.
		*(*string)(unsafe.Pointer(uintptr(p) + i)) = s
	}
	return e
}
