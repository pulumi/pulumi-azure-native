// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Outputs
{

    /// <summary>
    /// Kubernetes role storage resource
    /// </summary>
    [OutputType]
    public sealed class KubernetesRoleStorageResponse
    {
        /// <summary>
        /// Mount points of shares in role(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.MountPointMapResponse> Endpoints;
        /// <summary>
        /// Kubernetes storage class info.
        /// </summary>
        public readonly ImmutableArray<Outputs.KubernetesRoleStorageClassInfoResponse> StorageClasses;

        [OutputConstructor]
        private KubernetesRoleStorageResponse(
            ImmutableArray<Outputs.MountPointMapResponse> endpoints,

            ImmutableArray<Outputs.KubernetesRoleStorageClassInfoResponse> storageClasses)
        {
            Endpoints = endpoints;
            StorageClasses = storageClasses;
        }
    }
}
