


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomDomainResponse struct {
	Name             *string `pulumi:"name"`
	UseSubDomainName *bool   `pulumi:"useSubDomainName"`
}

type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDomainResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) Elem() CustomDomainResponseOutput {
	return o.ApplyT(func(v *CustomDomainResponse) CustomDomainResponse {
		if v != nil {
			return *v
		}
		var ret CustomDomainResponse
		return ret
	}).(CustomDomainResponseOutput)
}

func (o CustomDomainResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainResponsePtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type EndpointsResponse struct {
	Blob  *string `pulumi:"blob"`
	Queue *string `pulumi:"queue"`
	Table *string `pulumi:"table"`
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *string { return v.Blob }).(pulumi.StringPtrOutput)
}

func (o EndpointsResponseOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *string { return v.Queue }).(pulumi.StringPtrOutput)
}

func (o EndpointsResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

type EndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) Elem() EndpointsResponseOutput {
	return o.ApplyT(func(v *EndpointsResponse) EndpointsResponse {
		if v != nil {
			return *v
		}
		var ret EndpointsResponse
		return ret
	}).(EndpointsResponseOutput)
}

func (o EndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Queue
	}).(pulumi.StringPtrOutput)
}

func (o EndpointsResponsePtrOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Table
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(EndpointsResponsePtrOutput{})
}
