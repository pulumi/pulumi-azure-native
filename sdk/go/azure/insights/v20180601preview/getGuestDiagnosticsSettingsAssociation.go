


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context, args *LookupGuestDiagnosticsSettingsAssociationArgs, opts ...pulumi.InvokeOption) (*LookupGuestDiagnosticsSettingsAssociationResult, error) {
	var rv LookupGuestDiagnosticsSettingsAssociationResult
	err := ctx.Invoke("azure-native:insights/v20180601preview:getGuestDiagnosticsSettingsAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestDiagnosticsSettingsAssociationArgs struct {
	AssociationName string `pulumi:"associationName"`
	ResourceUri     string `pulumi:"resourceUri"`
}


type LookupGuestDiagnosticsSettingsAssociationResult struct {
	GuestDiagnosticSettingsName string            `pulumi:"guestDiagnosticSettingsName"`
	Id                          string            `pulumi:"id"`
	Location                    string            `pulumi:"location"`
	Name                        string            `pulumi:"name"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        string            `pulumi:"type"`
}

func LookupGuestDiagnosticsSettingsAssociationOutput(ctx *pulumi.Context, args LookupGuestDiagnosticsSettingsAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupGuestDiagnosticsSettingsAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestDiagnosticsSettingsAssociationResult, error) {
			args := v.(LookupGuestDiagnosticsSettingsAssociationArgs)
			r, err := LookupGuestDiagnosticsSettingsAssociation(ctx, &args, opts...)
			var s LookupGuestDiagnosticsSettingsAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestDiagnosticsSettingsAssociationResultOutput)
}

type LookupGuestDiagnosticsSettingsAssociationOutputArgs struct {
	AssociationName pulumi.StringInput `pulumi:"associationName"`
	ResourceUri     pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupGuestDiagnosticsSettingsAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestDiagnosticsSettingsAssociationArgs)(nil)).Elem()
}


type LookupGuestDiagnosticsSettingsAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupGuestDiagnosticsSettingsAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestDiagnosticsSettingsAssociationResult)(nil)).Elem()
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) ToLookupGuestDiagnosticsSettingsAssociationResultOutput() LookupGuestDiagnosticsSettingsAssociationResultOutput {
	return o
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) ToLookupGuestDiagnosticsSettingsAssociationResultOutputWithContext(ctx context.Context) LookupGuestDiagnosticsSettingsAssociationResultOutput {
	return o
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) GuestDiagnosticSettingsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestDiagnosticsSettingsAssociationResult) string { return v.GuestDiagnosticSettingsName }).(pulumi.StringOutput)
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestDiagnosticsSettingsAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestDiagnosticsSettingsAssociationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestDiagnosticsSettingsAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGuestDiagnosticsSettingsAssociationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGuestDiagnosticsSettingsAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestDiagnosticsSettingsAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestDiagnosticsSettingsAssociationResultOutput{})
}
