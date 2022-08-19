


package managednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedNetworkGroup(ctx *pulumi.Context, args *LookupManagedNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkGroupResult, error) {
	var rv LookupManagedNetworkGroupResult
	err := ctx.Invoke("azure-native:managednetwork:getManagedNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkGroupArgs struct {
	ManagedNetworkGroupName string `pulumi:"managedNetworkGroupName"`
	ManagedNetworkName      string `pulumi:"managedNetworkName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupManagedNetworkGroupResult struct {
	Etag              string               `pulumi:"etag"`
	Id                string               `pulumi:"id"`
	Kind              *string              `pulumi:"kind"`
	Location          *string              `pulumi:"location"`
	ManagementGroups  []ResourceIdResponse `pulumi:"managementGroups"`
	Name              string               `pulumi:"name"`
	ProvisioningState string               `pulumi:"provisioningState"`
	Subnets           []ResourceIdResponse `pulumi:"subnets"`
	Subscriptions     []ResourceIdResponse `pulumi:"subscriptions"`
	Type              string               `pulumi:"type"`
	VirtualNetworks   []ResourceIdResponse `pulumi:"virtualNetworks"`
}

func LookupManagedNetworkGroupOutput(ctx *pulumi.Context, args LookupManagedNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) LookupManagedNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedNetworkGroupResult, error) {
			args := v.(LookupManagedNetworkGroupArgs)
			r, err := LookupManagedNetworkGroup(ctx, &args, opts...)
			var s LookupManagedNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedNetworkGroupResultOutput)
}

type LookupManagedNetworkGroupOutputArgs struct {
	ManagedNetworkGroupName pulumi.StringInput `pulumi:"managedNetworkGroupName"`
	ManagedNetworkName      pulumi.StringInput `pulumi:"managedNetworkName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkGroupArgs)(nil)).Elem()
}


type LookupManagedNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (LookupManagedNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkGroupResult)(nil)).Elem()
}

func (o LookupManagedNetworkGroupResultOutput) ToLookupManagedNetworkGroupResultOutput() LookupManagedNetworkGroupResultOutput {
	return o
}

func (o LookupManagedNetworkGroupResultOutput) ToLookupManagedNetworkGroupResultOutputWithContext(ctx context.Context) LookupManagedNetworkGroupResultOutput {
	return o
}

func (o LookupManagedNetworkGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupManagedNetworkGroupResultOutput) ManagementGroups() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.ManagementGroups }).(ResourceIdResponseArrayOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Subnets() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.Subnets }).(ResourceIdResponseArrayOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Subscriptions() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.Subscriptions }).(ResourceIdResponseArrayOutput)
}

func (o LookupManagedNetworkGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkGroupResultOutput) VirtualNetworks() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.VirtualNetworks }).(ResourceIdResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedNetworkGroupResultOutput{})
}
