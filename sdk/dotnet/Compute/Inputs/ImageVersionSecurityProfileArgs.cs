// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// The security profile of a gallery image version
    /// </summary>
    public sealed class ImageVersionSecurityProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains UEFI settings for the image version.
        /// </summary>
        [Input("uefiSettings")]
        public Input<Inputs.GalleryImageVersionUefiSettingsArgs>? UefiSettings { get; set; }

        public ImageVersionSecurityProfileArgs()
        {
        }
        public static new ImageVersionSecurityProfileArgs Empty => new ImageVersionSecurityProfileArgs();
    }
}
