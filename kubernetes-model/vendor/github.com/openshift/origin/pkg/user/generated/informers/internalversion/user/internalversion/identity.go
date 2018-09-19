/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// This file was automatically generated by informer-gen

package internalversion

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	internalinterfaces "github.com/openshift/origin/pkg/user/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/user/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/user/generated/listers/user/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// IdentityInformer provides access to a shared informer and lister for
// Identities.
type IdentityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.IdentityLister
}

type identityInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIdentityInformer constructs a new informer for Identity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIdentityInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIdentityInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIdentityInformer constructs a new informer for Identity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIdentityInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.User().Identities().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.User().Identities().Watch(options)
			},
		},
		&user.Identity{},
		resyncPeriod,
		indexers,
	)
}

func (f *identityInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIdentityInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *identityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&user.Identity{}, f.defaultInformer)
}

func (f *identityInformer) Lister() internalversion.IdentityLister {
	return internalversion.NewIdentityLister(f.Informer().GetIndexer())
}
