


package v20161101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:importexport/v20161101:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobArgs struct {
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupJobResult struct {
	Id         string                   `pulumi:"id"`
	Identity   *IdentityDetailsResponse `pulumi:"identity"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties JobDetailsResponse       `pulumi:"properties"`
	SystemData SystemDataResponse       `pulumi:"systemData"`
	Tags       interface{}              `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}


func (val *LookupJobResult) Defaults() *LookupJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Identity = tmp.Identity.Defaults()

	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}


type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Identity() IdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *IdentityDetailsResponse { return v.Identity }).(IdentityDetailsResponsePtrOutput)
}

func (o LookupJobResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobResultOutput) Properties() JobDetailsResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobDetailsResponse { return v.Properties }).(JobDetailsResponseOutput)
}

func (o LookupJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupJobResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
