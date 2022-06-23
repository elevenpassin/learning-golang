package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want ""`, msg, err)
	}
}

func TestHellosName(t *testing.T) {
	names := []string{"Gladys", "Samantha"}

	msgs, err := Hellos(names)

	for _, msg := range msgs {
		want := regexp.MustCompile(msg)
		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hellos([]string{"Gladys", "Samantha"}) = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	}

}
