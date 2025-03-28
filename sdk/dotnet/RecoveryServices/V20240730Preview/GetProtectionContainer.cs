// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.V20240730Preview
{
    public static class GetProtectionContainer
    {
        /// <summary>
        /// Gets details of the specific container registered to your Recovery Services Vault.
        /// </summary>
        public static Task<GetProtectionContainerResult> InvokeAsync(GetProtectionContainerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProtectionContainerResult>("azure-native:recoveryservices/v20240730preview:getProtectionContainer", args ?? new GetProtectionContainerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of the specific container registered to your Recovery Services Vault.
        /// </summary>
        public static Output<GetProtectionContainerResult> Invoke(GetProtectionContainerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProtectionContainerResult>("azure-native:recoveryservices/v20240730preview:getProtectionContainer", args ?? new GetProtectionContainerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of the specific container registered to your Recovery Services Vault.
        /// </summary>
        public static Output<GetProtectionContainerResult> Invoke(GetProtectionContainerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProtectionContainerResult>("azure-native:recoveryservices/v20240730preview:getProtectionContainer", args ?? new GetProtectionContainerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProtectionContainerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the container whose details need to be fetched.
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// Name of the fabric where the container belongs.
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
        [Input("vaultName", required: true)]
        public string VaultName { get; set; } = null!;

        public GetProtectionContainerArgs()
        {
        }
        public static new GetProtectionContainerArgs Empty => new GetProtectionContainerArgs();
    }

    public sealed class GetProtectionContainerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the container whose details need to be fetched.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// Name of the fabric where the container belongs.
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
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        public GetProtectionContainerInvokeArgs()
        {
        }
        public static new GetProtectionContainerInvokeArgs Empty => new GetProtectionContainerInvokeArgs();
    }


    [OutputType]
    public sealed class GetProtectionContainerResult
    {
        /// <summary>
        /// Optional ETag.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// Resource Id represents the complete path to the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name associated with the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ProtectionContainerResource properties
        /// </summary>
        public readonly object Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProtectionContainerResult(
            string? eTag,

            string id,

            string? location,

            string name,

            object properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            ETag = eTag;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
