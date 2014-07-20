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
				"[fF]unction",
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
				"U?Int16Array",
				"(U?Int|Float)32Array",
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
				"(set|clear)Timeout",
				"(set|clear)Interval",
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
				"(Complex|Float|Integer)?Type",
				"Type1",
				"bool",
				"byte",
				"break",
				"continue",
				"complex(64|128)",
				"error",
				"float(32|64)",
				"string",
				"u?int(8|16|32|64)?",
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
				"a(nd|s)",
				"assert",
				"break",
				"class",
				"continue",
				"de[lf]",
				"el(se|if)",
				"ex(ec|cept)",
				"finally",
				"f?or",
				"from",
				"global",
				"if",
				"import",
				"i[ns]",
				"lambda",
				"not",
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
				"ca(se|tch)",
				"char",
				"class",
				"con(st|tinue)",
				"default",
				"do(uble)?",
				"else",
				"enum",
				"extends",
				"final(ly)?",
				"float",
				"for",
				"goto",
				"if",
				"im(plements|port)",
				"instanceof",
				"int(erface)?",
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
				"th(is|rows?)",
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
				"RUBY_(PLATFORM|RELEASE_DATE|VERSION)",
				"STD(ERR|IN|OUT)",
				"TOPLEVEL_BINDING",
				"a(lias|nd)",
				"abort",
				"begin",
				"break",
				"case",
				"class",
				"def(ined\\?)?",
				"do",
				"el(se|sif)",
				"en(d|sure)",
				"exit",
				"false",
				"for",
				"i[fn]",
				"module",
				"next",
				"nil",
				"not",
				"or",
				"print",
				"puts",
				"raise",
				"re(do|scue|try|turn)",
				"self",
				"super",
				"then",
				"trap",
				"true",
				"un(def|less|til)",
				"wh(en|ile)",
				"__FILE__",
				"__LINE__",
			})...),
		},
	}

	Languages["Node"] = Languages["Javascript"]
}
