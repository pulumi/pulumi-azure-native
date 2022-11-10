


package machinelearningexperimentation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:machinelearningexperimentation:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	AccountName       string `pulumi:"accountName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupProjectResult struct {
	AccountId         string            `pulumi:"accountId"`
	CreationDate      string            `pulumi:"creationDate"`
	Description       *string           `pulumi:"description"`
	FriendlyName      string            `pulumi:"friendlyName"`
	Gitrepo           *string           `pulumi:"gitrepo"`
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProjectId         string            `pulumi:"projectId"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
	WorkspaceId       string            `pulumi:"workspaceId"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}


type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Gitrepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Gitrepo }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
