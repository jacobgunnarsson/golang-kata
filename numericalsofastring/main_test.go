package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Hello, World!",func() {Expect(solve("Hello, World!")).To(Equal("1112111121311"))})
	It("Inconceivable!",func() {Expect(solve("Inconceivable!")).To(Equal("11112211111121"))})
	It("hello hello",func() {Expect(solve("hello hello")).To(Equal("11121122342"))})
	It("Hello",func() {Expect(solve("Hello")).To(Equal("11121"))})
	It("aaaaaaaaaaaa",func() {Expect(solve("aaaaaaaaaaaa")).To(Equal("123456789101112"))})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "numericalsofastring")
}
