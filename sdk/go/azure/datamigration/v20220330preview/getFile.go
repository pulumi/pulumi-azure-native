


package v20220330preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFile(ctx *pulumi.Context, args *LookupFileArgs, opts ...pulumi.InvokeOption) (*LookupFileResult, error) {
	var rv LookupFileResult
	err := ctx.Invoke("azure-native:datamigration/v20220330preview:getFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileArgs struct {
	FileName    string `pulumi:"fileName"`
	GroupName   string `pulumi:"groupName"`
	ProjectName string `pulumi:"projectName"`
	ServiceName string `pulumi:"serviceName"`
}


type LookupFileResult struct {
	Etag       *string                       `pulumi:"etag"`
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties ProjectFilePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupFileOutput(ctx *pulumi.Context, args LookupFileOutputArgs, opts ...pulumi.InvokeOption) LookupFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileResult, error) {
			args := v.(LookupFileArgs)
			r, err := LookupFile(ctx, &args, opts...)
			var s LookupFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileResultOutput)
}

type LookupFileOutputArgs struct {
	FileName    pulumi.StringInput `pulumi:"fileName"`
	GroupName   pulumi.StringInput `pulumi:"groupName"`
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileArgs)(nil)).Elem()
}


type LookupFileResultOutput struct{ *pulumi.OutputState }

func (LookupFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileResult)(nil)).Elem()
}

func (o LookupFileResultOutput) ToLookupFileResultOutput() LookupFileResultOutput {
	return o
}

func (o LookupFileResultOutput) ToLookupFileResultOutputWithContext(ctx context.Context) LookupFileResultOutput {
	return o
}

func (o LookupFileResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileResultOutput) Properties() ProjectFilePropertiesResponseOutput {
	return o.ApplyT(func(v LookupFileResult) ProjectFilePropertiesResponse { return v.Properties }).(ProjectFilePropertiesResponseOutput)
}

func (o LookupFileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileResultOutput{})
}
