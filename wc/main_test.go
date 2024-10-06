package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T){
	b := bytes.NewBufferString("Hello World Hello Mundo Bonjour Monde")

	expected := 6

	results := count(b, false, false)

	if results != expected {
		t.Errorf("Expected %d, got %d", expected,results)
	}

}

func TestCountLines(t *testing.T){
	b := bytes.NewBufferString("Hello World\nHello Mundo\nBonjour Monde")

	expected := 3

	results := count(b, true, false)

	if results != expected {
		t.Errorf("Expected %d, got %d", expected,results)
	}

}

func TestCountBytes(t *testing.T){
	b := bytes.NewBufferString("Hello World Hello Mundo Bonjour Monde")

	expected := 37

	results := count(b, false, true)

	if results != expected {
		t.Errorf("Expected %d, got %d", expected,results)
	}

}