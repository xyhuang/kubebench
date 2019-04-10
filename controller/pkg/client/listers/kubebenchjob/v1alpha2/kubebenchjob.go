// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kubeflow/kubebench/controller/pkg/apis/kubebenchjob/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubebenchJobLister helps list KubebenchJobs.
type KubebenchJobLister interface {
	// List lists all KubebenchJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.KubebenchJob, err error)
	// KubebenchJobs returns an object that can list and get KubebenchJobs.
	KubebenchJobs(namespace string) KubebenchJobNamespaceLister
	KubebenchJobListerExpansion
}

// kubebenchJobLister implements the KubebenchJobLister interface.
type kubebenchJobLister struct {
	indexer cache.Indexer
}

// NewKubebenchJobLister returns a new KubebenchJobLister.
func NewKubebenchJobLister(indexer cache.Indexer) KubebenchJobLister {
	return &kubebenchJobLister{indexer: indexer}
}

// List lists all KubebenchJobs in the indexer.
func (s *kubebenchJobLister) List(selector labels.Selector) (ret []*v1alpha2.KubebenchJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.KubebenchJob))
	})
	return ret, err
}

// KubebenchJobs returns an object that can list and get KubebenchJobs.
func (s *kubebenchJobLister) KubebenchJobs(namespace string) KubebenchJobNamespaceLister {
	return kubebenchJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubebenchJobNamespaceLister helps list and get KubebenchJobs.
type KubebenchJobNamespaceLister interface {
	// List lists all KubebenchJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.KubebenchJob, err error)
	// Get retrieves the KubebenchJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.KubebenchJob, error)
	KubebenchJobNamespaceListerExpansion
}

// kubebenchJobNamespaceLister implements the KubebenchJobNamespaceLister
// interface.
type kubebenchJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubebenchJobs in the indexer for a given namespace.
func (s kubebenchJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.KubebenchJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.KubebenchJob))
	})
	return ret, err
}

// Get retrieves the KubebenchJob from the indexer for a given namespace and name.
func (s kubebenchJobNamespaceLister) Get(name string) (*v1alpha2.KubebenchJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("kubebenchjob"), name)
	}
	return obj.(*v1alpha2.KubebenchJob), nil
}
