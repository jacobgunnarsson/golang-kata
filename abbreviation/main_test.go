package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(solve("Sam Harris")).To(Equal("S.H"))
		Expect(solve("Patrick Feenan")).To(Equal("P.F"))
		Expect(solve("Evan Cole")).To(Equal("E.C"))
		Expect(solve("P Favuzzi")).To(Equal("P.F"))
		Expect(solve("David Mendieta")).To(Equal("D.M"))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "abbreviation")
}
