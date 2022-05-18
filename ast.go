package pylot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// ef. ast_dump.py
const astDumpPy = `
import ast
import sys

if __name__ == '__main__':

    args = sys.argv
    if len(args) == 1:
        # not found : analyze target (*.py)
        print("not found : analyze target (*.py)\n")
        exit(1)

    target = args[1]

    with open(args[1], "r") as f:
        tree = ast.parse(f.read())
        print(ast.dump(tree))

    exit(0)

`

func DumpAstWithFilePath(filepath string) (string, error) {
	tmpF, _ := ioutil.TempFile("", "astDumpPy")
	defer os.Remove(tmpF.Name())

	_, err := tmpF.WriteString(astDumpPy)
	if err != nil {
		return "", err
	}

	out, err := exec.Command("python", tmpF.Name(), filepath).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func DumpAstWithCode(code string) (string, error) {
	tmpF, _ := ioutil.TempFile("", "astDumpPy")
	defer os.Remove(tmpF.Name())
	_, err := tmpF.WriteString(astDumpPy)
	if err != nil {
		return "", fmt.Errorf("failed to create analyzeFile.py: %v", err)
	}

	tmpPy, _ := ioutil.TempFile("", "codePy")
	defer os.Remove(tmpPy.Name())
	_, err = tmpPy.WriteString(code + "\n")
	if err != nil {
		return "", fmt.Errorf("failed to create targetFile.py: %v", err)
	}

	out, err := exec.Command("python", tmpF.Name(), tmpPy.Name()).Output()
	if err != nil {
		return "", fmt.Errorf("failed to analyze: %v", err)
	}
	return string(out), nil
}
