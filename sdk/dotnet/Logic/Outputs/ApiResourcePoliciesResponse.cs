// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.Outputs
{

    /// <summary>
    /// The API resource policies.
    /// </summary>
    [OutputType]
    public sealed class ApiResourcePoliciesResponse
    {
        /// <summary>
        /// The API level only policies XML as embedded content.
        /// </summary>
        public readonly string? Content;
        /// <summary>
        /// The content link to the policies.
        /// </summary>
        public readonly string? ContentLink;

        [OutputConstructor]
        private ApiResourcePoliciesResponse(
            string? content,

            string? contentLink)
        {
            Content = content;
            ContentLink = contentLink;
        }
    }
}
