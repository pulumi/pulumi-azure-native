


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppVnetConnection(ctx *pulumi.Context, args *LookupWebAppVnetConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppVnetConnectionResult, error) {
	var rv LookupWebAppVnetConnectionResult
	err := ctx.Invoke("azure-native:web/v20200901:getWebAppVnetConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppVnetConnectionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupWebAppVnetConnectionResult struct {
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

func LookupWebAppVnetConnectionOutput(ctx *pulumi.Context, args LookupWebAppVnetConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppVnetConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppVnetConnectionResult, error) {
			args := v.(LookupWebAppVnetConnectionArgs)
			r, err := LookupWebAppVnetConnection(ctx, &args, opts...)
			var s LookupWebAppVnetConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppVnetConnectionResultOutput)
}

type LookupWebAppVnetConnectionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VnetName          pulumi.StringInput `pulumi:"vnetName"`
}

func (LookupWebAppVnetConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppVnetConnectionArgs)(nil)).Elem()
}


type LookupWebAppVnetConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppVnetConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppVnetConnectionResult)(nil)).Elem()
}

func (o LookupWebAppVnetConnectionResultOutput) ToLookupWebAppVnetConnectionResultOutput() LookupWebAppVnetConnectionResultOutput {
	return o
}

func (o LookupWebAppVnetConnectionResultOutput) ToLookupWebAppVnetConnectionResultOutputWithContext(ctx context.Context) LookupWebAppVnetConnectionResultOutput {
	return o
}

func (o LookupWebAppVnetConnectionResultOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) *string { return v.CertBlob }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) CertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) string { return v.CertThumbprint }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) *string { return v.DnsServers }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) IsSwift() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) *bool { return v.IsSwift }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) ResyncRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) bool { return v.ResyncRequired }).(pulumi.BoolOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) []VnetRouteResponse { return v.Routes }).(VnetRouteResponseArrayOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppVnetConnectionResultOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppVnetConnectionResult) *string { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppVnetConnectionResultOutput{})
}
