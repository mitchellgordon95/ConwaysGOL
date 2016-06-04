package quadtree

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("quadNode", func() {
	var quad Node
	BeforeEach(func() {
		quad = EmptyTree(5)
	})

	It("returns the correct level", func() {
		Expect(quad.Level()).To(Equal(uint(4)))
	})
	It("returns errors for out of bounds", func() {
		_, err1 := quad.SetValue(100000000, 1, true)
		_, err2 := quad.GetValue(1, 1000000000)

		Expect(err1).To(HaveOccurred())
		Expect(err2).To(HaveOccurred())
	})
	It("sets and returns the correct value", func() {
		val, err := quad.GetValue(0, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(val).To(Equal(false))

		setAndGet(0, 0, quad)
		setAndGet(-150, 1, quad)
		setAndGet(-150, -10, quad)
		setAndGet(50, -15, quad)
	})
	It("doesn't affect the original value", func() {
		_, err := quad.SetValue(0, 0, true)
		Expect(err).ToNot(HaveOccurred())
		Expect(quad).To(Equal(EmptyTree(5)))
	})
	It("properly caches previously seen nodes", func() {
		qn := quad.(*quadNode)
		changed, err := quad.SetValue(0, 0, true)
		Expect(err).ToNot(HaveOccurred())
		original, err := changed.SetValue(0, 0, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(qn).To(Equal(original.(*quadNode)))
	})
})

// Asserts a val is false.
// Sets a value to true and then gets it asserts that the result is true
func setAndGet(x, y int64, quad Node) {
	val, err := quad.GetValue(0, 0)
	Expect(err).ToNot(HaveOccurred())
	Expect(val).To(Equal(false))
	quad, err = quad.SetValue(0, 0, true)
	Expect(err).ToNot(HaveOccurred())
	val, err = quad.GetValue(0, 0)
	Expect(err).ToNot(HaveOccurred())
	Expect(val).To(Equal(true))
}

// tree := qt.EmptyTree(63)

// tree, _ = tree.SetValue(0, 0, true)
// tree, _ = tree.SetValue(0, -4, true)

// printTree(tree)

// tree, _ = tree.SetValue(7000000000, 0, true)

// val, err := tree.GetValue(7000000000, 0)

// if err != nil {
// 	fmt.Println("ERRFD fs")
// }
// fmt.Print(val)
