package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fixed Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
		Expect(PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
		Expect(PositiveSum([]int{})).To(Equal(0))
		Expect(PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
		Expect(PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "sumofpositive")
}
