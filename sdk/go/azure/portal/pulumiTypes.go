


package portal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConsoleCreateProperties struct {
	OsType            string  `pulumi:"osType"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Uri               *string `pulumi:"uri"`
}





type ConsoleCreatePropertiesInput interface {
	pulumi.Input

	ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput
	ToConsoleCreatePropertiesOutputWithContext(context.Context) ConsoleCreatePropertiesOutput
}

type ConsoleCreatePropertiesArgs struct {
	OsType            pulumi.StringInput    `pulumi:"osType"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Uri               pulumi.StringPtrInput `pulumi:"uri"`
}

func (ConsoleCreatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleCreateProperties)(nil)).Elem()
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput {
	return i.ToConsoleCreatePropertiesOutputWithContext(context.Background())
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesOutputWithContext(ctx context.Context) ConsoleCreatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesOutput)
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return i.ToConsoleCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesOutput).ToConsoleCreatePropertiesPtrOutputWithContext(ctx)
}









type ConsoleCreatePropertiesPtrInput interface {
	pulumi.Input

	ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput
	ToConsoleCreatePropertiesPtrOutputWithContext(context.Context) ConsoleCreatePropertiesPtrOutput
}

type consoleCreatePropertiesPtrType ConsoleCreatePropertiesArgs

func ConsoleCreatePropertiesPtr(v *ConsoleCreatePropertiesArgs) ConsoleCreatePropertiesPtrInput {
	return (*consoleCreatePropertiesPtrType)(v)
}

func (*consoleCreatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleCreateProperties)(nil)).Elem()
}

func (i *consoleCreatePropertiesPtrType) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return i.ToConsoleCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i *consoleCreatePropertiesPtrType) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesPtrOutput)
}

type ConsoleCreatePropertiesOutput struct{ *pulumi.OutputState }

func (ConsoleCreatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleCreateProperties)(nil)).Elem()
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput {
	return o
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesOutputWithContext(ctx context.Context) ConsoleCreatePropertiesOutput {
	return o
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return o.ToConsoleCreatePropertiesPtrOutputWithContext(context.Background())
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsoleCreateProperties) *ConsoleCreateProperties {
		return &v
	}).(ConsoleCreatePropertiesPtrOutput)
}

func (o ConsoleCreatePropertiesOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ConsoleCreatePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ConsoleCreatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConsoleCreatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleCreateProperties)(nil)).Elem()
}

func (o ConsoleCreatePropertiesPtrOutput) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return o
}

func (o ConsoleCreatePropertiesPtrOutput) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return o
}

func (o ConsoleCreatePropertiesPtrOutput) Elem() ConsoleCreatePropertiesOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) ConsoleCreateProperties {
		if v != nil {
			return *v
		}
		var ret ConsoleCreateProperties
		return ret
	}).(ConsoleCreatePropertiesOutput)
}

func (o ConsoleCreatePropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ConsolePropertiesResponse struct {
	OsType            string `pulumi:"osType"`
	ProvisioningState string `pulumi:"provisioningState"`
	Uri               string `pulumi:"uri"`
}





type ConsolePropertiesResponseInput interface {
	pulumi.Input

	ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput
	ToConsolePropertiesResponseOutputWithContext(context.Context) ConsolePropertiesResponseOutput
}

type ConsolePropertiesResponseArgs struct {
	OsType            pulumi.StringInput `pulumi:"osType"`
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
	Uri               pulumi.StringInput `pulumi:"uri"`
}

func (ConsolePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsolePropertiesResponse)(nil)).Elem()
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput {
	return i.ToConsolePropertiesResponseOutputWithContext(context.Background())
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponseOutputWithContext(ctx context.Context) ConsolePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsolePropertiesResponseOutput)
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return i.ToConsolePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsolePropertiesResponseOutput).ToConsolePropertiesResponsePtrOutputWithContext(ctx)
}









type ConsolePropertiesResponsePtrInput interface {
	pulumi.Input

	ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput
	ToConsolePropertiesResponsePtrOutputWithContext(context.Context) ConsolePropertiesResponsePtrOutput
}

type consolePropertiesResponsePtrType ConsolePropertiesResponseArgs

func ConsolePropertiesResponsePtr(v *ConsolePropertiesResponseArgs) ConsolePropertiesResponsePtrInput {
	return (*consolePropertiesResponsePtrType)(v)
}

func (*consolePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsolePropertiesResponse)(nil)).Elem()
}

func (i *consolePropertiesResponsePtrType) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return i.ToConsolePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *consolePropertiesResponsePtrType) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsolePropertiesResponsePtrOutput)
}

type ConsolePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConsolePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsolePropertiesResponse)(nil)).Elem()
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput {
	return o
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponseOutputWithContext(ctx context.Context) ConsolePropertiesResponseOutput {
	return o
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return o.ToConsolePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsolePropertiesResponse) *ConsolePropertiesResponse {
		return &v
	}).(ConsolePropertiesResponsePtrOutput)
}

func (o ConsolePropertiesResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ConsolePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConsolePropertiesResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ConsolePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConsolePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsolePropertiesResponse)(nil)).Elem()
}

func (o ConsolePropertiesResponsePtrOutput) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return o
}

func (o ConsolePropertiesResponsePtrOutput) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return o
}

func (o ConsolePropertiesResponsePtrOutput) Elem() ConsolePropertiesResponseOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) ConsolePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConsolePropertiesResponse
		return ret
	}).(ConsolePropertiesResponseOutput)
}

func (o ConsolePropertiesResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ConsolePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ConsolePropertiesResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type DashboardLens struct {
	Metadata map[string]interface{} `pulumi:"metadata"`
	Order    int                    `pulumi:"order"`
	Parts    []DashboardParts       `pulumi:"parts"`
}





type DashboardLensInput interface {
	pulumi.Input

	ToDashboardLensOutput() DashboardLensOutput
	ToDashboardLensOutputWithContext(context.Context) DashboardLensOutput
}

type DashboardLensArgs struct {
	Metadata pulumi.MapInput          `pulumi:"metadata"`
	Order    pulumi.IntInput          `pulumi:"order"`
	Parts    DashboardPartsArrayInput `pulumi:"parts"`
}

func (DashboardLensArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLens)(nil)).Elem()
}

func (i DashboardLensArgs) ToDashboardLensOutput() DashboardLensOutput {
	return i.ToDashboardLensOutputWithContext(context.Background())
}

func (i DashboardLensArgs) ToDashboardLensOutputWithContext(ctx context.Context) DashboardLensOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensOutput)
}





type DashboardLensArrayInput interface {
	pulumi.Input

	ToDashboardLensArrayOutput() DashboardLensArrayOutput
	ToDashboardLensArrayOutputWithContext(context.Context) DashboardLensArrayOutput
}

type DashboardLensArray []DashboardLensInput

func (DashboardLensArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardLens)(nil)).Elem()
}

func (i DashboardLensArray) ToDashboardLensArrayOutput() DashboardLensArrayOutput {
	return i.ToDashboardLensArrayOutputWithContext(context.Background())
}

func (i DashboardLensArray) ToDashboardLensArrayOutputWithContext(ctx context.Context) DashboardLensArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensArrayOutput)
}

type DashboardLensOutput struct{ *pulumi.OutputState }

func (DashboardLensOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLens)(nil)).Elem()
}

func (o DashboardLensOutput) ToDashboardLensOutput() DashboardLensOutput {
	return o
}

func (o DashboardLensOutput) ToDashboardLensOutputWithContext(ctx context.Context) DashboardLensOutput {
	return o
}

func (o DashboardLensOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardLens) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardLensOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardLens) int { return v.Order }).(pulumi.IntOutput)
}

func (o DashboardLensOutput) Parts() DashboardPartsArrayOutput {
	return o.ApplyT(func(v DashboardLens) []DashboardParts { return v.Parts }).(DashboardPartsArrayOutput)
}

type DashboardLensArrayOutput struct{ *pulumi.OutputState }

func (DashboardLensArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardLens)(nil)).Elem()
}

func (o DashboardLensArrayOutput) ToDashboardLensArrayOutput() DashboardLensArrayOutput {
	return o
}

func (o DashboardLensArrayOutput) ToDashboardLensArrayOutputWithContext(ctx context.Context) DashboardLensArrayOutput {
	return o
}

func (o DashboardLensArrayOutput) Index(i pulumi.IntInput) DashboardLensOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DashboardLens {
		return vs[0].([]DashboardLens)[vs[1].(int)]
	}).(DashboardLensOutput)
}

type DashboardLensResponse struct {
	Metadata map[string]interface{}   `pulumi:"metadata"`
	Order    int                      `pulumi:"order"`
	Parts    []DashboardPartsResponse `pulumi:"parts"`
}





type DashboardLensResponseInput interface {
	pulumi.Input

	ToDashboardLensResponseOutput() DashboardLensResponseOutput
	ToDashboardLensResponseOutputWithContext(context.Context) DashboardLensResponseOutput
}

type DashboardLensResponseArgs struct {
	Metadata pulumi.MapInput                  `pulumi:"metadata"`
	Order    pulumi.IntInput                  `pulumi:"order"`
	Parts    DashboardPartsResponseArrayInput `pulumi:"parts"`
}

func (DashboardLensResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLensResponse)(nil)).Elem()
}

func (i DashboardLensResponseArgs) ToDashboardLensResponseOutput() DashboardLensResponseOutput {
	return i.ToDashboardLensResponseOutputWithContext(context.Background())
}

func (i DashboardLensResponseArgs) ToDashboardLensResponseOutputWithContext(ctx context.Context) DashboardLensResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensResponseOutput)
}





type DashboardLensResponseArrayInput interface {
	pulumi.Input

	ToDashboardLensResponseArrayOutput() DashboardLensResponseArrayOutput
	ToDashboardLensResponseArrayOutputWithContext(context.Context) DashboardLensResponseArrayOutput
}

type DashboardLensResponseArray []DashboardLensResponseInput

func (DashboardLensResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardLensResponse)(nil)).Elem()
}

func (i DashboardLensResponseArray) ToDashboardLensResponseArrayOutput() DashboardLensResponseArrayOutput {
	return i.ToDashboardLensResponseArrayOutputWithContext(context.Background())
}

func (i DashboardLensResponseArray) ToDashboardLensResponseArrayOutputWithContext(ctx context.Context) DashboardLensResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensResponseArrayOutput)
}

