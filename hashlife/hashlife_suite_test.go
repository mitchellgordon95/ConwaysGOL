package hashlife_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHashlife(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hashlife Suite")
}
