package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Should return the correct values!", func() {
		Expect(solve("a")).To(Equal(0))
		Expect(solve("cbd")).To(Equal(9))
		Expect(solve("zea")).To(Equal(26))
		Expect(solve("az")).To(Equal(26))
		Expect(solve("baz")).To(Equal(26))
		Expect(solve("aeiou")).To(Equal(0))
		Expect(solve("uaoczei")).To(Equal(29))
		Expect(solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")).To(Equal(143))
		Expect(solve("codewars")).To(Equal(37))
		Expect(solve("bup")).To(Equal(16))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tablepattern")
}
