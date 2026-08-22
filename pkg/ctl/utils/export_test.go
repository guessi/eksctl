package utils

import (
	"context"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
)

func ValidateLoggingFlags(toEnable, toDisable []string) error {
	return validateLoggingFlags(toEnable, toDisable)
}

// DeletionProtectionUpdater is exported for testing.
type DeletionProtectionUpdater = deletionProtectionUpdater

// UpdateDeletionProtection is exported for testing.
func UpdateDeletionProtection(ctx context.Context, updater DeletionProtectionUpdater, cfg *api.ClusterConfig, current *bool, planMode bool) error {
	return updateDeletionProtection(ctx, updater, cfg, current, planMode)
}
