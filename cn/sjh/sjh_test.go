package sjh_test

import (
	"testing"

	"github.com/eynstudio/gox/cn/sjh"
)

func TestValidAndMask(t *testing.T) {
	s := sjh.ValidAndMask("13750419848")
	t.Log(s)
}
