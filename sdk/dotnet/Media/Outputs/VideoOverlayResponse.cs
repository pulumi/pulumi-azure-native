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
    /// Describes the properties of a video overlay.
    /// </summary>
    [OutputType]
    public sealed class VideoOverlayResponse
    {
        /// <summary>
        /// The gain level of audio in the overlay. The value should be in the range [0, 1.0]. The default is 1.0.
        /// </summary>
        public readonly double? AudioGainLevel;
        /// <summary>
        /// An optional rectangular window used to crop the overlay image or video.
        /// </summary>
        public readonly Outputs.RectangleResponse? CropRectangle;
        /// <summary>
        /// The end position, with reference to the input video, at which the overlay ends. The value should be in ISO 8601 format. For example, PT30S to end the overlay at 30 seconds into the input video. If not specified or the value is greater than the input video duration, the overlay will be applied until the end of the input video if the overlay media duration is greater than the input video duration, else the overlay will last as long as the overlay media duration.
        /// </summary>
        public readonly string? End;
        /// <summary>
        /// The duration over which the overlay fades in onto the input video. The value should be in ISO 8601 duration format. If not specified the default behavior is to have no fade in (same as PT0S).
        /// </summary>
        public readonly string? FadeInDuration;
        /// <summary>
        /// The duration over which the overlay fades out of the input video. The value should be in ISO 8601 duration format. If not specified the default behavior is to have no fade out (same as PT0S).
        /// </summary>
        public readonly string? FadeOutDuration;
        /// <summary>
        /// The label of the job input which is to be used as an overlay. The Input must specify exactly one file. You can specify an image file in JPG, PNG, GIF or BMP format, or an audio file (such as a WAV, MP3, WMA or M4A file), or a video file. See https://aka.ms/mesformats for the complete list of supported audio and video file formats.
        /// </summary>
        public readonly string InputLabel;
        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.VideoOverlay'.
        /// </summary>
        public readonly string OdataType;
        /// <summary>
        /// The opacity of the overlay. This is a value in the range [0 - 1.0]. Default is 1.0 which mean the overlay is opaque.
        /// </summary>
        public readonly double? Opacity;
        /// <summary>
        /// The location in the input video where the overlay is applied.
        /// </summary>
        public readonly Outputs.RectangleResponse? Position;
        /// <summary>
        /// The start position, with reference to the input video, at which the overlay starts. The value should be in ISO 8601 format. For example, PT05S to start the overlay at 5 seconds into the input video. If not specified the overlay starts from the beginning of the input video.
        /// </summary>
        public readonly string? Start;

        [OutputConstructor]
        private VideoOverlayResponse(
            double? audioGainLevel,

            Outputs.RectangleResponse? cropRectangle,

            string? end,

            string? fadeInDuration,

            string? fadeOutDuration,

            string inputLabel,

            string odataType,

            double? opacity,

            Outputs.RectangleResponse? position,

            string? start)
        {
            AudioGainLevel = audioGainLevel;
            CropRectangle = cropRectangle;
            End = end;
            FadeInDuration = fadeInDuration;
            FadeOutDuration = fadeOutDuration;
            InputLabel = inputLabel;
            OdataType = odataType;
            Opacity = opacity;
            Position = position;
            Start = start;
        }
    }
}
