


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubRouteTableV2(ctx *pulumi.Context, args *LookupVirtualHubRouteTableV2Args, opts ...pulumi.InvokeOption) (*LookupVirtualHubRouteTableV2Result, error) {
	var rv LookupVirtualHubRouteTableV2Result
	err := ctx.Invoke("azure-native:network/v20220701:getVirtualHubRouteTableV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubRouteTableV2Args struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteTableName    string `pulumi:"routeTableName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubRouteTableV2Result struct {
	AttachedConnections []string                    `pulumi:"attachedConnections"`
	Etag                string                      `pulumi:"etag"`
	Id                  *string                     `pulumi:"id"`
	Name                *string                     `pulumi:"name"`
	ProvisioningState   string                      `pulumi:"provisioningState"`
	Routes              []VirtualHubRouteV2Response `pulumi:"routes"`
}

func LookupVirtualHubRouteTableV2Output(ctx *pulumi.Context, args LookupVirtualHubRouteTableV2OutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubRouteTableV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubRouteTableV2Result, error) {
			args := v.(LookupVirtualHubRouteTableV2Args)
			r, err := LookupVirtualHubRouteTableV2(ctx, &args, opts...)
			var s LookupVirtualHubRouteTableV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubRouteTableV2ResultOutput)
}

type LookupVirtualHubRouteTableV2OutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteTableName    pulumi.StringInput `pulumi:"routeTableName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubRouteTableV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubRouteTableV2Args)(nil)).Elem()
}


type LookupVirtualHubRouteTableV2ResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubRouteTableV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubRouteTableV2Result)(nil)).Elem()
}

func (o LookupVirtualHubRouteTableV2ResultOutput) ToLookupVirtualHubRouteTableV2ResultOutput() LookupVirtualHubRouteTableV2ResultOutput {
	return o
}

func (o LookupVirtualHubRouteTableV2ResultOutput) ToLookupVirtualHubRouteTableV2ResultOutputWithContext(ctx context.Context) LookupVirtualHubRouteTableV2ResultOutput {
	return o
}

func (o LookupVirtualHubRouteTableV2ResultOutput) AttachedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) []string { return v.AttachedConnections }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualHubRouteTableV2ResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualHubRouteTableV2ResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubRouteTableV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubRouteTableV2ResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubRouteTableV2ResultOutput) Routes() VirtualHubRouteV2ResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubRouteTableV2Result) []VirtualHubRouteV2Response { return v.Routes }).(VirtualHubRouteV2ResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubRouteTableV2ResultOutput{})
}
