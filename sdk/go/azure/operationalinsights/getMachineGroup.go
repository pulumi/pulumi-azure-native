


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineGroup(ctx *pulumi.Context, args *LookupMachineGroupArgs, opts ...pulumi.InvokeOption) (*LookupMachineGroupResult, error) {
	var rv LookupMachineGroupResult
	err := ctx.Invoke("azure-native:operationalinsights:getMachineGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineGroupArgs struct {
	EndTime           *string `pulumi:"endTime"`
	MachineGroupName  string  `pulumi:"machineGroupName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	StartTime         *string `pulumi:"startTime"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type LookupMachineGroupResult struct {
	Count       *int                                `pulumi:"count"`
	DisplayName string                              `pulumi:"displayName"`
	Etag        *string                             `pulumi:"etag"`
	GroupType   *string                             `pulumi:"groupType"`
	Id          string                              `pulumi:"id"`
	Kind        string                              `pulumi:"kind"`
	Machines    []MachineReferenceWithHintsResponse `pulumi:"machines"`
	Name        string                              `pulumi:"name"`
	Type        string                              `pulumi:"type"`
}

func LookupMachineGroupOutput(ctx *pulumi.Context, args LookupMachineGroupOutputArgs, opts ...pulumi.InvokeOption) LookupMachineGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineGroupResult, error) {
			args := v.(LookupMachineGroupArgs)
			r, err := LookupMachineGroup(ctx, &args, opts...)
			var s LookupMachineGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineGroupResultOutput)
}

type LookupMachineGroupOutputArgs struct {
	EndTime           pulumi.StringPtrInput `pulumi:"endTime"`
	MachineGroupName  pulumi.StringInput    `pulumi:"machineGroupName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	StartTime         pulumi.StringPtrInput `pulumi:"startTime"`
	WorkspaceName     pulumi.StringInput    `pulumi:"workspaceName"`
}

func (LookupMachineGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineGroupArgs)(nil)).Elem()
}


type LookupMachineGroupResultOutput struct{ *pulumi.OutputState }

func (LookupMachineGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineGroupResult)(nil)).Elem()
}

func (o LookupMachineGroupResultOutput) ToLookupMachineGroupResultOutput() LookupMachineGroupResultOutput {
	return o
}

func (o LookupMachineGroupResultOutput) ToLookupMachineGroupResultOutputWithContext(ctx context.Context) LookupMachineGroupResultOutput {
	return o
}

func (o LookupMachineGroupResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LookupMachineGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupMachineGroupResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMachineGroupResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

func (o LookupMachineGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineGroupResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMachineGroupResultOutput) Machines() MachineReferenceWithHintsResponseArrayOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) []MachineReferenceWithHintsResponse { return v.Machines }).(MachineReferenceWithHintsResponseArrayOutput)
}

func (o LookupMachineGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineGroupResultOutput{})
}
