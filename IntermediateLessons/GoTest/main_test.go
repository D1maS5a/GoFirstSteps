package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com"
	if !IsValudEmail(data) {
		t.Errorf("IsValidEmail(%v) = false, want true ", data)
	}
}

func TestIsValidEmailTable(t *testing.T) {
	table := []struct {
		email string
		want  bool
	}{
		{"missing@tld", true},
		{"email@example.com", false},
	}

	for _, data := range table {
		result := IsValudEmail(data.email)
		if result != data.want {
			t.Errorf("%v: %t, want: %t", data.email, result, data.want)
		}
	}
}
