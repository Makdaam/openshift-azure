package fakerp

import (
	"context"
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/openshift/openshift-azure/test/clients/azure"
)

var _ = Describe("Checking and changing plugin loglevel [LogLevelChanger][Fake][EveryPR]", func() {
	var (
		cli *azure.Client
	)

	BeforeEach(func() {
		var err error
		cli, err = azure.NewClientFromEnvironment(false)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should allow an SRE to fetch and change the current loglevel of the plugin", func() {
		By("Using the OSA admin client to fetch the loglevel")
		status, err := cli.OpenShiftManagedClustersAdmin.GetLogLevel(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(status).NotTo(BeNil())
		By("unmarshalling the result")
		var level string
		err = json.Unmarshal(status.Items, &level)
		Expect(err).NotTo(HaveOccurred())
		By("returned loglevel equal to 'debug' at first")
		Expect(level).To(Equal("debug"))
		By("changing log level to 'Warning'")
		changestatus, err := cli.OpenShiftManagedClustersAdmin.ChangeLogLevelAndWait(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), "Warn")
		Expect(err).NotTo(HaveOccurred())
		Expect(changestatus).NotTo(BeNil())

		By("fetching the new log level")
		status, err = cli.OpenShiftManagedClustersAdmin.GetLogLevel(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(status).NotTo(BeNil())
		By("unmarshalling the result")
		level = ""
		err = json.Unmarshal(status.Items, &level)
		Expect(err).NotTo(HaveOccurred())
		By("retured loglevel equal to 'warning' after change")
		Expect(level).To(Equal("warning"))

		By("changing log level back to 'Debug'")
		changestatus, err = cli.OpenShiftManagedClustersAdmin.ChangeLogLevelAndWait(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), "DEBUG")
		Expect(err).NotTo(HaveOccurred())
		Expect(changestatus).NotTo(BeNil())

		By("fetching the new log level")
		status, err = cli.OpenShiftManagedClustersAdmin.GetLogLevel(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(status).NotTo(BeNil())
		By("unmarshalling the result")
		level = ""
		err = json.Unmarshal(status.Items, &level)
		Expect(err).NotTo(HaveOccurred())
		By("returned loglevel equal to 'debug' after change")
		Expect(level).To(Equal("debug"))
	})
})
