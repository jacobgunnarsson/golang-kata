package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(solve("codewarriors")).To(Equal(2))
		Expect(solve("suoidea")).To(Equal(3))
		Expect(solve("ultrarevolutionariees")).To(Equal(3))
		Expect(solve("strengthlessnesses")).To(Equal(1))
		Expect(solve("cuboideonavicuare")).To(Equal(2))
		Expect(solve("chrononhotonthuooaos")).To(Equal(5))
		Expect(solve("iiihoovaeaaaoougjyaw")).To(Equal(8))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "longestvowelchain")
}
