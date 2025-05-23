// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    /// <summary>
    /// Specify either the userName or autoUser property, but not both.
    /// </summary>
    [OutputType]
    public sealed class UserIdentityResponse
    {
        /// <summary>
        /// The userName and autoUser properties are mutually exclusive; you must specify one but not both.
        /// </summary>
        public readonly Outputs.AutoUserSpecificationResponse? AutoUser;
        /// <summary>
        /// The userName and autoUser properties are mutually exclusive; you must specify one but not both.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private UserIdentityResponse(
            Outputs.AutoUserSpecificationResponse? autoUser,

            string? userName)
        {
            AutoUser = autoUser;
            UserName = userName;
        }
    }
}
