// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices
{
    public static class GetReplicationStorageClassificationMapping
    {
        /// <summary>
        /// Gets the details of the specified storage classification mapping.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetReplicationStorageClassificationMappingResult> InvokeAsync(GetReplicationStorageClassificationMappingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReplicationStorageClassificationMappingResult>("azure-native:recoveryservices:getReplicationStorageClassificationMapping", args ?? new GetReplicationStorageClassificationMappingArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the specified storage classification mapping.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetReplicationStorageClassificationMappingResult> Invoke(GetReplicationStorageClassificationMappingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplicationStorageClassificationMappingResult>("azure-native:recoveryservices:getReplicationStorageClassificationMapping", args ?? new GetReplicationStorageClassificationMappingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the specified storage classification mapping.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetReplicationStorageClassificationMappingResult> Invoke(GetReplicationStorageClassificationMappingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplicationStorageClassificationMappingResult>("azure-native:recoveryservices:getReplicationStorageClassificationMapping", args ?? new GetReplicationStorageClassificationMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplicationStorageClassificationMappingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public string FabricName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the recovery services vault is present.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the recovery services vault.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        /// <summary>
        /// Storage classification mapping name.
        /// </summary>
        [Input("storageClassificationMappingName", required: true)]
        public string StorageClassificationMappingName { get; set; } = null!;

        /// <summary>
        /// Storage classification name.
        /// </summary>
        [Input("storageClassificationName", required: true)]
        public string StorageClassificationName { get; set; } = null!;

        public GetReplicationStorageClassificationMappingArgs()
        {
        }
        public static new GetReplicationStorageClassificationMappingArgs Empty => new GetReplicationStorageClassificationMappingArgs();
    }

    public sealed class GetReplicationStorageClassificationMappingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public Input<string> FabricName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the recovery services vault is present.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the recovery services vault.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// Storage classification mapping name.
        /// </summary>
        [Input("storageClassificationMappingName", required: true)]
        public Input<string> StorageClassificationMappingName { get; set; } = null!;

        /// <summary>
        /// Storage classification name.
        /// </summary>
        [Input("storageClassificationName", required: true)]
        public Input<string> StorageClassificationName { get; set; } = null!;

        public GetReplicationStorageClassificationMappingInvokeArgs()
        {
        }
        public static new GetReplicationStorageClassificationMappingInvokeArgs Empty => new GetReplicationStorageClassificationMappingInvokeArgs();
    }


    [OutputType]
    public sealed class GetReplicationStorageClassificationMappingResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the storage mapping object.
        /// </summary>
        public readonly Outputs.StorageClassificationMappingPropertiesResponse Properties;
        /// <summary>
        /// Resource Type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetReplicationStorageClassificationMappingResult(
            string azureApiVersion,

            string id,

            string? location,

            string name,

            Outputs.StorageClassificationMappingPropertiesResponse properties,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
