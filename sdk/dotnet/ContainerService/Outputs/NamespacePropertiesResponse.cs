// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Properties of a namespace managed by ARM
    /// </summary>
    [OutputType]
    public sealed class NamespacePropertiesResponse
    {
        /// <summary>
        /// Action if Kubernetes namespace with same name already exists.
        /// </summary>
        public readonly string? AdoptionPolicy;
        /// <summary>
        /// The annotations of managed namespace.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// The default network policy enforced upon the namespace. Customers can have other Kubernetes network policy objects under the namespace. All the network policies will be enforced.
        /// </summary>
        public readonly Outputs.NetworkPoliciesResponse? DefaultNetworkPolicy;
        /// <summary>
        /// The default resource quota enforced upon the namespace. Customers can have other Kubernetes resource quota objects under the namespace. All the resource quotas will be enforced.
        /// </summary>
        public readonly Outputs.ResourceQuotaResponse? DefaultResourceQuota;
        /// <summary>
        /// Delete options of a namespace.
        /// </summary>
        public readonly string? DeletePolicy;
        /// <summary>
        /// The labels of managed namespace.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// The current provisioning state of the namespace.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private NamespacePropertiesResponse(
            string? adoptionPolicy,

            ImmutableDictionary<string, string>? annotations,

            Outputs.NetworkPoliciesResponse? defaultNetworkPolicy,

            Outputs.ResourceQuotaResponse? defaultResourceQuota,

            string? deletePolicy,

            ImmutableDictionary<string, string>? labels,

            string provisioningState)
        {
            AdoptionPolicy = adoptionPolicy;
            Annotations = annotations;
            DefaultNetworkPolicy = defaultNetworkPolicy;
            DefaultResourceQuota = defaultResourceQuota;
            DeletePolicy = deletePolicy;
            Labels = labels;
            ProvisioningState = provisioningState;
        }
    }
}
