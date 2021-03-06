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

package envs

import (
	"github.com/lastbackend/lastbackend/pkg/api/client"
	"github.com/lastbackend/lastbackend/pkg/storage"
	"github.com/lastbackend/lastbackend/pkg/api/cache"
)

var e Env

type Env struct {
	storage storage.Storage
	client  client.Client
	cache   *cache.Cache
}

func Get() *Env {
	return &e
}



func (c *Env) SetClient(client client.Client) {
	c.client = client
}

func (c *Env) GetClient() client.Client {
	return c.client
}

func (c *Env) SetStorage(storage storage.Storage) {
	c.storage = storage
}

func (c *Env) GetStorage() storage.Storage {
	return c.storage
}

func (c *Env) SetCache(ch *cache.Cache) {
	c.cache = ch
}

func (c *Env) GetCache() *cache.Cache {
	return c.cache
}