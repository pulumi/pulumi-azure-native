// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// Deployment Configuration
    /// </summary>
    public sealed class DeploymentConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("scaleUnits", required: true)]
        private InputList<Inputs.ScaleUnitsArgs>? _scaleUnits;

        /// <summary>
        /// Scale units will contains list of deployment data
        /// </summary>
        public InputList<Inputs.ScaleUnitsArgs> ScaleUnits
        {
            get => _scaleUnits ?? (_scaleUnits = new InputList<Inputs.ScaleUnitsArgs>());
            set => _scaleUnits = value;
        }

        /// <summary>
        /// deployment template version 
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DeploymentConfigurationArgs()
        {
        }
        public static new DeploymentConfigurationArgs Empty => new DeploymentConfigurationArgs();
    }
}
