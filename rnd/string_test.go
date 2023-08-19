package rnd

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	length := 10
	chars := "abc"
	result := RandomString(length, chars)
	if len(result) != length {
		t.Errorf("RandomString(%d, %s) = %s, want length %d", length, chars, result, length)
	}

	for _, char := range result {
		if !strings.Contains(chars, string(char)) {
			t.Errorf("RandomString(%d, %s) = %s, want only chars from %s", length, chars, result, chars)
		}
	}
}

func TestRandomLowerAlpha(t *testing.T) {
	length := 10
	result := RandomLowerAlpha(length)
	if len(result) != length {
		t.Errorf("RandomLowerAlpha(%d) = %s, want length %d", length, result, length)
	}

	for _, char := range result {
		if !strings.Contains(LowerAlpha, string(char)) {
			t.Errorf("RandomLowerAlpha(%d) = %s, want only chars from %s", length, result, LowerAlpha)
		}
	}
}

func TestRandomUpperAlpha(t *testing.T) {
	length := 10
	result := RandomUpperAlpha(length)
	if len(result) != length {
		t.Errorf("RandomUpperAlpha(%d) = %s, want length %d", length, result, length)
	}

	for _, char := range result {
		if !strings.Contains(UpperAlpha, string(char)) {
			t.Errorf("RandomUpperAlpha(%d) = %s, want only chars from %s", length, result, UpperAlpha)
		}
	}
}

func TestRandomAlpha(t *testing.T) {
	length := 10
	result := RandomAlpha(length)
	if len(result) != length {
		t.Errorf("RandomAlpha(%d) = %s, want length %d", length, result, length)
	}

	for _, char := range result {
		if !strings.Contains(LowerAlpha+UpperAlpha, string(char)) {
			t.Errorf("RandomAlpha(%d) = %s, want only chars from %s", length, result, LowerAlpha+UpperAlpha)
		}
	}
}

func TestRandomLowerAlphaNumeric(t *testing.T) {
	length := 10
	result := RandomLowerAlphaNumeric(length)
	if len(result) != length {
		t.Errorf("RandomLowerAlphaNumeric(%d) = %s, want length %d", length, result, length)
	}

	for _, char := range result {
		if !strings.Contains(LowerAlpha+Numbers, string(char)) {
			t.Errorf("RandomLowerAlphaNumeric(%d) = %s, want only chars from %s", length, result, LowerAlpha+Numbers)
		}
	}
}

func TestRandomUpperAlphaNumeric(t *testing.T) {
	length := 10
	result := RandomUpperAlphaNumeric(length)
	if len(result) != length {
		t.Errorf("RandomUpperAlphaNumeric(%d) = %s, want length %d", length, result, length)
	}

	for _, char := range result {
		if !strings.Contains(UpperAlpha+Numbers, string(char)) {
			t.Errorf("RandomUpperAlphaNumeric(%d) = %s, want only chars from %s", length, result, UpperAlpha+Numbers)
		}
	}
}

func TestRandomAlphaNumeric(t *testing.T) {
	length := 10
	result := RandomAlphaNumeric(length)
	if len(result) != length {
		t.Errorf("RandomAlphaNumeric(%d) = %s, want length %d", length, result, length)
	}

	for _, char := range result {
		if !strings.Contains(LowerAlpha+UpperAlpha+Numbers, string(char)) {
			t.Errorf("RandomAlphaNumeric(%d) = %s, want only chars from %s", length, result, LowerAlpha+UpperAlpha+Numbers)
		}
	}
}
