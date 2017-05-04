package gomega_matchers_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGomegaMatchers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GomegaMatchers Suite")
}

func readFileContents(filePath string) []byte {
	f := openFile(filePath)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("failed to read file contents: %v", err))
	}
	return b
}

func openFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %v", err))
	}
	return f
}
