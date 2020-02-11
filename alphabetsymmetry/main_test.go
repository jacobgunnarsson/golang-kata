package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(solve([]string{"abode", "ABc", "xyzD"})).To(Equal([]int{4, 3, 1}))
		Expect(solve([]string{"abide", "ABc", "xyz"})).To(Equal([]int{4, 3, 0}))
		Expect(solve([]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"})).To(Equal([]int{6, 5, 7}))
		Expect(solve([]string{"encode", "abc", "xyzD", "ABmD"})).To(Equal([]int{1, 3, 1, 3}))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "alphabetsymmetry")
}
