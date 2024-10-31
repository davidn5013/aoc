// Package skeleton makes skeletons to be filled out with solutions.
package skeleton

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/davidn5013/aoc/util"
)

//go:embed sol_tmpls/*.go
var fsSol embed.FS

// Run makes a skeleton main.go and main_test.go file for the given day and year
func RunSol(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	ts, err := template.ParseFS(fsSol, "sol_tmpls/*.go")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	y2 := year - 2000
	mainFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("sol/d%d%02d/d%d%02d.go", y2, day, y2, day))
	testFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("sol/d%d%02d/d%d%02d_test.go", y2, day, y2, day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}
	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "dYYDD.go", nil)
	ts.ExecuteTemplate(testFile, "dYYDD_test.go", nil)
	fmt.Printf("templates made for %d-day%d\n", year, day)
}
