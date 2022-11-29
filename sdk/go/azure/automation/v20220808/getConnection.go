


package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("azure-native:automation/v20220808:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ConnectionName        string `pulumi:"connectionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupConnectionResult struct {
	ConnectionType        *ConnectionTypeAssociationPropertyResponse `pulumi:"connectionType"`
	CreationTime          string                                     `pulumi:"creationTime"`
	Description           *string                                    `pulumi:"description"`
	FieldDefinitionValues map[string]string                          `pulumi:"fieldDefinitionValues"`
	Id                    string                                     `pulumi:"id"`
	LastModifiedTime      string                                     `pulumi:"lastModifiedTime"`
	Name                  string                                     `pulumi:"name"`
	Type                  string                                     `pulumi:"type"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionResult, error) {
			args := v.(LookupConnectionArgs)
			r, err := LookupConnection(ctx, &args, opts...)
			var s LookupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ConnectionName        pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}


type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ConnectionType() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *ConnectionTypeAssociationPropertyResponse { return v.ConnectionType }).(ConnectionTypeAssociationPropertyResponsePtrOutput)
}

func (o LookupConnectionResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionResultOutput) FieldDefinitionValues() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.FieldDefinitionValues }).(pulumi.StringMapOutput)
}

func (o LookupConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
