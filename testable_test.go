package testable

import (
	"fmt"
	"os"
)

func ExampleDoubleEmptyLines() {
	capture := CaptureStdout()

	fmt.Println("a\n\n\nb")

	os.Stdout.Write(Normalize(capture()))
	// Output:
	// a
	//
	//
	// b
}

func ExampleCRLF() {
	capture := CaptureStdout()

	fmt.Println("a\r\nb")

	os.Stdout.Write(Normalize(capture()))
	// Output:
	// a
	// b
}

func ExampleTrailingWhitespace() {
	capture := CaptureStdout()

	fmt.Println("a\f\t \nb")

	os.Stdout.Write(Normalize(capture()))
	// Output:
	// a
	// b
}

func ExampleCarriageReturn() {
	capture := CaptureStdout()

	fmt.Println("aaa\rbbb")

	os.Stdout.Write(Normalize(capture()))
	// Output:
	// aaabbb
}

func ExampleFixExampleOutput() {
	done := FixExampleOutput()
	defer done()

	fmt.Println("aaa\rbbb")
	// Output:
	// aaabbb
}
