// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "open-cluster-management.io/api/client/cluster/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AddOnPlacementScores returns a AddOnPlacementScoreInformer.
	AddOnPlacementScores() AddOnPlacementScoreInformer
	// ClusterClaims returns a ClusterClaimInformer.
	ClusterClaims() ClusterClaimInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AddOnPlacementScores returns a AddOnPlacementScoreInformer.
func (v *version) AddOnPlacementScores() AddOnPlacementScoreInformer {
	return &addOnPlacementScoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterClaims returns a ClusterClaimInformer.
func (v *version) ClusterClaims() ClusterClaimInformer {
	return &clusterClaimInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
