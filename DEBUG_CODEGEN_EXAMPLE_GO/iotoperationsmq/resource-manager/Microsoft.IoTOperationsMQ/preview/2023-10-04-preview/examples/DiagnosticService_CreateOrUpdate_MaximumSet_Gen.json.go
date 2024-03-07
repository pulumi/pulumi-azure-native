package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsmq.NewDiagnosticService(ctx, "diagnosticService", &iotoperationsmq.DiagnosticServiceArgs{
			DataExportFrequencySeconds: pulumi.Int(26084),
			DiagnosticServiceName:      pulumi.String("73-1El3-1"),
			ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
				Name: pulumi.String("an"),
				Type: pulumi.String("CustomLocation"),
			},
			Image: &iotoperationsmq.ContainerImageArgs{
				PullPolicy:  pulumi.String("imfuzvqxgbdwliqnn"),
				PullSecrets: pulumi.String("klnqimxqsrdwhcqldjvdtsrs"),
				Repository:  pulumi.String("m"),
				Tag:         pulumi.String("jygfdiamhhm"),
			},
			Location:                         pulumi.String("sbhavoiabxjpuq"),
			LogFormat:                        pulumi.String("i"),
			LogLevel:                         pulumi.String("aomqhmpa"),
			MaxDataStorageSize:               pulumi.Float64(3757017229),
			MetricsPort:                      pulumi.Int(37109),
			MqName:                           pulumi.String("6RCAgs-XQ-Y2HsUF2"),
			OpenTelemetryTracesCollectorAddr: pulumi.String("ggqmprmjlmmkfdpb"),
			ResourceGroupName:                pulumi.String("rgiotoperationsmq"),
			StaleDataTimeoutSeconds:          pulumi.Int(51616),
			Tags:                             nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
