// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiCenter.Outputs
{

    /// <summary>
    /// The license information for the API.
    /// </summary>
    [OutputType]
    public sealed class LicenseResponse
    {
        /// <summary>
        /// SPDX license information for the API. The identifier field is mutually
        /// exclusive of the URL field.
        /// </summary>
        public readonly string? Identifier;
        /// <summary>
        /// Name of the license.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// URL pointing to the license details. The URL field is mutually exclusive of the
        /// identifier field.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private LicenseResponse(
            string? identifier,

            string? name,

            string? url)
        {
            Identifier = identifier;
            Name = name;
            Url = url;
        }
    }
}
