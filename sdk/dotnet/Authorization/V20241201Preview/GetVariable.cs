// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.V20241201Preview
{
    public static class GetVariable
    {
        /// <summary>
        /// This operation retrieves a single variable, given its name and the subscription it was created at.
        /// </summary>
        public static Task<GetVariableResult> InvokeAsync(GetVariableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVariableResult>("azure-native:authorization/v20241201preview:getVariable", args ?? new GetVariableArgs(), options.WithDefaults());

        /// <summary>
        /// This operation retrieves a single variable, given its name and the subscription it was created at.
        /// </summary>
        public static Output<GetVariableResult> Invoke(GetVariableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableResult>("azure-native:authorization/v20241201preview:getVariable", args ?? new GetVariableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This operation retrieves a single variable, given its name and the subscription it was created at.
        /// </summary>
        public static Output<GetVariableResult> Invoke(GetVariableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableResult>("azure-native:authorization/v20241201preview:getVariable", args ?? new GetVariableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVariableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the variable to operate on.
        /// </summary>
        [Input("variableName", required: true)]
        public string VariableName { get; set; } = null!;

        public GetVariableArgs()
        {
        }
        public static new GetVariableArgs Empty => new GetVariableArgs();
    }

    public sealed class GetVariableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the variable to operate on.
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        public GetVariableInvokeArgs()
        {
        }
        public static new GetVariableInvokeArgs Empty => new GetVariableInvokeArgs();
    }


    [OutputType]
    public sealed class GetVariableResult
    {
        /// <summary>
        /// Variable column definitions.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyVariableColumnResponse> Columns;
        /// <summary>
        /// The ID of the variable.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the variable.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource (Microsoft.Authorization/variables).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVariableResult(
            ImmutableArray<Outputs.PolicyVariableColumnResponse> columns,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Columns = columns;
            Id = id;
            Name = name;
            SystemData = systemData;
            Type = type;
        }
    }
}
