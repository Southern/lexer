package lexer

import (
  "github.com/Southern/scanner"
  "regexp"
)

var regex = map[string]map[string]scanner.Definition{
  "comments": map[string]scanner.Definition{
    // Hash comments
    "hash": scanner.Definition{regexp.MustCompile("^#(.+)"), "COMMENT"},

    // Single line comments
    "oneline": scanner.Definition{regexp.MustCompile("^\\/{2,}\\s*(.+)"), "COMMENT"},

    // Multi-line comments
    "multiline": scanner.Definition{regexp.MustCompile("^/\\*([^*]|[\r\n]|(\\*+([^*/]|[\r\n])))*\\*+/"), "COMMENT"},
  },

  "strings": map[string]scanner.Definition{
    // Double quote strings
    "double": scanner.Definition{regexp.MustCompile("^\"([^\"\\n]|\\\")+\""), "STRING"},

    // Single quote strings
    "single": scanner.Definition{regexp.MustCompile("^'([^'\\n]|\\')+'"), "STRING"},
  },

  "operators": map[string]scanner.Definition{
    // Most common operators in languages: ++, --, -=, +=, /=, %=, ==, =
    "common": scanner.Definition{regexp.MustCompile("^(\\+{1,2}|-{1,2}|[+\\-=%\\/]=?)"), "OPERATOR"},
  },
}
