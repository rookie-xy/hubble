package scanner_test

import (
	"errors"
	"io"
	"strings"
	"testing"
	"github.com/rookie-xy/hubble/scanner"
)

const smallMaxTokenSize = 256 // Much smaller for more efficient testing.


var testError = errors.New("testError")

// Test the correct error is returned when the split function errors out.
func TestSplitError(t *testing.T) {
	// Create a split function that delivers a little data, then a predictable error.
	numSplits := 0
	const okCount = 7
	errorSplit := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF {
			panic("didn't get enough data")
		}
		if numSplits >= okCount {
			return 0, nil, testError
		}
		numSplits++
		return 1, data[0:1], nil
	}
	// Read the data.
	const text = "abcdefghijklmnopqrstuvwxyz"
	buf := strings.NewReader(text)
	s := scanner.New(&slowReader{1, buf})
	s.Split(errorSplit)
	var i int
	for i = 0; s.Scan(); i++ {
		if len(s.Bytes()) != 1 || text[i] != s.Bytes()[0] {
			t.Errorf("#%d: expected %q got %q", i, text[i], s.Bytes()[0])
		}
	}
	// Check correct termination location and error.
	if i != okCount {
		t.Errorf("unexpected termination; expected %d tokens got %d", okCount, i)
	}
	err := s.Err()
	if err != testError {
		t.Fatalf("expected %q got %v", testError, err)
	}
}

// Test that an EOF is overridden by a user-generated scan error.
func TestErrAtEOF(t *testing.T) {
	s := scanner.New(strings.NewReader("1 2 33"))
	// This splitter will fail on last entry, after s.err==EOF.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = ScanWords(data, atEOF)
		if len(token) > 1 {
			if s.ErrOrEOF() != io.EOF {
				t.Fatal("not testing EOF")
			}
			err = testError
		}
		return
	}
	s.Split(split)
	for s.Scan() {
	}
	if s.Err() != testError {
		t.Fatal("wrong error:", s.Err())
	}
}

// Test for issue 5268.
type alwaysError struct{}

func (alwaysError) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestNonEOFWithEmptyRead(t *testing.T) {
	scanner := scanner.New(alwaysError{})
	for scanner.Scan() {
		t.Fatal("read should fail")
	}
	err := scanner.Err()
	if err != io.ErrUnexpectedEOF {
		t.Errorf("unexpected error: %v", err)
	}
}

// Test that Scan finishes if we have endless empty reads.
type endlessZeros struct{}

func (endlessZeros) Read(p []byte) (int, error) {
	return 0, nil
}

func TestBadReader(t *testing.T) {
	scanner := scanner.New(endlessZeros{})
	for scanner.Scan() {
		t.Fatal("read should fail")
	}
	err := scanner.Err()
	if err != io.ErrNoProgress {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestScanWordsExcessiveWhiteSpace(t *testing.T) {
	const word = "ipsum"
	s := strings.Repeat(" ", 4*smallMaxTokenSize) + word
	scanner := scanner.New(strings.NewReader(s))
	scanner.MaxTokenSize(smallMaxTokenSize)
	scanner.Split(ScanWords)
	if !scanner.Scan() {
		t.Fatalf("scan failed: %v", scanner.Err())
	}
	if token := scanner.Text(); token != word {
		t.Fatalf("unexpected token: %v", token)
	}
}

// Test that empty tokens, including at end of line or end of file, are found by the scanner.
// Issue 8672: Could miss final empty token.

func commaSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	return 0, data, ErrFinalToken
}

func testEmptyTokens(t *testing.T, text string, values []string) {
	s := scanner.New(strings.NewReader(text))
	s.Split(commaSplit)
	var i int
	for i = 0; s.Scan(); i++ {
		if i >= len(values) {
			t.Fatalf("got %d fields, expected %d", i+1, len(values))
		}
		if s.Text() != values[i] {
			t.Errorf("%d: expected %q got %q", i, values[i], s.Text())
		}
	}
	if i != len(values) {
		t.Fatalf("got %d fields, expected %d", i, len(values))
	}
	if err := s.Err(); err != nil {
		t.Fatal(err)
	}
}

func TestEmptyTokens(t *testing.T) {
	testEmptyTokens(t, "1,2,3,", []string{"1", "2", "3", ""})
}

func TestWithNoEmptyTokens(t *testing.T) {
	testEmptyTokens(t, "1,2,3", []string{"1", "2", "3"})
}

func loopAtEOFSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) > 0 {
		return 1, data[:1], nil
	}
	return 0, data, nil
}

func TestDontLoopForever(t *testing.T) {
	s := scanner.New(strings.NewReader("abc"))
	s.Split(loopAtEOFSplit)
	// Expect a panic
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("should have panicked")
		}
		if msg, ok := err.(string); !ok || !strings.Contains(msg, "empty tokens") {
			panic(err)
		}
	}()
	for count := 0; s.Scan(); count++ {
		if count > 1000 {
			t.Fatal("looping")
		}
	}
	if s.Err() != nil {
		t.Fatal("after scan:", s.Err())
	}
}

func TestBlankLines(t *testing.T) {
	s := scanner.New(strings.NewReader(strings.Repeat("\n", 1000)))
	for count := 0; s.Scan(); count++ {
		if count > 2000 {
			t.Fatal("looping")
		}
	}
	if s.Err() != nil {
		t.Fatal("after scan:", s.Err())
	}
}

type countdown int

func (c *countdown) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if *c > 0 {
		*c--
		return 1, data[:1], nil
	}
	return 0, nil, nil
}

// Check that the looping-at-EOF check doesn't trigger for merely empty tokens.
func TestEmptyLinesOK(t *testing.T) {
	c := countdown(10000)
	s := scanner.New(strings.NewReader(strings.Repeat("\n", 10000)))
	s.Split(c.split)
	for s.Scan() {
	}
	if s.Err() != nil {
		t.Fatal("after scan:", s.Err())
	}
	if c != 0 {
		t.Fatalf("stopped with %d left to process", c)
	}
}

// Make sure we can read a huge token if a big enough buffer is provided.
func TestHugeBuffer(t *testing.T) {
	text := strings.Repeat("x", 2*MaxScanTokenSize)
	s := scanner.New(strings.NewReader(text + "\n"))
	s.Buffer(make([]byte, 100), 3*MaxScanTokenSize)
	for s.Scan() {
		token := s.Text()
		if token != text {
			t.Errorf("scan got incorrect token of length %d", len(token))
		}
	}
	if s.Err() != nil {
		t.Fatal("after scan:", s.Err())
	}
}
