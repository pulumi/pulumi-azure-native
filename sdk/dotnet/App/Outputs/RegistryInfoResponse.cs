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
    /// Container App registry information.
    /// </summary>
    [OutputType]
    public sealed class RegistryInfoResponse
    {
        /// <summary>
        /// registry server Url.
        /// </summary>
        public readonly string? RegistryUrl;
        /// <summary>
        /// registry username.
        /// </summary>
        public readonly string? RegistryUserName;

        [OutputConstructor]
        private RegistryInfoResponse(
            string? registryUrl,

            string? registryUserName)
        {
            RegistryUrl = registryUrl;
            RegistryUserName = registryUserName;
        }
    }
}
