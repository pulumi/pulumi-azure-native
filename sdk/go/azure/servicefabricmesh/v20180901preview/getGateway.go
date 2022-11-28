


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180901preview:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayArgs struct {
	GatewayResourceName string `pulumi:"gatewayResourceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupGatewayResult struct {
	Description        *string              `pulumi:"description"`
	DestinationNetwork NetworkRefResponse   `pulumi:"destinationNetwork"`
	Http               []HttpConfigResponse `pulumi:"http"`
	Id                 string               `pulumi:"id"`
	IpAddress          string               `pulumi:"ipAddress"`
	Location           string               `pulumi:"location"`
	Name               string               `pulumi:"name"`
	ProvisioningState  string               `pulumi:"provisioningState"`
	SourceNetwork      NetworkRefResponse   `pulumi:"sourceNetwork"`
	Status             string               `pulumi:"status"`
	StatusDetails      string               `pulumi:"statusDetails"`
	Tags               map[string]string    `pulumi:"tags"`
	Tcp                []TcpConfigResponse  `pulumi:"tcp"`
	Type               string               `pulumi:"type"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

type LookupGatewayOutputArgs struct {
	GatewayResourceName pulumi.StringInput `pulumi:"gatewayResourceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}


type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayResultOutput) DestinationNetwork() NetworkRefResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) NetworkRefResponse { return v.DestinationNetwork }).(NetworkRefResponseOutput)
}

func (o LookupGatewayResultOutput) Http() HttpConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupGatewayResult) []HttpConfigResponse { return v.Http }).(HttpConfigResponseArrayOutput)
}

func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) SourceNetwork() NetworkRefResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) NetworkRefResponse { return v.SourceNetwork }).(NetworkRefResponseOutput)
}

func (o LookupGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.StatusDetails }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGatewayResultOutput) Tcp() TcpConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupGatewayResult) []TcpConfigResponse { return v.Tcp }).(TcpConfigResponseArrayOutput)
}

func (o LookupGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
