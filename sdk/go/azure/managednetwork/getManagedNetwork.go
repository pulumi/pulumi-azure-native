


package managednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedNetwork(ctx *pulumi.Context, args *LookupManagedNetworkArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkResult, error) {
	var rv LookupManagedNetworkResult
	err := ctx.Invoke("azure-native:managednetwork:getManagedNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkArgs struct {
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupManagedNetworkResult struct {
	Connectivity      ConnectivityCollectionResponse `pulumi:"connectivity"`
	Etag              string                         `pulumi:"etag"`
	Id                string                         `pulumi:"id"`
	Location          string                         `pulumi:"location"`
	Name              string                         `pulumi:"name"`
	ProvisioningState string                         `pulumi:"provisioningState"`
	Scope             *ScopeResponse                 `pulumi:"scope"`
	Tags              map[string]string              `pulumi:"tags"`
	Type              string                         `pulumi:"type"`
}

func LookupManagedNetworkOutput(ctx *pulumi.Context, args LookupManagedNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupManagedNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedNetworkResult, error) {
			args := v.(LookupManagedNetworkArgs)
			r, err := LookupManagedNetwork(ctx, &args, opts...)
			var s LookupManagedNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedNetworkResultOutput)
}

type LookupManagedNetworkOutputArgs struct {
	ManagedNetworkName pulumi.StringInput `pulumi:"managedNetworkName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkArgs)(nil)).Elem()
}


type LookupManagedNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupManagedNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkResult)(nil)).Elem()
}

func (o LookupManagedNetworkResultOutput) ToLookupManagedNetworkResultOutput() LookupManagedNetworkResultOutput {
	return o
}

func (o LookupManagedNetworkResultOutput) ToLookupManagedNetworkResultOutputWithContext(ctx context.Context) LookupManagedNetworkResultOutput {
	return o
}

func (o LookupManagedNetworkResultOutput) Connectivity() ConnectivityCollectionResponseOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) ConnectivityCollectionResponse { return v.Connectivity }).(ConnectivityCollectionResponseOutput)
}

func (o LookupManagedNetworkResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedNetworkResultOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o LookupManagedNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedNetworkResultOutput{})
}
