package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("The solution should return the correct values for the sample test cases!", func() {
		Expect(solve(2)).To(Equal("1 sheep...2 sheep..."))
		Expect(solve(0)).To(Equal(""))
		Expect(solve(1)).To(Equal("1 sheep..."))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "countsheep")
}
