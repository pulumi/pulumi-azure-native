


package v20201101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHubRouteTable(ctx *pulumi.Context, args *LookupHubRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupHubRouteTableResult, error) {
	var rv LookupHubRouteTableResult
	err := ctx.Invoke("azure-native:network/v20201101:getHubRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubRouteTableArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteTableName    string `pulumi:"routeTableName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupHubRouteTableResult struct {
	AssociatedConnections  []string           `pulumi:"associatedConnections"`
	Etag                   string             `pulumi:"etag"`
	Id                     *string            `pulumi:"id"`
	Labels                 []string           `pulumi:"labels"`
	Name                   *string            `pulumi:"name"`
	PropagatingConnections []string           `pulumi:"propagatingConnections"`
	ProvisioningState      string             `pulumi:"provisioningState"`
	Routes                 []HubRouteResponse `pulumi:"routes"`
	Type                   string             `pulumi:"type"`
}

func LookupHubRouteTableOutput(ctx *pulumi.Context, args LookupHubRouteTableOutputArgs, opts ...pulumi.InvokeOption) LookupHubRouteTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHubRouteTableResult, error) {
			args := v.(LookupHubRouteTableArgs)
			r, err := LookupHubRouteTable(ctx, &args, opts...)
			var s LookupHubRouteTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHubRouteTableResultOutput)
}

type LookupHubRouteTableOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteTableName    pulumi.StringInput `pulumi:"routeTableName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupHubRouteTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubRouteTableArgs)(nil)).Elem()
}


type LookupHubRouteTableResultOutput struct{ *pulumi.OutputState }

func (LookupHubRouteTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubRouteTableResult)(nil)).Elem()
}

func (o LookupHubRouteTableResultOutput) ToLookupHubRouteTableResultOutput() LookupHubRouteTableResultOutput {
	return o
}

func (o LookupHubRouteTableResultOutput) ToLookupHubRouteTableResultOutputWithContext(ctx context.Context) LookupHubRouteTableResultOutput {
	return o
}

func (o LookupHubRouteTableResultOutput) AssociatedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) []string { return v.AssociatedConnections }).(pulumi.StringArrayOutput)
}

func (o LookupHubRouteTableResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupHubRouteTableResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupHubRouteTableResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupHubRouteTableResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupHubRouteTableResultOutput) PropagatingConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) []string { return v.PropagatingConnections }).(pulumi.StringArrayOutput)
}

func (o LookupHubRouteTableResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHubRouteTableResultOutput) Routes() HubRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) []HubRouteResponse { return v.Routes }).(HubRouteResponseArrayOutput)
}

func (o LookupHubRouteTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubRouteTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHubRouteTableResultOutput{})
}
