


package v20210517preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachine(ctx *pulumi.Context, args *LookupMachineArgs, opts ...pulumi.InvokeOption) (*LookupMachineResult, error) {
	var rv LookupMachineResult
	err := ctx.Invoke("azure-native:hybridcompute/v20210517preview:getMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	MachineName       string  `pulumi:"machineName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupMachineResult struct {
	Id         string                    `pulumi:"id"`
	Identity   *IdentityResponse         `pulumi:"identity"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties MachinePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}

func LookupMachineOutput(ctx *pulumi.Context, args LookupMachineOutputArgs, opts ...pulumi.InvokeOption) LookupMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineResult, error) {
			args := v.(LookupMachineArgs)
			r, err := LookupMachine(ctx, &args, opts...)
			var s LookupMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineResultOutput)
}

type LookupMachineOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	MachineName       pulumi.StringInput    `pulumi:"machineName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineArgs)(nil)).Elem()
}


type LookupMachineResultOutput struct{ *pulumi.OutputState }

func (LookupMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineResult)(nil)).Elem()
}

func (o LookupMachineResultOutput) ToLookupMachineResultOutput() LookupMachineResultOutput {
	return o
}

func (o LookupMachineResultOutput) ToLookupMachineResultOutputWithContext(ctx context.Context) LookupMachineResultOutput {
	return o
}

func (o LookupMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineResultOutput) Properties() MachinePropertiesResponseOutput {
	return o.ApplyT(func(v LookupMachineResult) MachinePropertiesResponse { return v.Properties }).(MachinePropertiesResponseOutput)
}

func (o LookupMachineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMachineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineResultOutput{})
}
