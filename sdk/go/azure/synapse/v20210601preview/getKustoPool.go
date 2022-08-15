


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoPool(ctx *pulumi.Context, args *LookupKustoPoolArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolResult, error) {
	var rv LookupKustoPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupKustoPoolArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupKustoPoolResult struct {
	DataIngestionUri      string                         `pulumi:"dataIngestionUri"`
	EnablePurge           *bool                          `pulumi:"enablePurge"`
	EnableStreamingIngest *bool                          `pulumi:"enableStreamingIngest"`
	Etag                  string                         `pulumi:"etag"`
	Id                    string                         `pulumi:"id"`
	LanguageExtensions    LanguageExtensionsListResponse `pulumi:"languageExtensions"`
	Location              string                         `pulumi:"location"`
	Name                  string                         `pulumi:"name"`
	OptimizedAutoscale    *OptimizedAutoscaleResponse    `pulumi:"optimizedAutoscale"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	Sku                   AzureSkuResponse               `pulumi:"sku"`
	State                 string                         `pulumi:"state"`
	StateReason           string                         `pulumi:"stateReason"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tags                  map[string]string              `pulumi:"tags"`
	Type                  string                         `pulumi:"type"`
	Uri                   string                         `pulumi:"uri"`
	WorkspaceUID          *string                        `pulumi:"workspaceUID"`
}


func (val *LookupKustoPoolResult) Defaults() *LookupKustoPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnablePurge) {
		enablePurge_ := false
		tmp.EnablePurge = &enablePurge_
	}
	if isZero(tmp.EnableStreamingIngest) {
		enableStreamingIngest_ := false
		tmp.EnableStreamingIngest = &enableStreamingIngest_
	}
	return &tmp
}

func LookupKustoPoolOutput(ctx *pulumi.Context, args LookupKustoPoolOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolResult, error) {
			args := v.(LookupKustoPoolArgs)
			r, err := LookupKustoPool(ctx, &args, opts...)
			var s LookupKustoPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoPoolResultOutput)
}

type LookupKustoPoolOutputArgs struct {
	KustoPoolName     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolArgs)(nil)).Elem()
}


type LookupKustoPoolResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolResult)(nil)).Elem()
}

func (o LookupKustoPoolResultOutput) ToLookupKustoPoolResultOutput() LookupKustoPoolResultOutput {
	return o
}

func (o LookupKustoPoolResultOutput) ToLookupKustoPoolResultOutputWithContext(ctx context.Context) LookupKustoPoolResultOutput {
	return o
}

func (o LookupKustoPoolResultOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.DataIngestionUri }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) EnablePurge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) *bool { return v.EnablePurge }).(pulumi.BoolPtrOutput)
}

func (o LookupKustoPoolResultOutput) EnableStreamingIngest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) *bool { return v.EnableStreamingIngest }).(pulumi.BoolPtrOutput)
}

func (o LookupKustoPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) LanguageExtensions() LanguageExtensionsListResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) LanguageExtensionsListResponse { return v.LanguageExtensions }).(LanguageExtensionsListResponseOutput)
}

func (o LookupKustoPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) OptimizedAutoscale() OptimizedAutoscaleResponsePtrOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) *OptimizedAutoscaleResponse { return v.OptimizedAutoscale }).(OptimizedAutoscaleResponsePtrOutput)
}

func (o LookupKustoPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) AzureSkuResponse { return v.Sku }).(AzureSkuResponseOutput)
}

func (o LookupKustoPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.StateReason }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKustoPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupKustoPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) string { return v.Uri }).(pulumi.StringOutput)
}

func (o LookupKustoPoolResultOutput) WorkspaceUID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolResult) *string { return v.WorkspaceUID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolResultOutput{})
}
