// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Chaos.Inputs
{

    /// <summary>
    /// Model that represents the Experiment properties model.
    /// </summary>
    public sealed class ExperimentPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional customer-managed Storage account where Experiment schema will be stored.
        /// </summary>
        [Input("customerDataStorage")]
        public Input<Inputs.CustomerDataStoragePropertiesArgs>? CustomerDataStorage { get; set; }

        [Input("selectors", required: true)]
        private InputList<Union<Inputs.ListSelectorArgs, Inputs.QuerySelectorArgs>>? _selectors;

        /// <summary>
        /// List of selectors.
        /// </summary>
        public InputList<Union<Inputs.ListSelectorArgs, Inputs.QuerySelectorArgs>> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Union<Inputs.ListSelectorArgs, Inputs.QuerySelectorArgs>>());
            set => _selectors = value;
        }

        [Input("steps", required: true)]
        private InputList<Inputs.StepArgs>? _steps;

        /// <summary>
        /// List of steps.
        /// </summary>
        public InputList<Inputs.StepArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.StepArgs>());
            set => _steps = value;
        }

        public ExperimentPropertiesArgs()
        {
        }
        public static new ExperimentPropertiesArgs Empty => new ExperimentPropertiesArgs();
    }
}
