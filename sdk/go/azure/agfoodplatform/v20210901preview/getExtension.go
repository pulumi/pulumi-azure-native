


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:agfoodplatform/v20210901preview:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	ExtensionId           string `pulumi:"extensionId"`
	FarmBeatsResourceName string `pulumi:"farmBeatsResourceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	ETag                      string             `pulumi:"eTag"`
	ExtensionApiDocsLink      string             `pulumi:"extensionApiDocsLink"`
	ExtensionAuthLink         string             `pulumi:"extensionAuthLink"`
	ExtensionCategory         string             `pulumi:"extensionCategory"`
	ExtensionId               string             `pulumi:"extensionId"`
	Id                        string             `pulumi:"id"`
	InstalledExtensionVersion string             `pulumi:"installedExtensionVersion"`
	Name                      string             `pulumi:"name"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
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
	ExtensionId           pulumi.StringInput `pulumi:"extensionId"`
	FarmBeatsResourceName pulumi.StringInput `pulumi:"farmBeatsResourceName"`
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

func (o LookupExtensionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ExtensionApiDocsLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionApiDocsLink }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ExtensionAuthLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionAuthLink }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ExtensionCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionCategory }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) ExtensionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ExtensionId }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) InstalledExtensionVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.InstalledExtensionVersion }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExtensionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
