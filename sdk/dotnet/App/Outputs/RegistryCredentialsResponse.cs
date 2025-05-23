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
    /// Container App Private Registry
    /// </summary>
    [OutputType]
    public sealed class RegistryCredentialsResponse
    {
        /// <summary>
        /// A Managed Identity to use to authenticate with Azure Container Registry. For user-assigned identities, use the full user-assigned identity Resource ID. For system-assigned identities, use 'system'
        /// </summary>
        public readonly string? Identity;
        /// <summary>
        /// The name of the Secret that contains the registry login password
        /// </summary>
        public readonly string? PasswordSecretRef;
        /// <summary>
        /// Container Registry Server
        /// </summary>
        public readonly string? Server;
        /// <summary>
        /// Container Registry Username
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private RegistryCredentialsResponse(
            string? identity,

            string? passwordSecretRef,

            string? server,

            string? username)
        {
            Identity = identity;
            PasswordSecretRef = passwordSecretRef;
            Server = server;
            Username = username;
        }
    }
}
