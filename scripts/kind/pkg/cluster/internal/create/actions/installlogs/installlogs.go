/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package installcni implements the install CNI action
package installlogs

import (
	"sigs.k8s.io/kind/pkg/errors"

	"sigs.k8s.io/kind/pkg/cluster/internal/create/actions"
	"sigs.k8s.io/kind/pkg/cluster/nodeutils"
)

type action struct{}

// NewAction returns a new action for installing default CNI
func NewAction() actions.Action {
	return &action{}
}

// Execute runs the action
func (a *action) Execute(ctx *actions.ActionContext) error {
	ctx.Status.Start("Installing logs db 🔌")
	defer ctx.Status.End(false)

	allNodes, err := ctx.Nodes()
	if err != nil {
		return err
	}

	// get the target node for this task
	controlPlanes, err := nodeutils.ControlPlaneNodes(allNodes)
	if err != nil {
		return err
	}
	node := controlPlanes[0] // kind expects at least one always

	// install the logs db
	if err := node.Command(
		"helm",
		"repo",
		"add",
		"loki",
		"https://grafana.github.io/loki/charts",
	).Run(); err != nil {
		return errors.Wrap(err, "failed to install logs db")
	}

	if err := node.Command(
		"helm",
		"repo",
		"update",
	).Run(); err != nil {
		return errors.Wrap(err, "failed to install logs db")
	}

	if err := node.Command(
		"helm",
		"upgrade",
		"--install",
		"loki",
		"loki/loki",
	).Run(); err != nil {
		return errors.Wrap(err, "failed to install logs db")
	}

	// mark success
	ctx.Status.End(true)
	return nil
}
