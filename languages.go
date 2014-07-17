package lexer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Southern/scanner"
)

func init() {
	Languages = map[string]*Language{
		"Javascript": {
			Extensions: []string{"js"},
			Map: append([]scanner.Definition{
				// Single line comments
				regex["comments"]["oneline"],

				// Multi-line comments
				regex["comments"]["multiline"],

				// Double quote strings
				regex["strings"]["double"],

				// Single quote strings
				regex["strings"]["single"],

				// ~, ===, .
				{regexp.MustCompile("^(~|={3}|\\.)"), "OPERATOR"},

				// Operators
				regex["operators"]["common"],

				{regexp.MustCompile(
					fmt.Sprintf("^(?i)([a-z0-9_][a-z0-9'_%s]+|[%s]{2,})",
						scanner.Unicode(), scanner.Unicode()),
				), "WORD"},
			}, scanner.Map()...),
			Modify: [][][]string{
				{
					{"CHAR", "{"},
					{"BLOCKSTART"},
				},
				{
					{"CHAR", "}"},
					{"BLOCKEND"},
				},
				{
					{"CHAR", "("},
					{"ARGSTART"},
				},
				{
					{"CHAR", ")"},
					{"ARGEND"},
				},
				{
					{"CHAR", "["},
					{"ARRAYSTART"},
				},
				{
					{"CHAR", "]"},
					{"ARRAYEND"},
				},
				{
					{"CHAR", ";"},
					{"END"},
				},
			},
		},

		"Go": {
			Extensions: []string{"go"},
			Map: append([]scanner.Definition{
				// Single line comments
				regex["comments"]["oneline"],

				// Multi-line comments
				regex["comments"]["multiline"],

				// Double quote strings
				regex["strings"]["double"],

				// Single quote strings, for byte and char types.
				regex["strings"]["single"],

				// <-, :=, :, ..., .
				{regexp.MustCompile("^(<-|:=?|\\.{3}|\\.)"), "OPERATOR"},

				// &&, &^, &^=
				{regexp.MustCompile("^(&{2}|(&\\^?)=?)"), "OPERATOR"},

				// Common operators
				regex["operators"]["common"],

				{regexp.MustCompile(
					fmt.Sprintf("^(?i)([a-z0-9_][a-z0-9'_%s]+|[%s]{2,})",
						scanner.Unicode(), scanner.Unicode()),
				), "WORD"},
			}, scanner.Map()...),
			Modify: [][][]string{
				{
					{"CHAR", "{"},
					{"BLOCKSTART"},
				},
				{
					{"CHAR", "}"},
					{"BLOCKEND"},
				},
				{
					{"CHAR", "("},
					{"ARGSTART"},
				},
				{
					{"CHAR", ")"},
					{"ARGEND"},
				},
			},
		},

		"Python": {
			Extensions: []string{"py"},
			Map: append([]scanner.Definition{
				// Comments
				{regexp.MustCompile("^#.+"), "COMMENT"},

				// Heredoc
				{regexp.MustCompile(`^"{3}(\"?[^"])*"{3}`), "DOCSTRING"},
				{regexp.MustCompile(`^'{3}(\'?[^'])*'{3}`), "DOCSTRING"},

				// Double quote strings
				regex["strings"]["double"],

				// Single quote strings
				regex["strings"]["single"],

				// Decorators
				{regexp.MustCompile("^@[^\\s]+"), "DECORATOR"},

				// //, **
				{regexp.MustCompile("^([*/]{2})"), "OPERATOR"},

				// Common operators
				regex["operators"]["common"],

				{regexp.MustCompile(
					fmt.Sprintf("^(?i)([a-z0-9_][a-z0-9'_%s]+|[%s]{2,})",
						scanner.Unicode(), scanner.Unicode()),
				), "WORD"},
			}, scanner.Map()...),
		},

		"Java": {
			Extensions: []string{"java"},
			Map: append([]scanner.Definition{
				// Single line comments
				regex["comments"]["oneline"],

				// Multi-line comments
				regex["comments"]["multiline"],

				// Double quote strings
				regex["strings"]["double"],

				// Single quote "string"
				regex["strings"]["single"],

				// Decorators
				{regexp.MustCompile("^@[^\\s]+"), "DECORATOR"},

				// ~, ?:, .
				{regexp.MustCompile("^(~|\\?:|\\.)"), "OPERATOR"},

				// Common operators
				regex["operators"]["common"],

				{regexp.MustCompile(
					fmt.Sprintf("^(?i)([a-z0-9_][a-z0-9'_%s]+|[%s]{2,})",
						scanner.Unicode(), scanner.Unicode()),
				), "WORD"},
			}, scanner.Map()...),
			Modify: [][][]string{
				{
					{"CHAR", "{"},
					{"BLOCKSTART"},
				},
				{
					{"CHAR", "}"},
					{"BLOCKEND"},
				},
				{
					{"CHAR", "("},
					{"ARGSTART"},
				},
				{
					{"CHAR", ")"},
					{"ARGEND"},
				},
				{
					{"CHAR", ";"},
					{"END"},
				},
			},
		},

		"Ruby": {
			Extensions: []string{"rb"},
			Map: append([]scanner.Definition{
				// Comments
				{regexp.MustCompile("^#.+"), "COMMENT"},

				// Double quote strings
				regex["strings"]["double"],

				// Single quote "string"
				regex["strings"]["single"],

				// ::, .., ..., ., <=>, ===, =~, !~
				{regexp.MustCompile("^(:{2}|\\.{2,3}|\\.|<=>|={3}|=[~>]|!~)"), "OPERATOR"},

				// Common operators
				regex["operators"]["common"],

				{regexp.MustCompile(
					fmt.Sprintf("^(?i)([a-z0-9_][a-z0-9'_%s]+|[%s]{2,})",
						scanner.Unicode(), scanner.Unicode()),
				), "WORD"},

				// Restricted Ruby variables
				{regexp.MustCompile(
					strings.Join([]string{
						"^\\$(",
						strings.Join([]string{
							"[!@\\/\\\\,;.<>0$?:&`'+_~]",
							"DEBUG",
							"defout",
							"F(ILENAME)?",
							"LOAD_PATH",
							"SAFE",
							"std(in|out|err)",
							"VERBOSE",
							"-[0adFiIlp]",
							"[1-9]+",
						}, "|"),
						")",
					}, "") + "\\b",
				), "IDENT"},
			}, scanner.Map()...),
			Modify: [][][]string{
				{
					{"CHAR", "{"},
					{"BLOCKSTART"},
				},
				{
					{"CHAR", "}"},
					{"BLOCKEND"},
				},
				{
					{"CHAR", "("},
					{"ARGSTART"},
				},
				{
					{"CHAR", ")"},
					{"ARGEND"},
				},
				{
					{"CHAR", "["},
					{"ARRAYSTART"},
				},
				{
					{"CHAR", "]"},
					{"ARRAYEND"},
				},
			},
		},
	}

	Languages["Node"] = Languages["Javascript"]
}
