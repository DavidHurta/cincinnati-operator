package cincinnati

import (
	cv1alpha1 "github.com/openshift/cincinnati-operator/pkg/apis/cincinnati/v1alpha1"
)

const (
	// NameContainerGraphBuilder is the Name property of the graph builder container
	NameContainerGraphBuilder string = "graph-builder"
	// NameContainerPolicyEngine is the Name property of the policy engine container
	NameContainerPolicyEngine string = "policy-engine"
	// NameInitContainerGraphData is the Name property of the graph data container
	NameInitContainerGraphData string = "graph-data"
	// openshiftConfigNamespace is the name of openshift's configuration namespace
	openshiftConfigNamespace = "openshift-config"
)

func nameDeployment(instance *cv1alpha1.Cincinnati) string {
	return instance.Name
}

func namePodDisruptionBudget(instance *cv1alpha1.Cincinnati) string {
	return instance.Name
}

func nameEnvConfig(instance *cv1alpha1.Cincinnati) string {
	return instance.Name + "-env"
}

func nameConfig(instance *cv1alpha1.Cincinnati) string {
	return instance.Name + "-config"
}

func namePolicyEngineService(instance *cv1alpha1.Cincinnati) string {
	return instance.Name + "-policy-engine"
}

func nameGraphBuilderService(instance *cv1alpha1.Cincinnati) string {
	return instance.Name + "-graph-builder"
}

func nameAdditionalTrustedCA(instance *cv1alpha1.Cincinnati) string {
	return instance.Name + "-trusted-ca"
}

func nameDeploymentTrustedCA() string {
	return "trusted-ca"
}