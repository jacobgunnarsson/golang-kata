package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(abbreviation("Sam Harris")).To(Equal("S.H"))
		Expect(abbreviation("Patrick Feenan")).To(Equal("P.F"))
		Expect(abbreviation("Evan Cole")).To(Equal("E.C"))
		Expect(abbreviation("P Favuzzi")).To(Equal("P.F"))
		Expect(abbreviation("David Mendieta")).To(Equal("D.M"))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "abbreviation")
}
