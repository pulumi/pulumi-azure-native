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
    /// See [disable AAD Pod Identity for a specific Pod/Application](https://azure.github.io/aad-pod-identity/docs/configure/application_exception/) for more details.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterPodIdentityExceptionResponse
    {
        /// <summary>
        /// The name of the pod identity exception.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The namespace of the pod identity exception.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// The pod labels to match.
        /// </summary>
        public readonly ImmutableDictionary<string, string> PodLabels;

        [OutputConstructor]
        private ManagedClusterPodIdentityExceptionResponse(
            string name,

            string @namespace,

            ImmutableDictionary<string, string> podLabels)
        {
            Name = name;
            Namespace = @namespace;
            PodLabels = podLabels;
        }
    }
}
