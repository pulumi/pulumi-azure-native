// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// A header to add to the WebTest.
    /// </summary>
    [OutputType]
    public sealed class HeaderFieldResponse
    {
        /// <summary>
        /// The name of the header.
        /// </summary>
        public readonly string? HeaderFieldName;
        /// <summary>
        /// The value of the header.
        /// </summary>
        public readonly string? HeaderFieldValue;

        [OutputConstructor]
        private HeaderFieldResponse(
            string? headerFieldName,

            string? headerFieldValue)
        {
            HeaderFieldName = headerFieldName;
            HeaderFieldValue = headerFieldValue;
        }
    }
}
