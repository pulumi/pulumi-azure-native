


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMigrateProjectsControllerMigrateProject(ctx *pulumi.Context, args *LookupMigrateProjectsControllerMigrateProjectArgs, opts ...pulumi.InvokeOption) (*LookupMigrateProjectsControllerMigrateProjectResult, error) {
	var rv LookupMigrateProjectsControllerMigrateProjectResult
	err := ctx.Invoke("azure-native:migrate/v20200501:getMigrateProjectsControllerMigrateProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateProjectsControllerMigrateProjectArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupMigrateProjectsControllerMigrateProjectResult struct {
	ETag       *string                          `pulumi:"eTag"`
	Id         string                           `pulumi:"id"`
	Location   *string                          `pulumi:"location"`
	Name       string                           `pulumi:"name"`
	Properties MigrateProjectPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse               `pulumi:"systemData"`
	Type       string                           `pulumi:"type"`
}

func LookupMigrateProjectsControllerMigrateProjectOutput(ctx *pulumi.Context, args LookupMigrateProjectsControllerMigrateProjectOutputArgs, opts ...pulumi.InvokeOption) LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrateProjectsControllerMigrateProjectResult, error) {
			args := v.(LookupMigrateProjectsControllerMigrateProjectArgs)
			r, err := LookupMigrateProjectsControllerMigrateProject(ctx, &args, opts...)
			var s LookupMigrateProjectsControllerMigrateProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrateProjectsControllerMigrateProjectResultOutput)
}

type LookupMigrateProjectsControllerMigrateProjectOutputArgs struct {
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMigrateProjectsControllerMigrateProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectsControllerMigrateProjectArgs)(nil)).Elem()
}


type LookupMigrateProjectsControllerMigrateProjectResultOutput struct{ *pulumi.OutputState }

func (LookupMigrateProjectsControllerMigrateProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectsControllerMigrateProjectResult)(nil)).Elem()
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ToLookupMigrateProjectsControllerMigrateProjectResultOutput() LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ToLookupMigrateProjectsControllerMigrateProjectResultOutputWithContext(ctx context.Context) LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) MigrateProjectPropertiesResponse {
		return v.Properties
	}).(MigrateProjectPropertiesResponseOutput)
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrateProjectsControllerMigrateProjectResultOutput{})
}
