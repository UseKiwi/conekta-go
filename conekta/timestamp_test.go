package conekta

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timestamp", func() {

	Describe("NewTimestamp", func() {
		It("Returns a pointer to a timestamp", func() {
			t := NewTimestamp("2015-01-01")
			Expect(t.Month()).To(Equal(time.January))
			Expect(t.Day()).To(Equal(1))
		})
	})
})
