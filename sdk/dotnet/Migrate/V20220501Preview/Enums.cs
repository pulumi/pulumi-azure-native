// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Migrate.V20220501Preview
{
    /// <summary>
    /// Gets or sets the status of automation artifacts.
    /// </summary>
    [EnumType]
    public readonly struct AutomationArtifactStatus : IEquatable<AutomationArtifactStatus>
    {
        private readonly string _value;

        private AutomationArtifactStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AutomationArtifactStatus NotGenerated { get; } = new AutomationArtifactStatus("NotGenerated");
        public static AutomationArtifactStatus Generated { get; } = new AutomationArtifactStatus("Generated");

        public static bool operator ==(AutomationArtifactStatus left, AutomationArtifactStatus right) => left.Equals(right);
        public static bool operator !=(AutomationArtifactStatus left, AutomationArtifactStatus right) => !left.Equals(right);

        public static explicit operator string(AutomationArtifactStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AutomationArtifactStatus other && Equals(other);
        public bool Equals(AutomationArtifactStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the configuration type.
    /// </summary>
    [EnumType]
    public readonly struct ConfigurationType : IEquatable<ConfigurationType>
    {
        private readonly string _value;

        private ConfigurationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConfigurationType IISConnectionString { get; } = new ConfigurationType("IISConnectionString");
        public static ConfigurationType IISAuthentication { get; } = new ConfigurationType("IISAuthentication");
        public static ConfigurationType ApacheTomcatContextResource { get; } = new ConfigurationType("ApacheTomcatContextResource");

        public static bool operator ==(ConfigurationType left, ConfigurationType right) => left.Equals(right);
        public static bool operator !=(ConfigurationType left, ConfigurationType right) => !left.Equals(right);

        public static explicit operator string(ConfigurationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConfigurationType other && Equals(other);
        public bool Equals(ConfigurationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the load balancer type.
    /// </summary>
    [EnumType]
    public readonly struct LoadBalancerType : IEquatable<LoadBalancerType>
    {
        private readonly string _value;

        private LoadBalancerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LoadBalancerType Private { get; } = new LoadBalancerType("Private");
        public static LoadBalancerType Public { get; } = new LoadBalancerType("Public");

        public static bool operator ==(LoadBalancerType left, LoadBalancerType right) => left.Equals(right);
        public static bool operator !=(LoadBalancerType left, LoadBalancerType right) => !left.Equals(right);

        public static explicit operator string(LoadBalancerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LoadBalancerType other && Equals(other);
        public bool Equals(LoadBalancerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct OperatingSystemType : IEquatable<OperatingSystemType>
    {
        private readonly string _value;

        private OperatingSystemType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OperatingSystemType Windows { get; } = new OperatingSystemType("Windows");
        public static OperatingSystemType Linux { get; } = new OperatingSystemType("Linux");

        public static bool operator ==(OperatingSystemType left, OperatingSystemType right) => left.Equals(right);
        public static bool operator !=(OperatingSystemType left, OperatingSystemType right) => !left.Equals(right);

        public static explicit operator string(OperatingSystemType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OperatingSystemType other && Equals(other);
        public bool Equals(OperatingSystemType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ResourceIdentityTypes : IEquatable<ResourceIdentityTypes>
    {
        private readonly string _value;

        private ResourceIdentityTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceIdentityTypes None { get; } = new ResourceIdentityTypes("None");
        public static ResourceIdentityTypes SystemAssigned { get; } = new ResourceIdentityTypes("SystemAssigned");
        public static ResourceIdentityTypes UserAssigned { get; } = new ResourceIdentityTypes("UserAssigned");

        public static bool operator ==(ResourceIdentityTypes left, ResourceIdentityTypes right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityTypes left, ResourceIdentityTypes right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityTypes other && Equals(other);
        public bool Equals(ResourceIdentityTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct SecretStoreType : IEquatable<SecretStoreType>
    {
        private readonly string _value;

        private SecretStoreType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecretStoreType None { get; } = new SecretStoreType("None");
        public static SecretStoreType KubeSecret { get; } = new SecretStoreType("KubeSecret");
        public static SecretStoreType KeyVaultSecret { get; } = new SecretStoreType("KeyVaultSecret");
        public static SecretStoreType AppServiceAppSettings { get; } = new SecretStoreType("AppServiceAppSettings");

        public static bool operator ==(SecretStoreType left, SecretStoreType right) => left.Equals(right);
        public static bool operator !=(SecretStoreType left, SecretStoreType right) => !left.Equals(right);

        public static explicit operator string(SecretStoreType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecretStoreType other && Equals(other);
        public bool Equals(SecretStoreType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the storage provider type on the target.
    /// Applicable when StorageProjectionType is not ContainerFileSystem.
    /// </summary>
    [EnumType]
    public readonly struct TargetHydrationStorageProviderType : IEquatable<TargetHydrationStorageProviderType>
    {
        private readonly string _value;

        private TargetHydrationStorageProviderType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetHydrationStorageProviderType AzureFileShare { get; } = new TargetHydrationStorageProviderType("AzureFileShare");

        public static bool operator ==(TargetHydrationStorageProviderType left, TargetHydrationStorageProviderType right) => left.Equals(right);
        public static bool operator !=(TargetHydrationStorageProviderType left, TargetHydrationStorageProviderType right) => !left.Equals(right);

        public static explicit operator string(TargetHydrationStorageProviderType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetHydrationStorageProviderType other && Equals(other);
        public bool Equals(TargetHydrationStorageProviderType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the target storage access type.
    /// </summary>
    [EnumType]
    public readonly struct TargetStorageAccessType : IEquatable<TargetStorageAccessType>
    {
        private readonly string _value;

        private TargetStorageAccessType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetStorageAccessType Shared { get; } = new TargetStorageAccessType("Shared");
        public static TargetStorageAccessType Exclusive { get; } = new TargetStorageAccessType("Exclusive");

        public static bool operator ==(TargetStorageAccessType left, TargetStorageAccessType right) => left.Equals(right);
        public static bool operator !=(TargetStorageAccessType left, TargetStorageAccessType right) => !left.Equals(right);

        public static explicit operator string(TargetStorageAccessType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetStorageAccessType other && Equals(other);
        public bool Equals(TargetStorageAccessType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the target projection type.
    /// </summary>
    [EnumType]
    public readonly struct TargetStorageProjectionType : IEquatable<TargetStorageProjectionType>
    {
        private readonly string _value;

        private TargetStorageProjectionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TargetStorageProjectionType ContainerFileSystem { get; } = new TargetStorageProjectionType("ContainerFileSystem");
        public static TargetStorageProjectionType PersistentVolume { get; } = new TargetStorageProjectionType("PersistentVolume");

        public static bool operator ==(TargetStorageProjectionType left, TargetStorageProjectionType right) => left.Equals(right);
        public static bool operator !=(TargetStorageProjectionType left, TargetStorageProjectionType right) => !left.Equals(right);

        public static explicit operator string(TargetStorageProjectionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TargetStorageProjectionType other && Equals(other);
        public bool Equals(TargetStorageProjectionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the deployment target platform.
    /// </summary>
    [EnumType]
    public readonly struct WorkloadDeploymentTarget : IEquatable<WorkloadDeploymentTarget>
    {
        private readonly string _value;

        private WorkloadDeploymentTarget(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkloadDeploymentTarget AzureKubernetesService { get; } = new WorkloadDeploymentTarget("AzureKubernetesService");
        public static WorkloadDeploymentTarget AzureAppServiceContainer { get; } = new WorkloadDeploymentTarget("AzureAppServiceContainer");
        public static WorkloadDeploymentTarget AzureAppServiceNative { get; } = new WorkloadDeploymentTarget("AzureAppServiceNative");

        public static bool operator ==(WorkloadDeploymentTarget left, WorkloadDeploymentTarget right) => left.Equals(right);
        public static bool operator !=(WorkloadDeploymentTarget left, WorkloadDeploymentTarget right) => !left.Equals(right);

        public static explicit operator string(WorkloadDeploymentTarget value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkloadDeploymentTarget other && Equals(other);
        public bool Equals(WorkloadDeploymentTarget other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the instance type.
    /// </summary>
    [EnumType]
    public readonly struct WorkloadDeploymentType : IEquatable<WorkloadDeploymentType>
    {
        private readonly string _value;

        private WorkloadDeploymentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkloadDeploymentType IISAKSWorkloadDeployment { get; } = new WorkloadDeploymentType("IISAKSWorkloadDeployment");
        public static WorkloadDeploymentType ApacheTomcatAKSWorkloadDeployment { get; } = new WorkloadDeploymentType("ApacheTomcatAKSWorkloadDeployment");

        public static bool operator ==(WorkloadDeploymentType left, WorkloadDeploymentType right) => left.Equals(right);
        public static bool operator !=(WorkloadDeploymentType left, WorkloadDeploymentType right) => !left.Equals(right);

        public static explicit operator string(WorkloadDeploymentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkloadDeploymentType other && Equals(other);
        public bool Equals(WorkloadDeploymentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the instance type.
    /// </summary>
    [EnumType]
    public readonly struct WorkloadType : IEquatable<WorkloadType>
    {
        private readonly string _value;

        private WorkloadType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkloadType IISWorkload { get; } = new WorkloadType("IISWorkload");
        public static WorkloadType ApacheTomcatWorkload { get; } = new WorkloadType("ApacheTomcatWorkload");

        public static bool operator ==(WorkloadType left, WorkloadType right) => left.Equals(right);
        public static bool operator !=(WorkloadType left, WorkloadType right) => !left.Equals(right);

        public static explicit operator string(WorkloadType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkloadType other && Equals(other);
        public bool Equals(WorkloadType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
