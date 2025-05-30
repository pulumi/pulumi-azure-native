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
    /// Error handling options upon a build failure
    /// </summary>
    [OutputType]
    public sealed class ImageTemplatePropertiesResponseErrorHandling
    {
        /// <summary>
        /// If there is a customizer error and this field is set to 'cleanup', the build VM and associated network resources will be cleaned up. This is the default behavior. If there is a customizer error and this field is set to 'abort', the build VM will be preserved.
        /// </summary>
        public readonly string? OnCustomizerError;
        /// <summary>
        /// If there is a validation error and this field is set to 'cleanup', the build VM and associated network resources will be cleaned up. This is the default behavior. If there is a validation error and this field is set to 'abort', the build VM will be preserved.
        /// </summary>
        public readonly string? OnValidationError;

        [OutputConstructor]
        private ImageTemplatePropertiesResponseErrorHandling(
            string? onCustomizerError,

            string? onValidationError)
        {
            OnCustomizerError = onCustomizerError;
            OnValidationError = onValidationError;
        }
    }
}
