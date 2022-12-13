


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAFDOrigin(ctx *pulumi.Context, args *LookupAFDOriginArgs, opts ...pulumi.InvokeOption) (*LookupAFDOriginResult, error) {
	var rv LookupAFDOriginResult
	err := ctx.Invoke("azure-native:cdn/v20221101preview:getAFDOrigin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAFDOriginArgs struct {
	OriginGroupName   string `pulumi:"originGroupName"`
	OriginName        string `pulumi:"originName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAFDOriginResult struct {
	AzureOrigin                 *ResourceReferenceResponse                   `pulumi:"azureOrigin"`
	DeploymentStatus            string                                       `pulumi:"deploymentStatus"`
	EnabledState                *string                                      `pulumi:"enabledState"`
	EnforceCertificateNameCheck *bool                                        `pulumi:"enforceCertificateNameCheck"`
	HostName                    string                                       `pulumi:"hostName"`
	HttpPort                    *int                                         `pulumi:"httpPort"`
	HttpsPort                   *int                                         `pulumi:"httpsPort"`
	Id                          string                                       `pulumi:"id"`
	Name                        string                                       `pulumi:"name"`
	OriginGroupName             string                                       `pulumi:"originGroupName"`
	OriginHostHeader            *string                                      `pulumi:"originHostHeader"`
	Priority                    *int                                         `pulumi:"priority"`
	ProvisioningState           string                                       `pulumi:"provisioningState"`
	SharedPrivateLinkResource   *SharedPrivateLinkResourcePropertiesResponse `pulumi:"sharedPrivateLinkResource"`
	SystemData                  SystemDataResponse                           `pulumi:"systemData"`
	Type                        string                                       `pulumi:"type"`
	Weight                      *int                                         `pulumi:"weight"`
}


func (val *LookupAFDOriginResult) Defaults() *LookupAFDOriginResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		enforceCertificateNameCheck_ := true
		tmp.EnforceCertificateNameCheck = &enforceCertificateNameCheck_
	}
	return &tmp
}

func LookupAFDOriginOutput(ctx *pulumi.Context, args LookupAFDOriginOutputArgs, opts ...pulumi.InvokeOption) LookupAFDOriginResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDOriginResult, error) {
			args := v.(LookupAFDOriginArgs)
			r, err := LookupAFDOrigin(ctx, &args, opts...)
			var s LookupAFDOriginResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDOriginResultOutput)
}

type LookupAFDOriginOutputArgs struct {
	OriginGroupName   pulumi.StringInput `pulumi:"originGroupName"`
	OriginName        pulumi.StringInput `pulumi:"originName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDOriginOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginArgs)(nil)).Elem()
}


type LookupAFDOriginResultOutput struct{ *pulumi.OutputState }

func (LookupAFDOriginResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginResult)(nil)).Elem()
}

func (o LookupAFDOriginResultOutput) ToLookupAFDOriginResultOutput() LookupAFDOriginResultOutput {
	return o
}

func (o LookupAFDOriginResultOutput) ToLookupAFDOriginResultOutputWithContext(ctx context.Context) LookupAFDOriginResultOutput {
	return o
}

func (o LookupAFDOriginResultOutput) AzureOrigin() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *ResourceReferenceResponse { return v.AzureOrigin }).(ResourceReferenceResponsePtrOutput)
}

func (o LookupAFDOriginResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o LookupAFDOriginResultOutput) EnforceCertificateNameCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *bool { return v.EnforceCertificateNameCheck }).(pulumi.BoolPtrOutput)
}

func (o LookupAFDOriginResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o LookupAFDOriginResultOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o LookupAFDOriginResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) OriginGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.OriginGroupName }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *string { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o LookupAFDOriginResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupAFDOriginResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) SharedPrivateLinkResource() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *SharedPrivateLinkResourcePropertiesResponse {
		return v.SharedPrivateLinkResource
	}).(SharedPrivateLinkResourcePropertiesResponsePtrOutput)
}

func (o LookupAFDOriginResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAFDOriginResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAFDOriginResultOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginResult) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDOriginResultOutput{})
}
