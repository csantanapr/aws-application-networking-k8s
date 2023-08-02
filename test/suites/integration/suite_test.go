package integration

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-application-networking-k8s/test/pkg/test"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testFramework *test.Framework
var ctx context.Context

func TestIntegration(t *testing.T) {
	ctx = test.NewContext(t)
	testFramework = test.NewFramework(ctx)
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(180 * time.Second)
	SetDefaultEventuallyPollingInterval(10 * time.Second)
	BeforeEach(func() { testFramework.ExpectToBeClean(ctx) })
	AfterEach(func() { testFramework.ExpectToBeClean(ctx) })
	RunSpecs(t, "Integration")
}