type DashboardLensResponseOutput struct{ *pulumi.OutputState }

func (DashboardLensResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLensResponse)(nil)).Elem()
}

func (o DashboardLensResponseOutput) ToDashboardLensResponseOutput() DashboardLensResponseOutput {
	return o
}

func (o DashboardLensResponseOutput) ToDashboardLensResponseOutputWithContext(ctx context.Context) DashboardLensResponseOutput {
	return o
}

func (o DashboardLensResponseOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardLensResponse) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardLensResponseOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardLensResponse) int { return v.Order }).(pulumi.IntOutput)
}

func (o DashboardLensResponseOutput) Parts() DashboardPartsResponseArrayOutput {
	return o.ApplyT(func(v DashboardLensResponse) []DashboardPartsResponse { return v.Parts }).(DashboardPartsResponseArrayOutput)
}

type DashboardLensResponseArrayOutput struct{ *pulumi.OutputState }

func (DashboardLensResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardLensResponse)(nil)).Elem()
}

func (o DashboardLensResponseArrayOutput) ToDashboardLensResponseArrayOutput() DashboardLensResponseArrayOutput {
	return o
}

func (o DashboardLensResponseArrayOutput) ToDashboardLensResponseArrayOutputWithContext(ctx context.Context) DashboardLensResponseArrayOutput {
	return o
}

func (o DashboardLensResponseArrayOutput) Index(i pulumi.IntInput) DashboardLensResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DashboardLensResponse {
		return vs[0].([]DashboardLensResponse)[vs[1].(int)]
	}).(DashboardLensResponseOutput)
}

type DashboardParts struct {
	Metadata *MarkdownPartMetadata  `pulumi:"metadata"`
	Position DashboardPartsPosition `pulumi:"position"`
}





type DashboardPartsInput interface {
	pulumi.Input

	ToDashboardPartsOutput() DashboardPartsOutput
	ToDashboardPartsOutputWithContext(context.Context) DashboardPartsOutput
}

type DashboardPartsArgs struct {
	Metadata MarkdownPartMetadataPtrInput `pulumi:"metadata"`
	Position DashboardPartsPositionInput  `pulumi:"position"`
}

func (DashboardPartsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardParts)(nil)).Elem()
}

func (i DashboardPartsArgs) ToDashboardPartsOutput() DashboardPartsOutput {
	return i.ToDashboardPartsOutputWithContext(context.Background())
}

func (i DashboardPartsArgs) ToDashboardPartsOutputWithContext(ctx context.Context) DashboardPartsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsOutput)
}





type DashboardPartsArrayInput interface {
	pulumi.Input

	ToDashboardPartsArrayOutput() DashboardPartsArrayOutput
	ToDashboardPartsArrayOutputWithContext(context.Context) DashboardPartsArrayOutput
}

type DashboardPartsArray []DashboardPartsInput

func (DashboardPartsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardParts)(nil)).Elem()
}

func (i DashboardPartsArray) ToDashboardPartsArrayOutput() DashboardPartsArrayOutput {
	return i.ToDashboardPartsArrayOutputWithContext(context.Background())
}

func (i DashboardPartsArray) ToDashboardPartsArrayOutputWithContext(ctx context.Context) DashboardPartsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsArrayOutput)
}

type DashboardPartsOutput struct{ *pulumi.OutputState }

func (DashboardPartsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardParts)(nil)).Elem()
}

func (o DashboardPartsOutput) ToDashboardPartsOutput() DashboardPartsOutput {
	return o
}

func (o DashboardPartsOutput) ToDashboardPartsOutputWithContext(ctx context.Context) DashboardPartsOutput {
	return o
}

func (o DashboardPartsOutput) Metadata() MarkdownPartMetadataPtrOutput {
	return o.ApplyT(func(v DashboardParts) *MarkdownPartMetadata { return v.Metadata }).(MarkdownPartMetadataPtrOutput)
}

func (o DashboardPartsOutput) Position() DashboardPartsPositionOutput {
	return o.ApplyT(func(v DashboardParts) DashboardPartsPosition { return v.Position }).(DashboardPartsPositionOutput)
}

type DashboardPartsArrayOutput struct{ *pulumi.OutputState }

func (DashboardPartsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardParts)(nil)).Elem()
}

func (o DashboardPartsArrayOutput) ToDashboardPartsArrayOutput() DashboardPartsArrayOutput {
	return o
}

func (o DashboardPartsArrayOutput) ToDashboardPartsArrayOutputWithContext(ctx context.Context) DashboardPartsArrayOutput {
	return o
}

func (o DashboardPartsArrayOutput) Index(i pulumi.IntInput) DashboardPartsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DashboardParts {
		return vs[0].([]DashboardParts)[vs[1].(int)]
	}).(DashboardPartsOutput)
}

type DashboardPartsPosition struct {
	ColSpan  int                    `pulumi:"colSpan"`
	Metadata map[string]interface{} `pulumi:"metadata"`
	RowSpan  int                    `pulumi:"rowSpan"`
	X        int                    `pulumi:"x"`
	Y        int                    `pulumi:"y"`
}





type DashboardPartsPositionInput interface {
	pulumi.Input

	ToDashboardPartsPositionOutput() DashboardPartsPositionOutput
	ToDashboardPartsPositionOutputWithContext(context.Context) DashboardPartsPositionOutput
}

type DashboardPartsPositionArgs struct {
	ColSpan  pulumi.IntInput `pulumi:"colSpan"`
	Metadata pulumi.MapInput `pulumi:"metadata"`
	RowSpan  pulumi.IntInput `pulumi:"rowSpan"`
	X        pulumi.IntInput `pulumi:"x"`
	Y        pulumi.IntInput `pulumi:"y"`
}

func (DashboardPartsPositionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsPosition)(nil)).Elem()
}

func (i DashboardPartsPositionArgs) ToDashboardPartsPositionOutput() DashboardPartsPositionOutput {
	return i.ToDashboardPartsPositionOutputWithContext(context.Background())
}

func (i DashboardPartsPositionArgs) ToDashboardPartsPositionOutputWithContext(ctx context.Context) DashboardPartsPositionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsPositionOutput)
}

type DashboardPartsPositionOutput struct{ *pulumi.OutputState }

func (DashboardPartsPositionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsPosition)(nil)).Elem()
}

func (o DashboardPartsPositionOutput) ToDashboardPartsPositionOutput() DashboardPartsPositionOutput {
	return o
}

func (o DashboardPartsPositionOutput) ToDashboardPartsPositionOutputWithContext(ctx context.Context) DashboardPartsPositionOutput {
	return o
}

func (o DashboardPartsPositionOutput) ColSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.ColSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsPositionOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardPartsPosition) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardPartsPositionOutput) RowSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.RowSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsPositionOutput) X() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.X }).(pulumi.IntOutput)
}

func (o DashboardPartsPositionOutput) Y() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.Y }).(pulumi.IntOutput)
}

type DashboardPartsResponse struct {
	Metadata *MarkdownPartMetadataResponse  `pulumi:"metadata"`
	Position DashboardPartsResponsePosition `pulumi:"position"`
}





type DashboardPartsResponseInput interface {
	pulumi.Input

	ToDashboardPartsResponseOutput() DashboardPartsResponseOutput
	ToDashboardPartsResponseOutputWithContext(context.Context) DashboardPartsResponseOutput
}

type DashboardPartsResponseArgs struct {
	Metadata MarkdownPartMetadataResponsePtrInput `pulumi:"metadata"`
	Position DashboardPartsResponsePositionInput  `pulumi:"position"`
}

func (DashboardPartsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponse)(nil)).Elem()
}

func (i DashboardPartsResponseArgs) ToDashboardPartsResponseOutput() DashboardPartsResponseOutput {
	return i.ToDashboardPartsResponseOutputWithContext(context.Background())
}

func (i DashboardPartsResponseArgs) ToDashboardPartsResponseOutputWithContext(ctx context.Context) DashboardPartsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsResponseOutput)
}





type DashboardPartsResponseArrayInput interface {
	pulumi.Input

	ToDashboardPartsResponseArrayOutput() DashboardPartsResponseArrayOutput
	ToDashboardPartsResponseArrayOutputWithContext(context.Context) DashboardPartsResponseArrayOutput
}

type DashboardPartsResponseArray []DashboardPartsResponseInput

func (DashboardPartsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardPartsResponse)(nil)).Elem()
}

func (i DashboardPartsResponseArray) ToDashboardPartsResponseArrayOutput() DashboardPartsResponseArrayOutput {
	return i.ToDashboardPartsResponseArrayOutputWithContext(context.Background())
}

func (i DashboardPartsResponseArray) ToDashboardPartsResponseArrayOutputWithContext(ctx context.Context) DashboardPartsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsResponseArrayOutput)
}

type DashboardPartsResponseOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponse)(nil)).Elem()
}

func (o DashboardPartsResponseOutput) ToDashboardPartsResponseOutput() DashboardPartsResponseOutput {
	return o
}

func (o DashboardPartsResponseOutput) ToDashboardPartsResponseOutputWithContext(ctx context.Context) DashboardPartsResponseOutput {
	return o
}

func (o DashboardPartsResponseOutput) Metadata() MarkdownPartMetadataResponsePtrOutput {
	return o.ApplyT(func(v DashboardPartsResponse) *MarkdownPartMetadataResponse { return v.Metadata }).(MarkdownPartMetadataResponsePtrOutput)
}

func (o DashboardPartsResponseOutput) Position() DashboardPartsResponsePositionOutput {
	return o.ApplyT(func(v DashboardPartsResponse) DashboardPartsResponsePosition { return v.Position }).(DashboardPartsResponsePositionOutput)
}

type DashboardPartsResponseArrayOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardPartsResponse)(nil)).Elem()
}

func (o DashboardPartsResponseArrayOutput) ToDashboardPartsResponseArrayOutput() DashboardPartsResponseArrayOutput {
	return o
}

func (o DashboardPartsResponseArrayOutput) ToDashboardPartsResponseArrayOutputWithContext(ctx context.Context) DashboardPartsResponseArrayOutput {
	return o
}

func (o DashboardPartsResponseArrayOutput) Index(i pulumi.IntInput) DashboardPartsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DashboardPartsResponse {
		return vs[0].([]DashboardPartsResponse)[vs[1].(int)]
	}).(DashboardPartsResponseOutput)
}

