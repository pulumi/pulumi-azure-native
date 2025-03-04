// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20240802Preview.Outputs
{

    /// <summary>
    /// Dapr component metadata.
    /// </summary>
    [OutputType]
    public sealed class DaprServiceBindMetadataResponse
    {
        /// <summary>
        /// Service bind metadata property name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Service bind metadata property value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private DaprServiceBindMetadataResponse(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
