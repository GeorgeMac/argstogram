package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh/terminal"
)

type Package struct {
	Name  string
	Files map[string]map[string]int
}

func Parse(dir string, pkgs chan<- Package) error {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	parsed, err := parser.ParseDir(fset, dir, func(os.FileInfo) bool { return true }, 0)
	if err != nil {
		return err
	}

	for name, pkg := range parsed {
		pack := Package{
			Name:  name,
			Files: make(map[string]map[string]int),
		}

		for name, file := range pkg.Files {
			fileCounts := make(map[string]int)

			for _, decl := range file.Decls {
				if fdcl, ok := decl.(*ast.FuncDecl); ok {
					fileCounts[fdcl.Name.Name] = fdcl.Type.Params.NumFields()
				}
			}

			pack.Files[name] = fileCounts
		}

		pkgs <- pack
	}

	return nil
}

func histAdd(hist []int, count int) ([]int, int) {
	var newHist []int
	if count > len(hist)-1 {
		newHist = make([]int, count+1)
		copy(newHist, hist)
	} else {
		newHist = hist
	}

	newHist[count]++

	return newHist, newHist[count]
}

func printHistogram(histogram []int, width, maxWidth int) {
	padding := int(math.Log10(float64(maxWidth))) + 1
	format := fmt.Sprintf("(%%02d) [%%0%dd] ", padding)
	for i, count := range histogram {
		fmt.Printf(format, i, count)

		// normalise
		length := count * (width - 20) / maxWidth
		for j := 0; j < length; j++ {
			fmt.Print("=")
		}
		fmt.Println()
	}
}

func main() {
	packages := make(chan Package)
	go func() {
		if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if err := Parse(path, packages); err != nil {
					panic(err)
				}
			}

			return nil
		}); err != nil {
			panic(err)
		}

		close(packages)
	}()

	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	histogram := []int{}
	var maxWidth int
	for pkg := range packages {
		for _, fi := range pkg.Files {
			for _, fnCount := range fi {
				var cur int
				histogram, cur = histAdd(histogram, fnCount)
				if maxWidth < cur {
					maxWidth = cur
				}
			}
		}
	}

	printHistogram(histogram, w, maxWidth)
}
