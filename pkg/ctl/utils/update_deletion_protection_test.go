package utils_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/ctl/utils"
)

type updateDeletionProtectionEntry struct {
	args        []string
	expectedErr string
}

var _ = DescribeTable("invalid usage of deletion-protection", func(e updateDeletionProtectionEntry) {
	cmd := newMockCmd(append([]string{"deletion-protection"}, e.args...)...)
	_, err := cmd.execute()
	Expect(err).To(MatchError(ContainSubstring(e.expectedErr)))
},
	Entry("missing --cluster option", updateDeletionProtectionEntry{
		args:        []string{"--enabled=true"},
		expectedErr: "--cluster must be set",
	}),

	Entry("missing --enabled option", updateDeletionProtectionEntry{
		args:        []string{"--cluster", "test"},
		expectedErr: "--enabled must be set",
	}),
)

var _ = Describe("deletion-protection with a config file", func() {
	writeConfigFile := func(extra string) string {
		path := filepath.Join(GinkgoT().TempDir(), "config.yaml")
		contents := `apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: test
  region: us-west-2
` + extra
		Expect(os.WriteFile(path, []byte(contents), 0600)).To(Succeed())
		return path
	}

	// load runs the loader for `eksctl utils deletion-protection` directly, without going
	// through the command handler that talks to AWS. Load reassigns Cmd.ClusterConfig when a
	// config file is used, so the config is read back off the Cmd rather than captured up front.
	load := func(configFile string) (*api.ClusterConfig, error) {
		enabled := false
		cmd := &cmdutils.Cmd{
			ClusterConfig:     api.NewClusterConfig(),
			ClusterConfigFile: configFile,
			CobraCommand: &cobra.Command{
				Use: "deletion-protection",
				Run: func(_ *cobra.Command, _ []string) {},
			},
		}
		err := cmdutils.NewUtilsUpdateDeletionProtectionLoader(cmd, &enabled).Load()
		return cmd.ClusterConfig, err
	}

	When("deletionProtection is set to true", func() {
		It("loads it from the config file", func() {
			cfg, err := load(writeConfigFile("deletionProtection: true\n"))
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg.DeletionProtection).NotTo(BeNil())
			Expect(*cfg.DeletionProtection).To(BeTrue())
		})
	})

	When("deletionProtection is set to false", func() {
		It("loads it from the config file", func() {
			cfg, err := load(writeConfigFile("deletionProtection: false\n"))
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg.DeletionProtection).NotTo(BeNil())
			Expect(*cfg.DeletionProtection).To(BeFalse())
		})
	})

	When("deletionProtection is omitted", func() {
		It("returns an error rather than silently doing nothing", func() {
			_, err := load(writeConfigFile(""))
			Expect(err).To(MatchError("deletionProtection must be set in the config file"))
		})
	})

	When("--enabled is passed alongside a config file", func() {
		It("rejects the combination", func() {
			configFile := writeConfigFile("deletionProtection: true\n")
			cmd := newMockCmd("deletion-protection", "--config-file", configFile, "--enabled=true")
			_, err := cmd.execute()
			Expect(err).To(MatchError(ContainSubstring("cannot use --enabled")))
		})
	})
})

type fakeDeletionProtectionUpdater struct {
	calls int
	cfg   *api.ClusterConfig
	err   error
}

func (f *fakeDeletionProtectionUpdater) UpdateClusterConfigForDeletionProtection(_ context.Context, cfg *api.ClusterConfig) error {
	f.calls++
	f.cfg = cfg
	return f.err
}

var _ = Describe("applying deletion protection", func() {
	var (
		updater *fakeDeletionProtectionUpdater
		cfg     *api.ClusterConfig
	)

	BeforeEach(func() {
		updater = &fakeDeletionProtectionUpdater{}
		cfg = api.NewClusterConfig()
		cfg.Metadata.Name = "test"
		cfg.Metadata.Region = "us-west-2"
	})

	apply := func(current *bool, planMode bool) error {
		return utils.UpdateDeletionProtection(context.Background(), updater, cfg, current, planMode)
	}

	When("the change is not approved", func() {
		It("does not call the EKS API", func() {
			cfg.DeletionProtection = api.Enabled()
			Expect(apply(api.Disabled(), true)).To(Succeed())
			Expect(updater.calls).To(BeZero())
		})
	})

	When("the change is approved", func() {
		It("calls the EKS API with the desired setting", func() {
			cfg.DeletionProtection = api.Enabled()
			Expect(apply(api.Disabled(), false)).To(Succeed())
			Expect(updater.calls).To(Equal(1))
			Expect(*updater.cfg.DeletionProtection).To(BeTrue())
		})

		It("can disable deletion protection", func() {
			cfg.DeletionProtection = api.Disabled()
			Expect(apply(api.Enabled(), false)).To(Succeed())
			Expect(updater.calls).To(Equal(1))
			Expect(*updater.cfg.DeletionProtection).To(BeFalse())
		})

		It("updates when the current setting is unknown", func() {
			cfg.DeletionProtection = api.Enabled()
			Expect(apply(nil, false)).To(Succeed())
			Expect(updater.calls).To(Equal(1))
		})

		It("propagates an API error", func() {
			cfg.DeletionProtection = api.Enabled()
			updater.err = errors.New("boom")
			Expect(apply(api.Disabled(), false)).To(MatchError("boom"))
		})
	})

	When("the cluster is already in the desired state", func() {
		It("does not call the EKS API", func() {
			cfg.DeletionProtection = api.Enabled()
			Expect(apply(api.Enabled(), false)).To(Succeed())
			Expect(updater.calls).To(BeZero())
		})
	})
})
