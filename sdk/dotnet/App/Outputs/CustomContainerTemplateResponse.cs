// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Custom container configuration.
    /// </summary>
    [OutputType]
    public sealed class CustomContainerTemplateResponse
    {
        /// <summary>
        /// List of container definitions for the sessions of the session pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.SessionContainerResponse> Containers;
        /// <summary>
        /// Session pool ingress configuration.
        /// </summary>
        public readonly Outputs.SessionIngressResponse? Ingress;
        /// <summary>
        /// Private container registry credentials for containers used by the sessions of the session pool.
        /// </summary>
        public readonly Outputs.SessionRegistryCredentialsResponse? RegistryCredentials;

        [OutputConstructor]
        private CustomContainerTemplateResponse(
            ImmutableArray<Outputs.SessionContainerResponse> containers,

            Outputs.SessionIngressResponse? ingress,

            Outputs.SessionRegistryCredentialsResponse? registryCredentials)
        {
            Containers = containers;
            Ingress = ingress;
            RegistryCredentials = registryCredentials;
        }
    }
}
