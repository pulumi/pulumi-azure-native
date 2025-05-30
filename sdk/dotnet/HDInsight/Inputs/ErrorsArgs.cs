// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// The error message associated with the cluster creation.
    /// </summary>
    public sealed class ErrorsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The error code.
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// The error message.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public ErrorsArgs()
        {
        }
        public static new ErrorsArgs Empty => new ErrorsArgs();
    }
}
