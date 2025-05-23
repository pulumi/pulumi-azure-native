// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Impact.Outputs
{

    /// <summary>
    /// Article details of the insight like title, description etc
    /// </summary>
    [OutputType]
    public sealed class ContentResponse
    {
        /// <summary>
        /// Description of the insight
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Title of the insight
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private ContentResponse(
            string description,

            string title)
        {
            Description = description;
            Title = title;
        }
    }
}
