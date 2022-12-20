


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVCenter(ctx *pulumi.Context, args *LookupVCenterArgs, opts ...pulumi.InvokeOption) (*LookupVCenterResult, error) {
	var rv LookupVCenterResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20201001preview:getVCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVCenterArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VcenterName       string `pulumi:"vcenterName"`
}


type LookupVCenterResult struct {
	ConnectionStatus   string                    `pulumi:"connectionStatus"`
	Credentials        *VICredentialResponse     `pulumi:"credentials"`
	CustomResourceName string                    `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Fqdn               string                    `pulumi:"fqdn"`
	Id                 string                    `pulumi:"id"`
	InstanceUuid       string                    `pulumi:"instanceUuid"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	Name               string                    `pulumi:"name"`
	Port               *int                      `pulumi:"port"`
	ProvisioningState  string                    `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse  `pulumi:"statuses"`
	SystemData         SystemDataResponse        `pulumi:"systemData"`
	Tags               map[string]string         `pulumi:"tags"`
	Type               string                    `pulumi:"type"`
	Uuid               string                    `pulumi:"uuid"`
	Version            string                    `pulumi:"version"`
}

func LookupVCenterOutput(ctx *pulumi.Context, args LookupVCenterOutputArgs, opts ...pulumi.InvokeOption) LookupVCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVCenterResult, error) {
			args := v.(LookupVCenterArgs)
			r, err := LookupVCenter(ctx, &args, opts...)
			var s LookupVCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVCenterResultOutput)
}

type LookupVCenterOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VcenterName       pulumi.StringInput `pulumi:"vcenterName"`
}

func (LookupVCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVCenterArgs)(nil)).Elem()
}


type LookupVCenterResultOutput struct{ *pulumi.OutputState }

func (LookupVCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVCenterResult)(nil)).Elem()
}

func (o LookupVCenterResultOutput) ToLookupVCenterResultOutput() LookupVCenterResultOutput {
	return o
}

func (o LookupVCenterResultOutput) ToLookupVCenterResultOutputWithContext(ctx context.Context) LookupVCenterResultOutput {
	return o
}

func (o LookupVCenterResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Credentials() VICredentialResponsePtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *VICredentialResponse { return v.Credentials }).(VICredentialResponsePtrOutput)
}

func (o LookupVCenterResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVCenterResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) InstanceUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.InstanceUuid }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupVCenterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVCenterResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupVCenterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVCenterResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupVCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVCenterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVCenterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupVCenterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVCenterResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVCenterResultOutput{})
}
