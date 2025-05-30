// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Execute spark job activity.
    /// </summary>
    [OutputType]
    public sealed class SynapseSparkJobDefinitionActivityResponse
    {
        /// <summary>
        /// User specified arguments to SynapseSparkJobDefinitionActivity.
        /// </summary>
        public readonly ImmutableArray<object> Arguments;
        /// <summary>
        /// The fully-qualified identifier or the main class that is in the main definition file, which will override the 'className' of the spark job definition you provide. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ClassName;
        /// <summary>
        /// Spark configuration properties, which will override the 'conf' of the spark job definition you provide.
        /// </summary>
        public readonly object? Conf;
        /// <summary>
        /// The type of the spark config.
        /// </summary>
        public readonly string? ConfigurationType;
        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityDependencyResponse> DependsOn;
        /// <summary>
        /// Activity description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Number of core and memory to be used for driver allocated in the specified Spark pool for the job, which will be used for overriding 'driverCores' and 'driverMemory' of the spark job definition you provide. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? DriverSize;
        /// <summary>
        /// Number of core and memory to be used for executors allocated in the specified Spark pool for the job, which will be used for overriding 'executorCores' and 'executorMemory' of the spark job definition you provide. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? ExecutorSize;
        /// <summary>
        /// The main file used for the job, which will override the 'file' of the spark job definition you provide. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object? File;
        /// <summary>
        /// (Deprecated. Please use pythonCodeReference and filesV2) Additional files used for reference in the main definition file, which will override the 'files' of the spark job definition you provide.
        /// </summary>
        public readonly ImmutableArray<object> Files;
        /// <summary>
        /// Additional files used for reference in the main definition file, which will override the 'jars' and 'files' of the spark job definition you provide.
        /// </summary>
        public readonly ImmutableArray<object> FilesV2;
        /// <summary>
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? LinkedServiceName;
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of executors to launch for this job, which will override the 'numExecutors' of the spark job definition you provide. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly object? NumExecutors;
        /// <summary>
        /// Status result of the activity when the state is set to Inactive. This is an optional property and if not provided when the activity is inactive, the status will be Succeeded by default.
        /// </summary>
        public readonly string? OnInactiveMarkAs;
        /// <summary>
        /// Activity policy.
        /// </summary>
        public readonly Outputs.ActivityPolicyResponse? Policy;
        /// <summary>
        /// Additional python code files used for reference in the main definition file, which will override the 'pyFiles' of the spark job definition you provide.
        /// </summary>
        public readonly ImmutableArray<object> PythonCodeReference;
        /// <summary>
        /// Scanning subfolders from the root folder of the main definition file, these files will be added as reference files. The folders named 'jars', 'pyFiles', 'files' or 'archives' will be scanned, and the folders name are case sensitive. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? ScanFolder;
        /// <summary>
        /// Spark configuration property.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? SparkConfig;
        /// <summary>
        /// Synapse spark job reference.
        /// </summary>
        public readonly Outputs.SynapseSparkJobReferenceResponse SparkJob;
        /// <summary>
        /// Activity state. This is an optional property and if not provided, the state will be Active by default.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The name of the big data pool which will be used to execute the spark batch job, which will override the 'targetBigDataPool' of the spark job definition you provide.
        /// </summary>
        public readonly Outputs.BigDataPoolParametrizationReferenceResponse? TargetBigDataPool;
        /// <summary>
        /// The spark configuration of the spark job.
        /// </summary>
        public readonly Outputs.SparkConfigurationParametrizationReferenceResponse? TargetSparkConfiguration;
        /// <summary>
        /// Type of activity.
        /// Expected value is 'SparkJob'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Activity user properties.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserPropertyResponse> UserProperties;

        [OutputConstructor]
        private SynapseSparkJobDefinitionActivityResponse(
            ImmutableArray<object> arguments,

            object? className,

            object? conf,

            string? configurationType,

            ImmutableArray<Outputs.ActivityDependencyResponse> dependsOn,

            string? description,

            object? driverSize,

            object? executorSize,

            object? file,

            ImmutableArray<object> files,

            ImmutableArray<object> filesV2,

            Outputs.LinkedServiceReferenceResponse? linkedServiceName,

            string name,

            object? numExecutors,

            string? onInactiveMarkAs,

            Outputs.ActivityPolicyResponse? policy,

            ImmutableArray<object> pythonCodeReference,

            object? scanFolder,

            ImmutableDictionary<string, object>? sparkConfig,

            Outputs.SynapseSparkJobReferenceResponse sparkJob,

            string? state,

            Outputs.BigDataPoolParametrizationReferenceResponse? targetBigDataPool,

            Outputs.SparkConfigurationParametrizationReferenceResponse? targetSparkConfiguration,

            string type,

            ImmutableArray<Outputs.UserPropertyResponse> userProperties)
        {
            Arguments = arguments;
            ClassName = className;
            Conf = conf;
            ConfigurationType = configurationType;
            DependsOn = dependsOn;
            Description = description;
            DriverSize = driverSize;
            ExecutorSize = executorSize;
            File = file;
            Files = files;
            FilesV2 = filesV2;
            LinkedServiceName = linkedServiceName;
            Name = name;
            NumExecutors = numExecutors;
            OnInactiveMarkAs = onInactiveMarkAs;
            Policy = policy;
            PythonCodeReference = pythonCodeReference;
            ScanFolder = scanFolder;
            SparkConfig = sparkConfig;
            SparkJob = sparkJob;
            State = state;
            TargetBigDataPool = targetBigDataPool;
            TargetSparkConfiguration = targetSparkConfiguration;
            Type = type;
            UserProperties = userProperties;
        }
    }
}
