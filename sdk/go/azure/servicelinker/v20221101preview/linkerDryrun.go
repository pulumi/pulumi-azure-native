


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkerDryrun struct {
	pulumi.CustomResourceState

	Name                pulumi.StringOutput                             `pulumi:"name"`
	OperationPreviews   DryrunOperationPreviewResponseArrayOutput       `pulumi:"operationPreviews"`
	Parameters          CreateOrUpdateDryrunParametersResponsePtrOutput `pulumi:"parameters"`
	PrerequisiteResults pulumi.ArrayOutput                              `pulumi:"prerequisiteResults"`
	ProvisioningState   pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                pulumi.StringOutput                             `pulumi:"type"`
}


func NewLinkerDryrun(ctx *pulumi.Context,
	name string, args *LinkerDryrunArgs, opts ...pulumi.ResourceOption) (*LinkerDryrun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	var resource LinkerDryrun
	err := ctx.RegisterResource("azure-native:servicelinker/v20221101preview:LinkerDryrun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkerDryrun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkerDryrunState, opts ...pulumi.ResourceOption) (*LinkerDryrun, error) {
	var resource LinkerDryrun
	err := ctx.ReadResource("azure-native:servicelinker/v20221101preview:LinkerDryrun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkerDryrunState struct {
}

type LinkerDryrunState struct {
}

func (LinkerDryrunState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkerDryrunState)(nil)).Elem()
}

type linkerDryrunArgs struct {
	DryrunName  *string                         `pulumi:"dryrunName"`
	Parameters  *CreateOrUpdateDryrunParameters `pulumi:"parameters"`
	ResourceUri string                          `pulumi:"resourceUri"`
}


type LinkerDryrunArgs struct {
	DryrunName  pulumi.StringPtrInput
	Parameters  CreateOrUpdateDryrunParametersPtrInput
	ResourceUri pulumi.StringInput
}

func (LinkerDryrunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkerDryrunArgs)(nil)).Elem()
}

type LinkerDryrunInput interface {
	pulumi.Input

	ToLinkerDryrunOutput() LinkerDryrunOutput
	ToLinkerDryrunOutputWithContext(ctx context.Context) LinkerDryrunOutput
}

func (*LinkerDryrun) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkerDryrun)(nil)).Elem()
}

func (i *LinkerDryrun) ToLinkerDryrunOutput() LinkerDryrunOutput {
	return i.ToLinkerDryrunOutputWithContext(context.Background())
}

func (i *LinkerDryrun) ToLinkerDryrunOutputWithContext(ctx context.Context) LinkerDryrunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkerDryrunOutput)
}

type LinkerDryrunOutput struct{ *pulumi.OutputState }

func (LinkerDryrunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkerDryrun)(nil)).Elem()
}

func (o LinkerDryrunOutput) ToLinkerDryrunOutput() LinkerDryrunOutput {
	return o
}

func (o LinkerDryrunOutput) ToLinkerDryrunOutputWithContext(ctx context.Context) LinkerDryrunOutput {
	return o
}

func (o LinkerDryrunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkerDryrunOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v *LinkerDryrun) DryrunOperationPreviewResponseArrayOutput { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

func (o LinkerDryrunOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v *LinkerDryrun) CreateOrUpdateDryrunParametersResponsePtrOutput { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

func (o LinkerDryrunOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.ArrayOutput { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

func (o LinkerDryrunOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LinkerDryrunOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LinkerDryrun) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LinkerDryrunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkerDryrun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkerDryrunOutput{})
}
