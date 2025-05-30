// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Inputs
{

    /// <summary>
    /// Class to specify configurations of PlayReady in Streaming Policy
    /// </summary>
    public sealed class StreamingPolicyPlayReadyConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template for the URL of the custom service delivering licenses to end user players.  Not required when using Azure Media Services for issuing licenses.  The template supports replaceable tokens that the service will update at runtime with the value specific to the request.  The currently supported token values are {AlternativeMediaId}, which is replaced with the value of StreamingLocatorId.AlternativeMediaId, and {ContentKeyId}, which is replaced with the value of identifier of the key being requested.
        /// </summary>
        [Input("customLicenseAcquisitionUrlTemplate")]
        public Input<string>? CustomLicenseAcquisitionUrlTemplate { get; set; }

        /// <summary>
        /// Custom attributes for PlayReady
        /// </summary>
        [Input("playReadyCustomAttributes")]
        public Input<string>? PlayReadyCustomAttributes { get; set; }

        public StreamingPolicyPlayReadyConfigurationArgs()
        {
        }
        public static new StreamingPolicyPlayReadyConfigurationArgs Empty => new StreamingPolicyPlayReadyConfigurationArgs();
    }
}
