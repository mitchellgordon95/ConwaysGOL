package quadtree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestQuadtree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quadtree Suite")
}
