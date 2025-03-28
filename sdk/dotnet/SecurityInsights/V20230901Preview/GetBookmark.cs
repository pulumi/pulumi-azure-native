// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230901Preview
{
    public static class GetBookmark
    {
        /// <summary>
        /// Gets a bookmark.
        /// </summary>
        public static Task<GetBookmarkResult> InvokeAsync(GetBookmarkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBookmarkResult>("azure-native:securityinsights/v20230901preview:getBookmark", args ?? new GetBookmarkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a bookmark.
        /// </summary>
        public static Output<GetBookmarkResult> Invoke(GetBookmarkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBookmarkResult>("azure-native:securityinsights/v20230901preview:getBookmark", args ?? new GetBookmarkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a bookmark.
        /// </summary>
        public static Output<GetBookmarkResult> Invoke(GetBookmarkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBookmarkResult>("azure-native:securityinsights/v20230901preview:getBookmark", args ?? new GetBookmarkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBookmarkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bookmark ID
        /// </summary>
        [Input("bookmarkId", required: true)]
        public string BookmarkId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetBookmarkArgs()
        {
        }
        public static new GetBookmarkArgs Empty => new GetBookmarkArgs();
    }

    public sealed class GetBookmarkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bookmark ID
        /// </summary>
        [Input("bookmarkId", required: true)]
        public Input<string> BookmarkId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetBookmarkInvokeArgs()
        {
        }
        public static new GetBookmarkInvokeArgs Empty => new GetBookmarkInvokeArgs();
    }


    [OutputType]
    public sealed class GetBookmarkResult
    {
        /// <summary>
        /// The time the bookmark was created
        /// </summary>
        public readonly string? Created;
        /// <summary>
        /// Describes a user that created the bookmark
        /// </summary>
        public readonly Outputs.UserInfoResponse? CreatedBy;
        /// <summary>
        /// The display name of the bookmark
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Describes the entity mappings of the bookmark
        /// </summary>
        public readonly ImmutableArray<Outputs.BookmarkEntityMappingsResponse> EntityMappings;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The bookmark event time
        /// </summary>
        public readonly string? EventTime;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Describes an incident that relates to bookmark
        /// </summary>
        public readonly Outputs.IncidentInfoResponse? IncidentInfo;
        /// <summary>
        /// List of labels relevant to this bookmark
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The notes of the bookmark
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// The query of the bookmark.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// The end time for the query
        /// </summary>
        public readonly string? QueryEndTime;
        /// <summary>
        /// The query result of the bookmark.
        /// </summary>
        public readonly string? QueryResult;
        /// <summary>
        /// The start time for the query
        /// </summary>
        public readonly string? QueryStartTime;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// A list of relevant mitre attacks
        /// </summary>
        public readonly ImmutableArray<string> Tactics;
        /// <summary>
        /// A list of relevant mitre techniques
        /// </summary>
        public readonly ImmutableArray<string> Techniques;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The last time the bookmark was updated
        /// </summary>
        public readonly string? Updated;
        /// <summary>
        /// Describes a user that updated the bookmark
        /// </summary>
        public readonly Outputs.UserInfoResponse? UpdatedBy;

        [OutputConstructor]
        private GetBookmarkResult(
            string? created,

            Outputs.UserInfoResponse? createdBy,

            string displayName,

            ImmutableArray<Outputs.BookmarkEntityMappingsResponse> entityMappings,

            string? etag,

            string? eventTime,

            string id,

            Outputs.IncidentInfoResponse? incidentInfo,

            ImmutableArray<string> labels,

            string name,

            string? notes,

            string query,

            string? queryEndTime,

            string? queryResult,

            string? queryStartTime,

            Outputs.SystemDataResponse systemData,

            ImmutableArray<string> tactics,

            ImmutableArray<string> techniques,

            string type,

            string? updated,

            Outputs.UserInfoResponse? updatedBy)
        {
            Created = created;
            CreatedBy = createdBy;
            DisplayName = displayName;
            EntityMappings = entityMappings;
            Etag = etag;
            EventTime = eventTime;
            Id = id;
            IncidentInfo = incidentInfo;
            Labels = labels;
            Name = name;
            Notes = notes;
            Query = query;
            QueryEndTime = queryEndTime;
            QueryResult = queryResult;
            QueryStartTime = queryStartTime;
            SystemData = systemData;
            Tactics = tactics;
            Techniques = techniques;
            Type = type;
            Updated = updated;
            UpdatedBy = updatedBy;
        }
    }
}
