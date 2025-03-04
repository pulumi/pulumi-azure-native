// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiCenter.V20240601Preview.Inputs
{

    /// <summary>
    /// Server
    /// </summary>
    public sealed class DeploymentServerArgs : global::Pulumi.ResourceArgs
    {
        [Input("runtimeUri")]
        private InputList<string>? _runtimeUri;

        /// <summary>
        /// Base runtime URLs for this deployment.
        /// </summary>
        public InputList<string> RuntimeUri
        {
            get => _runtimeUri ?? (_runtimeUri = new InputList<string>());
            set => _runtimeUri = value;
        }

        public DeploymentServerArgs()
        {
        }
        public static new DeploymentServerArgs Empty => new DeploymentServerArgs();
    }
}
