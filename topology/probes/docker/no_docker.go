// +build !linux

/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package docker

import (
	"github.com/skydive-project/skydive/common"
	ns "github.com/skydive-project/skydive/topology/probes/netns"
)

// ProbeHandler describes a Docker topology graph that enhance the graph
type ProbeHandler struct {
}

// Start the probe handler
func (p *ProbeHandler) Start() {
}

// Stop the probe handler
func (p *ProbeHandler) Stop() {
}

// NewProbeHandler creates a new topology Docker probe
func NewProbeHandler(nsProbe *ns.Probe, dockerURL, netnsRunPath string) (*ProbeHandler, error) {
	return nil, common.ErrNotImplemented
}
