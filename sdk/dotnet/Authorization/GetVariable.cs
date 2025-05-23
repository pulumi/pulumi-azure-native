// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization
{
    public static class GetVariable
    {
        /// <summary>
        /// This operation retrieves a single variable, given its name and the subscription it was created at.
        /// 
        /// Uses Azure REST API version 2022-08-01-preview.
        /// 
        /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVariableResult> InvokeAsync(GetVariableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVariableResult>("azure-native:authorization:getVariable", args ?? new GetVariableArgs(), options.WithDefaults());

        /// <summary>
        /// This operation retrieves a single variable, given its name and the subscription it was created at.
        /// 
        /// Uses Azure REST API version 2022-08-01-preview.
        /// 
        /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVariableResult> Invoke(GetVariableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableResult>("azure-native:authorization:getVariable", args ?? new GetVariableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This operation retrieves a single variable, given its name and the subscription it was created at.
        /// 
        /// Uses Azure REST API version 2022-08-01-preview.
        /// 
        /// Other available API versions: 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVariableResult> Invoke(GetVariableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVariableResult>("azure-native:authorization:getVariable", args ?? new GetVariableInvokeArgs(), options.WithDefaults());
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
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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
            string azureApiVersion,

            ImmutableArray<Outputs.PolicyVariableColumnResponse> columns,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Columns = columns;
            Id = id;
            Name = name;
            SystemData = systemData;
            Type = type;
        }
    }
}
