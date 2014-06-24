package lexer

import (
  "fmt"
  "github.com/Southern/scanner"
  "io/ioutil"
  "strings"
)

type Language struct {
  Extensions []string
  Map        []scanner.Definition
  Modify     [][][]string
}

var Languages map[string]*Language

type Lexer struct {
  Language string

  Scanner  scanner.Scanner
  language *Language
}

func New() Lexer {
  return Lexer{
    Language: "plaintext",
    Scanner:  scanner.New(),
  }
}

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
      l.language = Languages[l.Language]
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

func (l Lexer) ReadFile(filename string) (error, Lexer) {
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return err, l
  }

  return l.Parse(data)
}
