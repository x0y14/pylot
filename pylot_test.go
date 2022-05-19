package pylot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJson(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect string
	}{
		{
			"print hello",
			`print("hello")`,
			"",
		},
		{
			"class",
			`class Human:
	def __init__(self, name: str):
		self.name: str = name
	def hello(self):
		print(self.name + "hello")
	def addMark(self, mark: str) -> str:
		return self.name + mark`,
			"",
		},
		{
			"class2",
			`class Human:
    def __init__(self, name: str):
        self.name = name

    def say(self, text: str) -> None:
        print(self.name + "< " + text)
    
    def mr(self) -> str:
        return "Mr."+self.name
    
    def my_name(self) -> str:
        return "myname"`,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j, err := ToJson(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, j)
		})
	}
}
