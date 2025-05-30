// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// Describes the bind options for the container
    /// </summary>
    [OutputType]
    public sealed class BindOptionsResponse
    {
        /// <summary>
        /// Indicate whether to create host path.
        /// </summary>
        public readonly bool? CreateHostPath;
        /// <summary>
        /// Type of Bind Option
        /// </summary>
        public readonly string? Propagation;
        /// <summary>
        /// Mention the selinux options.
        /// </summary>
        public readonly string? Selinux;

        [OutputConstructor]
        private BindOptionsResponse(
            bool? createHostPath,

            string? propagation,

            string? selinux)
        {
            CreateHostPath = createHostPath;
            Propagation = propagation;
            Selinux = selinux;
        }
    }
}
