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
// This file was automatically generated by lister-gen

package internalversion

import (
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OAuthAccessTokenLister helps list OAuthAccessTokens.
type OAuthAccessTokenLister interface {
	// List lists all OAuthAccessTokens in the indexer.
	List(selector labels.Selector) (ret []*oauth.OAuthAccessToken, err error)
	// Get retrieves the OAuthAccessToken from the index for a given name.
	Get(name string) (*oauth.OAuthAccessToken, error)
	OAuthAccessTokenListerExpansion
}

// oAuthAccessTokenLister implements the OAuthAccessTokenLister interface.
type oAuthAccessTokenLister struct {
	indexer cache.Indexer
}

// NewOAuthAccessTokenLister returns a new OAuthAccessTokenLister.
func NewOAuthAccessTokenLister(indexer cache.Indexer) OAuthAccessTokenLister {
	return &oAuthAccessTokenLister{indexer: indexer}
}

// List lists all OAuthAccessTokens in the indexer.
func (s *oAuthAccessTokenLister) List(selector labels.Selector) (ret []*oauth.OAuthAccessToken, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*oauth.OAuthAccessToken))
	})
	return ret, err
}

// Get retrieves the OAuthAccessToken from the index for a given name.
func (s *oAuthAccessTokenLister) Get(name string) (*oauth.OAuthAccessToken, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(oauth.Resource("oauthaccesstoken"), name)
	}
	return obj.(*oauth.OAuthAccessToken), nil
}
