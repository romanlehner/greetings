package greetings

import "testing"

func TestHelloWorldEnglish(t *testing.T) {
	want := "Hello World!"
	if got := HelloWorldEnglish(); got != want {
		t.Errorf("HelloWorld() = %q, want %q", got, want)
	}
}

func TestHelloWorldJapanese(t *testing.T) {
	want := "こんにちは 世界!"
	if got := HelloWorldJapanese(); got != want {
		t.Errorf("HelloWorldJapanese() = %q, want %q", got, want)
	}
}

func TestHelloWorldGerman(t *testing.T) {
	want := "Hallo Welt!"
	if got := HelloWorldGerman(); got != want {
		t.Errorf("HelloWorldGerman() = %q, want %q", got, want)
	}
}
