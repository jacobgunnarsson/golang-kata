package kata

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	var game SnakesLadders

	game.NewGame()

	It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",1, 1,"Player 1 is on square 38"),func() {Expect(game.Play(1, 1)).To(Equal("Player 1 is on square 38"))} )
	It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",1, 5,"Player 1 is on square 44"),func() {Expect(game.Play(1, 5)).To(Equal("Player 1 is on square 44"))} )
	It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",6, 2,"Player 2 is on square 31"),func() {Expect(game.Play(6, 2)).To(Equal("Player 2 is on square 31"))} )
	It( fmt.Sprintf("Testing Roll with (%v, %v) to obtain \"%v\"",1, 1,"Player 1 is on square 25"),func() {Expect(game.Play(1, 1)).To(Equal("Player 1 is on square 25"))} )
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "snakesandladders")
}
