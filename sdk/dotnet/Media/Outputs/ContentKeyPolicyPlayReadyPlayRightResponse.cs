// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Outputs
{

    /// <summary>
    /// Configures the Play Right in the PlayReady license.
    /// </summary>
    [OutputType]
    public sealed class ContentKeyPolicyPlayReadyPlayRightResponse
    {
        /// <summary>
        /// Configures Automatic Gain Control (AGC) and Color Stripe in the license. Must be between 0 and 3 inclusive.
        /// </summary>
        public readonly int? AgcAndColorStripeRestriction;
        /// <summary>
        /// Configures Unknown output handling settings of the license.
        /// </summary>
        public readonly string AllowPassingVideoContentToUnknownOutput;
        /// <summary>
        /// Specifies the output protection level for compressed digital audio.
        /// </summary>
        public readonly int? AnalogVideoOpl;
        /// <summary>
        /// Specifies the output protection level for compressed digital audio.
        /// </summary>
        public readonly int? CompressedDigitalAudioOpl;
        /// <summary>
        /// Specifies the output protection level for compressed digital video.
        /// </summary>
        public readonly int? CompressedDigitalVideoOpl;
        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        public readonly bool DigitalVideoOnlyContentRestriction;
        /// <summary>
        /// Configures the Explicit Analog Television Output Restriction in the license. Configuration data must be between 0 and 3 inclusive.
        /// </summary>
        public readonly Outputs.ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse? ExplicitAnalogTelevisionOutputRestriction;
        /// <summary>
        /// The amount of time that the license is valid after the license is first used to play content.
        /// </summary>
        public readonly string? FirstPlayExpiration;
        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        public readonly bool ImageConstraintForAnalogComponentVideoRestriction;
        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        public readonly bool ImageConstraintForAnalogComputerMonitorRestriction;
        /// <summary>
        /// Configures the Serial Copy Management System (SCMS) in the license. Must be between 0 and 3 inclusive.
        /// </summary>
        public readonly int? ScmsRestriction;
        /// <summary>
        /// Specifies the output protection level for uncompressed digital audio.
        /// </summary>
        public readonly int? UncompressedDigitalAudioOpl;
        /// <summary>
        /// Specifies the output protection level for uncompressed digital video.
        /// </summary>
        public readonly int? UncompressedDigitalVideoOpl;

        [OutputConstructor]
        private ContentKeyPolicyPlayReadyPlayRightResponse(
            int? agcAndColorStripeRestriction,

            string allowPassingVideoContentToUnknownOutput,

            int? analogVideoOpl,

            int? compressedDigitalAudioOpl,

            int? compressedDigitalVideoOpl,

            bool digitalVideoOnlyContentRestriction,

            Outputs.ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse? explicitAnalogTelevisionOutputRestriction,

            string? firstPlayExpiration,

            bool imageConstraintForAnalogComponentVideoRestriction,

            bool imageConstraintForAnalogComputerMonitorRestriction,

            int? scmsRestriction,

            int? uncompressedDigitalAudioOpl,

            int? uncompressedDigitalVideoOpl)
        {
            AgcAndColorStripeRestriction = agcAndColorStripeRestriction;
            AllowPassingVideoContentToUnknownOutput = allowPassingVideoContentToUnknownOutput;
            AnalogVideoOpl = analogVideoOpl;
            CompressedDigitalAudioOpl = compressedDigitalAudioOpl;
            CompressedDigitalVideoOpl = compressedDigitalVideoOpl;
            DigitalVideoOnlyContentRestriction = digitalVideoOnlyContentRestriction;
            ExplicitAnalogTelevisionOutputRestriction = explicitAnalogTelevisionOutputRestriction;
            FirstPlayExpiration = firstPlayExpiration;
            ImageConstraintForAnalogComponentVideoRestriction = imageConstraintForAnalogComponentVideoRestriction;
            ImageConstraintForAnalogComputerMonitorRestriction = imageConstraintForAnalogComputerMonitorRestriction;
            ScmsRestriction = scmsRestriction;
            UncompressedDigitalAudioOpl = uncompressedDigitalAudioOpl;
            UncompressedDigitalVideoOpl = uncompressedDigitalVideoOpl;
        }
    }
}
