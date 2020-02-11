package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("should test that the solution returns the correct value", func() {
		// Only one cat is in the yard
		Expect(solve([]string {
		"------------",
		"------------",
		"-L----------",
		"------------",
		"------------",
		"------------"}, 10)).To(Equal(true))

		// There are two cats in the yard, and they are closer than the minimum distance
		Expect(solve([]string {
		"------------",
		"---M--------",
		"------------",
		"------------",
		"-------R----",
		"------------"}, 6)).To(Equal(false))

		// All three cats are in the yard, all further apart than or equal to the minimum distance
		Expect(solve([]string {
		"-----------L",
		"--R---------",
		"------------",
		"------------",
		"------------",
		"--M---------"}, 4)).To(Equal(true))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "catkata1")
}
