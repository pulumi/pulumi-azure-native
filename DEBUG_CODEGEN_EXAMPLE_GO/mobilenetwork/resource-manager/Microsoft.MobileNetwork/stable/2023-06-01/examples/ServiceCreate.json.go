package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewService(ctx, "service", &mobilenetwork.ServiceArgs{
			Location:          pulumi.String("eastus"),
			MobileNetworkName: pulumi.String("testMobileNetwork"),
			PccRules: []mobilenetwork.PccRuleConfigurationArgs{
				{
					RuleName:       pulumi.String("default-rule"),
					RulePrecedence: pulumi.Int(255),
					RuleQosPolicy: {
						AllocationAndRetentionPriorityLevel: pulumi.Int(9),
						FiveQi:                              pulumi.Int(9),
						MaximumBitRate: {
							Downlink: pulumi.String("1 Gbps"),
							Uplink:   pulumi.String("500 Mbps"),
						},
						PreemptionCapability:    pulumi.String("NotPreempt"),
						PreemptionVulnerability: pulumi.String("Preemptable"),
					},
					ServiceDataFlowTemplates: mobilenetwork.ServiceDataFlowTemplateArray{
						{
							Direction: pulumi.String("Uplink"),
							Ports:     pulumi.StringArray{},
							Protocol: pulumi.StringArray{
								pulumi.String("ip"),
							},
							RemoteIpList: pulumi.StringArray{
								pulumi.String("10.3.4.0/24"),
							},
							TemplateName: pulumi.String("IP-to-server"),
						},
					},
					TrafficControl: pulumi.String("Enabled"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("TestService"),
			ServicePrecedence: pulumi.Int(255),
			ServiceQosPolicy: mobilenetwork.QosPolicyResponse{
				AllocationAndRetentionPriorityLevel: pulumi.Int(9),
				FiveQi:                              pulumi.Int(9),
				MaximumBitRate: &mobilenetwork.AmbrArgs{
					Downlink: pulumi.String("1 Gbps"),
					Uplink:   pulumi.String("500 Mbps"),
				},
				PreemptionCapability:    pulumi.String("NotPreempt"),
				PreemptionVulnerability: pulumi.String("Preemptable"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
