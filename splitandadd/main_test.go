package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Tests", func() {
		Expect(solve([]int{1, 2, 3, 4, 5}, 2)).To(Equal([]int{5, 10}))
		Expect(solve([]int{1, 2, 3, 4, 5}, 3)).To(Equal([]int{15}))
		Expect(solve([]int{15}, 3)).To(Equal([]int{15}))
		Expect(solve([]int{1, 2, 3, 4}, 1)).To(Equal([]int{4, 6}))
		Expect(solve([]int{1, 2, 3, 4, 5, 6}, 20)).To(Equal([]int{21}))
		Expect(solve([]int{32, 45, 43, 23, 54, 23, 54, 34}, 2)).To(Equal([]int{183, 125}))
		Expect(solve([]int{32, 45, 43, 23, 54, 23, 54, 34}, 0)).To(Equal([]int{32, 45, 43, 23, 54, 23, 54, 34}))
		Expect(solve([]int{3, 234, 25, 345, 45, 34, 234, 235, 345}, 3)).To(Equal([]int{305, 1195}))
		Expect(solve([]int{3, 234, 25, 345, 45, 34, 234, 235, 345, 34, 534, 45, 645, 645, 645, 4656, 45, 3}, 4)).To(Equal([]int{1040, 7712}))
		Expect(solve([]int{23, 345, 345, 345, 34536, 567, 568, 6, 34536, 54, 7546, 456}, 20)).To(Equal([]int{79327}))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "splitandadd")
}
