package cmd

import (
	"bufio"
	"github.com/grugrut/gorgasia/pkg/org"
	"os"
	"regexp"
)

var nodePrefix = regexp.MustCompile(`(^\*+)`)

// FromFile parse org file to Root struct.
func FromFile(filename string) (*org.Root, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	root := org.Root{}

	s := bufio.NewScanner(f)

	parent := root
	for s.Scan() {
		// 正規表現は遅いらしいがとりあえず正規表現で実装
		if nodePrefix.Match(s.Bytes()) {

		}
	}
	return &parent, nil
}
