package e2e

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"k8s.io/klog/v2"
)

var (
	// TestFlags passed in from the command line.
	TestFlags = struct {
		Project            string
		ResourcePrefix     string
		ServiceAccountName string
		RunID              string
	}{
		Project:            "",
		ResourcePrefix:     "k8scp-",
		ServiceAccountName: "",
	}
	RunID string
)

func init() {
	klog.InitFlags(flag.CommandLine)

	flag.StringVar(&TestFlags.Project, "project", TestFlags.Project, "GCP Project ID")
	flag.StringVar(&TestFlags.ResourcePrefix, "resourcePrefix", TestFlags.ResourcePrefix, "Prefix used to name all resources created in the tests. Any resources with this prefix will be removed during cleanup.")
	flag.StringVar(&TestFlags.ServiceAccountName, "sa-name", TestFlags.ServiceAccountName, "Name of the Service Account to impersonate")
	flag.StringVar(&RunID, "run-id", fmt.Sprintf("%0x", rand.Int63()&0xffff), "The RunID to use when managing resources.")
}
func ParseFlagsOrDie() {
	flag.Parse()

	if TestFlags.Project == "" {
		fmt.Println("-project must be set")
		os.Exit(1)
	}
}
