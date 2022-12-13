


package v20150504preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("azure-native:network/v20150504preview:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupZoneArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ZoneName          string `pulumi:"zoneName"`
}


type LookupZoneResult struct {
	Etag       *string                `pulumi:"etag"`
	Id         string                 `pulumi:"id"`
	Location   string                 `pulumi:"location"`
	Name       string                 `pulumi:"name"`
	Properties ZonePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string      `pulumi:"tags"`
	Type       string                 `pulumi:"type"`
}

func LookupZoneOutput(ctx *pulumi.Context, args LookupZoneOutputArgs, opts ...pulumi.InvokeOption) LookupZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZoneResult, error) {
			args := v.(LookupZoneArgs)
			r, err := LookupZone(ctx, &args, opts...)
			var s LookupZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZoneResultOutput)
}

type LookupZoneOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ZoneName          pulumi.StringInput `pulumi:"zoneName"`
}

func (LookupZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneArgs)(nil)).Elem()
}


type LookupZoneResultOutput struct{ *pulumi.OutputState }

func (LookupZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZoneResult)(nil)).Elem()
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutput() LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) ToLookupZoneResultOutputWithContext(ctx context.Context) LookupZoneResultOutput {
	return o
}

func (o LookupZoneResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupZoneResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupZoneResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupZoneResultOutput) Properties() ZonePropertiesResponseOutput {
	return o.ApplyT(func(v LookupZoneResult) ZonePropertiesResponse { return v.Properties }).(ZonePropertiesResponseOutput)
}

func (o LookupZoneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupZoneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupZoneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupZoneResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZoneResultOutput{})
}
