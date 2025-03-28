// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OffAzure
{
    public static class ListHypervSitesControllerHealthSummary
    {
        /// <summary>
        /// Method to get site health summary.
        /// 
        /// Uses Azure REST API version 2023-06-06.
        /// 
        /// Other available API versions: 2023-10-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Task<ListHypervSitesControllerHealthSummaryResult> InvokeAsync(ListHypervSitesControllerHealthSummaryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListHypervSitesControllerHealthSummaryResult>("azure-native:offazure:listHypervSitesControllerHealthSummary", args ?? new ListHypervSitesControllerHealthSummaryArgs(), options.WithDefaults());

        /// <summary>
        /// Method to get site health summary.
        /// 
        /// Uses Azure REST API version 2023-06-06.
        /// 
        /// Other available API versions: 2023-10-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<ListHypervSitesControllerHealthSummaryResult> Invoke(ListHypervSitesControllerHealthSummaryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListHypervSitesControllerHealthSummaryResult>("azure-native:offazure:listHypervSitesControllerHealthSummary", args ?? new ListHypervSitesControllerHealthSummaryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Method to get site health summary.
        /// 
        /// Uses Azure REST API version 2023-06-06.
        /// 
        /// Other available API versions: 2023-10-01-preview, 2024-05-01-preview.
        /// </summary>
        public static Output<ListHypervSitesControllerHealthSummaryResult> Invoke(ListHypervSitesControllerHealthSummaryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListHypervSitesControllerHealthSummaryResult>("azure-native:offazure:listHypervSitesControllerHealthSummary", args ?? new ListHypervSitesControllerHealthSummaryInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListHypervSitesControllerHealthSummaryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Site name
        /// </summary>
        [Input("siteName", required: true)]
        public string SiteName { get; set; } = null!;

        public ListHypervSitesControllerHealthSummaryArgs()
        {
        }
        public static new ListHypervSitesControllerHealthSummaryArgs Empty => new ListHypervSitesControllerHealthSummaryArgs();
    }

    public sealed class ListHypervSitesControllerHealthSummaryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Site name
        /// </summary>
        [Input("siteName", required: true)]
        public Input<string> SiteName { get; set; } = null!;

        public ListHypervSitesControllerHealthSummaryInvokeArgs()
        {
        }
        public static new ListHypervSitesControllerHealthSummaryInvokeArgs Empty => new ListHypervSitesControllerHealthSummaryInvokeArgs();
    }


    [OutputType]
    public sealed class ListHypervSitesControllerHealthSummaryResult
    {
        /// <summary>
        /// Gets the value of next link.
        /// </summary>
        public readonly string NextLink;
        /// <summary>
        /// Gets the list of SiteHealthSummary.
        /// </summary>
        public readonly ImmutableArray<Outputs.SiteHealthSummaryResponse> Value;

        [OutputConstructor]
        private ListHypervSitesControllerHealthSummaryResult(
            string nextLink,

            ImmutableArray<Outputs.SiteHealthSummaryResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
