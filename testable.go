package testable

import (
	"go/ast"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
)

// Normalize the actual example output before comparison. This is necessary when the
// actual example output contains whitespace characters that would be removed in the expected
// example output.
//
// The consequence is that example output will be matched leniently.
func Normalize(buf []byte) []byte {
	// remove \r as documented for ast.Comment
	str := string(buf)
	str = strings.ReplaceAll(str, "\r", "")

	cg := ast.CommentGroup{}
	for _, line := range strings.Split(str, "\n") {
		cg.List = append(cg.List, &ast.Comment{
			Text: "// " + strings.TrimSpace(line), // Trim \f, because trimming in CommentGroup.Text() doesn't
		})
	}
	return []byte(cg.Text())
}

// CaptureStdout replaces os.Stdout to capture all output.
// The returned function will restore os.Stdout and return the captured output.
func CaptureStdout() func() []byte {
	tmp, err := ioutil.TempFile("", "CaptureStdout")
	if err != nil {
		panic(err)
	}

	if err = syscall.Unlink(tmp.Name()); err != nil {
		panic(err)
	}

	out := os.Stdout
	os.Stdout = tmp

	return func() []byte {
		defer func() {
			os.Stdout = out
			tmp.Close()
		}()

		if _, err := tmp.Seek(0, 0); err != nil {
			panic(err)
		}

		buf, err := ioutil.ReadAll(tmp)
		if err != nil {
			panic(err)
		}

		return buf
	}
}

// FixExampleOutput is a convenience function to capture stdout and make it testable
func FixExampleOutput() func() {
	output := CaptureStdout()
	return func() {
		os.Stdout.Write(Normalize(output()))
	}
}
