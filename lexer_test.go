package lexer_test

import (
	"github.com/Southern/lexer"
	"io/ioutil"
	"strings"
	"testing"
)

var l = lexer.New()

func TestParse(t *testing.T) {
	scan, err := l.Parse("Javascript", `/*

  Animal can be used as a base for different types of animals.

*/
function Animal(name) {
  this.name = name;
}

// Make our animal say hello!
Animal.prototype.sayHello = function() {
  return 'Hello from ' + this.name;
};

// Make our animal make some noise!
Animal.prototype.makeNoise = function() {
  return this.noise || '<chirp>';
};

/*

  Aw, look. It's a cute little dog.

*/
function Dog(name, breed) {
  this.name = name;
  this.breed = breed;
  this.noise = 'Woof!';
}

// Inherit Animal
Dog.prototype = new Animal();

// Our dog is smart. He can say hello AND his breed.
Dog.prototype.sayExtendedHello = function() {
  return this.sayHello() + ', ' + this.breed;
};

// Our dog can also bark. Not as impressive.
Dog.prototype.bark = function() {
  return this.noise;
};

// Expose our Animal and Dog to the outside world.
module.exports = {
  Animal: Animal,
  Dog: Dog,
};
`)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	Status("Scan: %+v", scan)
}

func TestParseHex(t *testing.T) {
	str := "0x00 0xFF 0xFFZZ"
	expects := [][]string{
		{"HEX", "0x00"},
		{"WHITESPACE", " "},
		{"HEX", "0xFF"},
		{"WHITESPACE", " "},
		{"WORD", "0xFFZZ"},
	}

	for lang := range lexer.Languages {
		scan, err := l.Parse(lang, str)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		for i, expect := range expects {
			if i > len(scan.Scanner.Tokens)-1 {
				t.Fatalf("Excpected more output: %+v\n", expects[i:])
			}
			if scan.Scanner.Tokens[i][0] != expect[0] ||
				scan.Scanner.Tokens[i][1] != expect[1] {
				t.Fatalf("Expected %+v, got %+v", expect, scan.Scanner.Tokens[i])
			}
		}

		Status("Scan: %+v", scan)
	}
}

func TestParseNoDataError(t *testing.T) {
	_, err := l.Parse()

	if err == nil {
		t.Fatalf("Expected error.")
	}

	Status("Got error: %+v", err)
}

func TestParseStringFirstError(t *testing.T) {
	_, err := l.Parse([]int{1, 2, 3, 4, 5}, "Test")

	if err == nil {
		t.Fatal("Expected error.")
	}

	Status("Got error: %+v", err)
}

func TestParseScannerError(t *testing.T) {
	_, err := l.Parse([]int{1, 2, 3, 4, 5})

	if err == nil {
		t.Fatalf("Expected error.")
	}

	Status("Got error: %+v", err)
}

func TestReadFile(t *testing.T) {
	Status("Reading all files in testdata directory")
	files, err := ioutil.ReadDir("testdata")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	Status("Scanning all files found in testdata directory")
	for len(files) > 0 {
		file := strings.Join([]string{"testdata", files[0].Name()}, "/")
		Status("Scanning file: %s", file)

		scan, err := l.ReadFile(file)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		Status("Scanned: %+v", scan)
		files = files[1:]
	}
}

func TestReadFileInvalidFileError(t *testing.T) {
	_, err := l.ReadFile("testdata/idontexist")

	if err == nil {
		t.Fatal("Expected error.")
	}

	Status("Got error: %+v", err)
}
