package checker_test

import (
	"github.com/MalteHerrmann/evmos-checker/checker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating a new checker", func() {
	var evmosCheck *checker.Checker

	BeforeEach(func() {
		evmosCheck = checker.NewChecker()
	})

	It("should return a valid checker instance", func() {
		Expect(evmosCheck).ToNot(BeNil(), "evmosCheck should not be nil")
	})

	It("should return an empty list of checks", func() {
		availableChecks := evmosCheck.GetChecks()
		Expect(availableChecks).To(BeEmpty(), "checks should be empty")
	})

	It("should start the log with empty checks", func() {
		Expect(evmosCheck.Run()).To(Succeed(), "starting the checker should succeed")
	})
})
