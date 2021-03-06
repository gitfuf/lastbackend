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

package views

import (
	"time"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
)

// Ingress - default node structure
type Ingress struct {
	Meta   IngressMeta   `json:"meta"`
	Status IngressStatus `json:"status"`
}

// IngressList - node map list
type IngressList map[string]*Ingress

// IngressMeta - node metadata structure
type IngressMeta struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

// IngressStatus - node state struct
type IngressStatus struct {
	Ready bool `json:"ready"`
}


type IngressSpec struct {
	Routes  map[string]types.RouteSpec   `json:"routes"`
}
