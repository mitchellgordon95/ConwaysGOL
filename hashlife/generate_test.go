package hashlife_test

import (
	"github.com/mitchellgordon95/ConwaysGOL/common"
	. "github.com/mitchellgordon95/ConwaysGOL/hashlife"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NextGeneration", func() {
	var hl common.GolBoard

	BeforeEach(func() {
		hl = NewHashLifeBoard()
	})

	It("handles a blinker", func() {
		hl = loadBoard(hl, [][]int64{{0, 0}, {0, 1}, {0, 2}})
		assertAlive(hl, [][]int64{{0, 0}, {0, 1}, {0, 2}})
		assertDead(hl, [][]int64{{-1, 1}, {1, 1}})
		hl = hl.Step()
		assertAlive(hl, [][]int64{{-1, 1}, {0, 1}, {1, 1}})
		assertDead(hl, [][]int64{{0, 0}, {0, 2}})
		hl = hl.Step()
		assertAlive(hl, [][]int64{{0, 0}, {0, 1}, {0, 2}})
		assertDead(hl, [][]int64{{-1, 1}, {1, 1}})
	})

})

func loadBoard(board common.GolBoard, alive [][]int64) common.GolBoard {
	for _, cell := range alive {
		board = board.AddCell(cell[0], cell[1])
	}
	return board
}

func assertAlive(board common.GolBoard, alive [][]int64) {
	for _, cell := range alive {
		Expect(board.IsAlive(cell[0], cell[1])).To(BeTrue())
	}
}
func assertDead(board common.GolBoard, dead [][]int64) {
	for _, cell := range dead {
		Expect(board.IsAlive(cell[0], cell[1])).To(BeFalse())
	}
}