type DashboardPartsResponsePosition struct {
	ColSpan  int                    `pulumi:"colSpan"`
	Metadata map[string]interface{} `pulumi:"metadata"`
	RowSpan  int                    `pulumi:"rowSpan"`
	X        int                    `pulumi:"x"`
	Y        int                    `pulumi:"y"`
}





type DashboardPartsResponsePositionInput interface {
	pulumi.Input

	ToDashboardPartsResponsePositionOutput() DashboardPartsResponsePositionOutput
	ToDashboardPartsResponsePositionOutputWithContext(context.Context) DashboardPartsResponsePositionOutput
}

type DashboardPartsResponsePositionArgs struct {
	ColSpan  pulumi.IntInput `pulumi:"colSpan"`
	Metadata pulumi.MapInput `pulumi:"metadata"`
	RowSpan  pulumi.IntInput `pulumi:"rowSpan"`
	X        pulumi.IntInput `pulumi:"x"`
	Y        pulumi.IntInput `pulumi:"y"`
}

func (DashboardPartsResponsePositionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponsePosition)(nil)).Elem()
}

func (i DashboardPartsResponsePositionArgs) ToDashboardPartsResponsePositionOutput() DashboardPartsResponsePositionOutput {
	return i.ToDashboardPartsResponsePositionOutputWithContext(context.Background())
}

func (i DashboardPartsResponsePositionArgs) ToDashboardPartsResponsePositionOutputWithContext(ctx context.Context) DashboardPartsResponsePositionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsResponsePositionOutput)
}

type DashboardPartsResponsePositionOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponsePositionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponsePosition)(nil)).Elem()
}

func (o DashboardPartsResponsePositionOutput) ToDashboardPartsResponsePositionOutput() DashboardPartsResponsePositionOutput {
	return o
}

func (o DashboardPartsResponsePositionOutput) ToDashboardPartsResponsePositionOutputWithContext(ctx context.Context) DashboardPartsResponsePositionOutput {
	return o
}

func (o DashboardPartsResponsePositionOutput) ColSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.ColSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsResponsePositionOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardPartsResponsePositionOutput) RowSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.RowSpan }).(pulumi.IntOutput)
}

func (o DashboardPartsResponsePositionOutput) X() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.X }).(pulumi.IntOutput)
}

func (o DashboardPartsResponsePositionOutput) Y() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.Y }).(pulumi.IntOutput)
}

type MarkdownPartMetadata struct {
	Inputs   []interface{}                 `pulumi:"inputs"`
	Settings *MarkdownPartMetadataSettings `pulumi:"settings"`
	Type     string                        `pulumi:"type"`
}





type MarkdownPartMetadataInput interface {
	pulumi.Input

	ToMarkdownPartMetadataOutput() MarkdownPartMetadataOutput
	ToMarkdownPartMetadataOutputWithContext(context.Context) MarkdownPartMetadataOutput
}

type MarkdownPartMetadataArgs struct {
	Inputs   pulumi.ArrayInput                    `pulumi:"inputs"`
	Settings MarkdownPartMetadataSettingsPtrInput `pulumi:"settings"`
	Type     pulumi.StringInput                   `pulumi:"type"`
}

func (MarkdownPartMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadata)(nil)).Elem()
}

func (i MarkdownPartMetadataArgs) ToMarkdownPartMetadataOutput() MarkdownPartMetadataOutput {
	return i.ToMarkdownPartMetadataOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataArgs) ToMarkdownPartMetadataOutputWithContext(ctx context.Context) MarkdownPartMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataOutput)
}

func (i MarkdownPartMetadataArgs) ToMarkdownPartMetadataPtrOutput() MarkdownPartMetadataPtrOutput {
	return i.ToMarkdownPartMetadataPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataArgs) ToMarkdownPartMetadataPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataOutput).ToMarkdownPartMetadataPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataPtrOutput() MarkdownPartMetadataPtrOutput
	ToMarkdownPartMetadataPtrOutputWithContext(context.Context) MarkdownPartMetadataPtrOutput
}

type markdownPartMetadataPtrType MarkdownPartMetadataArgs

func MarkdownPartMetadataPtr(v *MarkdownPartMetadataArgs) MarkdownPartMetadataPtrInput {
	return (*markdownPartMetadataPtrType)(v)
}

func (*markdownPartMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadata)(nil)).Elem()
}

func (i *markdownPartMetadataPtrType) ToMarkdownPartMetadataPtrOutput() MarkdownPartMetadataPtrOutput {
	return i.ToMarkdownPartMetadataPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataPtrType) ToMarkdownPartMetadataPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataPtrOutput)
}

type MarkdownPartMetadataOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadata)(nil)).Elem()
}

func (o MarkdownPartMetadataOutput) ToMarkdownPartMetadataOutput() MarkdownPartMetadataOutput {
	return o
}

func (o MarkdownPartMetadataOutput) ToMarkdownPartMetadataOutputWithContext(ctx context.Context) MarkdownPartMetadataOutput {
	return o
}

func (o MarkdownPartMetadataOutput) ToMarkdownPartMetadataPtrOutput() MarkdownPartMetadataPtrOutput {
	return o.ToMarkdownPartMetadataPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataOutput) ToMarkdownPartMetadataPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadata) *MarkdownPartMetadata {
		return &v
	}).(MarkdownPartMetadataPtrOutput)
}

func (o MarkdownPartMetadataOutput) Inputs() pulumi.ArrayOutput {
	return o.ApplyT(func(v MarkdownPartMetadata) []interface{} { return v.Inputs }).(pulumi.ArrayOutput)
}

func (o MarkdownPartMetadataOutput) Settings() MarkdownPartMetadataSettingsPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadata) *MarkdownPartMetadataSettings { return v.Settings }).(MarkdownPartMetadataSettingsPtrOutput)
}

func (o MarkdownPartMetadataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MarkdownPartMetadata) string { return v.Type }).(pulumi.StringOutput)
}

type MarkdownPartMetadataPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadata)(nil)).Elem()
}

func (o MarkdownPartMetadataPtrOutput) ToMarkdownPartMetadataPtrOutput() MarkdownPartMetadataPtrOutput {
	return o
}

func (o MarkdownPartMetadataPtrOutput) ToMarkdownPartMetadataPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataPtrOutput {
	return o
}

func (o MarkdownPartMetadataPtrOutput) Elem() MarkdownPartMetadataOutput {
	return o.ApplyT(func(v *MarkdownPartMetadata) MarkdownPartMetadata {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadata
		return ret
	}).(MarkdownPartMetadataOutput)
}

func (o MarkdownPartMetadataPtrOutput) Inputs() pulumi.ArrayOutput {
	return o.ApplyT(func(v *MarkdownPartMetadata) []interface{} {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(pulumi.ArrayOutput)
}

func (o MarkdownPartMetadataPtrOutput) Settings() MarkdownPartMetadataSettingsPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadata) *MarkdownPartMetadataSettings {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(MarkdownPartMetadataSettingsPtrOutput)
}

func (o MarkdownPartMetadataPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadata) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MarkdownPartMetadataContent struct {
	Settings *MarkdownPartMetadataSettingsSettings `pulumi:"settings"`
}





type MarkdownPartMetadataContentInput interface {
	pulumi.Input

	ToMarkdownPartMetadataContentOutput() MarkdownPartMetadataContentOutput
	ToMarkdownPartMetadataContentOutputWithContext(context.Context) MarkdownPartMetadataContentOutput
}

type MarkdownPartMetadataContentArgs struct {
	Settings MarkdownPartMetadataSettingsSettingsPtrInput `pulumi:"settings"`
}

func (MarkdownPartMetadataContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataContent)(nil)).Elem()
}

func (i MarkdownPartMetadataContentArgs) ToMarkdownPartMetadataContentOutput() MarkdownPartMetadataContentOutput {
	return i.ToMarkdownPartMetadataContentOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataContentArgs) ToMarkdownPartMetadataContentOutputWithContext(ctx context.Context) MarkdownPartMetadataContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataContentOutput)
}

func (i MarkdownPartMetadataContentArgs) ToMarkdownPartMetadataContentPtrOutput() MarkdownPartMetadataContentPtrOutput {
	return i.ToMarkdownPartMetadataContentPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataContentArgs) ToMarkdownPartMetadataContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataContentOutput).ToMarkdownPartMetadataContentPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataContentPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataContentPtrOutput() MarkdownPartMetadataContentPtrOutput
	ToMarkdownPartMetadataContentPtrOutputWithContext(context.Context) MarkdownPartMetadataContentPtrOutput
}

type markdownPartMetadataContentPtrType MarkdownPartMetadataContentArgs

func MarkdownPartMetadataContentPtr(v *MarkdownPartMetadataContentArgs) MarkdownPartMetadataContentPtrInput {
	return (*markdownPartMetadataContentPtrType)(v)
}

func (*markdownPartMetadataContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataContent)(nil)).Elem()
}

func (i *markdownPartMetadataContentPtrType) ToMarkdownPartMetadataContentPtrOutput() MarkdownPartMetadataContentPtrOutput {
	return i.ToMarkdownPartMetadataContentPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataContentPtrType) ToMarkdownPartMetadataContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataContentPtrOutput)
}

type MarkdownPartMetadataContentOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataContent)(nil)).Elem()
}

func (o MarkdownPartMetadataContentOutput) ToMarkdownPartMetadataContentOutput() MarkdownPartMetadataContentOutput {
	return o
}

func (o MarkdownPartMetadataContentOutput) ToMarkdownPartMetadataContentOutputWithContext(ctx context.Context) MarkdownPartMetadataContentOutput {
	return o
}

func (o MarkdownPartMetadataContentOutput) ToMarkdownPartMetadataContentPtrOutput() MarkdownPartMetadataContentPtrOutput {
	return o.ToMarkdownPartMetadataContentPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataContentOutput) ToMarkdownPartMetadataContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataContent) *MarkdownPartMetadataContent {
		return &v
	}).(MarkdownPartMetadataContentPtrOutput)
}

func (o MarkdownPartMetadataContentOutput) Settings() MarkdownPartMetadataSettingsSettingsPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataContent) *MarkdownPartMetadataSettingsSettings { return v.Settings }).(MarkdownPartMetadataSettingsSettingsPtrOutput)
}

type MarkdownPartMetadataContentPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataContent)(nil)).Elem()
}

