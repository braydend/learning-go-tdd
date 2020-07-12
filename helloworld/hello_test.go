package main

import "testing"

func assertSameString(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected: %s. Received: %s", expected, actual)
	}
}

func TestHello(t *testing.T) {
	t.Run("test hello with names", func(t *testing.T) {
		expected := "Hello Bray!"
		result := Hello("Bray", "")

		assertSameString(t, result, expected)
	})

	t.Run("test hello world when empty string passed", func(t *testing.T) {
		expected := "Hello world!"
		result := Hello("", "")

		assertSameString(t, result, expected)
	})

	t.Run("In spanish", func(t *testing.T) {
		expected := "Hola Bray!"
		result := Hello("Bray", "Spanish")

		assertSameString(t, result, expected)
	})

	t.Run("In french", func(t *testing.T) {
		expected := "Bonjour Bray!"
		result := Hello("Bray", "French")

		assertSameString(t, result, expected)
	})
}
