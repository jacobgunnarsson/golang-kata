package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(solve(2015)).To(Equal(3))
		Expect(solve(1986)).To(Equal(1))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "unluckydays")
}
