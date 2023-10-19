package a

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStackit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stackit Suite")
}

var _ = Describe("f", func() {
	It("no 'invalid pointer found on stack' please", func() {
		g, _ := f()
		Expect(errors.New("not even related to the call to f")).NotTo(HaveOccurred())
		Expect(g.Listeners).NotTo(BeNil()) // accessing Listeners would panic if we ever got here because g is nil
	})
})
