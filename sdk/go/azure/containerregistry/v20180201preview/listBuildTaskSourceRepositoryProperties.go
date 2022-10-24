


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBuildTaskSourceRepositoryProperties(ctx *pulumi.Context, args *ListBuildTaskSourceRepositoryPropertiesArgs, opts ...pulumi.InvokeOption) (*ListBuildTaskSourceRepositoryPropertiesResult, error) {
	var rv ListBuildTaskSourceRepositoryPropertiesResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:listBuildTaskSourceRepositoryProperties", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListBuildTaskSourceRepositoryPropertiesArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBuildTaskSourceRepositoryPropertiesResult struct {
	IsCommitTriggerEnabled      *bool                          `pulumi:"isCommitTriggerEnabled"`
	RepositoryUrl               string                         `pulumi:"repositoryUrl"`
	SourceControlAuthProperties *SourceControlAuthInfoResponse `pulumi:"sourceControlAuthProperties"`
	SourceControlType           string                         `pulumi:"sourceControlType"`
}


func (val *ListBuildTaskSourceRepositoryPropertiesResult) Defaults() *ListBuildTaskSourceRepositoryPropertiesResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsCommitTriggerEnabled) {
		isCommitTriggerEnabled_ := false
		tmp.IsCommitTriggerEnabled = &isCommitTriggerEnabled_
	}
	return &tmp
}

func ListBuildTaskSourceRepositoryPropertiesOutput(ctx *pulumi.Context, args ListBuildTaskSourceRepositoryPropertiesOutputArgs, opts ...pulumi.InvokeOption) ListBuildTaskSourceRepositoryPropertiesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBuildTaskSourceRepositoryPropertiesResult, error) {
			args := v.(ListBuildTaskSourceRepositoryPropertiesArgs)
			r, err := ListBuildTaskSourceRepositoryProperties(ctx, &args, opts...)
			var s ListBuildTaskSourceRepositoryPropertiesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBuildTaskSourceRepositoryPropertiesResultOutput)
}

type ListBuildTaskSourceRepositoryPropertiesOutputArgs struct {
	BuildTaskName     pulumi.StringInput `pulumi:"buildTaskName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBuildTaskSourceRepositoryPropertiesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildTaskSourceRepositoryPropertiesArgs)(nil)).Elem()
}


type ListBuildTaskSourceRepositoryPropertiesResultOutput struct{ *pulumi.OutputState }

func (ListBuildTaskSourceRepositoryPropertiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildTaskSourceRepositoryPropertiesResult)(nil)).Elem()
}

func (o ListBuildTaskSourceRepositoryPropertiesResultOutput) ToListBuildTaskSourceRepositoryPropertiesResultOutput() ListBuildTaskSourceRepositoryPropertiesResultOutput {
	return o
}

func (o ListBuildTaskSourceRepositoryPropertiesResultOutput) ToListBuildTaskSourceRepositoryPropertiesResultOutputWithContext(ctx context.Context) ListBuildTaskSourceRepositoryPropertiesResultOutput {
	return o
}

func (o ListBuildTaskSourceRepositoryPropertiesResultOutput) IsCommitTriggerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListBuildTaskSourceRepositoryPropertiesResult) *bool { return v.IsCommitTriggerEnabled }).(pulumi.BoolPtrOutput)
}

func (o ListBuildTaskSourceRepositoryPropertiesResultOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ListBuildTaskSourceRepositoryPropertiesResult) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func (o ListBuildTaskSourceRepositoryPropertiesResultOutput) SourceControlAuthProperties() SourceControlAuthInfoResponsePtrOutput {
	return o.ApplyT(func(v ListBuildTaskSourceRepositoryPropertiesResult) *SourceControlAuthInfoResponse {
		return v.SourceControlAuthProperties
	}).(SourceControlAuthInfoResponsePtrOutput)
}

func (o ListBuildTaskSourceRepositoryPropertiesResultOutput) SourceControlType() pulumi.StringOutput {
	return o.ApplyT(func(v ListBuildTaskSourceRepositoryPropertiesResult) string { return v.SourceControlType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBuildTaskSourceRepositoryPropertiesResultOutput{})
}
