package weirdstringcase

import "testing"

func TestWeirdString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"abc def", "AbC DeF"},
		{"This11 is a test Looks like you passed", "ThIs11 Is A TeSt LoOkS LiKe YoU PaSsEd"},
		{"", ""},
	}

	for _, c := range cases {
		got := WeirdString(c.in)

		if got != c.want {
			t.Errorf("WeirdString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
