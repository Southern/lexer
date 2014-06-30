package lexer

import (
  "github.com/Southern/scanner"
  "regexp"
  "strings"
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
    // Common operators found in almost all languages
    "common": scanner.Definition{regexp.MustCompile(strings.Join([]string{
      "^(",
      strings.Join([]string{
        // ++, --
        "([+\\-&|]{2})",
        // !, !=, <, <<, <=, <<=, >, >>, >>>, >=, >>=, >>>=, ^, ^=, +, +=, -,
        // -=, %, %=, *, *=
        "((<{2}|>{2,3}|[!|&<>^+\\-=%/*])=?)",
      }, "|"),
      ")",
    }, "")), "OPERATOR"},
  },
}
