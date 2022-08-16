


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementGroupNetworkManagerConnection(ctx *pulumi.Context, args *LookupManagementGroupNetworkManagerConnectionArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupNetworkManagerConnectionResult, error) {
	var rv LookupManagementGroupNetworkManagerConnectionResult
	err := ctx.Invoke("azure-native:network/v20210501preview:getManagementGroupNetworkManagerConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupNetworkManagerConnectionArgs struct {
	ManagementGroupId            string `pulumi:"managementGroupId"`
	NetworkManagerConnectionName string `pulumi:"networkManagerConnectionName"`
}


type LookupManagementGroupNetworkManagerConnectionResult struct {
	Description      *string            `pulumi:"description"`
	Etag             string             `pulumi:"etag"`
	Id               string             `pulumi:"id"`
	Name             string             `pulumi:"name"`
	NetworkManagerId *string            `pulumi:"networkManagerId"`
	SystemData       SystemDataResponse `pulumi:"systemData"`
	Type             string             `pulumi:"type"`
}

func LookupManagementGroupNetworkManagerConnectionOutput(ctx *pulumi.Context, args LookupManagementGroupNetworkManagerConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupNetworkManagerConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupNetworkManagerConnectionResult, error) {
			args := v.(LookupManagementGroupNetworkManagerConnectionArgs)
			r, err := LookupManagementGroupNetworkManagerConnection(ctx, &args, opts...)
			var s LookupManagementGroupNetworkManagerConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupNetworkManagerConnectionResultOutput)
}

type LookupManagementGroupNetworkManagerConnectionOutputArgs struct {
	ManagementGroupId            pulumi.StringInput `pulumi:"managementGroupId"`
	NetworkManagerConnectionName pulumi.StringInput `pulumi:"networkManagerConnectionName"`
}

func (LookupManagementGroupNetworkManagerConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupNetworkManagerConnectionArgs)(nil)).Elem()
}


type LookupManagementGroupNetworkManagerConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupNetworkManagerConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupNetworkManagerConnectionResult)(nil)).Elem()
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) ToLookupManagementGroupNetworkManagerConnectionResultOutput() LookupManagementGroupNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) ToLookupManagementGroupNetworkManagerConnectionResultOutputWithContext(ctx context.Context) LookupManagementGroupNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) *string { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupNetworkManagerConnectionResultOutput{})
}
