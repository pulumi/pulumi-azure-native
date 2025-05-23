// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// A voice receiver.
    /// </summary>
    public sealed class VoiceReceiverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The country code of the voice receiver.
        /// </summary>
        [Input("countryCode", required: true)]
        public Input<string> CountryCode { get; set; } = null!;

        /// <summary>
        /// The name of the voice receiver. Names must be unique across all receivers within a tenant action group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The phone number of the voice receiver.
        /// </summary>
        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        public VoiceReceiverArgs()
        {
        }
        public static new VoiceReceiverArgs Empty => new VoiceReceiverArgs();
    }
}
