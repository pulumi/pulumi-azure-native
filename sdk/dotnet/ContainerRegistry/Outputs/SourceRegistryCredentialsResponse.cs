// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// Describes the credential parameters for accessing the source registry.
    /// </summary>
    [OutputType]
    public sealed class SourceRegistryCredentialsResponse
    {
        /// <summary>
        /// The authentication mode which determines the source registry login scope. The credentials for the source registry
        /// will be generated using the given scope. These credentials will be used to login to
        /// the source registry during the run.
        /// </summary>
        public readonly string? LoginMode;

        [OutputConstructor]
        private SourceRegistryCredentialsResponse(string? loginMode)
        {
            LoginMode = loginMode;
        }
    }
}
