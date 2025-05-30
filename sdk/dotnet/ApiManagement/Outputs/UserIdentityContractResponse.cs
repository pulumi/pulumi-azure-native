// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Outputs
{

    /// <summary>
    /// User identity details.
    /// </summary>
    [OutputType]
    public sealed class UserIdentityContractResponse
    {
        /// <summary>
        /// Identifier value within provider.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Identity provider name.
        /// </summary>
        public readonly string? Provider;

        [OutputConstructor]
        private UserIdentityContractResponse(
            string? id,

            string? provider)
        {
            Id = id;
            Provider = provider;
        }
    }
}
