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
    /// Describes the properties for an output ISO MP4 file.
    /// </summary>
    public sealed class Mp4FormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The file naming pattern used for the creation of output files. The following macros are supported in the file name: {Basename} - An expansion macro that will use the name of the input video file. If the base name(the file suffix is not included) of the input video file is less than 32 characters long, the base name of input video files will be used. If the length of base name of the input video file exceeds 32 characters, the base name is truncated to the first 32 characters in total length. {Extension} - The appropriate extension for this format. {Label} - The label assigned to the codec/layer. {Index} - A unique index for thumbnails. Only applicable to thumbnails. {AudioStream} - string "Audio" plus audio stream number(start from 1). {Bitrate} - The audio/video bitrate in kbps. Not applicable to thumbnails. {Codec} - The type of the audio/video codec. {Resolution} - The video resolution. Any unsubstituted macros will be collapsed and removed from the filename.
        /// </summary>
        [Input("filenamePattern", required: true)]
        public Input<string> FilenamePattern { get; set; } = null!;

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.Mp4Format'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        [Input("outputFiles")]
        private InputList<Inputs.OutputFileArgs>? _outputFiles;

        /// <summary>
        /// The list of output files to produce.  Each entry in the list is a set of audio and video layer labels to be muxed together .
        /// </summary>
        public InputList<Inputs.OutputFileArgs> OutputFiles
        {
            get => _outputFiles ?? (_outputFiles = new InputList<Inputs.OutputFileArgs>());
            set => _outputFiles = value;
        }

        public Mp4FormatArgs()
        {
        }
        public static new Mp4FormatArgs Empty => new Mp4FormatArgs();
    }
}
