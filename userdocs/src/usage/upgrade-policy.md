# Cluster Upgrade Policy

This document describes how to configure the upgrade policy for your EKS cluster using eksctl.

## Overview

The `upgradePolicy` field allows you to specify the support type for your EKS cluster. This determines the level of support AWS provides for your cluster version.

## Support Types

- **STANDARD**: The default support type that provides standard AWS support for the cluster
- **EXTENDED**: Provides extended support for older Kubernetes versions beyond the standard support period

## Configuration

You can specify the upgrade policy in your cluster configuration file:

```yaml
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

metadata:
  name: my-cluster
  region: us-west-2
  version: "1.34"
  upgradePolicy:
    supportType: "EXTENDED"  # or "STANDARD"

managedNodeGroups:
  - name: mng-1
    instanceType: m5.large
    desiredCapacity: 1
```

## Command Line Usage

When creating a cluster with a specific upgrade policy:

```bash
eksctl create cluster --config-file=cluster-config.yaml
```

## Examples

### Extended Support (Default)
```yaml
metadata:
  name: extended-cluster
  region: us-west-2
  upgradePolicy:
    supportType: "EXTENDED"
```

### Standard Support
```yaml
metadata:
  name: standard-cluster
  region: us-west-2
  upgradePolicy:
    supportType: "STANDARD"
```

### No Upgrade Policy (Uses AWS Default)
```yaml
metadata:
  name: default-cluster
  region: us-west-2
  # No upgradePolicy specified - uses AWS default behavior
```

## Notes

- If no `upgradePolicy` is specified, AWS will use its default behavior
- The upgrade policy can only be set during cluster creation
- Extended support may incur additional costs - refer to AWS EKS pricing documentation
