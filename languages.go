package lexer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Southern/scanner"
)

func makeKeywords(keys []string) [][][]string {
	k := make([][][]string, len(keys))
	for i, key := range keys {
		k[i] = [][]string{
			{"WORD", key},
			{"KEYWORD"},
		}
	}
	return k
}

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

				// Hexadecimal
				regex["numbers"]["hex"],

				// ~, ===, .
				{regexp.MustCompile("^(~|={3}|\\.)"), "OPERATOR"},

				// Operators
				regex["operators"]["common"],

				{regexp.MustCompile(
					fmt.Sprintf("^(?i)([a-z0-9_][a-z0-9'_%s]+|[%s]{2,})",
						scanner.Unicode(), scanner.Unicode()),
				), "WORD"},
			}, scanner.Map()...),
			Modify: append([][][]string{
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
			}, makeKeywords([]string{
				"Object",
				"Function",
				"function",
				"Boolean",
				"Error",
				"EvalError",
				"InternalError",
				"RangeError",
				"ReferenceError",
				"SyntaxError",
				"TypeError",
				"URIError",
				"Number",
				"Math",
				"Date",
				"String",
				"RegExp",
				"Array",
				"U?Int8Array",
				"UInt8ClampedArray",
				"UInt16Array",
				"Int16Array",
				"UInt32Array",
				"Int32Array",
				"Float32Array",
				"Float64Array",
				"ArrayBuffer",
				"DataView",
				"JSON",
				"Infinity",
				"NaN",
				"undefined",
				"null",
				"__proto__",
				"prototype",
				"constructor",
				"new",
				"true",
				"false",
				"for",
				"while",
				"setTimeout",
				"clearTimeout",
				"setInterval",
				"clearInterval",
				"if",
				"else",
				"continue",
				"break",
				"return",
			})...),
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

				// Hexadecimal
				regex["numbers"]["hex"],

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
			Modify: append([][][]string{
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
			}, makeKeywords([]string{
				"ComplexType",
				"FloatType",
				"IntegerType",
				"Type1",
				"bool",
				"byte",
				"break",
				"continue",
				"complex64",
				"complex128",
				"error",
				"float32",
				"float64",
				"string",
				"int8",
				"int16",
				"int32",
				"int64",
				"uint8",
				"uint16",
				"uint32",
				"uint64",
				"uintptr",
				"true",
				"false",
				"iota",
				"func",
				"type",
				"struct",
				"chan",
				"for",
				"if",
				"else",
				"map",
				"switch",
				"type",
				"case",
				"return",
			})...),
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

				// Hexadecimal
				regex["numbers"]["hex"],

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
			Modify: makeKeywords([]string{
				"and",
				"as",
				"assert",
				"break",
				"class",
				"continue",
				"del",
				"def",
				"else",
				"elif",
				"except",
				"exec",
				"finally",
				"for",
				"from",
				"global",
				"if",
				"import",
				"in",
				"is",
				"lambda",
				"not",
				"or",
				"pass",
				"print",
				"raise",
				"return",
				"try",
				"while",
				"with",
				"yield",
				"None",
			}),
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

				// Hexadecimal
				regex["numbers"]["hex"],

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
			Modify: append([][][]string{
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
			}, makeKeywords([]string{
				"abstract",
				"assert",
				"boolean",
				"break",
				"byte",
				"case",
				"catch",
				"char",
				"class",
				"const",
				"continue",
				"default",
				"do",
				"double",
				"else",
				"enum",
				"extends",
				"final",
				"finally",
				"float",
				"for",
				"goto",
				"if",
				"implements",
				"import",
				"instanceof",
				"int",
				"interface",
				"long",
				"native",
				"new",
				"package",
				"private",
				"protected",
				"public",
				"return",
				"short",
				"static",
				"strictfp",
				"super",
				"switch",
				"syncrhonized",
				"this",
				"throw",
				"throws",
				"transient",
				"try",
				"void",
				"volatile",
				"while",
				"false",
				"null",
				"true",
			})...),
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

				// Hexadecimal
				regex["numbers"]["hex"],

				// :symbol
				{regexp.MustCompile("^(?i):[a-z_][a-z0-9_]+"), "SYMBOL"},

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
			Modify: append([][][]string{
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
			}, makeKeywords([]string{
				"BEGIN",
				"END",
				"TRUE",
				"FALSE",
				"NIL",
				"ARGF",
				"ARGV",
				"DATA",
				"ENV",
				"RUBY_PLATFORM",
				"RUBY_RELEASE_DATE",
				"RUBY_VERSION",
				"STDERR",
				"STDIN",
				"STDOUT",
				"TOPLEVEL_BINDING",
				"alias",
				"and",
				"abort",
				"begin",
				"break",
				"case",
				"class",
				"def",
				"defined?",
				"do",
				"else",
				"elsif",
				"end",
				"ensure",
				"exit",
				"false",
				"for",
				"if",
				"in",
				"module",
				"next",
				"nil",
				"not",
				"or",
				"print",
				"puts",
				"raise",
				"redo",
				"rescue",
				"retry",
				"return",
				"self",
				"super",
				"then",
				"trap",
				"true",
				"undef",
				"unless",
				"until",
				"when",
				"while",
				"__FILE__",
				"__LINE__",
			})...),
		},
	}

	Languages["Node"] = Languages["Javascript"]
}
