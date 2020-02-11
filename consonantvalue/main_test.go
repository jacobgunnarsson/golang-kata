package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Hello, World!",func() {Expect(Numericals("Hello, World!")).To(Equal("1112111121311"))})
	It("Inconceivable!",func() {Expect(Numericals("Inconceivable!")).To(Equal("11112211111121"))})
	It("hello hello",func() {Expect(Numericals("hello hello")).To(Equal("11121122342"))})
	It("Hello",func() {Expect(Numericals("Hello")).To(Equal("11121"))})
	It("aaaaaaaaaaaa",func() {Expect(Numericals("aaaaaaaaaaaa")).To(Equal("123456789101112"))})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tablepattern")
}
