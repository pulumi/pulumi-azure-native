// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor
{
    public static class GetDataCollectionRuleAssociation
    {
        /// <summary>
        /// Definition of generic ARM proxy resource.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// </summary>
        public static Task<GetDataCollectionRuleAssociationResult> InvokeAsync(GetDataCollectionRuleAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataCollectionRuleAssociationResult>("azure-native:monitor:getDataCollectionRuleAssociation", args ?? new GetDataCollectionRuleAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of generic ARM proxy resource.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// </summary>
        public static Output<GetDataCollectionRuleAssociationResult> Invoke(GetDataCollectionRuleAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataCollectionRuleAssociationResult>("azure-native:monitor:getDataCollectionRuleAssociation", args ?? new GetDataCollectionRuleAssociationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of generic ARM proxy resource.
        /// 
        /// Uses Azure REST API version 2022-06-01.
        /// </summary>
        public static Output<GetDataCollectionRuleAssociationResult> Invoke(GetDataCollectionRuleAssociationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataCollectionRuleAssociationResult>("azure-native:monitor:getDataCollectionRuleAssociation", args ?? new GetDataCollectionRuleAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataCollectionRuleAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the association. The name is case insensitive.
        /// </summary>
        [Input("associationName", required: true)]
        public string AssociationName { get; set; } = null!;

        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetDataCollectionRuleAssociationArgs()
        {
        }
        public static new GetDataCollectionRuleAssociationArgs Empty => new GetDataCollectionRuleAssociationArgs();
    }

    public sealed class GetDataCollectionRuleAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the association. The name is case insensitive.
        /// </summary>
        [Input("associationName", required: true)]
        public Input<string> AssociationName { get; set; } = null!;

        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetDataCollectionRuleAssociationInvokeArgs()
        {
        }
        public static new GetDataCollectionRuleAssociationInvokeArgs Empty => new GetDataCollectionRuleAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataCollectionRuleAssociationResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The resource ID of the data collection endpoint that is to be associated.
        /// </summary>
        public readonly string? DataCollectionEndpointId;
        /// <summary>
        /// The resource ID of the data collection rule that is to be associated.
        /// </summary>
        public readonly string? DataCollectionRuleId;
        /// <summary>
        /// Description of the association.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Resource entity tag (ETag).
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified ID of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Metadata about the resource
        /// </summary>
        public readonly Outputs.DataCollectionRuleAssociationResponseMetadata Metadata;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData SystemData;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataCollectionRuleAssociationResult(
            string azureApiVersion,

            string? dataCollectionEndpointId,

            string? dataCollectionRuleId,

            string? description,

            string etag,

            string id,

            Outputs.DataCollectionRuleAssociationResponseMetadata metadata,

            string name,

            string provisioningState,

            Outputs.DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DataCollectionEndpointId = dataCollectionEndpointId;
            DataCollectionRuleId = dataCollectionRuleId;
            Description = description;
            Etag = etag;
            Id = id;
            Metadata = metadata;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
