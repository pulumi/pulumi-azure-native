// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Inputs
{

    /// <summary>
    /// The resource selector to filter policies by resource properties.
    /// </summary>
    public sealed class ResourceSelectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource selector.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("selectors")]
        private InputList<Inputs.SelectorArgs>? _selectors;

        /// <summary>
        /// The list of the selector expressions.
        /// </summary>
        public InputList<Inputs.SelectorArgs> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Inputs.SelectorArgs>());
            set => _selectors = value;
        }

        public ResourceSelectorArgs()
        {
        }
        public static new ResourceSelectorArgs Empty => new ResourceSelectorArgs();
    }
}
