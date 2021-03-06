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

package service

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/controller/envs"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/log"
)

type Controller struct {
	spec   chan *types.Service
	status chan *types.Service
	active bool
}

// Watch services spec changes
func (sc *Controller) WatchSpec() {

	var (
		stg = envs.Get().GetStorage()
	)

	log.Debug("controller:service:controller: start watch service spec")
	go func() {
		for {
			select {
			case s := <-sc.spec:
				{
					if !sc.active {
						log.Debug("controller:service:controller: skip management course it is in slave mode")
						continue
					}

					if s == nil {
						log.Debug("controller:service:controller: skip because service is nil")
						continue
					}

					log.Debugf("controller:service:controller: Service needs to be provisioned: %s:%s", s.Meta.Namespace, s.Meta.Name)
					if err := Provision(s); err != nil {
						log.Errorf("controller:service:controller: service provision: %s err: %s", s.Meta.Name, err.Error())
					}
				}
			}
		}
	}()

	stg.Service().WatchSpec(context.Background(), sc.spec)
}

// Watch services spec changes
func (sc *Controller) WatchStatus() {

	var (
		stg = envs.Get().GetStorage()
		msg = "controller:service:watch_status:"
	)

	log.Debugf("%s> start watch service status", msg)
	go func() {
		for {
			select {
			case s := <-sc.status:
				{
					if !sc.active {
						log.Debugf("%s> skip management course it is in slave mode", msg)
						continue
					}

					if s == nil {
						log.Debugf("%s> skip because service is nil", msg)
						continue
					}

					log.Debugf("%s> Service needs to be provisioned: %s", msg, s.SelfLink())
					if err := HandleStatus(s); err != nil {
						log.Errorf("%s> service provision: %s err: %s", msg, s.SelfLink(), err.Error())
					}
				}
			}
		}
	}()

	stg.Service().WatchStatus(context.Background(), sc.status)
}

// Pause service controller because not lead
func (sc *Controller) Pause() {
	sc.active = false
}

// Resume service controller management
func (sc *Controller) Resume() {

	var (
		stg = envs.Get().GetStorage()
		msg = "controller:service:resume:"
	)

	sc.active = true

	log.Debugf("%s> start check services states", msg)
	nss, err := stg.Namespace().List(context.Background())
	if err != nil {
		log.Errorf("%s> get namespaces list err: %s", msg, err.Error())
	}

	for _, ns := range nss {
		svcs, err := stg.Service().ListByNamespace(context.Background(), ns.Meta.Name)
		if err != nil {
			log.Errorf("%s> get services list err: %s", msg, err.Error())
		}

		for _, svc := range svcs {
			svc, err := stg.Service().Get(context.Background(), svc.Meta.Namespace, svc.Meta.Name)
			if err != nil {
				log.Errorf("%s> get service err: %s", msg, err.Error())
			}
			sc.spec <- svc
		}

		for _, svc := range svcs {
			log.Debugf("%s> check service [%s] status", msg, svc.SelfLink())
			sc.status <- svc
		}
	}
}

// NewServiceController return new controller instance
func NewServiceController(_ context.Context) *Controller {
	sc := new(Controller)
	sc.active = false
	sc.spec = make(chan *types.Service)
	sc.status = make(chan *types.Service)
	return sc
}
