package fmtp

import (
	"reflect"
	"testing"
)

func TestRedParseFmtp(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected FMTP
	}{
		"OneParam": {
			input: "111",
			expected: &redFMTP{
				parameters: []string{
					"111",
				},
			},
		},
		"OneParamWithWhiteSpeces": {
			input: "\t111 ",
			expected: &redFMTP{
				parameters: []string{
					"111",
				},
			},
		},
		"TwoParams": {
			input: "111/112",
			expected: &redFMTP{
				parameters: []string{
					"111",
					"112",
				},
			},
		},
		"TwoParamsWithWhiteSpeces": {
			input: "\n\t111/112 ",
			expected: &redFMTP{
				parameters: []string{
					"111",
					"112",
				},
			},
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			f := Parse("audio/red", testCase.input)
			if !reflect.DeepEqual(testCase.expected, f) {
				t.Errorf("Expected Fmtp params: %v, got: %v", testCase.expected, f)
			}

			if f.MimeType() != "audio/red" {
				t.Errorf("Expected MimeType of red, got: %s", f.MimeType())
			}
		})
	}
}

func TestRedFmtpCompare(t *testing.T) {
	consistString := map[bool]string{true: "consist", false: "inconsist"}

	testCases := map[string]struct {
		a, b    string
		consist bool
	}{
		"Equal": {
			a:       "111/111",
			b:       "111/111",
			consist: true,
		},
		"EqualWithWhitespaceVariants": {
			a:       "111/111",
			b:       "  \n 111/111\t\n",
			consist: true,
		},
		"Inconsistent": {
			a:       "111/112",
			b:       "112/112",
			consist: false,
		},
		"Inconsistent_OneHasExtraParam": {
			a:       "111",
			b:       "111/111",
			consist: false,
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		check := func(t *testing.T, a, b string) {
			aa := Parse("audio/red", a)
			bb := Parse("audio/red", b)
			c := aa.Match(bb)
			if c != testCase.consist {
				t.Errorf(
					"'%s' and '%s' are expected to be %s, but treated as %s",
					a, b, consistString[testCase.consist], consistString[c],
				)
			}

			// test reverse case here
			c = bb.Match(aa)
			if c != testCase.consist {
				t.Errorf(
					"'%s' and '%s' are expected to be %s, but treated as %s",
					a, b, consistString[testCase.consist], consistString[c],
				)
			}
		}
		t.Run(name, func(t *testing.T) {
			check(t, testCase.a, testCase.b)
		})
	}
}
