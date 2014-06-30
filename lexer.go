package lexer

import (
  "fmt"
  "github.com/Southern/scanner"
  "io/ioutil"
  "strings"
)

/*

Language is used for each individual language to define file extensions,
scanner definitions, and modifications that need to be made during lexing.

*/
type Language struct {
  Extensions []string
  Map        []scanner.Definition
  Modify     [][][]string
}

/*

Languages is used to map each Language to a string that can be accessible
when providing a language argument to Lexer.Parse().

*/
var Languages map[string]*Language

/*

Lexer is a single instance of a lexer, with a single language loaded, that
will allow the tokenization of the data sent to it.

*/
type Lexer struct {
  Language string

  Scanner  scanner.Scanner
  language *Language
}

/*

New returns a new instance of Lexer with the default Language being
"plaintext". This will also create a new scanner instance specifically for
this lexer.

*/
func New() Lexer {
  return Lexer{
    "plaintext",
    scanner.New(),
    &Language{},
  }
}

/*

Parse is used to parse all data sent to it.

If the number of arguments sent to Parse are more than 1, the first argument
is assumed to be the language that you wish this data to be parsed as.
Otherwise, the language will never be switched and your data will be parsed as
plaintext, since that is the default setting.

*/
func (l Lexer) Parse(data ...interface{}) (error, Lexer) {
  var toScan interface{}

  if len(data) == 0 && len(l.Scanner.Tokens) == 0 {
    return fmt.Errorf("No data to parse."), l
  }

  if len(data) > 1 {
    switch data[0].(type) {
    case string:
      l.Language = data[0].(string)
    default:
      return fmt.Errorf("Expected a string as the first argument."), l
    }

    toScan = data[1]
  } else if len(data) == 1 {
    toScan = data[0]
  }

  for lang, data := range Languages {
    if strings.ToLower(l.Language) == strings.ToLower(lang) {
      l.language = Languages[lang]
      if len(data.Map) > 0 {
        l.Scanner.Map = data.Map
      }
    }
  }

  err, scan := l.Scanner.Parse(toScan)
  if err != nil {
    return err, l
  }

  for i := 0; i < len(l.language.Modify); i++ {
  next:
    for j := 0; j < len(scan.Tokens); j++ {
      matches := make([]bool, len(scan.Tokens[j]))
      for k := 0; k < len(l.language.Modify[i][0]); k++ {
        if scan.Tokens[j][k] == l.language.Modify[i][0][k] {
          matches[k] = true
        }
      }

      for k := 0; k < len(matches); k++ {
        if !matches[k] {
          continue next
        }
      }

      for k := 0; k < len(l.language.Modify[i][1]); k++ {
        if len(l.language.Modify[i][1][k]) > 0 {
          scan.Tokens[j][k] = l.language.Modify[i][1][k]
        }
      }
    }
  }

  l.Scanner = scan

  return nil, l
}

/*

ReadFile is used to read in an entire file, and attempt to guess what language
we are going to be parsing. It does this by checking the file extensions that
have been defined for each language in the lexer. If there are multiple
languages that use the same extension, the first language that matches will be
used.

For example, let's say you have a Node.js file. It will parse the file as
Javascript rather than Node, because the map keys in Go are in alphabetical
order. This gives the J in Javascript a higher index than the N in Node, and
therefore Javascript matches first.

*/
func (l Lexer) ReadFile(filename string) (error, Lexer) {
  file, err := ioutil.ReadFile(filename)
  if err != nil {
    return err, l
  }

  split, langs := strings.Split(filename, "."), make([]string, 0)
  if len(split) > 1 {
    ext := strings.ToLower(split[len(split)-1])
    for lang, data := range Languages {
      if len(data.Extensions) > 0 {
        for _, e := range data.Extensions {
          if ext == strings.ToLower(e) {
            langs = append(langs, lang)
          }
        }
      }
    }
  }

  if len(langs) > 0 {
    return l.Parse(langs[0], file)
  }

  return l.Parse(file)
}