func (o MarkdownPartMetadataContentPtrOutput) ToMarkdownPartMetadataContentPtrOutput() MarkdownPartMetadataContentPtrOutput {
	return o
}

func (o MarkdownPartMetadataContentPtrOutput) ToMarkdownPartMetadataContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataContentPtrOutput {
	return o
}

func (o MarkdownPartMetadataContentPtrOutput) Elem() MarkdownPartMetadataContentOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataContent) MarkdownPartMetadataContent {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataContent
		return ret
	}).(MarkdownPartMetadataContentOutput)
}

func (o MarkdownPartMetadataContentPtrOutput) Settings() MarkdownPartMetadataSettingsSettingsPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataContent) *MarkdownPartMetadataSettingsSettings {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(MarkdownPartMetadataSettingsSettingsPtrOutput)
}

type MarkdownPartMetadataResponse struct {
	Inputs   []interface{}                         `pulumi:"inputs"`
	Settings *MarkdownPartMetadataResponseSettings `pulumi:"settings"`
	Type     string                                `pulumi:"type"`
}





type MarkdownPartMetadataResponseInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseOutput() MarkdownPartMetadataResponseOutput
	ToMarkdownPartMetadataResponseOutputWithContext(context.Context) MarkdownPartMetadataResponseOutput
}

type MarkdownPartMetadataResponseArgs struct {
	Inputs   pulumi.ArrayInput                            `pulumi:"inputs"`
	Settings MarkdownPartMetadataResponseSettingsPtrInput `pulumi:"settings"`
	Type     pulumi.StringInput                           `pulumi:"type"`
}

func (MarkdownPartMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponse)(nil)).Elem()
}

func (i MarkdownPartMetadataResponseArgs) ToMarkdownPartMetadataResponseOutput() MarkdownPartMetadataResponseOutput {
	return i.ToMarkdownPartMetadataResponseOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseArgs) ToMarkdownPartMetadataResponseOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseOutput)
}

func (i MarkdownPartMetadataResponseArgs) ToMarkdownPartMetadataResponsePtrOutput() MarkdownPartMetadataResponsePtrOutput {
	return i.ToMarkdownPartMetadataResponsePtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseArgs) ToMarkdownPartMetadataResponsePtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseOutput).ToMarkdownPartMetadataResponsePtrOutputWithContext(ctx)
}









type MarkdownPartMetadataResponsePtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponsePtrOutput() MarkdownPartMetadataResponsePtrOutput
	ToMarkdownPartMetadataResponsePtrOutputWithContext(context.Context) MarkdownPartMetadataResponsePtrOutput
}

type markdownPartMetadataResponsePtrType MarkdownPartMetadataResponseArgs

func MarkdownPartMetadataResponsePtr(v *MarkdownPartMetadataResponseArgs) MarkdownPartMetadataResponsePtrInput {
	return (*markdownPartMetadataResponsePtrType)(v)
}

func (*markdownPartMetadataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponse)(nil)).Elem()
}

func (i *markdownPartMetadataResponsePtrType) ToMarkdownPartMetadataResponsePtrOutput() MarkdownPartMetadataResponsePtrOutput {
	return i.ToMarkdownPartMetadataResponsePtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataResponsePtrType) ToMarkdownPartMetadataResponsePtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponsePtrOutput)
}

type MarkdownPartMetadataResponseOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponse)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseOutput) ToMarkdownPartMetadataResponseOutput() MarkdownPartMetadataResponseOutput {
	return o
}

func (o MarkdownPartMetadataResponseOutput) ToMarkdownPartMetadataResponseOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseOutput {
	return o
}

func (o MarkdownPartMetadataResponseOutput) ToMarkdownPartMetadataResponsePtrOutput() MarkdownPartMetadataResponsePtrOutput {
	return o.ToMarkdownPartMetadataResponsePtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataResponseOutput) ToMarkdownPartMetadataResponsePtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataResponse) *MarkdownPartMetadataResponse {
		return &v
	}).(MarkdownPartMetadataResponsePtrOutput)
}

func (o MarkdownPartMetadataResponseOutput) Inputs() pulumi.ArrayOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponse) []interface{} { return v.Inputs }).(pulumi.ArrayOutput)
}

func (o MarkdownPartMetadataResponseOutput) Settings() MarkdownPartMetadataResponseSettingsPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponse) *MarkdownPartMetadataResponseSettings { return v.Settings }).(MarkdownPartMetadataResponseSettingsPtrOutput)
}

func (o MarkdownPartMetadataResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponse) string { return v.Type }).(pulumi.StringOutput)
}

type MarkdownPartMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponse)(nil)).Elem()
}

func (o MarkdownPartMetadataResponsePtrOutput) ToMarkdownPartMetadataResponsePtrOutput() MarkdownPartMetadataResponsePtrOutput {
	return o
}

func (o MarkdownPartMetadataResponsePtrOutput) ToMarkdownPartMetadataResponsePtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponsePtrOutput {
	return o
}

func (o MarkdownPartMetadataResponsePtrOutput) Elem() MarkdownPartMetadataResponseOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponse) MarkdownPartMetadataResponse {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataResponse
		return ret
	}).(MarkdownPartMetadataResponseOutput)
}

func (o MarkdownPartMetadataResponsePtrOutput) Inputs() pulumi.ArrayOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(pulumi.ArrayOutput)
}

func (o MarkdownPartMetadataResponsePtrOutput) Settings() MarkdownPartMetadataResponseSettingsPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponse) *MarkdownPartMetadataResponseSettings {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(MarkdownPartMetadataResponseSettingsPtrOutput)
}

func (o MarkdownPartMetadataResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MarkdownPartMetadataResponseContent struct {
	Settings *MarkdownPartMetadataResponseSettingsSettings `pulumi:"settings"`
}





type MarkdownPartMetadataResponseContentInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseContentOutput() MarkdownPartMetadataResponseContentOutput
	ToMarkdownPartMetadataResponseContentOutputWithContext(context.Context) MarkdownPartMetadataResponseContentOutput
}

type MarkdownPartMetadataResponseContentArgs struct {
	Settings MarkdownPartMetadataResponseSettingsSettingsPtrInput `pulumi:"settings"`
}

func (MarkdownPartMetadataResponseContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponseContent)(nil)).Elem()
}

func (i MarkdownPartMetadataResponseContentArgs) ToMarkdownPartMetadataResponseContentOutput() MarkdownPartMetadataResponseContentOutput {
	return i.ToMarkdownPartMetadataResponseContentOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseContentArgs) ToMarkdownPartMetadataResponseContentOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseContentOutput)
}

func (i MarkdownPartMetadataResponseContentArgs) ToMarkdownPartMetadataResponseContentPtrOutput() MarkdownPartMetadataResponseContentPtrOutput {
	return i.ToMarkdownPartMetadataResponseContentPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseContentArgs) ToMarkdownPartMetadataResponseContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseContentOutput).ToMarkdownPartMetadataResponseContentPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataResponseContentPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseContentPtrOutput() MarkdownPartMetadataResponseContentPtrOutput
	ToMarkdownPartMetadataResponseContentPtrOutputWithContext(context.Context) MarkdownPartMetadataResponseContentPtrOutput
}

type markdownPartMetadataResponseContentPtrType MarkdownPartMetadataResponseContentArgs

func MarkdownPartMetadataResponseContentPtr(v *MarkdownPartMetadataResponseContentArgs) MarkdownPartMetadataResponseContentPtrInput {
	return (*markdownPartMetadataResponseContentPtrType)(v)
}

func (*markdownPartMetadataResponseContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponseContent)(nil)).Elem()
}

func (i *markdownPartMetadataResponseContentPtrType) ToMarkdownPartMetadataResponseContentPtrOutput() MarkdownPartMetadataResponseContentPtrOutput {
	return i.ToMarkdownPartMetadataResponseContentPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataResponseContentPtrType) ToMarkdownPartMetadataResponseContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseContentPtrOutput)
}

type MarkdownPartMetadataResponseContentOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponseContent)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseContentOutput) ToMarkdownPartMetadataResponseContentOutput() MarkdownPartMetadataResponseContentOutput {
	return o
}

func (o MarkdownPartMetadataResponseContentOutput) ToMarkdownPartMetadataResponseContentOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseContentOutput {
	return o
}

func (o MarkdownPartMetadataResponseContentOutput) ToMarkdownPartMetadataResponseContentPtrOutput() MarkdownPartMetadataResponseContentPtrOutput {
	return o.ToMarkdownPartMetadataResponseContentPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataResponseContentOutput) ToMarkdownPartMetadataResponseContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataResponseContent) *MarkdownPartMetadataResponseContent {
		return &v
	}).(MarkdownPartMetadataResponseContentPtrOutput)
}

func (o MarkdownPartMetadataResponseContentOutput) Settings() MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseContent) *MarkdownPartMetadataResponseSettingsSettings {
		return v.Settings
	}).(MarkdownPartMetadataResponseSettingsSettingsPtrOutput)
}

type MarkdownPartMetadataResponseContentPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponseContent)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseContentPtrOutput) ToMarkdownPartMetadataResponseContentPtrOutput() MarkdownPartMetadataResponseContentPtrOutput {
	return o
}

func (o MarkdownPartMetadataResponseContentPtrOutput) ToMarkdownPartMetadataResponseContentPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseContentPtrOutput {
	return o
}

func (o MarkdownPartMetadataResponseContentPtrOutput) Elem() MarkdownPartMetadataResponseContentOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseContent) MarkdownPartMetadataResponseContent {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataResponseContent
		return ret
	}).(MarkdownPartMetadataResponseContentOutput)
}

func (o MarkdownPartMetadataResponseContentPtrOutput) Settings() MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseContent) *MarkdownPartMetadataResponseSettingsSettings {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(MarkdownPartMetadataResponseSettingsSettingsPtrOutput)
}

type MarkdownPartMetadataResponseSettings struct {
	Content *MarkdownPartMetadataResponseContent `pulumi:"content"`
}





type MarkdownPartMetadataResponseSettingsInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseSettingsOutput() MarkdownPartMetadataResponseSettingsOutput
	ToMarkdownPartMetadataResponseSettingsOutputWithContext(context.Context) MarkdownPartMetadataResponseSettingsOutput
}

type MarkdownPartMetadataResponseSettingsArgs struct {
	Content MarkdownPartMetadataResponseContentPtrInput `pulumi:"content"`
}

func (MarkdownPartMetadataResponseSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponseSettings)(nil)).Elem()
}

func (i MarkdownPartMetadataResponseSettingsArgs) ToMarkdownPartMetadataResponseSettingsOutput() MarkdownPartMetadataResponseSettingsOutput {
	return i.ToMarkdownPartMetadataResponseSettingsOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseSettingsArgs) ToMarkdownPartMetadataResponseSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseSettingsOutput)
}

func (i MarkdownPartMetadataResponseSettingsArgs) ToMarkdownPartMetadataResponseSettingsPtrOutput() MarkdownPartMetadataResponseSettingsPtrOutput {
	return i.ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseSettingsArgs) ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseSettingsOutput).ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataResponseSettingsPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseSettingsPtrOutput() MarkdownPartMetadataResponseSettingsPtrOutput
	ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(context.Context) MarkdownPartMetadataResponseSettingsPtrOutput
}

type markdownPartMetadataResponseSettingsPtrType MarkdownPartMetadataResponseSettingsArgs

func MarkdownPartMetadataResponseSettingsPtr(v *MarkdownPartMetadataResponseSettingsArgs) MarkdownPartMetadataResponseSettingsPtrInput {
	return (*markdownPartMetadataResponseSettingsPtrType)(v)
}

func (*markdownPartMetadataResponseSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponseSettings)(nil)).Elem()
}

func (i *markdownPartMetadataResponseSettingsPtrType) ToMarkdownPartMetadataResponseSettingsPtrOutput() MarkdownPartMetadataResponseSettingsPtrOutput {
	return i.ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataResponseSettingsPtrType) ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseSettingsPtrOutput)
}

type MarkdownPartMetadataResponseSettingsOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponseSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseSettingsOutput) ToMarkdownPartMetadataResponseSettingsOutput() MarkdownPartMetadataResponseSettingsOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsOutput) ToMarkdownPartMetadataResponseSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsOutput) ToMarkdownPartMetadataResponseSettingsPtrOutput() MarkdownPartMetadataResponseSettingsPtrOutput {
	return o.ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataResponseSettingsOutput) ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataResponseSettings) *MarkdownPartMetadataResponseSettings {
		return &v
	}).(MarkdownPartMetadataResponseSettingsPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsOutput) Content() MarkdownPartMetadataResponseContentPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseSettings) *MarkdownPartMetadataResponseContent { return v.Content }).(MarkdownPartMetadataResponseContentPtrOutput)
}

type MarkdownPartMetadataResponseSettingsPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponseSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseSettingsPtrOutput) ToMarkdownPartMetadataResponseSettingsPtrOutput() MarkdownPartMetadataResponseSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsPtrOutput) ToMarkdownPartMetadataResponseSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsPtrOutput) Elem() MarkdownPartMetadataResponseSettingsOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettings) MarkdownPartMetadataResponseSettings {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataResponseSettings
		return ret
	}).(MarkdownPartMetadataResponseSettingsOutput)
}

func (o MarkdownPartMetadataResponseSettingsPtrOutput) Content() MarkdownPartMetadataResponseContentPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettings) *MarkdownPartMetadataResponseContent {
		if v == nil {
			return nil
		}
		return v.Content
	}).(MarkdownPartMetadataResponseContentPtrOutput)
}

type MarkdownPartMetadataResponseSettingsSettings struct {
	Content        *string `pulumi:"content"`
	MarkdownSource *int    `pulumi:"markdownSource"`
	MarkdownUri    *string `pulumi:"markdownUri"`
	Subtitle       *string `pulumi:"subtitle"`
	Title          *string `pulumi:"title"`
}





type MarkdownPartMetadataResponseSettingsSettingsInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseSettingsSettingsOutput() MarkdownPartMetadataResponseSettingsSettingsOutput
	ToMarkdownPartMetadataResponseSettingsSettingsOutputWithContext(context.Context) MarkdownPartMetadataResponseSettingsSettingsOutput
}

type MarkdownPartMetadataResponseSettingsSettingsArgs struct {
	Content        pulumi.StringPtrInput `pulumi:"content"`
	MarkdownSource pulumi.IntPtrInput    `pulumi:"markdownSource"`
	MarkdownUri    pulumi.StringPtrInput `pulumi:"markdownUri"`
	Subtitle       pulumi.StringPtrInput `pulumi:"subtitle"`
	Title          pulumi.StringPtrInput `pulumi:"title"`
}

func (MarkdownPartMetadataResponseSettingsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponseSettingsSettings)(nil)).Elem()
}

func (i MarkdownPartMetadataResponseSettingsSettingsArgs) ToMarkdownPartMetadataResponseSettingsSettingsOutput() MarkdownPartMetadataResponseSettingsSettingsOutput {
	return i.ToMarkdownPartMetadataResponseSettingsSettingsOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseSettingsSettingsArgs) ToMarkdownPartMetadataResponseSettingsSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseSettingsSettingsOutput)
}

func (i MarkdownPartMetadataResponseSettingsSettingsArgs) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutput() MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return i.ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataResponseSettingsSettingsArgs) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseSettingsSettingsOutput).ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataResponseSettingsSettingsPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataResponseSettingsSettingsPtrOutput() MarkdownPartMetadataResponseSettingsSettingsPtrOutput
	ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(context.Context) MarkdownPartMetadataResponseSettingsSettingsPtrOutput
}

type markdownPartMetadataResponseSettingsSettingsPtrType MarkdownPartMetadataResponseSettingsSettingsArgs

func MarkdownPartMetadataResponseSettingsSettingsPtr(v *MarkdownPartMetadataResponseSettingsSettingsArgs) MarkdownPartMetadataResponseSettingsSettingsPtrInput {
	return (*markdownPartMetadataResponseSettingsSettingsPtrType)(v)
}

func (*markdownPartMetadataResponseSettingsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponseSettingsSettings)(nil)).Elem()
}

func (i *markdownPartMetadataResponseSettingsSettingsPtrType) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutput() MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return i.ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataResponseSettingsSettingsPtrType) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataResponseSettingsSettingsPtrOutput)
}

type MarkdownPartMetadataResponseSettingsSettingsOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseSettingsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataResponseSettingsSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) ToMarkdownPartMetadataResponseSettingsSettingsOutput() MarkdownPartMetadataResponseSettingsSettingsOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) ToMarkdownPartMetadataResponseSettingsSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsSettingsOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutput() MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return o.ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataResponseSettingsSettings) *MarkdownPartMetadataResponseSettingsSettings {
		return &v
	}).(MarkdownPartMetadataResponseSettingsSettingsPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseSettingsSettings) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) MarkdownSource() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseSettingsSettings) *int { return v.MarkdownSource }).(pulumi.IntPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) MarkdownUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseSettingsSettings) *string { return v.MarkdownUri }).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) Subtitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseSettingsSettings) *string { return v.Subtitle }).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataResponseSettingsSettings) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type MarkdownPartMetadataResponseSettingsSettingsPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataResponseSettingsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataResponseSettingsSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutput() MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) ToMarkdownPartMetadataResponseSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataResponseSettingsSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) Elem() MarkdownPartMetadataResponseSettingsSettingsOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettingsSettings) MarkdownPartMetadataResponseSettingsSettings {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataResponseSettingsSettings
		return ret
	}).(MarkdownPartMetadataResponseSettingsSettingsOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) MarkdownSource() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettingsSettings) *int {
		if v == nil {
			return nil
		}
		return v.MarkdownSource
	}).(pulumi.IntPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) MarkdownUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.MarkdownUri
	}).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) Subtitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Subtitle
	}).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataResponseSettingsSettingsPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataResponseSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type MarkdownPartMetadataSettings struct {
	Content *MarkdownPartMetadataContent `pulumi:"content"`
}





type MarkdownPartMetadataSettingsInput interface {
	pulumi.Input

	ToMarkdownPartMetadataSettingsOutput() MarkdownPartMetadataSettingsOutput
	ToMarkdownPartMetadataSettingsOutputWithContext(context.Context) MarkdownPartMetadataSettingsOutput
}

type MarkdownPartMetadataSettingsArgs struct {
	Content MarkdownPartMetadataContentPtrInput `pulumi:"content"`
}

func (MarkdownPartMetadataSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataSettings)(nil)).Elem()
}

func (i MarkdownPartMetadataSettingsArgs) ToMarkdownPartMetadataSettingsOutput() MarkdownPartMetadataSettingsOutput {
	return i.ToMarkdownPartMetadataSettingsOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataSettingsArgs) ToMarkdownPartMetadataSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataSettingsOutput)
}

func (i MarkdownPartMetadataSettingsArgs) ToMarkdownPartMetadataSettingsPtrOutput() MarkdownPartMetadataSettingsPtrOutput {
	return i.ToMarkdownPartMetadataSettingsPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataSettingsArgs) ToMarkdownPartMetadataSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataSettingsOutput).ToMarkdownPartMetadataSettingsPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataSettingsPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataSettingsPtrOutput() MarkdownPartMetadataSettingsPtrOutput
	ToMarkdownPartMetadataSettingsPtrOutputWithContext(context.Context) MarkdownPartMetadataSettingsPtrOutput
}

type markdownPartMetadataSettingsPtrType MarkdownPartMetadataSettingsArgs

func MarkdownPartMetadataSettingsPtr(v *MarkdownPartMetadataSettingsArgs) MarkdownPartMetadataSettingsPtrInput {
	return (*markdownPartMetadataSettingsPtrType)(v)
}

func (*markdownPartMetadataSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataSettings)(nil)).Elem()
}

func (i *markdownPartMetadataSettingsPtrType) ToMarkdownPartMetadataSettingsPtrOutput() MarkdownPartMetadataSettingsPtrOutput {
	return i.ToMarkdownPartMetadataSettingsPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataSettingsPtrType) ToMarkdownPartMetadataSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataSettingsPtrOutput)
}

type MarkdownPartMetadataSettingsOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataSettingsOutput) ToMarkdownPartMetadataSettingsOutput() MarkdownPartMetadataSettingsOutput {
	return o
}

func (o MarkdownPartMetadataSettingsOutput) ToMarkdownPartMetadataSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsOutput {
	return o
}

