


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVmmServer(ctx *pulumi.Context, args *LookupVmmServerArgs, opts ...pulumi.InvokeOption) (*LookupVmmServerResult, error) {
	var rv LookupVmmServerResult
	err := ctx.Invoke("azure-native:scvmm/v20200605preview:getVmmServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVmmServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VmmServerName     string `pulumi:"vmmServerName"`
}


type LookupVmmServerResult struct {
	ConnectionStatus  string                                  `pulumi:"connectionStatus"`
	Credentials       *VMMServerPropertiesResponseCredentials `pulumi:"credentials"`
	ErrorMessage      string                                  `pulumi:"errorMessage"`
	ExtendedLocation  ExtendedLocationResponse                `pulumi:"extendedLocation"`
	Fqdn              string                                  `pulumi:"fqdn"`
	Id                string                                  `pulumi:"id"`
	Location          string                                  `pulumi:"location"`
	Name              string                                  `pulumi:"name"`
	Port              *int                                    `pulumi:"port"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                      `pulumi:"systemData"`
	Tags              map[string]string                       `pulumi:"tags"`
	Type              string                                  `pulumi:"type"`
	Uuid              string                                  `pulumi:"uuid"`
	Version           string                                  `pulumi:"version"`
}

func LookupVmmServerOutput(ctx *pulumi.Context, args LookupVmmServerOutputArgs, opts ...pulumi.InvokeOption) LookupVmmServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVmmServerResult, error) {
			args := v.(LookupVmmServerArgs)
			r, err := LookupVmmServer(ctx, &args, opts...)
			var s LookupVmmServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVmmServerResultOutput)
}

type LookupVmmServerOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VmmServerName     pulumi.StringInput `pulumi:"vmmServerName"`
}

func (LookupVmmServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmmServerArgs)(nil)).Elem()
}


type LookupVmmServerResultOutput struct{ *pulumi.OutputState }

func (LookupVmmServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmmServerResult)(nil)).Elem()
}

func (o LookupVmmServerResultOutput) ToLookupVmmServerResultOutput() LookupVmmServerResultOutput {
	return o
}

func (o LookupVmmServerResultOutput) ToLookupVmmServerResultOutputWithContext(ctx context.Context) LookupVmmServerResultOutput {
	return o
}

func (o LookupVmmServerResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Credentials() VMMServerPropertiesResponseCredentialsPtrOutput {
	return o.ApplyT(func(v LookupVmmServerResult) *VMMServerPropertiesResponseCredentials { return v.Credentials }).(VMMServerPropertiesResponseCredentialsPtrOutput)
}

func (o LookupVmmServerResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVmmServerResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o LookupVmmServerResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVmmServerResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupVmmServerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVmmServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVmmServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVmmServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVmmServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupVmmServerResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmmServerResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmmServerResultOutput{})
}
