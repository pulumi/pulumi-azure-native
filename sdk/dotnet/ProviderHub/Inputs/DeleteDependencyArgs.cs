// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class DeleteDependencyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Linked property.
        /// </summary>
        [Input("linkedProperty")]
        public Input<string>? LinkedProperty { get; set; }

        /// <summary>
        /// Linked type.
        /// </summary>
        [Input("linkedType")]
        public Input<string>? LinkedType { get; set; }

        [Input("requiredFeatures")]
        private InputList<string>? _requiredFeatures;

        /// <summary>
        /// Required features.
        /// </summary>
        public InputList<string> RequiredFeatures
        {
            get => _requiredFeatures ?? (_requiredFeatures = new InputList<string>());
            set => _requiredFeatures = value;
        }

        public DeleteDependencyArgs()
        {
        }
        public static new DeleteDependencyArgs Empty => new DeleteDependencyArgs();
    }
}
