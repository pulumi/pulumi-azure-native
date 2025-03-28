// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoIndexer.V20250301.Inputs
{

    /// <summary>
    /// The openAi services details
    /// </summary>
    public sealed class OpenAiServicesForPutRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The openAi services resource id
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The user assigned identity to be used to grant permissions
        /// </summary>
        [Input("userAssignedIdentity")]
        public Input<string>? UserAssignedIdentity { get; set; }

        public OpenAiServicesForPutRequestArgs()
        {
        }
        public static new OpenAiServicesForPutRequestArgs Empty => new OpenAiServicesForPutRequestArgs();
    }
}
