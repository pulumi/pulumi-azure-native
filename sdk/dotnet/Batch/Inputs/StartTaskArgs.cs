// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Inputs
{

    /// <summary>
    /// In some cases the start task may be re-run even though the node was not rebooted. Due to this, start tasks should be idempotent and exit gracefully if the setup they're performing has already been done. Special care should be taken to avoid start tasks which create breakaway process or install/launch services from the start task working directory, as this will block Batch from being able to re-run the start task.
    /// </summary>
    public sealed class StartTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command line does not run under a shell, and therefore cannot take advantage of shell features such as environment variable expansion. If you want to take advantage of such features, you should invoke the shell in the command line, for example using "cmd /c MyCommand" in Windows or "/bin/sh -c MyCommand" in Linux. Required if any other properties of the startTask are specified.
        /// </summary>
        [Input("commandLine")]
        public Input<string>? CommandLine { get; set; }

        /// <summary>
        /// When this is specified, all directories recursively below the AZ_BATCH_NODE_ROOT_DIR (the root of Azure Batch directories on the node) are mapped into the container, all task environment variables are mapped into the container, and the task command line is executed in the container.
        /// </summary>
        [Input("containerSettings")]
        public Input<Inputs.TaskContainerSettingsArgs>? ContainerSettings { get; set; }

        [Input("environmentSettings")]
        private InputList<Inputs.EnvironmentSettingArgs>? _environmentSettings;
        public InputList<Inputs.EnvironmentSettingArgs> EnvironmentSettings
        {
            get => _environmentSettings ?? (_environmentSettings = new InputList<Inputs.EnvironmentSettingArgs>());
            set => _environmentSettings = value;
        }

        /// <summary>
        /// The Batch service retries a task if its exit code is nonzero. Note that this value specifically controls the number of retries. The Batch service will try the task once, and may then retry up to this limit. For example, if the maximum retry count is 3, Batch tries the task up to 4 times (one initial try and 3 retries). If the maximum retry count is 0, the Batch service does not retry the task. If the maximum retry count is -1, the Batch service retries the task without limit. Default is 0
        /// </summary>
        [Input("maxTaskRetryCount")]
        public Input<int>? MaxTaskRetryCount { get; set; }

        [Input("resourceFiles")]
        private InputList<Inputs.ResourceFileArgs>? _resourceFiles;
        public InputList<Inputs.ResourceFileArgs> ResourceFiles
        {
            get => _resourceFiles ?? (_resourceFiles = new InputList<Inputs.ResourceFileArgs>());
            set => _resourceFiles = value;
        }

        /// <summary>
        /// If omitted, the task runs as a non-administrative user unique to the task.
        /// </summary>
        [Input("userIdentity")]
        public Input<Inputs.UserIdentityArgs>? UserIdentity { get; set; }

        /// <summary>
        /// If true and the start task fails on a compute node, the Batch service retries the start task up to its maximum retry count (maxTaskRetryCount). If the task has still not completed successfully after all retries, then the Batch service marks the compute node unusable, and will not schedule tasks to it. This condition can be detected via the node state and scheduling error detail. If false, the Batch service will not wait for the start task to complete. In this case, other tasks can start executing on the compute node while the start task is still running; and even if the start task fails, new tasks will continue to be scheduled on the node. The default is true.
        /// </summary>
        [Input("waitForSuccess")]
        public Input<bool>? WaitForSuccess { get; set; }

        public StartTaskArgs()
        {
            MaxTaskRetryCount = 0;
        }
        public static new StartTaskArgs Empty => new StartTaskArgs();
    }
}
