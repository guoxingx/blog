package webserver

import (
	"fmt"
	"testing"
)

func TestSignPassword(t *testing.T) {
	usernames := []string{
		"enigma",
		"hostess",
	}

	passwords := []string{
		"hidden",
		"hidden",
	}

	signed := []string{
		"05fc0490ef490761667ec82a4af82f4d",
		"9e91daf26edfad77f67ac347b4c0932e",
	}

	for i := range usernames {
		u := usernames[i]
		p := passwords[i]
		s := signPassword(u, p)
		if s != signed[i] {
			fmt.Printf("%d. signed: %s, should be: %s\n", i, s, signed[i])
			t.Fail()
		}
	}
}
