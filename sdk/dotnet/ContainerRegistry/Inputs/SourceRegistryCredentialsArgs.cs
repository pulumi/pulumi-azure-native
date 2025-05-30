// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Inputs
{

    /// <summary>
    /// Describes the credential parameters for accessing the source registry.
    /// </summary>
    public sealed class SourceRegistryCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication mode which determines the source registry login scope. The credentials for the source registry
        /// will be generated using the given scope. These credentials will be used to login to
        /// the source registry during the run.
        /// </summary>
        [Input("loginMode")]
        public InputUnion<string, Pulumi.AzureNative.ContainerRegistry.SourceRegistryLoginMode>? LoginMode { get; set; }

        public SourceRegistryCredentialsArgs()
        {
        }
        public static new SourceRegistryCredentialsArgs Empty => new SourceRegistryCredentialsArgs();
    }
}
