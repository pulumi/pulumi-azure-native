


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20190501:getWorkspace", args, &rv, opts...)
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
	ApplicationInsights *string           `pulumi:"applicationInsights"`
	ContainerRegistry   *string           `pulumi:"containerRegistry"`
	CreationTime        string            `pulumi:"creationTime"`
	Description         *string           `pulumi:"description"`
	DiscoveryUrl        *string           `pulumi:"discoveryUrl"`
	FriendlyName        *string           `pulumi:"friendlyName"`
	Id                  string            `pulumi:"id"`
	Identity            *IdentityResponse `pulumi:"identity"`
	KeyVault            *string           `pulumi:"keyVault"`
	Location            *string           `pulumi:"location"`
	Name                string            `pulumi:"name"`
	ProvisioningState   string            `pulumi:"provisioningState"`
	StorageAccount      *string           `pulumi:"storageAccount"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
	WorkspaceId         string            `pulumi:"workspaceId"`
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

func (o LookupWorkspaceResultOutput) ApplicationInsights() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ApplicationInsights }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) ContainerRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ContainerRegistry }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) KeyVault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.KeyVault }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
