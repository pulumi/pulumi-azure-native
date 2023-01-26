


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("azure-native:appplatform/v20220901preview:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppArgs struct {
	AppName           string  `pulumi:"appName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	SyncStatus        *string `pulumi:"syncStatus"`
}


type LookupAppResult struct {
	Id         string                             `pulumi:"id"`
	Identity   *ManagedIdentityPropertiesResponse `pulumi:"identity"`
	Location   *string                            `pulumi:"location"`
	Name       string                             `pulumi:"name"`
	Properties AppResourcePropertiesResponse      `pulumi:"properties"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	Type       string                             `pulumi:"type"`
}


func (val *LookupAppResult) Defaults() *LookupAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	AppName           pulumi.StringInput    `pulumi:"appName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput    `pulumi:"serviceName"`
	SyncStatus        pulumi.StringPtrInput `pulumi:"syncStatus"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}


type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) Identity() ManagedIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupAppResult) *ManagedIdentityPropertiesResponse { return v.Identity }).(ManagedIdentityPropertiesResponsePtrOutput)
}

func (o LookupAppResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppResultOutput) Properties() AppResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAppResult) AppResourcePropertiesResponse { return v.Properties }).(AppResourcePropertiesResponseOutput)
}

func (o LookupAppResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAppResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
