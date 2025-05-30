// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.Inputs
{

    /// <summary>
    /// Describes an image source that is a managed image in customer subscription. This image must reside in the same subscription and region as the Image Builder template.
    /// </summary>
    public sealed class ImageTemplateManagedImageSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARM resource id of the managed image in customer subscription
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// Specifies the type of source image you want to start with.
        /// Expected value is 'ManagedImage'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ImageTemplateManagedImageSourceArgs()
        {
        }
        public static new ImageTemplateManagedImageSourceArgs Empty => new ImageTemplateManagedImageSourceArgs();
    }
}
