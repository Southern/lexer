package lexer

import (
  "github.com/Southern/scanner"
  "regexp"
  "strings"
)

func init() {
  Languages = map[string]*Language{
    "Javascript": &Language{
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

        // Operators
        regex["operators"]["common"],

        // Restricted words
        scanner.Definition{regexp.MustCompile(
          strings.Join([]string{
            "^(",
            strings.Join([]string{
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
            }, "|"),
            ")",
          }, ""),
        ), "IDENT"},
      }, scanner.Map()...),
      Modify: [][][]string{
        [][]string{
          []string{"CHAR", "{"},
          []string{"BLOCKSTART"},
        },
        [][]string{
          []string{"CHAR", "}"},
          []string{"BLOCKEND"},
        },
        [][]string{
          []string{"CHAR", "("},
          []string{"ARGSTART"},
        },
        [][]string{
          []string{"CHAR", ")"},
          []string{"ARGEND"},
        },
        [][]string{
          []string{"CHAR", ";"},
          []string{"END"},
        },
      },
    },

    "Go": &Language{
      Extensions: []string{"go"},
      Map: append([]scanner.Definition{
        // Single line comments
        regex["comments"]["oneline"],

        // Multi-line comments
        regex["comments"]["multiline"],

        // Double quote strings
        regex["strings"]["double"],

        // Operators
        regex["operators"]["common"],

        // Restricted words
        scanner.Definition{regexp.MustCompile(
          strings.Join([]string{
            "^(",
            strings.Join([]string{
              "(Complex|Float|Integer)?Type",
              "Type1",
              "bool",
              "byte",
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
            }, "|"),
            ")",
          }, ""),
        ), "IDENT"},
      }, scanner.Map()...),
      Modify: [][][]string{
        [][]string{
          []string{"CHAR", "{"},
          []string{"BLOCKSTART"},
        },
        [][]string{
          []string{"CHAR", "}"},
          []string{"BLOCKEND"},
        },
        [][]string{
          []string{"CHAR", "("},
          []string{"ARGSTART"},
        },
        [][]string{
          []string{"CHAR", ")"},
          []string{"ARGEND"},
        },
      },
    },
  }

  Languages["Node"] = Languages["Javascript"]
  Languages["Node"].Map = append([]scanner.Definition{
    // Restricted Node words
    scanner.Definition{regexp.MustCompile(
      strings.Join([]string{
        "^(",
        strings.Join([]string{
          "module",
          "exports",
          "require",
          "global",
          "process",
          "console",
          "__dirname",
          "__filename",
        }, "|"),
        ")",
      }, ""),
    ), "IDENT"},
  }, Languages["Node"].Map...)
}
