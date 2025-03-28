// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20250101Preview.Inputs
{

    public sealed class JobResourcesArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;

        /// <summary>
        /// List of instance types to choose from.
        /// </summary>
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        public JobResourcesArgs()
        {
        }
        public static new JobResourcesArgs Empty => new JobResourcesArgs();
    }
}
