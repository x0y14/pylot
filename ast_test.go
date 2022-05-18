package pylot

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestAstDumpWithCode(t *testing.T) {
	code := `print("hello, python")` + "\n"
	astSTR, err := DumpAstWithCode(code)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Module(body=[Expr(value=Call(func=Name(id='print', ctx=Load()), args=[Constant(value='hello, python')], keywords=[]))], type_ignores=[])\n", astSTR)
}

func TestDumpAstWithFilePath(t *testing.T) {
	tmp, _ := ioutil.TempFile("", "testpy")
	defer os.Remove(tmp.Name())

	code := `print("hello, python")` + "\n"
	_, err := tmp.WriteString(code)
	if err != nil {
		t.Fatal(err)
	}

	astSTR, err := DumpAstWithFilePath(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Module(body=[Expr(value=Call(func=Name(id='print', ctx=Load()), args=[Constant(value='hello, python')], keywords=[]))], type_ignores=[])\n", astSTR)
}
