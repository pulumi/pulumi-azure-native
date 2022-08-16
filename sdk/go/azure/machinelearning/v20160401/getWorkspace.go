


package v20160401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearning/v20160401:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	CreationTime         string            `pulumi:"creationTime"`
	Id                   string            `pulumi:"id"`
	KeyVaultIdentifierId *string           `pulumi:"keyVaultIdentifierId"`
	Location             string            `pulumi:"location"`
	Name                 string            `pulumi:"name"`
	OwnerEmail           string            `pulumi:"ownerEmail"`
	StudioEndpoint       string            `pulumi:"studioEndpoint"`
	Tags                 map[string]string `pulumi:"tags"`
	Type                 string            `pulumi:"type"`
	UserStorageAccountId string            `pulumi:"userStorageAccountId"`
	WorkspaceId          string            `pulumi:"workspaceId"`
	WorkspaceState       string            `pulumi:"workspaceState"`
	WorkspaceType        string            `pulumi:"workspaceType"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

type LookupWorkspaceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}


type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) KeyVaultIdentifierId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.KeyVaultIdentifierId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) OwnerEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.OwnerEmail }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) StudioEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.StudioEndpoint }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) UserStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.UserStorageAccountId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceState }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
