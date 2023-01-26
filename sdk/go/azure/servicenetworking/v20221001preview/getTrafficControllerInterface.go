


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrafficControllerInterface(ctx *pulumi.Context, args *LookupTrafficControllerInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupTrafficControllerInterfaceResult, error) {
	var rv LookupTrafficControllerInterfaceResult
	err := ctx.Invoke("azure-native:servicenetworking/v20221001preview:getTrafficControllerInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficControllerInterfaceArgs struct {
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TrafficControllerName string `pulumi:"trafficControllerName"`
}


type LookupTrafficControllerInterfaceResult struct {
	Associations           []ResourceIDResponse `pulumi:"associations"`
	ConfigurationEndpoints []string             `pulumi:"configurationEndpoints"`
	Frontends              []ResourceIDResponse `pulumi:"frontends"`
	Id                     string               `pulumi:"id"`
	Location               string               `pulumi:"location"`
	Name                   string               `pulumi:"name"`
	ProvisioningState      string               `pulumi:"provisioningState"`
	SystemData             SystemDataResponse   `pulumi:"systemData"`
	Tags                   map[string]string    `pulumi:"tags"`
	Type                   string               `pulumi:"type"`
}

func LookupTrafficControllerInterfaceOutput(ctx *pulumi.Context, args LookupTrafficControllerInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupTrafficControllerInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrafficControllerInterfaceResult, error) {
			args := v.(LookupTrafficControllerInterfaceArgs)
			r, err := LookupTrafficControllerInterface(ctx, &args, opts...)
			var s LookupTrafficControllerInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrafficControllerInterfaceResultOutput)
}

type LookupTrafficControllerInterfaceOutputArgs struct {
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TrafficControllerName pulumi.StringInput `pulumi:"trafficControllerName"`
}

func (LookupTrafficControllerInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficControllerInterfaceArgs)(nil)).Elem()
}


type LookupTrafficControllerInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupTrafficControllerInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficControllerInterfaceResult)(nil)).Elem()
}

func (o LookupTrafficControllerInterfaceResultOutput) ToLookupTrafficControllerInterfaceResultOutput() LookupTrafficControllerInterfaceResultOutput {
	return o
}

func (o LookupTrafficControllerInterfaceResultOutput) ToLookupTrafficControllerInterfaceResultOutputWithContext(ctx context.Context) LookupTrafficControllerInterfaceResultOutput {
	return o
}

func (o LookupTrafficControllerInterfaceResultOutput) Associations() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) []ResourceIDResponse { return v.Associations }).(ResourceIDResponseArrayOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) ConfigurationEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) []string { return v.ConfigurationEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) Frontends() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) []ResourceIDResponse { return v.Frontends }).(ResourceIDResponseArrayOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTrafficControllerInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrafficControllerInterfaceResultOutput{})
}
