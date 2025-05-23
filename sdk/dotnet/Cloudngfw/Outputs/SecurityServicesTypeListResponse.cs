// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Outputs
{

    /// <summary>
    /// Security services type list
    /// </summary>
    [OutputType]
    public sealed class SecurityServicesTypeListResponse
    {
        /// <summary>
        /// list
        /// </summary>
        public readonly ImmutableArray<Outputs.NameDescriptionObjectResponse> Entry;
        /// <summary>
        /// security services type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private SecurityServicesTypeListResponse(
            ImmutableArray<Outputs.NameDescriptionObjectResponse> entry,

            string? type)
        {
            Entry = entry;
            Type = type;
        }
    }
}
