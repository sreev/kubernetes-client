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
/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1beta1

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EvictionLister helps list Evictions.
type EvictionLister interface {
	// List lists all Evictions in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Eviction, err error)
	// Evictions returns an object that can list and get Evictions.
	Evictions(namespace string) EvictionNamespaceLister
	EvictionListerExpansion
}

// evictionLister implements the EvictionLister interface.
type evictionLister struct {
	indexer cache.Indexer
}

// NewEvictionLister returns a new EvictionLister.
func NewEvictionLister(indexer cache.Indexer) EvictionLister {
	return &evictionLister{indexer: indexer}
}

// List lists all Evictions in the indexer.
func (s *evictionLister) List(selector labels.Selector) (ret []*v1beta1.Eviction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Eviction))
	})
	return ret, err
}

// Evictions returns an object that can list and get Evictions.
func (s *evictionLister) Evictions(namespace string) EvictionNamespaceLister {
	return evictionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EvictionNamespaceLister helps list and get Evictions.
type EvictionNamespaceLister interface {
	// List lists all Evictions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.Eviction, err error)
	// Get retrieves the Eviction from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.Eviction, error)
	EvictionNamespaceListerExpansion
}

// evictionNamespaceLister implements the EvictionNamespaceLister
// interface.
type evictionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Evictions in the indexer for a given namespace.
func (s evictionNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Eviction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Eviction))
	})
	return ret, err
}

// Get retrieves the Eviction from the indexer for a given namespace and name.
func (s evictionNamespaceLister) Get(name string) (*v1beta1.Eviction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("eviction"), name)
	}
	return obj.(*v1beta1.Eviction), nil
}
