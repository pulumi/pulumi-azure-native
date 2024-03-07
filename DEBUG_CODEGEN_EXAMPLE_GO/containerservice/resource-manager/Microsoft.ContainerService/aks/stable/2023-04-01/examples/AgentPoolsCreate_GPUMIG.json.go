package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
			AgentPoolName:      pulumi.String("agentpool1"),
			Count:              pulumi.Int(3),
			GpuInstanceProfile: pulumi.String("MIG2g"),
			KubeletConfig: &containerservice.KubeletConfigArgs{
				AllowedUnsafeSysctls: pulumi.StringArray{
					pulumi.String("kernel.msg*"),
					pulumi.String("net.core.somaxconn"),
				},
				CpuCfsQuota:           pulumi.Bool(true),
				CpuCfsQuotaPeriod:     pulumi.String("200ms"),
				CpuManagerPolicy:      pulumi.String("static"),
				FailSwapOn:            pulumi.Bool(false),
				ImageGcHighThreshold:  pulumi.Int(90),
				ImageGcLowThreshold:   pulumi.Int(70),
				TopologyManagerPolicy: pulumi.String("best-effort"),
			},
			LinuxOSConfig: &containerservice.LinuxOSConfigArgs{
				SwapFileSizeMB: pulumi.Int(1500),
				Sysctls: &containerservice.SysctlConfigArgs{
					KernelThreadsMax:        pulumi.Int(99999),
					NetCoreWmemDefault:      pulumi.Int(12345),
					NetIpv4IpLocalPortRange: pulumi.String("20000 60000"),
					NetIpv4TcpTwReuse:       pulumi.Bool(true),
				},
				TransparentHugePageDefrag:  pulumi.String("madvise"),
				TransparentHugePageEnabled: pulumi.String("always"),
			},
			OrchestratorVersion: pulumi.String(""),
			OsType:              pulumi.String("Linux"),
			ResourceGroupName:   pulumi.String("rg1"),
			ResourceName:        pulumi.String("clustername1"),
			VmSize:              pulumi.String("Standard_ND96asr_v4"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
