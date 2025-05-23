// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of Tier
    /// </summary>
    [OutputType]
    public sealed class TierResponse
    {
        /// <summary>
        /// The name of this environment tier.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of this environment tier.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private TierResponse(
            string? name,

            string? type,

            string? version)
        {
            Name = name;
            Type = type;
            Version = version;
        }
    }
}
