


package kubernetesconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	ClusterName           string `pulumi:"clusterName"`
	ClusterResourceName   string `pulumi:"clusterResourceName"`
	ClusterRp             string `pulumi:"clusterRp"`
	ExtensionInstanceName string `pulumi:"extensionInstanceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	AutoUpgradeMinorVersion        *bool                          `pulumi:"autoUpgradeMinorVersion"`
	ConfigurationProtectedSettings map[string]string              `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          map[string]string              `pulumi:"configurationSettings"`
	CreationTime                   string                         `pulumi:"creationTime"`
	ErrorInfo                      ErrorDefinitionResponse        `pulumi:"errorInfo"`
	ExtensionType                  *string                        `pulumi:"extensionType"`
	Id                             string                         `pulumi:"id"`
	Identity                       *ConfigurationIdentityResponse `pulumi:"identity"`
	InstallState                   string                         `pulumi:"installState"`
	LastModifiedTime               string                         `pulumi:"lastModifiedTime"`
	LastStatusTime                 string                         `pulumi:"lastStatusTime"`
	Name                           string                         `pulumi:"name"`
	ReleaseTrain                   *string                        `pulumi:"releaseTrain"`
	Scope                          *ScopeResponse                 `pulumi:"scope"`
	Statuses                       []ExtensionStatusResponse      `pulumi:"statuses"`
	SystemData                     *SystemDataResponse            `pulumi:"systemData"`
	Type                           string                         `pulumi:"type"`
	Version                        *string                        `pulumi:"version"`
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
	ClusterName           pulumi.StringInput `pulumi:"clusterName"`
	ClusterResourceName   pulumi.StringInput `pulumi:"clusterResourceName"`
	ClusterRp             pulumi.StringInput `pulumi:"clusterRp"`
	ExtensionInstanceName pulumi.StringInput `pulumi:"extensionInstanceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupExtensionResultOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) ConfigurationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.ConfigurationSettings }).(pulumi.StringMapOutput)
}

func (o LookupExtensionResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ErrorInfo() ErrorDefinitionResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) ErrorDefinitionResponse { return v.ErrorInfo }).(ErrorDefinitionResponseOutput)
}

func (o LookupExtensionResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Identity() ConfigurationIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *ConfigurationIdentityResponse { return v.Identity }).(ConfigurationIdentityResponsePtrOutput)
}

func (o LookupExtensionResultOutput) InstallState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.InstallState }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) LastStatusTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.LastStatusTime }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
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

func (o LookupExtensionResultOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *SystemDataResponse { return v.SystemData }).(SystemDataResponsePtrOutput)
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
