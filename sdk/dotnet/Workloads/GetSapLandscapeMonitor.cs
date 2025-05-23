// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads
{
    public static class GetSapLandscapeMonitor
    {
        /// <summary>
        /// Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group, and resource name.
        /// 
        /// Uses Azure REST API version 2024-02-01-preview.
        /// 
        /// Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetSapLandscapeMonitorResult> InvokeAsync(GetSapLandscapeMonitorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSapLandscapeMonitorResult>("azure-native:workloads:getSapLandscapeMonitor", args ?? new GetSapLandscapeMonitorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group, and resource name.
        /// 
        /// Uses Azure REST API version 2024-02-01-preview.
        /// 
        /// Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSapLandscapeMonitorResult> Invoke(GetSapLandscapeMonitorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSapLandscapeMonitorResult>("azure-native:workloads:getSapLandscapeMonitor", args ?? new GetSapLandscapeMonitorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group, and resource name.
        /// 
        /// Uses Azure REST API version 2024-02-01-preview.
        /// 
        /// Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetSapLandscapeMonitorResult> Invoke(GetSapLandscapeMonitorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSapLandscapeMonitorResult>("azure-native:workloads:getSapLandscapeMonitor", args ?? new GetSapLandscapeMonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSapLandscapeMonitorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the SAP monitor resource.
        /// </summary>
        [Input("monitorName", required: true)]
        public string MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSapLandscapeMonitorArgs()
        {
        }
        public static new GetSapLandscapeMonitorArgs Empty => new GetSapLandscapeMonitorArgs();
    }

    public sealed class GetSapLandscapeMonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the SAP monitor resource.
        /// </summary>
        [Input("monitorName", required: true)]
        public Input<string> MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetSapLandscapeMonitorInvokeArgs()
        {
        }
        public static new GetSapLandscapeMonitorInvokeArgs Empty => new GetSapLandscapeMonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetSapLandscapeMonitorResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Gets or sets the SID groupings by landscape and Environment.
        /// </summary>
        public readonly Outputs.SapLandscapeMonitorPropertiesGroupingResponse? Grouping;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of provisioning of the SAP monitor.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Gets or sets the list Top Metric Thresholds for SAP Landscape Monitor Dashboard
        /// </summary>
        public readonly ImmutableArray<Outputs.SapLandscapeMonitorMetricThresholdsResponse> TopMetricsThresholds;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSapLandscapeMonitorResult(
            string azureApiVersion,

            Outputs.SapLandscapeMonitorPropertiesGroupingResponse? grouping,

            string id,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableArray<Outputs.SapLandscapeMonitorMetricThresholdsResponse> topMetricsThresholds,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Grouping = grouping;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            TopMetricsThresholds = topMetricsThresholds;
            Type = type;
        }
    }
}
