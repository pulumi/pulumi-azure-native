// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearning.V20160501Preview.Inputs
{

    /// <summary>
    /// Sample input data for the service's input(s).
    /// </summary>
    public sealed class ExampleRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("globalParameters")]
        private InputMap<object>? _globalParameters;

        /// <summary>
        /// Sample input data for the web service's global parameters
        /// </summary>
        public InputMap<object> GlobalParameters
        {
            get => _globalParameters ?? (_globalParameters = new InputMap<object>());
            set => _globalParameters = value;
        }

        [Input("inputs")]
        private InputMap<ImmutableArray<ImmutableArray<object>>>? _inputs;

        /// <summary>
        /// Sample input data for the web service's input(s) given as an input name to sample input values matrix map.
        /// </summary>
        public InputMap<ImmutableArray<ImmutableArray<object>>> Inputs
        {
            get => _inputs ?? (_inputs = new InputMap<ImmutableArray<ImmutableArray<object>>>());
            set => _inputs = value;
        }

        public ExampleRequestArgs()
        {
        }
        public static new ExampleRequestArgs Empty => new ExampleRequestArgs();
    }
}
