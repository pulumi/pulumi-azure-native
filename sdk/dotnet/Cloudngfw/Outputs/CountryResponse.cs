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
    /// Country Description
    /// </summary>
    [OutputType]
    public sealed class CountryResponse
    {
        /// <summary>
        /// country code
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// code description
        /// </summary>
        public readonly string? Description;

        [OutputConstructor]
        private CountryResponse(
            string code,

            string? description)
        {
            Code = code;
            Description = description;
        }
    }
}
