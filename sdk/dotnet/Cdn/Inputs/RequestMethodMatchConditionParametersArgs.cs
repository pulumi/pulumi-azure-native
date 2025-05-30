// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Inputs
{

    /// <summary>
    /// Defines the parameters for RequestMethod match conditions
    /// </summary>
    public sealed class RequestMethodMatchConditionParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("matchValues")]
        private InputList<Union<string, Pulumi.AzureNative.Cdn.RequestMethodMatchValue>>? _matchValues;

        /// <summary>
        /// The match value for the condition of the delivery rule
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Cdn.RequestMethodMatchValue>> MatchValues
        {
            get => _matchValues ?? (_matchValues = new InputList<Union<string, Pulumi.AzureNative.Cdn.RequestMethodMatchValue>>());
            set => _matchValues = value;
        }

        /// <summary>
        /// Describes if this is negate condition or not
        /// </summary>
        [Input("negateCondition")]
        public Input<bool>? NegateCondition { get; set; }

        /// <summary>
        /// Describes operator to be matched
        /// </summary>
        [Input("operator", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Cdn.RequestMethodOperator> Operator { get; set; } = null!;

        [Input("transforms")]
        private InputList<Union<string, Pulumi.AzureNative.Cdn.Transform>>? _transforms;

        /// <summary>
        /// List of transforms
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Cdn.Transform>> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<Union<string, Pulumi.AzureNative.Cdn.Transform>>());
            set => _transforms = value;
        }

        /// <summary>
        /// 
        /// Expected value is 'DeliveryRuleRequestMethodConditionParameters'.
        /// </summary>
        [Input("typeName", required: true)]
        public Input<string> TypeName { get; set; } = null!;

        public RequestMethodMatchConditionParametersArgs()
        {
            NegateCondition = false;
        }
        public static new RequestMethodMatchConditionParametersArgs Empty => new RequestMethodMatchConditionParametersArgs();
    }
}
