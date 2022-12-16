


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:migrate/v20171111preview:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupProjectResult struct {
	CreatedTimestamp    string      `pulumi:"createdTimestamp"`
	CustomerWorkspaceId *string     `pulumi:"customerWorkspaceId"`
	DiscoveryStatus     string      `pulumi:"discoveryStatus"`
	ETag                *string     `pulumi:"eTag"`
	Id                  string      `pulumi:"id"`
	Location            *string     `pulumi:"location"`
	Name                string      `pulumi:"name"`
	NumberOfAssessments int         `pulumi:"numberOfAssessments"`
	NumberOfGroups      int         `pulumi:"numberOfGroups"`
	NumberOfMachines    int         `pulumi:"numberOfMachines"`
	ProvisioningState   *string     `pulumi:"provisioningState"`
	Tags                interface{} `pulumi:"tags"`
	Type                string      `pulumi:"type"`
	UpdatedTimestamp    string      `pulumi:"updatedTimestamp"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}


type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) DiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DiscoveryStatus }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) NumberOfAssessments() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NumberOfAssessments }).(pulumi.IntOutput)
}

func (o LookupProjectResultOutput) NumberOfGroups() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NumberOfGroups }).(pulumi.IntOutput)
}

func (o LookupProjectResultOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NumberOfMachines }).(pulumi.IntOutput)
}

func (o LookupProjectResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProjectResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
