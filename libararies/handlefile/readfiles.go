package handlefile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type filePath []string

var analysePaths filePath
var examplePaths filePath

func walkDirs(dirpath string) (anaPaths, examPaths filePath) {
	filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if strings.Contains(path, "tests") || strings.Contains(path, "examples") {
			if !strings.Contains(path, "vendor") {
				examplePaths = append(examplePaths, path)
			}
		} else if strings.HasSuffix(info.Name(), ".go") {
			analysePaths = append(analysePaths, path)
			fmt.Printf("visited file: %q\n", path)
		}
		return nil
	})
	return analysePaths, examplePaths
}

func readfile(path string) {
	// fset := token.NewFileSet()
	// astfile, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	// if err != nil {
		// fmt.Errorf("panic: %s", err)
	// }

}
