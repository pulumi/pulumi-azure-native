// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.Outputs
{

    /// <summary>
    /// Specifies optimization to be performed on image.
    /// </summary>
    [OutputType]
    public sealed class ImageTemplatePropertiesResponseOptimize
    {
        /// <summary>
        /// Optimization is applied on the image for a faster VM boot.
        /// </summary>
        public readonly Outputs.ImageTemplatePropertiesResponseVmBoot? VmBoot;

        [OutputConstructor]
        private ImageTemplatePropertiesResponseOptimize(Outputs.ImageTemplatePropertiesResponseVmBoot? vmBoot)
        {
            VmBoot = vmBoot;
        }
    }
}
