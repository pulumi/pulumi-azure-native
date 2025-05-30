// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization
{
    public static class GetVariableValue
    {
        /// <summary>
        /// This operation retrieves a single variable value; given its name, subscription it was created at and the variable it's created for.
        /// 
        /// Uses Azure REST API version 2022-08-01-preview.
        /// 
        /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVariableValueResult> InvokeAsync(GetVariableValueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVariableValueResult>("azure-native:authorization:getVariableValue", args ?? new GetVariableValueArgs(), options.WithDefaults());

        /// <summary>
        /// This operation retrieves a single variable value; given its name, subscription it was created at and the variable it's created for.
        /// 
        /// Uses Azure REST API version 2022-08-01-preview.
        /// 
        /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVariableValueResult> Invoke(GetVariableValueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableValueResult>("azure-native:authorization:getVariableValue", args ?? new GetVariableValueInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This operation retrieves a single variable value; given its name, subscription it was created at and the variable it's created for.
        /// 
        /// Uses Azure REST API version 2022-08-01-preview.
        /// 
        /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVariableValueResult> Invoke(GetVariableValueInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableValueResult>("azure-native:authorization:getVariableValue", args ?? new GetVariableValueInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVariableValueArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the variable to operate on.
        /// </summary>
        [Input("variableName", required: true)]
        public string VariableName { get; set; } = null!;

        /// <summary>
        /// The name of the variable value to operate on.
        /// </summary>
        [Input("variableValueName", required: true)]
        public string VariableValueName { get; set; } = null!;

        public GetVariableValueArgs()
        {
        }
        public static new GetVariableValueArgs Empty => new GetVariableValueArgs();
    }

    public sealed class GetVariableValueInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the variable to operate on.
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        /// <summary>
        /// The name of the variable value to operate on.
        /// </summary>
        [Input("variableValueName", required: true)]
        public Input<string> VariableValueName { get; set; } = null!;

        public GetVariableValueInvokeArgs()
        {
        }
        public static new GetVariableValueInvokeArgs Empty => new GetVariableValueInvokeArgs();
    }


    [OutputType]
    public sealed class GetVariableValueResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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
        /// The type of the resource (Microsoft.Authorization/variables/values).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Variable value column value array.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyVariableValueColumnValueResponse> Values;

        [OutputConstructor]
        private GetVariableValueResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string type,

            ImmutableArray<Outputs.PolicyVariableValueColumnValueResponse> values)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            SystemData = systemData;
            Type = type;
            Values = values;
        }
    }
}
