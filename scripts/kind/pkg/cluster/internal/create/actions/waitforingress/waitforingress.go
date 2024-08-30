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

// Package waitforready implements the wait for ready action
package waitforingress

import (
	"fmt"
	"strings"
	"time"

	"sigs.k8s.io/kind/pkg/cluster/internal/create/actions"
	"sigs.k8s.io/kind/pkg/cluster/nodes"
	"sigs.k8s.io/kind/pkg/cluster/nodeutils"
	"sigs.k8s.io/kind/pkg/exec"
)

// Action implements an action for waiting for the cluster to be ready
type Action struct {
	waitTime time.Duration
}

// NewAction returns a new action for waiting for the cluster to be ready
func NewAction(waitTime time.Duration) actions.Action {
	return &Action{
		waitTime: waitTime,
	}
}

// Execute runs the action
func (a *Action) Execute(ctx *actions.ActionContext) error {
	// skip entirely if the wait time is 0
	if a.waitTime == time.Duration(0) {
		return nil
	}
	ctx.Status.Start(
		fmt.Sprintf(
			"Waiting ≤ %s for Ingress Controller = Ready ⏳",
			formatDuration(a.waitTime),
		),
	)

	allNodes, err := ctx.Nodes()
	if err != nil {
		return err
	}
	// get a control plane node to use to check cluster status
	controlPlanes, err := nodeutils.ControlPlaneNodes(allNodes)
	if err != nil {
		return err
	}
	node := controlPlanes[0] // kind expects at least one always

	// Wait for the nodes to reach Ready status.
	startTime := time.Now()

	isReady := waitForReady(ctx, node, startTime.Add(a.waitTime))
	if !isReady {
		ctx.Status.End(false)
		ctx.Logger.V(0).Info(" • WARNING: Timed out waiting for Ready ⚠️")
		return nil
	}

	// mark success
	ctx.Status.End(true)
	ctx.Logger.V(0).Infof(" • Ready after %s 💚", formatDuration(time.Since(startTime)))
	return nil
}

// WaitForReady uses kubectl inside the "node" container to check if the
// clusterctl nodes are "Ready".
func waitForReady(ctx *actions.ActionContext, node nodes.Node, until time.Time) bool {
	return tryUntil(until, func() bool {
		cmd := node.Command(
			"kubectl",
			"get",
			"pods",
			"--no-headers",
			"-n=ingress-nginx",
			"-o=jsonpath='{.items[*]['status.conditions..status']}'",
		)
		lines, err := exec.OutputLines(cmd)
		if err != nil {
			return false
		}
		status := strings.Fields(lines[0])
		ctx.Logger.V(0).Infof("%s", status)
		for _, s := range status {
			if !strings.Contains(s, "True") {
				return false
			}
		}
		return true
	})
}

// helper that calls `try()`` in a loop until the deadline `until`
// has passed or `try()`returns true, returns whether try ever returned true
func tryUntil(until time.Time, try func() bool) bool {
	for until.After(time.Now()) {
		if try() {
			return true
		}
	}
	return false
}

func formatDuration(duration time.Duration) string {
	return duration.Round(time.Second).String()
}
