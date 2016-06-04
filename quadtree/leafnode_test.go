package quadtree

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("leafNode", func() {
	var leaf leafNode
	BeforeEach(func() {
		leaf = LeafNode(true).(leafNode)
	})

	It("returns a level of 0", func() {
		Expect(leaf.Level()).To(Equal(uint(0)))
	})
	It("returns errors for out of bounds", func() {
		_, err1 := leaf.SetValue(1, 1, true)
		_, err2 := leaf.GetValue(1, 1)

		Expect(err1).To(HaveOccurred())
		Expect(err2).To(HaveOccurred())
	})
	It("returns the correct value", func() {
		val1, err1 := leaf.GetValue(0, 0)
		Expect(err1).ToNot(HaveOccurred())
		Expect(val1).To(Equal(true))
	})
	It("doesn't affect the original value", func() {
		val, err := leaf.SetValue(0, 0, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(val).To(Equal(leafNode(false)))
		Expect(leaf).To(Equal(leafNode(true)))
	})
})
