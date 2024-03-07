package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := costmanagement.NewCostAllocationRule(ctx, "costAllocationRule", &costmanagement.CostAllocationRuleArgs{
BillingAccountId: pulumi.String("100"),
Properties: costmanagement.CostAllocationRulePropertiesResponse{
Description: pulumi.String("This is a testRule"),
Details: interface{}{
SourceResources: costmanagement.SourceCostAllocationResourceArray{
&costmanagement.SourceCostAllocationResourceArgs{
Name: pulumi.String("category"),
ResourceType: pulumi.String("Tag"),
Values: pulumi.StringArray{
pulumi.String("devops"),
},
},
},
TargetResources: costmanagement.TargetCostAllocationResourceArray{
interface{}{
Name: pulumi.String("ResourceGroupName"),
PolicyType: pulumi.String("FixedProportion"),
ResourceType: pulumi.String("Dimension"),
Values: costmanagement.CostAllocationProportionArray{
&costmanagement.CostAllocationProportionArgs{
Name: pulumi.String("destinationRG"),
Percentage: pulumi.Float64(33.33),
},
&costmanagement.CostAllocationProportionArgs{
Name: pulumi.String("destinationRG2"),
Percentage: pulumi.Float64(33.33),
},
&costmanagement.CostAllocationProportionArgs{
Name: pulumi.String("destinationRG3"),
Percentage: pulumi.Float64(33.34),
},
},
},
},
},
Status: pulumi.String("Active"),
},
RuleName: pulumi.String("testRule"),
})
if err != nil {
return err
}
return nil
})
}
