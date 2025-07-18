// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    /// <summary>
    /// The rest of the world group two region configuration.
    /// </summary>
    public sealed class DefaultRolloutSpecificationRestOfTheWorldGroupTwoArgs : global::Pulumi.ResourceArgs
    {
        [Input("regions")]
        private InputList<string>? _regions;
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// The wait duration.
        /// </summary>
        [Input("waitDuration")]
        public Input<string>? WaitDuration { get; set; }

        public DefaultRolloutSpecificationRestOfTheWorldGroupTwoArgs()
        {
        }
        public static new DefaultRolloutSpecificationRestOfTheWorldGroupTwoArgs Empty => new DefaultRolloutSpecificationRestOfTheWorldGroupTwoArgs();
    }
}
