package checker_test

import (
	"github.com/MalteHerrmann/evmos-checker/checker"
	"github.com/MalteHerrmann/evmos-checker/checks/lint"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating a new checker", func() {
	It("should return a valid checker instance with no checks", func() {
		evmosCheck := checker.NewChecker(checker.NewConfig(
			[]checker.Check{},
		))

		Expect(evmosCheck).ToNot(BeNil(), "evmosCheck should not be nil")
		Expect(evmosCheck.GetChecks()).To(BeEmpty(), "checks should be empty")
		Expect(evmosCheck.Run()).To(Succeed(), "running the checks should succeed")
	})

	It("should return a valid checker with lint checks", func() {
		evmosCheck := checker.NewChecker(checker.NewConfig(
			lint.NewLintChecks(),
		))

		Expect(evmosCheck).ToNot(BeNil(), "evmosCheck should not be nil")
		Expect(evmosCheck.GetChecks()).To(HaveLen(2), "checks should have length 2")
		Expect(evmosCheck.Run()).To(Succeed(), "running the checks should succeed")
	})
})
