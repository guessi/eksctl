# Cluster Deletion Protection

This document describes how to configure deletion protection for your EKS cluster using eksctl.

## Overview

The `deletionProtection` field allows you to enable deletion protection for your EKS cluster. This prevents accidental cluster deletion.

## Configuration

You can specify deletion protection in your cluster configuration file:

```yaml
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

metadata:
  name: my-cluster
  region: us-west-2

deletionProtection: true
```

## Command Line Usage

When creating a cluster with deletion protection:

```bash
eksctl create cluster --config-file=cluster-config.yaml
```

To update deletion protection on an existing cluster:

```bash
# Enable deletion protection
eksctl utils deletion-protection --cluster=my-cluster --enabled=true --approve

# Disable deletion protection
eksctl utils deletion-protection --cluster=my-cluster --enabled=false --approve
```

Without `--approve`, eksctl only logs the change it would make and does not call the EKS API. Rerun
with `--approve` once you are satisfied with the proposed change.

The same update can be driven from a config file, in which case `deletionProtection` is read from the
file and `--enabled` must not be passed:

```bash
eksctl utils deletion-protection --config-file=cluster-config.yaml --approve
```

## Notes

- If no `deletionProtection` is specified, AWS will use its default behavior (false)
- Deletion protection can be set during cluster creation and updated later
- When enabled, you must disable deletion protection before you can delete the cluster
- The setting only applies to clusters in an active state

