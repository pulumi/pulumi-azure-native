


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMigrateProject(ctx *pulumi.Context, args *LookupMigrateProjectArgs, opts ...pulumi.InvokeOption) (*LookupMigrateProjectResult, error) {
	var rv LookupMigrateProjectResult
	err := ctx.Invoke("azure-native:migrate/v20180901preview:getMigrateProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateProjectArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMigrateProjectResult struct {
	ETag       *string                          `pulumi:"eTag"`
	Id         string                           `pulumi:"id"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties MigrateProjectPropertiesResponse `pulumi:"properties"`
	Tags       *MigrateProjectResponseTags      `pulumi:"tags"`
	Type       string                           `pulumi:"type"`
}

func LookupMigrateProjectOutput(ctx *pulumi.Context, args LookupMigrateProjectOutputArgs, opts ...pulumi.InvokeOption) LookupMigrateProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrateProjectResult, error) {
			args := v.(LookupMigrateProjectArgs)
			r, err := LookupMigrateProject(ctx, &args, opts...)
			var s LookupMigrateProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrateProjectResultOutput)
}

type LookupMigrateProjectOutputArgs struct {
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMigrateProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectArgs)(nil)).Elem()
}


type LookupMigrateProjectResultOutput struct{ *pulumi.OutputState }

func (LookupMigrateProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectResult)(nil)).Elem()
}

func (o LookupMigrateProjectResultOutput) ToLookupMigrateProjectResultOutput() LookupMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectResultOutput) ToLookupMigrateProjectResultOutputWithContext(ctx context.Context) LookupMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupMigrateProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMigrateProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMigrateProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMigrateProjectResultOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) MigrateProjectPropertiesResponse { return v.Properties }).(MigrateProjectPropertiesResponseOutput)
}

func (o LookupMigrateProjectResultOutput) Tags() MigrateProjectResponseTagsPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) *MigrateProjectResponseTags { return v.Tags }).(MigrateProjectResponseTagsPtrOutput)
}

func (o LookupMigrateProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrateProjectResultOutput{})
}
