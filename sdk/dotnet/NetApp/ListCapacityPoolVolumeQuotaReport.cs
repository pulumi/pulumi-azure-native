// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp
{
    public static class ListCapacityPoolVolumeQuotaReport
    {
        /// <summary>
        /// Returns report of quotas for the volume
        /// 
        /// Uses Azure REST API version 2024-09-01-preview.
        /// 
        /// Other available API versions: 2024-03-01-preview, 2024-05-01-preview, 2024-07-01-preview, 2025-01-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListCapacityPoolVolumeQuotaReportResult> InvokeAsync(ListCapacityPoolVolumeQuotaReportArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListCapacityPoolVolumeQuotaReportResult>("azure-native:netapp:listCapacityPoolVolumeQuotaReport", args ?? new ListCapacityPoolVolumeQuotaReportArgs(), options.WithDefaults());

        /// <summary>
        /// Returns report of quotas for the volume
        /// 
        /// Uses Azure REST API version 2024-09-01-preview.
        /// 
        /// Other available API versions: 2024-03-01-preview, 2024-05-01-preview, 2024-07-01-preview, 2025-01-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListCapacityPoolVolumeQuotaReportResult> Invoke(ListCapacityPoolVolumeQuotaReportInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListCapacityPoolVolumeQuotaReportResult>("azure-native:netapp:listCapacityPoolVolumeQuotaReport", args ?? new ListCapacityPoolVolumeQuotaReportInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns report of quotas for the volume
        /// 
        /// Uses Azure REST API version 2024-09-01-preview.
        /// 
        /// Other available API versions: 2024-03-01-preview, 2024-05-01-preview, 2024-07-01-preview, 2025-01-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListCapacityPoolVolumeQuotaReportResult> Invoke(ListCapacityPoolVolumeQuotaReportInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListCapacityPoolVolumeQuotaReportResult>("azure-native:netapp:listCapacityPoolVolumeQuotaReport", args ?? new ListCapacityPoolVolumeQuotaReportInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListCapacityPoolVolumeQuotaReportArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capacity pool
        /// </summary>
        [Input("poolName", required: true)]
        public string PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public string VolumeName { get; set; } = null!;

        public ListCapacityPoolVolumeQuotaReportArgs()
        {
        }
        public static new ListCapacityPoolVolumeQuotaReportArgs Empty => new ListCapacityPoolVolumeQuotaReportArgs();
    }

    public sealed class ListCapacityPoolVolumeQuotaReportInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capacity pool
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public Input<string> VolumeName { get; set; } = null!;

        public ListCapacityPoolVolumeQuotaReportInvokeArgs()
        {
        }
        public static new ListCapacityPoolVolumeQuotaReportInvokeArgs Empty => new ListCapacityPoolVolumeQuotaReportInvokeArgs();
    }


    [OutputType]
    public sealed class ListCapacityPoolVolumeQuotaReportResult
    {
        /// <summary>
        /// URL to get the next set of results.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of volume quota report records
        /// </summary>
        public readonly ImmutableArray<Outputs.QuotaReportResponse> Value;

        [OutputConstructor]
        private ListCapacityPoolVolumeQuotaReportResult(
            string? nextLink,

            ImmutableArray<Outputs.QuotaReportResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
