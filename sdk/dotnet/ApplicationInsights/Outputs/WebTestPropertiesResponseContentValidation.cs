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
    /// The collection of content validation properties
    /// </summary>
    [OutputType]
    public sealed class WebTestPropertiesResponseContentValidation
    {
        /// <summary>
        /// Content to look for in the return of the WebTest.  Must not be null or empty.
        /// </summary>
        public readonly string? ContentMatch;
        /// <summary>
        /// When set, this value makes the ContentMatch validation case insensitive.
        /// </summary>
        public readonly bool? IgnoreCase;
        /// <summary>
        /// When true, validation will pass if there is a match for the ContentMatch string.  If false, validation will fail if there is a match
        /// </summary>
        public readonly bool? PassIfTextFound;

        [OutputConstructor]
        private WebTestPropertiesResponseContentValidation(
            string? contentMatch,

            bool? ignoreCase,

            bool? passIfTextFound)
        {
            ContentMatch = contentMatch;
            IgnoreCase = ignoreCase;
            PassIfTextFound = passIfTextFound;
        }
    }
}
