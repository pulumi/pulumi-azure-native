


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:iotsecurity/v20210201preview:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	Scope string `pulumi:"scope"`
}


type LookupSiteResult struct {
	DisplayName string             `pulumi:"displayName"`
	Id          string             `pulumi:"id"`
	Name        string             `pulumi:"name"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Tags        map[string]string  `pulumi:"tags"`
	Type        string             `pulumi:"type"`
}

func LookupSiteOutput(ctx *pulumi.Context, args LookupSiteOutputArgs, opts ...pulumi.InvokeOption) LookupSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteResult, error) {
			args := v.(LookupSiteArgs)
			r, err := LookupSite(ctx, &args, opts...)
			var s LookupSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteResultOutput)
}

type LookupSiteOutputArgs struct {
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteArgs)(nil)).Elem()
}


type LookupSiteResultOutput struct{ *pulumi.OutputState }

func (LookupSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteResult)(nil)).Elem()
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutput() LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutputWithContext(ctx context.Context) LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSiteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteResultOutput{})
}
