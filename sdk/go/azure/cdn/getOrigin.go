


package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOrigin(ctx *pulumi.Context, args *LookupOriginArgs, opts ...pulumi.InvokeOption) (*LookupOriginResult, error) {
	var rv LookupOriginResult
	err := ctx.Invoke("azure-native:cdn:getOrigin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	OriginName        string `pulumi:"originName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOriginResult struct {
	Enabled                    *bool              `pulumi:"enabled"`
	HostName                   string             `pulumi:"hostName"`
	HttpPort                   *int               `pulumi:"httpPort"`
	HttpsPort                  *int               `pulumi:"httpsPort"`
	Id                         string             `pulumi:"id"`
	Name                       string             `pulumi:"name"`
	OriginHostHeader           *string            `pulumi:"originHostHeader"`
	Priority                   *int               `pulumi:"priority"`
	PrivateEndpointStatus      string             `pulumi:"privateEndpointStatus"`
	PrivateLinkAlias           *string            `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string            `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string            `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string            `pulumi:"privateLinkResourceId"`
	ProvisioningState          string             `pulumi:"provisioningState"`
	ResourceState              string             `pulumi:"resourceState"`
	SystemData                 SystemDataResponse `pulumi:"systemData"`
	Type                       string             `pulumi:"type"`
	Weight                     *int               `pulumi:"weight"`
}

func LookupOriginOutput(ctx *pulumi.Context, args LookupOriginOutputArgs, opts ...pulumi.InvokeOption) LookupOriginResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOriginResult, error) {
			args := v.(LookupOriginArgs)
			r, err := LookupOrigin(ctx, &args, opts...)
			var s LookupOriginResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOriginResultOutput)
}

type LookupOriginOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	OriginName        pulumi.StringInput `pulumi:"originName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOriginOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginArgs)(nil)).Elem()
}


type LookupOriginResultOutput struct{ *pulumi.OutputState }

func (LookupOriginResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginResult)(nil)).Elem()
}

func (o LookupOriginResultOutput) ToLookupOriginResultOutput() LookupOriginResultOutput {
	return o
}

func (o LookupOriginResultOutput) ToLookupOriginResultOutputWithContext(ctx context.Context) LookupOriginResultOutput {
	return o
}

func (o LookupOriginResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupOriginResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o LookupOriginResultOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o LookupOriginResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *string { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o LookupOriginResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupOriginResultOutput) PrivateEndpointStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.PrivateEndpointStatus }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o LookupOriginResultOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o LookupOriginResultOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *string { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o LookupOriginResultOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupOriginResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOriginResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOriginResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOriginResultOutput{})
}
