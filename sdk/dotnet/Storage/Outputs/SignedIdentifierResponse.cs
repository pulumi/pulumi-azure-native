// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    [OutputType]
    public sealed class SignedIdentifierResponse
    {
        /// <summary>
        /// Access policy
        /// </summary>
        public readonly Outputs.AccessPolicyResponse? AccessPolicy;
        /// <summary>
        /// An unique identifier of the stored access policy.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private SignedIdentifierResponse(
            Outputs.AccessPolicyResponse? accessPolicy,

            string? id)
        {
            AccessPolicy = accessPolicy;
            Id = id;
        }
    }
}
