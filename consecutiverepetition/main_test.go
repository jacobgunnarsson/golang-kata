package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("it should work with the sample tests", func() {
		Expect(solve("aaaabb")).Should(Equal(Result{'a', 4}))
		Expect(solve("bbbaaabaaaa")).Should(Equal(Result{'a', 4}))
		Expect(solve("cbdeuuu900")).Should(Equal(Result{'u', 3}))
		Expect(solve("abbbbb")).Should(Equal(Result{'b', 5}))
		Expect(solve("aabb")).Should(Equal(Result{'a', 2}))
		Expect(solve("ba")).Should(Equal(Result{'b', 1}))
		Expect(solve("")).Should(Equal(Result{}))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "consecutiverepetition")
}
