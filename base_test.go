package base

import "testing"

func TestFrombase(t *testing.T) {
	got := Frombase("101", 2)
	exp := 5
	if got != exp {
		t.Fatalf("got %v, expected %v\n", got, exp)
	}

	got = Frombase("3A57", 12)
	exp = 6691
	if got != exp {
		t.Fatalf("got %v, expected %v\n", got, exp)
	}
}

func TestTobase(t *testing.T) {
	got := Tobase(5, 2)
	exp := "101"
	if got != exp {
		t.Fatalf("got %v, expected %v\n", got, exp)
	}

	got = Tobase(6691, 12)
	exp = "3A57"
	if got != exp {
		t.Fatalf("got %v, expected %v\n", got, exp)
	}
}

func TestBoth(t *testing.T) {
	got := Tobase(Frombase("3A57", 12), 6)
	exp := "50551"
	if got != exp {
		t.Fatalf("got %v, expected %v\n", got, exp)
	}
}
