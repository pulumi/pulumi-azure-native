// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.V20240815Preview.Inputs
{

    /// <summary>
    /// The properties of an ADR instance.
    /// </summary>
    public sealed class AdrPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This determines if the ADR service is enabled.
        /// </summary>
        [Input("state", required: true)]
        public InputUnion<string, Pulumi.AzureNative.IoTOperations.V20240815Preview.OperationalMode> State { get; set; } = null!;

        public AdrPropertiesArgs()
        {
        }
        public static new AdrPropertiesArgs Empty => new AdrPropertiesArgs();
    }
}
