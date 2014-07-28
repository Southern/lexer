package lexer

import (
	"github.com/Southern/scanner"
	"regexp"
	"strings"
)

var regex = map[string]map[string]scanner.Definition{
	"numbers": {
		"hex": {regexp.MustCompile("^(?i)0x[a-f0-9]+\\b"), "HEX"},
	},

	"comments": {
		// Hash comments
		"hash": {regexp.MustCompile("^#(.+)"), "COMMENT"},

		// Single line comments
		"oneline": {regexp.MustCompile("^\\/{2,}\\s*(.+)"), "COMMENT"},

		// Multi-line comments
		"multiline": {regexp.MustCompile("^/\\*([^*]|[\r\n]|(\\*+([^*/]|[\r\n])))*\\*+/"), "COMMENT"},
	},

	"strings": {
		// Double quote strings
		"double": {regexp.MustCompile("^\"([^\"\\n]|\\\")*\""), "STRING"},

		// Single quote strings
		"single": {regexp.MustCompile("^'([^'\\n]|\\')*'"), "STRING"},
	},

	"operators": {
		// Common operators found in almost all languages
		"common": {regexp.MustCompile(strings.Join([]string{
			"^(",
			strings.Join([]string{
				// &, |, ^
				"([&|\\^])",
				// ++, --, &&, ||
				"([+\\-&|]{2})",
				// !, !=, <, <<, <=, <<=, >, >>, >>>, >=, >>=, >>>=, ^, ^=, +, +=, -,
				// -=, %, %=, *, *=
				"((<{2}|>{2,3}|[!|&<>^+\\-=%/*])=?)",
			}, "|"),
			")",
		}, "")), "OPERATOR"},
	},
}
