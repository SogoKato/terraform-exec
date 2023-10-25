// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"os/exec"
)

// StateReplaceProvider represents the terraform state mv subcommand.
func (tf *Terraform) StateReplaceProvider(ctx context.Context, from_provider_fqn string, to_provider_fqn string) error {
	cmd, err := tf.stateReplaceProviderCmd(ctx, from_provider_fqn, to_provider_fqn)
	if err != nil {
		return err
	}
	return tf.runTerraformCmd(ctx, cmd)
}

func (tf *Terraform) stateReplaceProviderCmd(ctx context.Context, from_provider_fqn string, to_provider_fqn string) (*exec.Cmd, error) {
	args := []string{"state", "replace-provider", "-no-color", "-auto-approve"}

	// positional arguments
	args = append(args, from_provider_fqn)
	args = append(args, to_provider_fqn)

	return tf.buildTerraformCmd(ctx, nil, args...), nil
}
