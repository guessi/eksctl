package utils

import (
	"context"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
)

func updateClusterDeletionProtectionCmd(cmd *cmdutils.Cmd) {
	cfg := api.NewClusterConfig()
	cmd.ClusterConfig = cfg

	cmd.SetDescription("deletion-protection", "Update cluster deletion protection", "")

	var enabled bool

	cmd.CobraCommand.RunE = func(_ *cobra.Command, args []string) error {
		cmd.NameArg = cmdutils.GetNameArg(args)
		if err := cmdutils.NewUtilsUpdateDeletionProtectionLoader(cmd, &enabled).Load(); err != nil {
			return err
		}
		return doUpdateClusterDeletionProtection(cmd)
	}

	cmd.FlagSetGroup.InFlagSet("General", func(fs *pflag.FlagSet) {
		cmdutils.AddClusterFlag(fs, cfg.Metadata)
		cmdutils.AddRegionFlag(fs, &cmd.ProviderConfig)
		cmdutils.AddConfigFileFlag(fs, &cmd.ClusterConfigFile)
		cmdutils.AddApproveFlag(fs, cmd)
		fs.BoolVar(&enabled, "enabled", false, "Whether deletion protection is enabled for the cluster")
	})

	cmdutils.AddCommonFlagsForAWS(cmd, &cmd.ProviderConfig, false)
}

// A deletionProtectionUpdater updates a cluster's deletion protection setting.
type deletionProtectionUpdater interface {
	UpdateClusterConfigForDeletionProtection(ctx context.Context, cfg *api.ClusterConfig) error
}

func doUpdateClusterDeletionProtection(cmd *cmdutils.Cmd) error {
	ctx := context.Background()
	ctl, err := cmd.NewProviderForExistingCluster(ctx)
	if err != nil {
		return err
	}

	cfg := cmd.ClusterConfig
	logger.Info("using region %s", cfg.Metadata.Region)

	if ok, err := ctl.CanUpdate(cfg); !ok {
		return err
	}

	return updateDeletionProtection(ctx, ctl, cfg, ctl.Status.ClusterInfo.Cluster.DeletionProtection, cmd.Plan)
}

// updateDeletionProtection applies cfg.DeletionProtection to the cluster, unless it already holds
// the desired value or planMode is set, in which case the intended change is only logged.
func updateDeletionProtection(ctx context.Context, updater deletionProtectionUpdater, cfg *api.ClusterConfig, current *bool, planMode bool) error {
	meta := cfg.Metadata

	// The loader guarantees this is set, either from --enabled or from the config file.
	desired := *cfg.DeletionProtection
	updateRequired := current == nil || *current != desired

	if updateRequired {
		cmdutils.LogIntendedAction(planMode, "%s deletion protection for cluster %q in %q",
			enableOrDisable(desired), meta.Name, meta.Region)
		if !planMode {
			if err := updater.UpdateClusterConfigForDeletionProtection(ctx, cfg); err != nil {
				return err
			}
		}
	} else {
		logger.Success("deletion protection for cluster %q in %q is already %s",
			meta.Name, meta.Region, enabledOrDisabled(desired))
	}

	cmdutils.LogPlanModeWarning(planMode && updateRequired)
	return nil
}

func enableOrDisable(enabled bool) string {
	if enabled {
		return "enable"
	}
	return "disable"
}

func enabledOrDisabled(enabled bool) string {
	if enabled {
		return "enabled"
	}
	return "disabled"
}
