// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.V20230101Preview
{
    public static class GetCacheRule
    {
        /// <summary>
        /// Gets the properties of the specified cache rule resource.
        /// </summary>
        public static Task<GetCacheRuleResult> InvokeAsync(GetCacheRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCacheRuleResult>("azure-native:containerregistry/v20230101preview:getCacheRule", args ?? new GetCacheRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the specified cache rule resource.
        /// </summary>
        public static Output<GetCacheRuleResult> Invoke(GetCacheRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCacheRuleResult>("azure-native:containerregistry/v20230101preview:getCacheRule", args ?? new GetCacheRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the specified cache rule resource.
        /// </summary>
        public static Output<GetCacheRuleResult> Invoke(GetCacheRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCacheRuleResult>("azure-native:containerregistry/v20230101preview:getCacheRule", args ?? new GetCacheRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCacheRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cache rule.
        /// </summary>
        [Input("cacheRuleName", required: true)]
        public string CacheRuleName { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCacheRuleArgs()
        {
        }
        public static new GetCacheRuleArgs Empty => new GetCacheRuleArgs();
    }

    public sealed class GetCacheRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cache rule.
        /// </summary>
        [Input("cacheRuleName", required: true)]
        public Input<string> CacheRuleName { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetCacheRuleInvokeArgs()
        {
        }
        public static new GetCacheRuleInvokeArgs Empty => new GetCacheRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetCacheRuleResult
    {
        /// <summary>
        /// The creation date of the cache rule.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The ARM resource ID of the credential store which is associated with the cache rule.
        /// </summary>
        public readonly string? CredentialSetResourceId;
        /// <summary>
        /// The resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Source repository pulled from upstream.
        /// </summary>
        public readonly string? SourceRepository;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Target repository specified in docker pull command.
        /// Eg: docker pull myregistry.azurecr.io/{targetRepository}:{tag}
        /// </summary>
        public readonly string? TargetRepository;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCacheRuleResult(
            string creationDate,

            string? credentialSetResourceId,

            string id,

            string name,

            string provisioningState,

            string? sourceRepository,

            Outputs.SystemDataResponse systemData,

            string? targetRepository,

            string type)
        {
            CreationDate = creationDate;
            CredentialSetResourceId = credentialSetResourceId;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            SourceRepository = sourceRepository;
            SystemData = systemData;
            TargetRepository = targetRepository;
            Type = type;
        }
    }
}
