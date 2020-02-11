package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(alphabetsymmetry([]string{"abode", "ABc", "xyzD"})).To(Equal([]int{4, 3, 1}))
		Expect(alphabetsymmetry([]string{"abide", "ABc", "xyz"})).To(Equal([]int{4, 3, 0}))
		Expect(alphabetsymmetry([]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"})).To(Equal([]int{6, 5, 7}))
		Expect(alphabetsymmetry([]string{"encode", "abc", "xyzD", "ABmD"})).To(Equal([]int{1, 3, 1, 3}))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "alphabetsymmetry")
}