func (o MarkdownPartMetadataSettingsOutput) ToMarkdownPartMetadataSettingsPtrOutput() MarkdownPartMetadataSettingsPtrOutput {
	return o.ToMarkdownPartMetadataSettingsPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataSettingsOutput) ToMarkdownPartMetadataSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataSettings) *MarkdownPartMetadataSettings {
		return &v
	}).(MarkdownPartMetadataSettingsPtrOutput)
}

func (o MarkdownPartMetadataSettingsOutput) Content() MarkdownPartMetadataContentPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataSettings) *MarkdownPartMetadataContent { return v.Content }).(MarkdownPartMetadataContentPtrOutput)
}

type MarkdownPartMetadataSettingsPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataSettingsPtrOutput) ToMarkdownPartMetadataSettingsPtrOutput() MarkdownPartMetadataSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataSettingsPtrOutput) ToMarkdownPartMetadataSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataSettingsPtrOutput) Elem() MarkdownPartMetadataSettingsOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettings) MarkdownPartMetadataSettings {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataSettings
		return ret
	}).(MarkdownPartMetadataSettingsOutput)
}

func (o MarkdownPartMetadataSettingsPtrOutput) Content() MarkdownPartMetadataContentPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettings) *MarkdownPartMetadataContent {
		if v == nil {
			return nil
		}
		return v.Content
	}).(MarkdownPartMetadataContentPtrOutput)
}

type MarkdownPartMetadataSettingsSettings struct {
	Content        *string `pulumi:"content"`
	MarkdownSource *int    `pulumi:"markdownSource"`
	MarkdownUri    *string `pulumi:"markdownUri"`
	Subtitle       *string `pulumi:"subtitle"`
	Title          *string `pulumi:"title"`
}





type MarkdownPartMetadataSettingsSettingsInput interface {
	pulumi.Input

	ToMarkdownPartMetadataSettingsSettingsOutput() MarkdownPartMetadataSettingsSettingsOutput
	ToMarkdownPartMetadataSettingsSettingsOutputWithContext(context.Context) MarkdownPartMetadataSettingsSettingsOutput
}

type MarkdownPartMetadataSettingsSettingsArgs struct {
	Content        pulumi.StringPtrInput `pulumi:"content"`
	MarkdownSource pulumi.IntPtrInput    `pulumi:"markdownSource"`
	MarkdownUri    pulumi.StringPtrInput `pulumi:"markdownUri"`
	Subtitle       pulumi.StringPtrInput `pulumi:"subtitle"`
	Title          pulumi.StringPtrInput `pulumi:"title"`
}

func (MarkdownPartMetadataSettingsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataSettingsSettings)(nil)).Elem()
}

func (i MarkdownPartMetadataSettingsSettingsArgs) ToMarkdownPartMetadataSettingsSettingsOutput() MarkdownPartMetadataSettingsSettingsOutput {
	return i.ToMarkdownPartMetadataSettingsSettingsOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataSettingsSettingsArgs) ToMarkdownPartMetadataSettingsSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataSettingsSettingsOutput)
}

func (i MarkdownPartMetadataSettingsSettingsArgs) ToMarkdownPartMetadataSettingsSettingsPtrOutput() MarkdownPartMetadataSettingsSettingsPtrOutput {
	return i.ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(context.Background())
}

func (i MarkdownPartMetadataSettingsSettingsArgs) ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataSettingsSettingsOutput).ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(ctx)
}









type MarkdownPartMetadataSettingsSettingsPtrInput interface {
	pulumi.Input

	ToMarkdownPartMetadataSettingsSettingsPtrOutput() MarkdownPartMetadataSettingsSettingsPtrOutput
	ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(context.Context) MarkdownPartMetadataSettingsSettingsPtrOutput
}

type markdownPartMetadataSettingsSettingsPtrType MarkdownPartMetadataSettingsSettingsArgs

func MarkdownPartMetadataSettingsSettingsPtr(v *MarkdownPartMetadataSettingsSettingsArgs) MarkdownPartMetadataSettingsSettingsPtrInput {
	return (*markdownPartMetadataSettingsSettingsPtrType)(v)
}

func (*markdownPartMetadataSettingsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataSettingsSettings)(nil)).Elem()
}

func (i *markdownPartMetadataSettingsSettingsPtrType) ToMarkdownPartMetadataSettingsSettingsPtrOutput() MarkdownPartMetadataSettingsSettingsPtrOutput {
	return i.ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(context.Background())
}

func (i *markdownPartMetadataSettingsSettingsPtrType) ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkdownPartMetadataSettingsSettingsPtrOutput)
}

type MarkdownPartMetadataSettingsSettingsOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataSettingsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MarkdownPartMetadataSettingsSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataSettingsSettingsOutput) ToMarkdownPartMetadataSettingsSettingsOutput() MarkdownPartMetadataSettingsSettingsOutput {
	return o
}

func (o MarkdownPartMetadataSettingsSettingsOutput) ToMarkdownPartMetadataSettingsSettingsOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsSettingsOutput {
	return o
}

func (o MarkdownPartMetadataSettingsSettingsOutput) ToMarkdownPartMetadataSettingsSettingsPtrOutput() MarkdownPartMetadataSettingsSettingsPtrOutput {
	return o.ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(context.Background())
}

func (o MarkdownPartMetadataSettingsSettingsOutput) ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MarkdownPartMetadataSettingsSettings) *MarkdownPartMetadataSettingsSettings {
		return &v
	}).(MarkdownPartMetadataSettingsSettingsPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataSettingsSettings) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsOutput) MarkdownSource() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataSettingsSettings) *int { return v.MarkdownSource }).(pulumi.IntPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsOutput) MarkdownUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataSettingsSettings) *string { return v.MarkdownUri }).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsOutput) Subtitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataSettingsSettings) *string { return v.Subtitle }).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MarkdownPartMetadataSettingsSettings) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type MarkdownPartMetadataSettingsSettingsPtrOutput struct{ *pulumi.OutputState }

func (MarkdownPartMetadataSettingsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkdownPartMetadataSettingsSettings)(nil)).Elem()
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) ToMarkdownPartMetadataSettingsSettingsPtrOutput() MarkdownPartMetadataSettingsSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) ToMarkdownPartMetadataSettingsSettingsPtrOutputWithContext(ctx context.Context) MarkdownPartMetadataSettingsSettingsPtrOutput {
	return o
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) Elem() MarkdownPartMetadataSettingsSettingsOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettingsSettings) MarkdownPartMetadataSettingsSettings {
		if v != nil {
			return *v
		}
		var ret MarkdownPartMetadataSettingsSettings
		return ret
	}).(MarkdownPartMetadataSettingsSettingsOutput)
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) MarkdownSource() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettingsSettings) *int {
		if v == nil {
			return nil
		}
		return v.MarkdownSource
	}).(pulumi.IntPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) MarkdownUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.MarkdownUri
	}).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) Subtitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Subtitle
	}).(pulumi.StringPtrOutput)
}

func (o MarkdownPartMetadataSettingsSettingsPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkdownPartMetadataSettingsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type StorageProfile struct {
	DiskSizeInGB             *int    `pulumi:"diskSizeInGB"`
	FileShareName            *string `pulumi:"fileShareName"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	DiskSizeInGB             pulumi.IntPtrInput    `pulumi:"diskSizeInGB"`
	FileShareName            pulumi.StringPtrInput `pulumi:"fileShareName"`
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

func (o StorageProfileOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

func (o StorageProfileOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfilePtrOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.FileShareName
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfilePtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageProfileResponse struct {
	DiskSizeInGB             *int    `pulumi:"diskSizeInGB"`
	FileShareName            *string `pulumi:"fileShareName"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}





type StorageProfileResponseInput interface {
	pulumi.Input

	ToStorageProfileResponseOutput() StorageProfileResponseOutput
	ToStorageProfileResponseOutputWithContext(context.Context) StorageProfileResponseOutput
}

type StorageProfileResponseArgs struct {
	DiskSizeInGB             pulumi.IntPtrInput    `pulumi:"diskSizeInGB"`
	FileShareName            pulumi.StringPtrInput `pulumi:"fileShareName"`
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
}

func (StorageProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return i.ToStorageProfileResponseOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput)
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput).ToStorageProfileResponsePtrOutputWithContext(ctx)
}









type StorageProfileResponsePtrInput interface {
	pulumi.Input

	ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput
	ToStorageProfileResponsePtrOutputWithContext(context.Context) StorageProfileResponsePtrOutput
}

type storageProfileResponsePtrType StorageProfileResponseArgs

func StorageProfileResponsePtr(v *StorageProfileResponseArgs) StorageProfileResponsePtrInput {
	return (*storageProfileResponsePtrType)(v)
}

func (*storageProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponsePtrOutput)
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfileResponse) *StorageProfileResponse {
		return &v
	}).(StorageProfileResponsePtrOutput)
}

func (o StorageProfileResponseOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponseOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponseOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponsePtrOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.FileShareName
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type TerminalSettings struct {
	FontSize  *string `pulumi:"fontSize"`
	FontStyle *string `pulumi:"fontStyle"`
}





type TerminalSettingsInput interface {
	pulumi.Input

	ToTerminalSettingsOutput() TerminalSettingsOutput
	ToTerminalSettingsOutputWithContext(context.Context) TerminalSettingsOutput
}

type TerminalSettingsArgs struct {
	FontSize  pulumi.StringPtrInput `pulumi:"fontSize"`
	FontStyle pulumi.StringPtrInput `pulumi:"fontStyle"`
}

func (TerminalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettings)(nil)).Elem()
}

func (i TerminalSettingsArgs) ToTerminalSettingsOutput() TerminalSettingsOutput {
	return i.ToTerminalSettingsOutputWithContext(context.Background())
}

func (i TerminalSettingsArgs) ToTerminalSettingsOutputWithContext(ctx context.Context) TerminalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsOutput)
}

func (i TerminalSettingsArgs) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return i.ToTerminalSettingsPtrOutputWithContext(context.Background())
}

func (i TerminalSettingsArgs) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsOutput).ToTerminalSettingsPtrOutputWithContext(ctx)
}









type TerminalSettingsPtrInput interface {
	pulumi.Input

	ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput
	ToTerminalSettingsPtrOutputWithContext(context.Context) TerminalSettingsPtrOutput
}

type terminalSettingsPtrType TerminalSettingsArgs

func TerminalSettingsPtr(v *TerminalSettingsArgs) TerminalSettingsPtrInput {
	return (*terminalSettingsPtrType)(v)
}

func (*terminalSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettings)(nil)).Elem()
}

