package gen

import "testing"

func TestToLowerCamel(t *testing.T) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
		{"AnyKind.of-string", "anyKindOfString"},
		{"ID", "id"},
		{"TTL", "ttl"},
		{"_ref", "ref"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := toLowerCamel(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
