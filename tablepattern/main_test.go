package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Should return the correct values for the basic test cases!", func() {
		Expect(solve(4, 3, "Nice pattern")).To(Equal(
`+---+---+---+
| N | i | c |
+---+---+---+
| e |   | p |
+---+---+---+
| a | t | t |
+---+---+---+
| e | r | n |
+---+---+---+`))

		Expect(solve(3, 3, "codewars")).To(Equal(
`+---+---+---+
| c | o | d |
+---+---+---+
| e | w | a |
+---+---+---+
| r | s |   |
+---+---+---+`))

		Expect(solve(1, 1, "?")).To(Equal(
`+---+
| ? |
+---+`))

    Expect(solve(2, 1, "")).To(Equal(
`+---+
|   |
+---+
|   |
+---+`))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tablepattern")
}
