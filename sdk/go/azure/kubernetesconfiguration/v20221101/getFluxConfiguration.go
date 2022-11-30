


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFluxConfiguration(ctx *pulumi.Context, args *LookupFluxConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupFluxConfigurationResult, error) {
	var rv LookupFluxConfigurationResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20221101:getFluxConfiguration", args, &rv, opts...)
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
	AzureBlob                      *AzureBlobDefinitionResponse               `pulumi:"azureBlob"`
	Bucket                         *BucketDefinitionResponse                  `pulumi:"bucket"`
	ComplianceState                string                                     `pulumi:"complianceState"`
	ConfigurationProtectedSettings map[string]string                          `pulumi:"configurationProtectedSettings"`
	ErrorMessage                   string                                     `pulumi:"errorMessage"`
	GitRepository                  *GitRepositoryDefinitionResponse           `pulumi:"gitRepository"`
	Id                             string                                     `pulumi:"id"`
	Kustomizations                 map[string]KustomizationDefinitionResponse `pulumi:"kustomizations"`
	Name                           string                                     `pulumi:"name"`
	Namespace                      *string                                    `pulumi:"namespace"`
	ProvisioningState              string                                     `pulumi:"provisioningState"`
	RepositoryPublicKey            string                                     `pulumi:"repositoryPublicKey"`
	Scope                          *string                                    `pulumi:"scope"`
	SourceKind                     *string                                    `pulumi:"sourceKind"`
	SourceSyncedCommitId           string                                     `pulumi:"sourceSyncedCommitId"`
	SourceUpdatedAt                string                                     `pulumi:"sourceUpdatedAt"`
	StatusUpdatedAt                string                                     `pulumi:"statusUpdatedAt"`
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
	tmp.AzureBlob = tmp.AzureBlob.Defaults()

	tmp.Bucket = tmp.Bucket.Defaults()

	tmp.GitRepository = tmp.GitRepository.Defaults()

	if isZero(tmp.Namespace) {
		namespace_ := "default"
		tmp.Namespace = &namespace_
	}
	if isZero(tmp.SourceKind) {
		sourceKind_ := "GitRepository"
		tmp.SourceKind = &sourceKind_
	}
	if isZero(tmp.Suspend) {
		suspend_ := false
		tmp.Suspend = &suspend_
	}
	return &tmp
}

func LookupFluxConfigurationOutput(ctx *pulumi.Context, args LookupFluxConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupFluxConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFluxConfigurationResult, error) {
			args := v.(LookupFluxConfigurationArgs)
			r, err := LookupFluxConfiguration(ctx, &args, opts...)
			var s LookupFluxConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFluxConfigurationResultOutput)
}

type LookupFluxConfigurationOutputArgs struct {
	ClusterName           pulumi.StringInput `pulumi:"clusterName"`
	ClusterResourceName   pulumi.StringInput `pulumi:"clusterResourceName"`
	ClusterRp             pulumi.StringInput `pulumi:"clusterRp"`
	FluxConfigurationName pulumi.StringInput `pulumi:"fluxConfigurationName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFluxConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluxConfigurationArgs)(nil)).Elem()
}


type LookupFluxConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupFluxConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluxConfigurationResult)(nil)).Elem()
}

func (o LookupFluxConfigurationResultOutput) ToLookupFluxConfigurationResultOutput() LookupFluxConfigurationResultOutput {
	return o
}

func (o LookupFluxConfigurationResultOutput) ToLookupFluxConfigurationResultOutputWithContext(ctx context.Context) LookupFluxConfigurationResultOutput {
	return o
}

func (o LookupFluxConfigurationResultOutput) AzureBlob() AzureBlobDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *AzureBlobDefinitionResponse { return v.AzureBlob }).(AzureBlobDefinitionResponsePtrOutput)
}

func (o LookupFluxConfigurationResultOutput) Bucket() BucketDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *BucketDefinitionResponse { return v.Bucket }).(BucketDefinitionResponsePtrOutput)
}

func (o LookupFluxConfigurationResultOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.ComplianceState }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) map[string]string { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

func (o LookupFluxConfigurationResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) GitRepository() GitRepositoryDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *GitRepositoryDefinitionResponse { return v.GitRepository }).(GitRepositoryDefinitionResponsePtrOutput)
}

func (o LookupFluxConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) Kustomizations() KustomizationDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) map[string]KustomizationDefinitionResponse {
		return v.Kustomizations
	}).(KustomizationDefinitionResponseMapOutput)
}

func (o LookupFluxConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupFluxConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) RepositoryPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.RepositoryPublicKey }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupFluxConfigurationResultOutput) SourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.SourceKind }).(pulumi.StringPtrOutput)
}

func (o LookupFluxConfigurationResultOutput) SourceSyncedCommitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.SourceSyncedCommitId }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) SourceUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.SourceUpdatedAt }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) StatusUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.StatusUpdatedAt }).(pulumi.StringOutput)
}

func (o LookupFluxConfigurationResultOutput) Statuses() ObjectStatusDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) []ObjectStatusDefinitionResponse { return v.Statuses }).(ObjectStatusDefinitionResponseArrayOutput)
}

func (o LookupFluxConfigurationResultOutput) Suspend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *bool { return v.Suspend }).(pulumi.BoolPtrOutput)
}

func (o LookupFluxConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFluxConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFluxConfigurationResultOutput{})
}
