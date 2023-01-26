


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppVnetConnectionSlot(ctx *pulumi.Context, args *LookupWebAppVnetConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppVnetConnectionSlotResult, error) {
	var rv LookupWebAppVnetConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20200901:getWebAppVnetConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppVnetConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupWebAppVnetConnectionSlotResult struct {
	CertBlob       *string             `pulumi:"certBlob"`
	CertThumbprint string              `pulumi:"certThumbprint"`
	DnsServers     *string             `pulumi:"dnsServers"`
	Id             string              `pulumi:"id"`
	IsSwift        *bool               `pulumi:"isSwift"`
	Kind           *string             `pulumi:"kind"`
	Name           string              `pulumi:"name"`
	ResyncRequired bool                `pulumi:"resyncRequired"`
	Routes         []VnetRouteResponse `pulumi:"routes"`
	SystemData     SystemDataResponse  `pulumi:"systemData"`
	Type           string              `pulumi:"type"`
	VnetResourceId *string             `pulumi:"vnetResourceId"`
}

func LookupWebAppVnetConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppVnetConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppVnetConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppVnetConnectionSlotResult, error) {
			args := v.(LookupWebAppVnetConnectionSlotArgs)
			r, err := LookupWebAppVnetConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppVnetConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppVnetConnectionSlotResultOutput)
}

type LookupWebAppVnetConnectionSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
	VnetName          pulumi.StringInput `pulumi:"vnetName"`
}

func (LookupWebAppVnetConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppVnetConnectionSlotArgs)(nil)).Elem()
}


type LookupWebAppVnetConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppVnetConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppVnetConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppVnetConnectionSlotResultOutput) ToLookupWebAppVnetConnectionSlotResultOutput() LookupWebAppVnetConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppVnetConnectionSlotResultOutput) ToLookupWebAppVnetConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppVnetConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppVnetConnectionSlotResultOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.CertBlob }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) CertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.CertThumbprint }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.DnsServers }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) IsSwift() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *bool { return v.IsSwift }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) ResyncRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) bool { return v.ResyncRequired }).(pulumi.BoolOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) []VnetRouteResponse { return v.Routes }).(VnetRouteResponseArrayOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionSlotResultOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionSlotResult) *string { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppVnetConnectionSlotResultOutput{})
}
