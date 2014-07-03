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

        // ~, ===, .
        scanner.Definition{regexp.MustCompile("^(~|={3}|\\.)"), "OPERATOR"},

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
              "continue",
              "break",
              "return",
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
          []string{"CHAR", "["},
          []string{"ARRAYSTART"},
        },
        [][]string{
          []string{"CHAR", "]"},
          []string{"ARRAYEND"},
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

        // <-, :=, :, ..., .
        scanner.Definition{regexp.MustCompile("^(<-|:=?|\\.{3}|\\.)"), "OPERATOR"},

        // &&, &^, &^=
        scanner.Definition{regexp.MustCompile("^(&{2}|(&\\^?)=?)"), "OPERATOR"},

        // Common operators
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

    "Python": &Language{
      Extensions: []string{"py"},
      Map: append([]scanner.Definition{
        // Comments
        scanner.Definition{regexp.MustCompile("^#.+"), "COMMENT"},

        // Heredoc
        scanner.Definition{regexp.MustCompile(`^"{3}(\"?[^"])*"{3}`), "DOCSTRING"},
        scanner.Definition{regexp.MustCompile(`^'{3}(\'?[^'])*'{3}`), "DOCSTRING"},

        // Double quote strings
        regex["strings"]["double"],

        // Single quote strings
        regex["strings"]["single"],

        // Decorators
        scanner.Definition{regexp.MustCompile("^@[^\\s]+"), "DECORATOR"},

        // //, **
        scanner.Definition{regexp.MustCompile("^([*/]{2})"), "OPERATOR"},

        // Common operators
        regex["operators"]["common"],

        // Restricted words
        scanner.Definition{regexp.MustCompile(
          strings.Join([]string{
            "^(",
            strings.Join([]string{
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
            }, "|"),
            ")",
          }, ""),
        ), "IDENT"},
      }, scanner.Map()...),
    },

    "Java": &Language{
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
        scanner.Definition{regexp.MustCompile("^@[^\\s]+"), "DECORATOR"},

        // ~, ?:, .
        scanner.Definition{regexp.MustCompile("^(~|\\?:|\\.)"), "OPERATOR"},

        // Common operators
        regex["operators"]["common"],

        // Restricted Java words
        scanner.Definition{regexp.MustCompile(
          strings.Join([]string{
            "^(",
            strings.Join([]string{
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

    "Ruby": &Language{
      Extensions: []string{"rb"},
      Map: append([]scanner.Definition{
        // Comments
        scanner.Definition{regexp.MustCompile("^#.+"), "COMMENT"},

        // Double quote strings
        regex["strings"]["double"],

        // Single quote "string"
        regex["strings"]["single"],

        // ::, .., ..., ., <=>, ===, =~, !~
        scanner.Definition{regexp.MustCompile("^(:{2}|\\.{2,3}|\\.|<=>|={3}|=[~>]|!~)"), "OPERATOR"},

        // Common operators
        regex["operators"]["common"],

        // Restricted Ruby words
        scanner.Definition{regexp.MustCompile(
          strings.Join([]string{
            "^(",
            strings.Join([]string{
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
            }, "|"),
            ")",
          }, "")), "IDENT"},

        // Restricted Ruby variables
        scanner.Definition{regexp.MustCompile(
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
          }, "")), "IDENT"},
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
          []string{"CHAR", "["},
          []string{"ARRAYSTART"},
        },
        [][]string{
          []string{"CHAR", "]"},
          []string{"ARRAYEND"},
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
