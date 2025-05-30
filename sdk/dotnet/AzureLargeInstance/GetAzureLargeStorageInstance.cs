// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureLargeInstance
{
    public static class GetAzureLargeStorageInstance
    {
        /// <summary>
        /// Gets an Azure Large Storage instance for the specified subscription, resource
        /// group, and instance name.
        /// 
        /// Uses Azure REST API version 2024-08-01-preview.
        /// </summary>
        public static Task<GetAzureLargeStorageInstanceResult> InvokeAsync(GetAzureLargeStorageInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAzureLargeStorageInstanceResult>("azure-native:azurelargeinstance:getAzureLargeStorageInstance", args ?? new GetAzureLargeStorageInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an Azure Large Storage instance for the specified subscription, resource
        /// group, and instance name.
        /// 
        /// Uses Azure REST API version 2024-08-01-preview.
        /// </summary>
        public static Output<GetAzureLargeStorageInstanceResult> Invoke(GetAzureLargeStorageInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureLargeStorageInstanceResult>("azure-native:azurelargeinstance:getAzureLargeStorageInstance", args ?? new GetAzureLargeStorageInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an Azure Large Storage instance for the specified subscription, resource
        /// group, and instance name.
        /// 
        /// Uses Azure REST API version 2024-08-01-preview.
        /// </summary>
        public static Output<GetAzureLargeStorageInstanceResult> Invoke(GetAzureLargeStorageInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureLargeStorageInstanceResult>("azure-native:azurelargeinstance:getAzureLargeStorageInstance", args ?? new GetAzureLargeStorageInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAzureLargeStorageInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the AzureLargeStorageInstance.
        /// </summary>
        [Input("azureLargeStorageInstanceName", required: true)]
        public string AzureLargeStorageInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAzureLargeStorageInstanceArgs()
        {
        }
        public static new GetAzureLargeStorageInstanceArgs Empty => new GetAzureLargeStorageInstanceArgs();
    }

    public sealed class GetAzureLargeStorageInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the AzureLargeStorageInstance.
        /// </summary>
        [Input("azureLargeStorageInstanceName", required: true)]
        public Input<string> AzureLargeStorageInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAzureLargeStorageInstanceInvokeArgs()
        {
        }
        public static new GetAzureLargeStorageInstanceInvokeArgs Empty => new GetAzureLargeStorageInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetAzureLargeStorageInstanceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Specifies the AzureLargeStorageInstance unique ID.
        /// </summary>
        public readonly string? AzureLargeStorageInstanceUniqueIdentifier;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the storage properties for the AzureLargeStorage instance.
        /// </summary>
        public readonly Outputs.StoragePropertiesResponse? StorageProperties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAzureLargeStorageInstanceResult(
            string azureApiVersion,

            string? azureLargeStorageInstanceUniqueIdentifier,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Outputs.StoragePropertiesResponse? storageProperties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            AzureLargeStorageInstanceUniqueIdentifier = azureLargeStorageInstanceUniqueIdentifier;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            StorageProperties = storageProperties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
