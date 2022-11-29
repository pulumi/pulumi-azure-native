


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RuleResultsPropertiesResponse struct {
	Results [][]string `pulumi:"results"`
}

type RuleResultsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RuleResultsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResultsPropertiesResponse)(nil)).Elem()
}

func (o RuleResultsPropertiesResponseOutput) ToRuleResultsPropertiesResponseOutput() RuleResultsPropertiesResponseOutput {
	return o
}

func (o RuleResultsPropertiesResponseOutput) ToRuleResultsPropertiesResponseOutputWithContext(ctx context.Context) RuleResultsPropertiesResponseOutput {
	return o
}

func (o RuleResultsPropertiesResponseOutput) Results() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v RuleResultsPropertiesResponse) [][]string { return v.Results }).(pulumi.StringArrayArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(RuleResultsPropertiesResponseOutput{})
}