func (i *terminalSettingsPtrType) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return i.ToTerminalSettingsPtrOutputWithContext(context.Background())
}

func (i *terminalSettingsPtrType) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsPtrOutput)
}

type TerminalSettingsOutput struct{ *pulumi.OutputState }

func (TerminalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettings)(nil)).Elem()
}

func (o TerminalSettingsOutput) ToTerminalSettingsOutput() TerminalSettingsOutput {
	return o
}

func (o TerminalSettingsOutput) ToTerminalSettingsOutputWithContext(ctx context.Context) TerminalSettingsOutput {
	return o
}

func (o TerminalSettingsOutput) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return o.ToTerminalSettingsPtrOutputWithContext(context.Background())
}

func (o TerminalSettingsOutput) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TerminalSettings) *TerminalSettings {
		return &v
	}).(TerminalSettingsPtrOutput)
}

func (o TerminalSettingsOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

type TerminalSettingsPtrOutput struct{ *pulumi.OutputState }

func (TerminalSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettings)(nil)).Elem()
}

func (o TerminalSettingsPtrOutput) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return o
}

func (o TerminalSettingsPtrOutput) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return o
}

func (o TerminalSettingsPtrOutput) Elem() TerminalSettingsOutput {
	return o.ApplyT(func(v *TerminalSettings) TerminalSettings {
		if v != nil {
			return *v
		}
		var ret TerminalSettings
		return ret
	}).(TerminalSettingsOutput)
}

func (o TerminalSettingsPtrOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettings) *string {
		if v == nil {
			return nil
		}
		return v.FontSize
	}).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsPtrOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettings) *string {
		if v == nil {
			return nil
		}
		return v.FontStyle
	}).(pulumi.StringPtrOutput)
}

type TerminalSettingsResponse struct {
	FontSize  *string `pulumi:"fontSize"`
	FontStyle *string `pulumi:"fontStyle"`
}





type TerminalSettingsResponseInput interface {
	pulumi.Input

	ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput
	ToTerminalSettingsResponseOutputWithContext(context.Context) TerminalSettingsResponseOutput
}

type TerminalSettingsResponseArgs struct {
	FontSize  pulumi.StringPtrInput `pulumi:"fontSize"`
	FontStyle pulumi.StringPtrInput `pulumi:"fontStyle"`
}

func (TerminalSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettingsResponse)(nil)).Elem()
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput {
	return i.ToTerminalSettingsResponseOutputWithContext(context.Background())
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponseOutputWithContext(ctx context.Context) TerminalSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsResponseOutput)
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return i.ToTerminalSettingsResponsePtrOutputWithContext(context.Background())
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsResponseOutput).ToTerminalSettingsResponsePtrOutputWithContext(ctx)
}









type TerminalSettingsResponsePtrInput interface {
	pulumi.Input

	ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput
	ToTerminalSettingsResponsePtrOutputWithContext(context.Context) TerminalSettingsResponsePtrOutput
}

type terminalSettingsResponsePtrType TerminalSettingsResponseArgs

func TerminalSettingsResponsePtr(v *TerminalSettingsResponseArgs) TerminalSettingsResponsePtrInput {
	return (*terminalSettingsResponsePtrType)(v)
}

func (*terminalSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettingsResponse)(nil)).Elem()
}

func (i *terminalSettingsResponsePtrType) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return i.ToTerminalSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *terminalSettingsResponsePtrType) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsResponsePtrOutput)
}

type TerminalSettingsResponseOutput struct{ *pulumi.OutputState }

func (TerminalSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettingsResponse)(nil)).Elem()
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput {
	return o
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponseOutputWithContext(ctx context.Context) TerminalSettingsResponseOutput {
	return o
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return o.ToTerminalSettingsResponsePtrOutputWithContext(context.Background())
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TerminalSettingsResponse) *TerminalSettingsResponse {
		return &v
	}).(TerminalSettingsResponsePtrOutput)
}

func (o TerminalSettingsResponseOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsResponseOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

type TerminalSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (TerminalSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettingsResponse)(nil)).Elem()
}

func (o TerminalSettingsResponsePtrOutput) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return o
}

func (o TerminalSettingsResponsePtrOutput) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return o
}

func (o TerminalSettingsResponsePtrOutput) Elem() TerminalSettingsResponseOutput {
	return o.ApplyT(func(v *TerminalSettingsResponse) TerminalSettingsResponse {
		if v != nil {
			return *v
		}
		var ret TerminalSettingsResponse
		return ret
	}).(TerminalSettingsResponseOutput)
}

func (o TerminalSettingsResponsePtrOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FontSize
	}).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsResponsePtrOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FontStyle
	}).(pulumi.StringPtrOutput)
}

type UserProperties struct {
	PreferredLocation  string           `pulumi:"preferredLocation"`
	PreferredOsType    string           `pulumi:"preferredOsType"`
	PreferredShellType string           `pulumi:"preferredShellType"`
	StorageProfile     StorageProfile   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettings `pulumi:"terminalSettings"`
}





type UserPropertiesInput interface {
	pulumi.Input

	ToUserPropertiesOutput() UserPropertiesOutput
	ToUserPropertiesOutputWithContext(context.Context) UserPropertiesOutput
}

type UserPropertiesArgs struct {
	PreferredLocation  pulumi.StringInput    `pulumi:"preferredLocation"`
	PreferredOsType    pulumi.StringInput    `pulumi:"preferredOsType"`
	PreferredShellType pulumi.StringInput    `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileInput   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsInput `pulumi:"terminalSettings"`
}

func (UserPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProperties)(nil)).Elem()
}

func (i UserPropertiesArgs) ToUserPropertiesOutput() UserPropertiesOutput {
	return i.ToUserPropertiesOutputWithContext(context.Background())
}

func (i UserPropertiesArgs) ToUserPropertiesOutputWithContext(ctx context.Context) UserPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesOutput)
}

func (i UserPropertiesArgs) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return i.ToUserPropertiesPtrOutputWithContext(context.Background())
}

func (i UserPropertiesArgs) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesOutput).ToUserPropertiesPtrOutputWithContext(ctx)
}









type UserPropertiesPtrInput interface {
	pulumi.Input

	ToUserPropertiesPtrOutput() UserPropertiesPtrOutput
	ToUserPropertiesPtrOutputWithContext(context.Context) UserPropertiesPtrOutput
}

type userPropertiesPtrType UserPropertiesArgs

func UserPropertiesPtr(v *UserPropertiesArgs) UserPropertiesPtrInput {
	return (*userPropertiesPtrType)(v)
}

func (*userPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProperties)(nil)).Elem()
}

func (i *userPropertiesPtrType) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return i.ToUserPropertiesPtrOutputWithContext(context.Background())
}

func (i *userPropertiesPtrType) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesPtrOutput)
}

type UserPropertiesOutput struct{ *pulumi.OutputState }

func (UserPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProperties)(nil)).Elem()
}

func (o UserPropertiesOutput) ToUserPropertiesOutput() UserPropertiesOutput {
	return o
}

func (o UserPropertiesOutput) ToUserPropertiesOutputWithContext(ctx context.Context) UserPropertiesOutput {
	return o
}

func (o UserPropertiesOutput) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return o.ToUserPropertiesPtrOutputWithContext(context.Background())
}

func (o UserPropertiesOutput) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserProperties) *UserProperties {
		return &v
	}).(UserPropertiesPtrOutput)
}

func (o UserPropertiesOutput) PreferredLocation() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredLocation }).(pulumi.StringOutput)
}

func (o UserPropertiesOutput) PreferredOsType() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredOsType }).(pulumi.StringOutput)
}

func (o UserPropertiesOutput) PreferredShellType() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredShellType }).(pulumi.StringOutput)
}

func (o UserPropertiesOutput) StorageProfile() StorageProfileOutput {
	return o.ApplyT(func(v UserProperties) StorageProfile { return v.StorageProfile }).(StorageProfileOutput)
}

func (o UserPropertiesOutput) TerminalSettings() TerminalSettingsOutput {
	return o.ApplyT(func(v UserProperties) TerminalSettings { return v.TerminalSettings }).(TerminalSettingsOutput)
}

type UserPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProperties)(nil)).Elem()
}

func (o UserPropertiesPtrOutput) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return o
}

func (o UserPropertiesPtrOutput) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return o
}

func (o UserPropertiesPtrOutput) Elem() UserPropertiesOutput {
	return o.ApplyT(func(v *UserProperties) UserProperties {
		if v != nil {
			return *v
		}
		var ret UserProperties
		return ret
	}).(UserPropertiesOutput)
}

func (o UserPropertiesPtrOutput) PreferredLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredLocation
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesPtrOutput) PreferredOsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredOsType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesPtrOutput) PreferredShellType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShellType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesPtrOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v *UserProperties) *StorageProfile {
		if v == nil {
			return nil
		}
		return &v.StorageProfile
	}).(StorageProfilePtrOutput)
}

func (o UserPropertiesPtrOutput) TerminalSettings() TerminalSettingsPtrOutput {
	return o.ApplyT(func(v *UserProperties) *TerminalSettings {
		if v == nil {
			return nil
		}
		return &v.TerminalSettings
	}).(TerminalSettingsPtrOutput)
}

type UserPropertiesResponse struct {
	PreferredLocation  string                   `pulumi:"preferredLocation"`
	PreferredOsType    string                   `pulumi:"preferredOsType"`
	PreferredShellType string                   `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileResponse   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsResponse `pulumi:"terminalSettings"`
}





type UserPropertiesResponseInput interface {
	pulumi.Input

	ToUserPropertiesResponseOutput() UserPropertiesResponseOutput
	ToUserPropertiesResponseOutputWithContext(context.Context) UserPropertiesResponseOutput
}

