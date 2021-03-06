// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterOperatorLister helps list ClusterOperators.
type ClusterOperatorLister interface {
	// List lists all ClusterOperators in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterOperator, err error)
	// Get retrieves the ClusterOperator from the index for a given name.
	Get(name string) (*v1.ClusterOperator, error)
	ClusterOperatorListerExpansion
}

// clusterOperatorLister implements the ClusterOperatorLister interface.
type clusterOperatorLister struct {
	indexer cache.Indexer
}

// NewClusterOperatorLister returns a new ClusterOperatorLister.
func NewClusterOperatorLister(indexer cache.Indexer) ClusterOperatorLister {
	return &clusterOperatorLister{indexer: indexer}
}

// List lists all ClusterOperators in the indexer.
func (s *clusterOperatorLister) List(selector labels.Selector) (ret []*v1.ClusterOperator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterOperator))
	})
	return ret, err
}

// Get retrieves the ClusterOperator from the index for a given name.
func (s *clusterOperatorLister) Get(name string) (*v1.ClusterOperator, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusteroperator"), name)
	}
	return obj.(*v1.ClusterOperator), nil
}
