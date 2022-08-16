


package orbital

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContactProfile(ctx *pulumi.Context, args *LookupContactProfileArgs, opts ...pulumi.InvokeOption) (*LookupContactProfileResult, error) {
	var rv LookupContactProfileResult
	err := ctx.Invoke("azure-native:orbital:getContactProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactProfileArgs struct {
	ContactProfileName string `pulumi:"contactProfileName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupContactProfileResult struct {
	AutoTrackingConfiguration    *string                      `pulumi:"autoTrackingConfiguration"`
	Etag                         string                       `pulumi:"etag"`
	EventHubUri                  *string                      `pulumi:"eventHubUri"`
	Id                           string                       `pulumi:"id"`
	Links                        []ContactProfileLinkResponse `pulumi:"links"`
	Location                     string                       `pulumi:"location"`
	MinimumElevationDegrees      *float64                     `pulumi:"minimumElevationDegrees"`
	MinimumViableContactDuration *string                      `pulumi:"minimumViableContactDuration"`
	Name                         string                       `pulumi:"name"`
	SystemData                   SystemDataResponse           `pulumi:"systemData"`
	Tags                         map[string]string            `pulumi:"tags"`
	Type                         string                       `pulumi:"type"`
}

func LookupContactProfileOutput(ctx *pulumi.Context, args LookupContactProfileOutputArgs, opts ...pulumi.InvokeOption) LookupContactProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactProfileResult, error) {
			args := v.(LookupContactProfileArgs)
			r, err := LookupContactProfile(ctx, &args, opts...)
			var s LookupContactProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactProfileResultOutput)
}

type LookupContactProfileOutputArgs struct {
	ContactProfileName pulumi.StringInput `pulumi:"contactProfileName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContactProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactProfileArgs)(nil)).Elem()
}


type LookupContactProfileResultOutput struct{ *pulumi.OutputState }

func (LookupContactProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactProfileResult)(nil)).Elem()
}

func (o LookupContactProfileResultOutput) ToLookupContactProfileResultOutput() LookupContactProfileResultOutput {
	return o
}

func (o LookupContactProfileResultOutput) ToLookupContactProfileResultOutputWithContext(ctx context.Context) LookupContactProfileResultOutput {
	return o
}

func (o LookupContactProfileResultOutput) AutoTrackingConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactProfileResult) *string { return v.AutoTrackingConfiguration }).(pulumi.StringPtrOutput)
}

func (o LookupContactProfileResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactProfileResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupContactProfileResultOutput) EventHubUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactProfileResult) *string { return v.EventHubUri }).(pulumi.StringPtrOutput)
}

func (o LookupContactProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContactProfileResultOutput) Links() ContactProfileLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupContactProfileResult) []ContactProfileLinkResponse { return v.Links }).(ContactProfileLinkResponseArrayOutput)
}

func (o LookupContactProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupContactProfileResultOutput) MinimumElevationDegrees() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupContactProfileResult) *float64 { return v.MinimumElevationDegrees }).(pulumi.Float64PtrOutput)
}

func (o LookupContactProfileResultOutput) MinimumViableContactDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactProfileResult) *string { return v.MinimumViableContactDuration }).(pulumi.StringPtrOutput)
}

func (o LookupContactProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContactProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContactProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContactProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContactProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupContactProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactProfileResultOutput{})
}
