


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLogProfile(ctx *pulumi.Context, args *LookupLogProfileArgs, opts ...pulumi.InvokeOption) (*LookupLogProfileResult, error) {
	var rv LookupLogProfileResult
	err := ctx.Invoke("azure-native:insights/v20160301:getLogProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLogProfileArgs struct {
	LogProfileName string `pulumi:"logProfileName"`
}


type LookupLogProfileResult struct {
	Categories       []string                `pulumi:"categories"`
	Id               string                  `pulumi:"id"`
	Location         string                  `pulumi:"location"`
	Locations        []string                `pulumi:"locations"`
	Name             string                  `pulumi:"name"`
	RetentionPolicy  RetentionPolicyResponse `pulumi:"retentionPolicy"`
	ServiceBusRuleId *string                 `pulumi:"serviceBusRuleId"`
	StorageAccountId *string                 `pulumi:"storageAccountId"`
	Tags             map[string]string       `pulumi:"tags"`
	Type             string                  `pulumi:"type"`
}

func LookupLogProfileOutput(ctx *pulumi.Context, args LookupLogProfileOutputArgs, opts ...pulumi.InvokeOption) LookupLogProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogProfileResult, error) {
			args := v.(LookupLogProfileArgs)
			r, err := LookupLogProfile(ctx, &args, opts...)
			var s LookupLogProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogProfileResultOutput)
}

type LookupLogProfileOutputArgs struct {
	LogProfileName pulumi.StringInput `pulumi:"logProfileName"`
}

func (LookupLogProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogProfileArgs)(nil)).Elem()
}


type LookupLogProfileResultOutput struct{ *pulumi.OutputState }

func (LookupLogProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogProfileResult)(nil)).Elem()
}

func (o LookupLogProfileResultOutput) ToLookupLogProfileResultOutput() LookupLogProfileResultOutput {
	return o
}

func (o LookupLogProfileResultOutput) ToLookupLogProfileResultOutputWithContext(ctx context.Context) LookupLogProfileResultOutput {
	return o
}

func (o LookupLogProfileResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogProfileResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LookupLogProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLogProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLogProfileResultOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogProfileResult) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o LookupLogProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLogProfileResultOutput) RetentionPolicy() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v LookupLogProfileResult) RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponseOutput)
}

func (o LookupLogProfileResultOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogProfileResult) *string { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupLogProfileResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogProfileResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupLogProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLogProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLogProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogProfileResultOutput{})
}