type UserPropertiesResponseArgs struct {
	PreferredLocation  pulumi.StringInput            `pulumi:"preferredLocation"`
	PreferredOsType    pulumi.StringInput            `pulumi:"preferredOsType"`
	PreferredShellType pulumi.StringInput            `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileResponseInput   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsResponseInput `pulumi:"terminalSettings"`
}

func (UserPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPropertiesResponse)(nil)).Elem()
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponseOutput() UserPropertiesResponseOutput {
	return i.ToUserPropertiesResponseOutputWithContext(context.Background())
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponseOutputWithContext(ctx context.Context) UserPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesResponseOutput)
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return i.ToUserPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesResponseOutput).ToUserPropertiesResponsePtrOutputWithContext(ctx)
}









type UserPropertiesResponsePtrInput interface {
	pulumi.Input

	ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput
	ToUserPropertiesResponsePtrOutputWithContext(context.Context) UserPropertiesResponsePtrOutput
}

type userPropertiesResponsePtrType UserPropertiesResponseArgs

func UserPropertiesResponsePtr(v *UserPropertiesResponseArgs) UserPropertiesResponsePtrInput {
	return (*userPropertiesResponsePtrType)(v)
}

func (*userPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPropertiesResponse)(nil)).Elem()
}

func (i *userPropertiesResponsePtrType) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return i.ToUserPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *userPropertiesResponsePtrType) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesResponsePtrOutput)
}

type UserPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPropertiesResponse)(nil)).Elem()
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponseOutput() UserPropertiesResponseOutput {
	return o
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponseOutputWithContext(ctx context.Context) UserPropertiesResponseOutput {
	return o
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return o.ToUserPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserPropertiesResponse) *UserPropertiesResponse {
		return &v
	}).(UserPropertiesResponsePtrOutput)
}

func (o UserPropertiesResponseOutput) PreferredLocation() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredLocation }).(pulumi.StringOutput)
}

func (o UserPropertiesResponseOutput) PreferredOsType() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredOsType }).(pulumi.StringOutput)
}

func (o UserPropertiesResponseOutput) PreferredShellType() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredShellType }).(pulumi.StringOutput)
}

func (o UserPropertiesResponseOutput) StorageProfile() StorageProfileResponseOutput {
	return o.ApplyT(func(v UserPropertiesResponse) StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponseOutput)
}

func (o UserPropertiesResponseOutput) TerminalSettings() TerminalSettingsResponseOutput {
	return o.ApplyT(func(v UserPropertiesResponse) TerminalSettingsResponse { return v.TerminalSettings }).(TerminalSettingsResponseOutput)
}

type UserPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPropertiesResponse)(nil)).Elem()
}

func (o UserPropertiesResponsePtrOutput) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return o
}

func (o UserPropertiesResponsePtrOutput) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return o
}

func (o UserPropertiesResponsePtrOutput) Elem() UserPropertiesResponseOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) UserPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserPropertiesResponse
		return ret
	}).(UserPropertiesResponseOutput)
}

func (o UserPropertiesResponsePtrOutput) PreferredLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredLocation
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesResponsePtrOutput) PreferredOsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredOsType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesResponsePtrOutput) PreferredShellType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShellType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesResponsePtrOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *StorageProfileResponse {
		if v == nil {
			return nil
		}
		return &v.StorageProfile
	}).(StorageProfileResponsePtrOutput)
}

func (o UserPropertiesResponsePtrOutput) TerminalSettings() TerminalSettingsResponsePtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *TerminalSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.TerminalSettings
	}).(TerminalSettingsResponsePtrOutput)
}

type ViolationResponse struct {
	ErrorMessage string `pulumi:"errorMessage"`
	Id           string `pulumi:"id"`
	UserId       string `pulumi:"userId"`
}





type ViolationResponseInput interface {
	pulumi.Input

	ToViolationResponseOutput() ViolationResponseOutput
	ToViolationResponseOutputWithContext(context.Context) ViolationResponseOutput
}

type ViolationResponseArgs struct {
	ErrorMessage pulumi.StringInput `pulumi:"errorMessage"`
	Id           pulumi.StringInput `pulumi:"id"`
	UserId       pulumi.StringInput `pulumi:"userId"`
}

func (ViolationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ViolationResponse)(nil)).Elem()
}

func (i ViolationResponseArgs) ToViolationResponseOutput() ViolationResponseOutput {
	return i.ToViolationResponseOutputWithContext(context.Background())
}

func (i ViolationResponseArgs) ToViolationResponseOutputWithContext(ctx context.Context) ViolationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViolationResponseOutput)
}





type ViolationResponseArrayInput interface {
	pulumi.Input

	ToViolationResponseArrayOutput() ViolationResponseArrayOutput
	ToViolationResponseArrayOutputWithContext(context.Context) ViolationResponseArrayOutput
}

type ViolationResponseArray []ViolationResponseInput

func (ViolationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ViolationResponse)(nil)).Elem()
}

func (i ViolationResponseArray) ToViolationResponseArrayOutput() ViolationResponseArrayOutput {
	return i.ToViolationResponseArrayOutputWithContext(context.Background())
}

func (i ViolationResponseArray) ToViolationResponseArrayOutputWithContext(ctx context.Context) ViolationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViolationResponseArrayOutput)
}

type ViolationResponseOutput struct{ *pulumi.OutputState }

func (ViolationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ViolationResponse)(nil)).Elem()
}

func (o ViolationResponseOutput) ToViolationResponseOutput() ViolationResponseOutput {
	return o
}

func (o ViolationResponseOutput) ToViolationResponseOutputWithContext(ctx context.Context) ViolationResponseOutput {
	return o
}

func (o ViolationResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ViolationResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o ViolationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ViolationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ViolationResponseOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v ViolationResponse) string { return v.UserId }).(pulumi.StringOutput)
}

type ViolationResponseArrayOutput struct{ *pulumi.OutputState }

func (ViolationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ViolationResponse)(nil)).Elem()
}

func (o ViolationResponseArrayOutput) ToViolationResponseArrayOutput() ViolationResponseArrayOutput {
	return o
}

func (o ViolationResponseArrayOutput) ToViolationResponseArrayOutputWithContext(ctx context.Context) ViolationResponseArrayOutput {
	return o
}

func (o ViolationResponseArrayOutput) Index(i pulumi.IntInput) ViolationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ViolationResponse {
		return vs[0].([]ViolationResponse)[vs[1].(int)]
	}).(ViolationResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleCreatePropertiesInput)(nil)).Elem(), ConsoleCreatePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleCreatePropertiesPtrInput)(nil)).Elem(), ConsoleCreatePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsolePropertiesResponseInput)(nil)).Elem(), ConsolePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsolePropertiesResponsePtrInput)(nil)).Elem(), ConsolePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardLensInput)(nil)).Elem(), DashboardLensArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardLensArrayInput)(nil)).Elem(), DashboardLensArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardLensResponseInput)(nil)).Elem(), DashboardLensResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardLensResponseArrayInput)(nil)).Elem(), DashboardLensResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPartsInput)(nil)).Elem(), DashboardPartsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPartsArrayInput)(nil)).Elem(), DashboardPartsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPartsPositionInput)(nil)).Elem(), DashboardPartsPositionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPartsResponseInput)(nil)).Elem(), DashboardPartsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPartsResponseArrayInput)(nil)).Elem(), DashboardPartsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardPartsResponsePositionInput)(nil)).Elem(), DashboardPartsResponsePositionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataInput)(nil)).Elem(), MarkdownPartMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataPtrInput)(nil)).Elem(), MarkdownPartMetadataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataContentInput)(nil)).Elem(), MarkdownPartMetadataContentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataContentPtrInput)(nil)).Elem(), MarkdownPartMetadataContentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseInput)(nil)).Elem(), MarkdownPartMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponsePtrInput)(nil)).Elem(), MarkdownPartMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseContentInput)(nil)).Elem(), MarkdownPartMetadataResponseContentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseContentPtrInput)(nil)).Elem(), MarkdownPartMetadataResponseContentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseSettingsInput)(nil)).Elem(), MarkdownPartMetadataResponseSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseSettingsPtrInput)(nil)).Elem(), MarkdownPartMetadataResponseSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseSettingsSettingsInput)(nil)).Elem(), MarkdownPartMetadataResponseSettingsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataResponseSettingsSettingsPtrInput)(nil)).Elem(), MarkdownPartMetadataResponseSettingsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataSettingsInput)(nil)).Elem(), MarkdownPartMetadataSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataSettingsPtrInput)(nil)).Elem(), MarkdownPartMetadataSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataSettingsSettingsInput)(nil)).Elem(), MarkdownPartMetadataSettingsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MarkdownPartMetadataSettingsSettingsPtrInput)(nil)).Elem(), MarkdownPartMetadataSettingsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileInput)(nil)).Elem(), StorageProfileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfilePtrInput)(nil)).Elem(), StorageProfileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileResponseInput)(nil)).Elem(), StorageProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileResponsePtrInput)(nil)).Elem(), StorageProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsInput)(nil)).Elem(), TerminalSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsPtrInput)(nil)).Elem(), TerminalSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsResponseInput)(nil)).Elem(), TerminalSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsResponsePtrInput)(nil)).Elem(), TerminalSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesInput)(nil)).Elem(), UserPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesPtrInput)(nil)).Elem(), UserPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesResponseInput)(nil)).Elem(), UserPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesResponsePtrInput)(nil)).Elem(), UserPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViolationResponseInput)(nil)).Elem(), ViolationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViolationResponseArrayInput)(nil)).Elem(), ViolationResponseArray{})
	pulumi.RegisterOutputType(ConsoleCreatePropertiesOutput{})
	pulumi.RegisterOutputType(ConsoleCreatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConsolePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConsolePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DashboardLensOutput{})
	pulumi.RegisterOutputType(DashboardLensArrayOutput{})
	pulumi.RegisterOutputType(DashboardLensResponseOutput{})
	pulumi.RegisterOutputType(DashboardLensResponseArrayOutput{})
	pulumi.RegisterOutputType(DashboardPartsOutput{})
	pulumi.RegisterOutputType(DashboardPartsArrayOutput{})
	pulumi.RegisterOutputType(DashboardPartsPositionOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponseOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponseArrayOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponsePositionOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataPtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataContentOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataContentPtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseContentOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseContentPtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseSettingsOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseSettingsPtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseSettingsSettingsOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataResponseSettingsSettingsPtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataSettingsOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataSettingsPtrOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataSettingsSettingsOutput{})
	pulumi.RegisterOutputType(MarkdownPartMetadataSettingsSettingsPtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(TerminalSettingsOutput{})
	pulumi.RegisterOutputType(TerminalSettingsPtrOutput{})
	pulumi.RegisterOutputType(TerminalSettingsResponseOutput{})
	pulumi.RegisterOutputType(TerminalSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserPropertiesOutput{})
	pulumi.RegisterOutputType(UserPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ViolationResponseOutput{})
	pulumi.RegisterOutputType(ViolationResponseArrayOutput{})
}
