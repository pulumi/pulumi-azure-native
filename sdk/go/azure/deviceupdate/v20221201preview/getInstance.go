


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("azure-native:deviceupdate/v20221201preview:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	AccountName       string `pulumi:"accountName"`
	InstanceName      string `pulumi:"instanceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstanceResult struct {
	AccountName                 string                               `pulumi:"accountName"`
	DiagnosticStorageProperties *DiagnosticStoragePropertiesResponse `pulumi:"diagnosticStorageProperties"`
	EnableDiagnostics           *bool                                `pulumi:"enableDiagnostics"`
	Id                          string                               `pulumi:"id"`
	IotHubs                     []IotHubSettingsResponse             `pulumi:"iotHubs"`
	Location                    string                               `pulumi:"location"`
	Name                        string                               `pulumi:"name"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse                   `pulumi:"systemData"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	InstanceName      pulumi.StringInput `pulumi:"instanceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}


type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) DiagnosticStorageProperties() DiagnosticStoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *DiagnosticStoragePropertiesResponse {
		return v.DiagnosticStorageProperties
	}).(DiagnosticStoragePropertiesResponsePtrOutput)
}

func (o LookupInstanceResultOutput) EnableDiagnostics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.EnableDiagnostics }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) IotHubs() IotHubSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []IotHubSettingsResponse { return v.IotHubs }).(IotHubSettingsResponseArrayOutput)
}

func (o LookupInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
