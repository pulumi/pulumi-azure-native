// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.V20230701Preview
{
    public static class GetHuntRelation
    {
        /// <summary>
        /// Gets a hunt relation
        /// </summary>
        public static Task<GetHuntRelationResult> InvokeAsync(GetHuntRelationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHuntRelationResult>("azure-native:securityinsights/v20230701preview:getHuntRelation", args ?? new GetHuntRelationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a hunt relation
        /// </summary>
        public static Output<GetHuntRelationResult> Invoke(GetHuntRelationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHuntRelationResult>("azure-native:securityinsights/v20230701preview:getHuntRelation", args ?? new GetHuntRelationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a hunt relation
        /// </summary>
        public static Output<GetHuntRelationResult> Invoke(GetHuntRelationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetHuntRelationResult>("azure-native:securityinsights/v20230701preview:getHuntRelation", args ?? new GetHuntRelationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHuntRelationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hunt id (GUID)
        /// </summary>
        [Input("huntId", required: true)]
        public string HuntId { get; set; } = null!;

        /// <summary>
        /// The hunt relation id (GUID)
        /// </summary>
        [Input("huntRelationId", required: true)]
        public string HuntRelationId { get; set; } = null!;

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

        public GetHuntRelationArgs()
        {
        }
        public static new GetHuntRelationArgs Empty => new GetHuntRelationArgs();
    }

    public sealed class GetHuntRelationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hunt id (GUID)
        /// </summary>
        [Input("huntId", required: true)]
        public Input<string> HuntId { get; set; } = null!;

        /// <summary>
        /// The hunt relation id (GUID)
        /// </summary>
        [Input("huntRelationId", required: true)]
        public Input<string> HuntRelationId { get; set; } = null!;

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

        public GetHuntRelationInvokeArgs()
        {
        }
        public static new GetHuntRelationInvokeArgs Empty => new GetHuntRelationInvokeArgs();
    }


    [OutputType]
    public sealed class GetHuntRelationResult
    {
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of labels relevant to this hunt
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the related resource
        /// </summary>
        public readonly string RelatedResourceId;
        /// <summary>
        /// The resource that the relation is related to
        /// </summary>
        public readonly string RelatedResourceKind;
        /// <summary>
        /// The name of the related resource
        /// </summary>
        public readonly string RelatedResourceName;
        /// <summary>
        /// The type of the hunt relation
        /// </summary>
        public readonly string RelationType;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetHuntRelationResult(
            string? etag,

            string id,

            ImmutableArray<string> labels,

            string name,

            string relatedResourceId,

            string relatedResourceKind,

            string relatedResourceName,

            string relationType,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Etag = etag;
            Id = id;
            Labels = labels;
            Name = name;
            RelatedResourceId = relatedResourceId;
            RelatedResourceKind = relatedResourceKind;
            RelatedResourceName = relatedResourceName;
            RelationType = relationType;
            SystemData = systemData;
            Type = type;
        }
    }
}
