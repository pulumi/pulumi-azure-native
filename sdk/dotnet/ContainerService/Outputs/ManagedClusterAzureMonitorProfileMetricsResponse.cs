// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Metrics profile for the Azure Monitor managed service for Prometheus addon. Collect out-of-the-box Kubernetes infrastructure metrics to send to an Azure Monitor Workspace and configure additional scraping for custom targets. See aka.ms/AzureManagedPrometheus for an overview.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterAzureMonitorProfileMetricsResponse
    {
        /// <summary>
        /// Whether to enable or disable the Azure Managed Prometheus addon for Prometheus monitoring. See aka.ms/AzureManagedPrometheus-aks-enable for details on enabling and disabling.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Kube State Metrics profile for the Azure Managed Prometheus addon. These optional settings are for the kube-state-metrics pod that is deployed with the addon. See aka.ms/AzureManagedPrometheus-optional-parameters for details.
        /// </summary>
        public readonly Outputs.ManagedClusterAzureMonitorProfileKubeStateMetricsResponse? KubeStateMetrics;

        [OutputConstructor]
        private ManagedClusterAzureMonitorProfileMetricsResponse(
            bool enabled,

            Outputs.ManagedClusterAzureMonitorProfileKubeStateMetricsResponse? kubeStateMetrics)
        {
            Enabled = enabled;
            KubeStateMetrics = kubeStateMetrics;
        }
    }
}
