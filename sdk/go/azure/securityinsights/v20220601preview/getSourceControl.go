


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:securityinsights/v20220601preview:getSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SourceControlId   string `pulumi:"sourceControlId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSourceControlResult struct {
	ContentTypes           []string                        `pulumi:"contentTypes"`
	Description            *string                         `pulumi:"description"`
	DisplayName            string                          `pulumi:"displayName"`
	Etag                   *string                         `pulumi:"etag"`
	Id                     string                          `pulumi:"id"`
	LastDeploymentInfo     *DeploymentInfoResponse         `pulumi:"lastDeploymentInfo"`
	Name                   string                          `pulumi:"name"`
	RepoType               string                          `pulumi:"repoType"`
	Repository             RepositoryResponse              `pulumi:"repository"`
	RepositoryResourceInfo *RepositoryResourceInfoResponse `pulumi:"repositoryResourceInfo"`
	SystemData             SystemDataResponse              `pulumi:"systemData"`
	Type                   string                          `pulumi:"type"`
	Version                *string                         `pulumi:"version"`
}

func LookupSourceControlOutput(ctx *pulumi.Context, args LookupSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceControlResult, error) {
			args := v.(LookupSourceControlArgs)
			r, err := LookupSourceControl(ctx, &args, opts...)
			var s LookupSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceControlResultOutput)
}

type LookupSourceControlOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SourceControlId   pulumi.StringInput `pulumi:"sourceControlId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlArgs)(nil)).Elem()
}


type LookupSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlResult)(nil)).Elem()
}

func (o LookupSourceControlResultOutput) ToLookupSourceControlResultOutput() LookupSourceControlResultOutput {
	return o
}

func (o LookupSourceControlResultOutput) ToLookupSourceControlResultOutputWithContext(ctx context.Context) LookupSourceControlResultOutput {
	return o
}

func (o LookupSourceControlResultOutput) ContentTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSourceControlResult) []string { return v.ContentTypes }).(pulumi.StringArrayOutput)
}

func (o LookupSourceControlResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) LastDeploymentInfo() DeploymentInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *DeploymentInfoResponse { return v.LastDeploymentInfo }).(DeploymentInfoResponsePtrOutput)
}

func (o LookupSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) RepoType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.RepoType }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) Repository() RepositoryResponseOutput {
	return o.ApplyT(func(v LookupSourceControlResult) RepositoryResponse { return v.Repository }).(RepositoryResponseOutput)
}

func (o LookupSourceControlResultOutput) RepositoryResourceInfo() RepositoryResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *RepositoryResourceInfoResponse { return v.RepositoryResourceInfo }).(RepositoryResourceInfoResponsePtrOutput)
}

func (o LookupSourceControlResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSourceControlResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSourceControlResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceControlResultOutput{})
}
