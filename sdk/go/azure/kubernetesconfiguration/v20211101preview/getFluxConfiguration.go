


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFluxConfiguration(ctx *pulumi.Context, args *LookupFluxConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupFluxConfigurationResult, error) {
	var rv LookupFluxConfigurationResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20211101preview:getFluxConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFluxConfigurationArgs struct {
	ClusterName           string `pulumi:"clusterName"`
	ClusterResourceName   string `pulumi:"clusterResourceName"`
	ClusterRp             string `pulumi:"clusterRp"`
	FluxConfigurationName string `pulumi:"fluxConfigurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupFluxConfigurationResult struct {
	ComplianceState                string                                     `pulumi:"complianceState"`
	ConfigurationProtectedSettings map[string]string                          `pulumi:"configurationProtectedSettings"`
	ErrorMessage                   string                                     `pulumi:"errorMessage"`
	GitRepository                  *GitRepositoryDefinitionResponse           `pulumi:"gitRepository"`
	Id                             string                                     `pulumi:"id"`
	Kustomizations                 map[string]KustomizationDefinitionResponse `pulumi:"kustomizations"`
	LastSourceSyncedAt             string                                     `pulumi:"lastSourceSyncedAt"`
	LastSourceSyncedCommitId       string                                     `pulumi:"lastSourceSyncedCommitId"`
	Name                           string                                     `pulumi:"name"`
	Namespace                      *string                                    `pulumi:"namespace"`
	ProvisioningState              string                                     `pulumi:"provisioningState"`
	RepositoryPublicKey            string                                     `pulumi:"repositoryPublicKey"`
	Scope                          *string                                    `pulumi:"scope"`
	SourceKind                     *string                                    `pulumi:"sourceKind"`
	Statuses                       []ObjectStatusDefinitionResponse           `pulumi:"statuses"`
	Suspend                        *bool                                      `pulumi:"suspend"`
	SystemData                     SystemDataResponse                         `pulumi:"systemData"`
	Type                           string                                     `pulumi:"type"`
}


func (val *LookupFluxConfigurationResult) Defaults() *LookupFluxConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.GitRepository = tmp.GitRepository.Defaults()

	if isZero(tmp.Namespace) {
		namespace_ := "default"
		tmp.Namespace = &namespace_
	}
	if isZero(tmp.Suspend) {
		suspend_ := false
		tmp.Suspend = &suspend_
	}
	return &tmp
}
