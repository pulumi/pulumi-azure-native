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
    /// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
    /// </summary>
    [OutputType]
    public sealed class ApplicationInsightsComponentAnalyticsItemPropertiesResponse
    {
        /// <summary>
        /// A function alias, used when the type of the item is Function
        /// </summary>
        public readonly string? FunctionAlias;

        [OutputConstructor]
        private ApplicationInsightsComponentAnalyticsItemPropertiesResponse(string? functionAlias)
        {
            FunctionAlias = functionAlias;
        }
    }
}
