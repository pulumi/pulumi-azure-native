// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp
{
    public static class ListVolumeQuotaReport
    {
        /// <summary>
        /// Returns report of quotas for the volume
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-05-01-preview, 2024-07-01-preview, 2024-09-01-preview.
        /// </summary>
        public static Task<ListVolumeQuotaReportResult> InvokeAsync(ListVolumeQuotaReportArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListVolumeQuotaReportResult>("azure-native:netapp:listVolumeQuotaReport", args ?? new ListVolumeQuotaReportArgs(), options.WithDefaults());

        /// <summary>
        /// Returns report of quotas for the volume
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-05-01-preview, 2024-07-01-preview, 2024-09-01-preview.
        /// </summary>
        public static Output<ListVolumeQuotaReportResult> Invoke(ListVolumeQuotaReportInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListVolumeQuotaReportResult>("azure-native:netapp:listVolumeQuotaReport", args ?? new ListVolumeQuotaReportInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns report of quotas for the volume
        /// 
        /// Uses Azure REST API version 2024-03-01-preview.
        /// 
        /// Other available API versions: 2024-05-01-preview, 2024-07-01-preview, 2024-09-01-preview.
        /// </summary>
        public static Output<ListVolumeQuotaReportResult> Invoke(ListVolumeQuotaReportInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListVolumeQuotaReportResult>("azure-native:netapp:listVolumeQuotaReport", args ?? new ListVolumeQuotaReportInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListVolumeQuotaReportArgs : global::Pulumi.InvokeArgs
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

        public ListVolumeQuotaReportArgs()
        {
        }
        public static new ListVolumeQuotaReportArgs Empty => new ListVolumeQuotaReportArgs();
    }

    public sealed class ListVolumeQuotaReportInvokeArgs : global::Pulumi.InvokeArgs
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

        public ListVolumeQuotaReportInvokeArgs()
        {
        }
        public static new ListVolumeQuotaReportInvokeArgs Empty => new ListVolumeQuotaReportInvokeArgs();
    }


    [OutputType]
    public sealed class ListVolumeQuotaReportResult
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
        private ListVolumeQuotaReportResult(
            string? nextLink,

            ImmutableArray<Outputs.QuotaReportResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
