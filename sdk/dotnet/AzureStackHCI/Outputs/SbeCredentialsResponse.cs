// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// secrets used for solution builder extension (SBE) partner extensibility.
    /// </summary>
    [OutputType]
    public sealed class SbeCredentialsResponse
    {
        /// <summary>
        /// secret name expected for Enterprise Cloud Engine (ECE).
        /// </summary>
        public readonly string? EceSecretName;
        /// <summary>
        /// secret URI stored in keyvault.
        /// </summary>
        public readonly string? SecretLocation;
        /// <summary>
        /// secret name stored in keyvault.
        /// </summary>
        public readonly string? SecretName;

        [OutputConstructor]
        private SbeCredentialsResponse(
            string? eceSecretName,

            string? secretLocation,

            string? secretName)
        {
            EceSecretName = eceSecretName;
            SecretLocation = secretLocation;
            SecretName = secretName;
        }
    }
}
