//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package request

import "github.com/lastbackend/lastbackend/pkg/distribution/types"

type IngressMetaOptions struct {
	Meta *types.IngressUpdateMetaOptions `json:"meta"`
}

type IngressConnectOptions struct {
	Status types.IngressStatus `json:"status"`
}

type IngressStatusOptions struct {
	Ready bool `ready`
	// Pods statuses
	Routes map[string]*IngressRouteStatusOptions `json:"routes"`
}

type IngressRouteStatusOptions struct {
	// route status state
	State string `json:"state" yaml:"state"`
	// route status message
	Message string `json:"message" yaml:"message"`
}

type IngressRemoveOptions struct {
	Force bool `json:"force"`
}

type IngressLogsOptions struct {
	Follow bool `json:"follow"`
}
