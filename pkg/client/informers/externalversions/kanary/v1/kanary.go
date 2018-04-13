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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	kanary_v1 "github.com/etiennecoutaud/kanary/pkg/apis/kanary/v1"
	versioned "github.com/etiennecoutaud/kanary/pkg/client/clientset/versioned"
	internalinterfaces "github.com/etiennecoutaud/kanary/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/etiennecoutaud/kanary/pkg/client/listers/kanary/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KanaryInformer provides access to a shared informer and lister for
// Kanaries.
type KanaryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KanaryLister
}

type kanaryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKanaryInformer constructs a new informer for Kanary type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKanaryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKanaryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKanaryInformer constructs a new informer for Kanary type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKanaryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KanaryV1().Kanaries(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KanaryV1().Kanaries(namespace).Watch(options)
			},
		},
		&kanary_v1.Kanary{},
		resyncPeriod,
		indexers,
	)
}

func (f *kanaryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKanaryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kanaryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kanary_v1.Kanary{}, f.defaultInformer)
}

func (f *kanaryInformer) Lister() v1.KanaryLister {
	return v1.NewKanaryLister(f.Informer().GetIndexer())
}