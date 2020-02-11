package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(solve("a")).To(Equal("a"))
		Expect(solve("Z")).To(Equal("Z"))
		Expect(solve("coDe")).To(Equal("code"))
		Expect(solve("CODe")).To(Equal("CODE"))
		Expect(solve("coDE")).To(Equal("code"))
		Expect(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).To(Equal("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "fixstringcase")
}
