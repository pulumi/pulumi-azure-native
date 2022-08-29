


package v20190301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectorCollectionErrorInfoResponse struct {
	ErrorCode         string `pulumi:"errorCode"`
	ErrorInnerMessage string `pulumi:"errorInnerMessage"`
	ErrorMessage      string `pulumi:"errorMessage"`
	ErrorStartTime    string `pulumi:"errorStartTime"`
}

type ConnectorCollectionErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorCollectionErrorInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionErrorInfoResponseOutput) ToConnectorCollectionErrorInfoResponseOutput() ConnectorCollectionErrorInfoResponseOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponseOutput) ToConnectorCollectionErrorInfoResponseOutputWithContext(ctx context.Context) ConnectorCollectionErrorInfoResponseOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorInnerMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorInnerMessage }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorStartTime }).(pulumi.StringOutput)
}

type ConnectorCollectionErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorCollectionErrorInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ToConnectorCollectionErrorInfoResponsePtrOutput() ConnectorCollectionErrorInfoResponsePtrOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ToConnectorCollectionErrorInfoResponsePtrOutputWithContext(ctx context.Context) ConnectorCollectionErrorInfoResponsePtrOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) Elem() ConnectorCollectionErrorInfoResponseOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) ConnectorCollectionErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorCollectionErrorInfoResponse
		return ret
	}).(ConnectorCollectionErrorInfoResponseOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorCode
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorInnerMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorInnerMessage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorStartTime
	}).(pulumi.StringPtrOutput)
}

type ConnectorCollectionInfoResponse struct {
	Error             *ConnectorCollectionErrorInfoResponse `pulumi:"error"`
	LastChecked       string                                `pulumi:"lastChecked"`
	LastUpdated       string                                `pulumi:"lastUpdated"`
	SourceLastUpdated string                                `pulumi:"sourceLastUpdated"`
}

type ConnectorCollectionInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorCollectionInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionInfoResponseOutput) ToConnectorCollectionInfoResponseOutput() ConnectorCollectionInfoResponseOutput {
	return o
}

func (o ConnectorCollectionInfoResponseOutput) ToConnectorCollectionInfoResponseOutputWithContext(ctx context.Context) ConnectorCollectionInfoResponseOutput {
	return o
}

func (o ConnectorCollectionInfoResponseOutput) Error() ConnectorCollectionErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) *ConnectorCollectionErrorInfoResponse { return v.Error }).(ConnectorCollectionErrorInfoResponsePtrOutput)
}

func (o ConnectorCollectionInfoResponseOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.LastChecked }).(pulumi.StringOutput)
}

func (o ConnectorCollectionInfoResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o ConnectorCollectionInfoResponseOutput) SourceLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.SourceLastUpdated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorCollectionErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionInfoResponseOutput{})
}
