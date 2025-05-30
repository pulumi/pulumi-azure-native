// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AlertsManagement
{
    public static class ListIssueResources
    {
        /// <summary>
        /// List all resources in the issue - this method uses pagination to return all resources
        /// 
        /// Uses Azure REST API version 2025-03-01-preview.
        /// </summary>
        public static Task<ListIssueResourcesResult> InvokeAsync(ListIssueResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListIssueResourcesResult>("azure-native:alertsmanagement:listIssueResources", args ?? new ListIssueResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// List all resources in the issue - this method uses pagination to return all resources
        /// 
        /// Uses Azure REST API version 2025-03-01-preview.
        /// </summary>
        public static Output<ListIssueResourcesResult> Invoke(ListIssueResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListIssueResourcesResult>("azure-native:alertsmanagement:listIssueResources", args ?? new ListIssueResourcesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List all resources in the issue - this method uses pagination to return all resources
        /// 
        /// Uses Azure REST API version 2025-03-01-preview.
        /// </summary>
        public static Output<ListIssueResourcesResult> Invoke(ListIssueResourcesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListIssueResourcesResult>("azure-native:alertsmanagement:listIssueResources", args ?? new ListIssueResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListIssueResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The filter to apply on the operation. For example, to filter by relevance, use "$filter=relevance eq 'Relevant'"
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The name of the IssueResource
        /// </summary>
        [Input("issueName", required: true)]
        public string IssueName { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public ListIssueResourcesArgs()
        {
        }
        public static new ListIssueResourcesArgs Empty => new ListIssueResourcesArgs();
    }

    public sealed class ListIssueResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The filter to apply on the operation. For example, to filter by relevance, use "$filter=relevance eq 'Relevant'"
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the IssueResource
        /// </summary>
        [Input("issueName", required: true)]
        public Input<string> IssueName { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public ListIssueResourcesInvokeArgs()
        {
        }
        public static new ListIssueResourcesInvokeArgs Empty => new ListIssueResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class ListIssueResourcesResult
    {
        /// <summary>
        /// The link to the next page of items
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// The RelatedResource items on this page
        /// </summary>
        public readonly ImmutableArray<Outputs.RelatedResourceResponse> Value;

        [OutputConstructor]
        private ListIssueResourcesResult(
            string? nextLink,

            ImmutableArray<Outputs.RelatedResourceResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
