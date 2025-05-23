// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Inputs
{

    /// <summary>
    /// The configuration of the backend circuit breaker
    /// </summary>
    public sealed class BackendCircuitBreakerArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.CircuitBreakerRuleArgs>? _rules;

        /// <summary>
        /// The rules for tripping the backend.
        /// </summary>
        public InputList<Inputs.CircuitBreakerRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.CircuitBreakerRuleArgs>());
            set => _rules = value;
        }

        public BackendCircuitBreakerArgs()
        {
        }
        public static new BackendCircuitBreakerArgs Empty => new BackendCircuitBreakerArgs();
    }
}
