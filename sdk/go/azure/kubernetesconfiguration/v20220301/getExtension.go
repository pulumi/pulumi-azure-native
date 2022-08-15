


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20220301:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupExtensionArgs struct {
	ClusterName         string `pulumi:"clusterName"`
	ClusterResourceName string `pulumi:"clusterResourceName"`
	ClusterRp           string `pulumi:"clusterRp"`
	ExtensionName       string `pulumi:"extensionName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	AksAssignedIdentity            *ExtensionResponseAksAssignedIdentity `pulumi:"aksAssignedIdentity"`
	AutoUpgradeMinorVersion        *bool                                 `pulumi:"autoUpgradeMinorVersion"`
	ConfigurationProtectedSettings map[string]string                     `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          map[string]string                     `pulumi:"configurationSettings"`
	CustomLocationSettings         map[string]string                     `pulumi:"customLocationSettings"`
	ErrorInfo                      ErrorDetailResponse                   `pulumi:"errorInfo"`
	ExtensionType                  *string                               `pulumi:"extensionType"`
	Id                             string                                `pulumi:"id"`
	Identity                       *IdentityResponse                     `pulumi:"identity"`
	InstalledVersion               string                                `pulumi:"installedVersion"`
	Name                           string                                `pulumi:"name"`
	PackageUri                     string                                `pulumi:"packageUri"`
	ProvisioningState              string                                `pulumi:"provisioningState"`
	ReleaseTrain                   *string                               `pulumi:"releaseTrain"`
	Scope                          *ScopeResponse                        `pulumi:"scope"`
	Statuses                       []ExtensionStatusResponse             `pulumi:"statuses"`
	SystemData                     SystemDataResponse                    `pulumi:"systemData"`
	Type                           string                                `pulumi:"type"`
	Version                        *string                               `pulumi:"version"`
}


func (val *LookupExtensionResult) Defaults() *LookupExtensionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoUpgradeMinorVersion) {
		autoUpgradeMinorVersion_ := true
		tmp.AutoUpgradeMinorVersion = &autoUpgradeMinorVersion_
	}
	if isZero(tmp.ReleaseTrain) {
		releaseTrain_ := "Stable"
		tmp.ReleaseTrain = &releaseTrain_
	}
	return &tmp
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtensionResult, error) {
			args := v.(LookupExtensionArgs)
			r, err := LookupExtension(ctx, &args, opts...)
			var s LookupExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	ClusterName         pulumi.StringInput `pulumi:"clusterName"`
	ClusterResourceName pulumi.StringInput `pulumi:"clusterResourceName"`
	ClusterRp           pulumi.StringInput `pulumi:"clusterRp"`
	ExtensionName       pulumi.StringInput `pulumi:"extensionName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}


type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) AksAssignedIdentity() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *ExtensionResponseAksAssignedIdentity { return v.AksAssignedIdentity }).(ExtensionResponseAksAssignedIdentityPtrOutput)
}

func (o LookupExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupExtensionResultOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) ConfigurationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.ConfigurationSettings }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) CustomLocationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.CustomLocationSettings }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) ErrorInfo() ErrorDetailResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) ErrorDetailResponse { return v.ErrorInfo }).(ErrorDetailResponseOutput)
}

func (o LookupExtensionResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupExtensionResultOutput) InstalledVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.InstalledVersion }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) PackageUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.PackageUri }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.ReleaseTrain }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o LookupExtensionResultOutput) Statuses() ExtensionStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupExtensionResult) []ExtensionStatusResponse { return v.Statuses }).(ExtensionStatusResponseArrayOutput)
}

func (o LookupExtensionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
