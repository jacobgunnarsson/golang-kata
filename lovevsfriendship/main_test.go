package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Testing for attitude", func() { Expect(solve("attitude")).To(Equal(100)) })
	It("Testing for friends", func() { Expect(solve("friends")).To(Equal(75)) })
	It("Testing for family", func() { Expect(solve("family")).To(Equal(66)) })
	It("Testing for selfness", func() { Expect(solve("selfness")).To(Equal(99)) })
	It("Testing for knowledge", func() { Expect(solve("knowledge")).To(Equal(96)) })
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "longestvowelchain")
}
