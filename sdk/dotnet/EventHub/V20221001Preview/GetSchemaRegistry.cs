// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventHub.V20221001Preview
{
    public static class GetSchemaRegistry
    {
        /// <summary>
        /// Gets the details of an EventHub schema group.
        /// </summary>
        public static Task<GetSchemaRegistryResult> InvokeAsync(GetSchemaRegistryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSchemaRegistryResult>("azure-native:eventhub/v20221001preview:getSchemaRegistry", args ?? new GetSchemaRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of an EventHub schema group.
        /// </summary>
        public static Output<GetSchemaRegistryResult> Invoke(GetSchemaRegistryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaRegistryResult>("azure-native:eventhub/v20221001preview:getSchemaRegistry", args ?? new GetSchemaRegistryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of an EventHub schema group.
        /// </summary>
        public static Output<GetSchemaRegistryResult> Invoke(GetSchemaRegistryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaRegistryResult>("azure-native:eventhub/v20221001preview:getSchemaRegistry", args ?? new GetSchemaRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSchemaRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Schema Group name 
        /// </summary>
        [Input("schemaGroupName", required: true)]
        public string SchemaGroupName { get; set; } = null!;

        public GetSchemaRegistryArgs()
        {
        }
        public static new GetSchemaRegistryArgs Empty => new GetSchemaRegistryArgs();
    }

    public sealed class GetSchemaRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Schema Group name 
        /// </summary>
        [Input("schemaGroupName", required: true)]
        public Input<string> SchemaGroupName { get; set; } = null!;

        public GetSchemaRegistryInvokeArgs()
        {
        }
        public static new GetSchemaRegistryInvokeArgs Empty => new GetSchemaRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetSchemaRegistryResult
    {
        /// <summary>
        /// Exact time the Schema Group was created.
        /// </summary>
        public readonly string CreatedAtUtc;
        /// <summary>
        /// The ETag value.
        /// </summary>
        public readonly string ETag;
        /// <summary>
        /// dictionary object for SchemaGroup group properties
        /// </summary>
        public readonly ImmutableDictionary<string, string>? GroupProperties;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        public readonly string? SchemaCompatibility;
        public readonly string? SchemaType;
        /// <summary>
        /// The system meta data relating to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Exact time the Schema Group was updated
        /// </summary>
        public readonly string UpdatedAtUtc;

        [OutputConstructor]
        private GetSchemaRegistryResult(
            string createdAtUtc,

            string eTag,

            ImmutableDictionary<string, string>? groupProperties,

            string id,

            string location,

            string name,

            string? schemaCompatibility,

            string? schemaType,

            Outputs.SystemDataResponse systemData,

            string type,

            string updatedAtUtc)
        {
            CreatedAtUtc = createdAtUtc;
            ETag = eTag;
            GroupProperties = groupProperties;
            Id = id;
            Location = location;
            Name = name;
            SchemaCompatibility = schemaCompatibility;
            SchemaType = schemaType;
            SystemData = systemData;
            Type = type;
            UpdatedAtUtc = updatedAtUtc;
        }
    }
}
