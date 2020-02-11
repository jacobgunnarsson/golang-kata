package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fixed Tests", func() {
	It("The solution should return the correct values for the sample test cases!", func() {
		Expect(countSheep(2)).To(Equal("1 sheep...2 sheep..."))
		Expect(countSheep(0)).To(Equal(""))
		Expect(countSheep(1)).To(Equal("1 sheep..."))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "countsheep")
}
