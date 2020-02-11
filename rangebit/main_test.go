package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Testing 2 and 7",func() {Expect(solve(2,7)).To(Equal(11))})
	It("Testing 0 and 1",func() {Expect(solve(0,1)).To(Equal(1))})
	It("Testing 4 and 4",func() {Expect(solve(4,4)).To(Equal(1))})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "rangebit")
}
