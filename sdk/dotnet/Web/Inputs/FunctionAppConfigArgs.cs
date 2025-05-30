// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Function app configuration.
    /// </summary>
    public sealed class FunctionAppConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Function app deployment configuration.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.FunctionsDeploymentArgs>? Deployment { get; set; }

        /// <summary>
        /// Function app runtime settings.
        /// </summary>
        [Input("runtime")]
        public Input<Inputs.FunctionsRuntimeArgs>? Runtime { get; set; }

        /// <summary>
        /// Function app scale and concurrency settings.
        /// </summary>
        [Input("scaleAndConcurrency")]
        public Input<Inputs.FunctionsScaleAndConcurrencyArgs>? ScaleAndConcurrency { get; set; }

        public FunctionAppConfigArgs()
        {
        }
        public static new FunctionAppConfigArgs Empty => new FunctionAppConfigArgs();
    }
}
