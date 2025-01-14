package deepl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeepl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deepl Suite")
}
