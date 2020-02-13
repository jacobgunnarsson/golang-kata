package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Test 1",func() {Expect(solve([]int{2,3,4,5,6})).To(Equal(true))})
	It("Test 2",func() {Expect(solve([]int{6,2,3,4,5})).To(Equal(true))})
	It("Test 3",func() {Expect(solve([]int{3,2,4,5,6})).To(Equal(false))})
	It("Test 4",func() {Expect(solve([]int{5,6,54,435,888,99,-8,1,2})).To(Equal(false))})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "circularlysortedarray")
}
