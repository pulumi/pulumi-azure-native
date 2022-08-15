


package v20210827

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedPrivateEndpoint(ctx *pulumi.Context, args *LookupManagedPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupManagedPrivateEndpointResult, error) {
	var rv LookupManagedPrivateEndpointResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getManagedPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPrivateEndpointArgs struct {
	ClusterName                string `pulumi:"clusterName"`
	ManagedPrivateEndpointName string `pulumi:"managedPrivateEndpointName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupManagedPrivateEndpointResult struct {
	GroupId                   string             `pulumi:"groupId"`
	Id                        string             `pulumi:"id"`
	Name                      string             `pulumi:"name"`
	PrivateLinkResourceId     string             `pulumi:"privateLinkResourceId"`
	PrivateLinkResourceRegion *string            `pulumi:"privateLinkResourceRegion"`
	ProvisioningState         string             `pulumi:"provisioningState"`
	RequestMessage            *string            `pulumi:"requestMessage"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
}

func LookupManagedPrivateEndpointOutput(ctx *pulumi.Context, args LookupManagedPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupManagedPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedPrivateEndpointResult, error) {
			args := v.(LookupManagedPrivateEndpointArgs)
			r, err := LookupManagedPrivateEndpoint(ctx, &args, opts...)
			var s LookupManagedPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedPrivateEndpointResultOutput)
}

type LookupManagedPrivateEndpointOutputArgs struct {
	ClusterName                pulumi.StringInput `pulumi:"clusterName"`
	ManagedPrivateEndpointName pulumi.StringInput `pulumi:"managedPrivateEndpointName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointArgs)(nil)).Elem()
}


type LookupManagedPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupManagedPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointResult)(nil)).Elem()
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutput() LookupManagedPrivateEndpointResultOutput {
	return o
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutputWithContext(ctx context.Context) LookupManagedPrivateEndpointResultOutput {
	return o
}

func (o LookupManagedPrivateEndpointResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) PrivateLinkResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.PrivateLinkResourceRegion }).(pulumi.StringPtrOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedPrivateEndpointResultOutput{})
}
