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
    /// Definition of Output
    /// </summary>
    [OutputType]
    public sealed class OutputResponse
    {
        /// <summary>
        /// Property description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Property exportName
        /// </summary>
        public readonly string? ExportName;
        /// <summary>
        /// Property outputKey
        /// </summary>
        public readonly string? OutputKey;
        /// <summary>
        /// Property outputValue
        /// </summary>
        public readonly string? OutputValue;

        [OutputConstructor]
        private OutputResponse(
            string? description,

            string? exportName,

            string? outputKey,

            string? outputValue)
        {
            Description = description;
            ExportName = exportName;
            OutputKey = outputKey;
            OutputValue = outputValue;
        }
    }
}
