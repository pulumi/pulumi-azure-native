// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.HDInsight
{
    /// <summary>
    /// User to specify which type of Autoscale to be implemented - Scheduled Based or Load Based.
    /// </summary>
    [EnumType]
    public readonly struct AutoscaleType : IEquatable<AutoscaleType>
    {
        private readonly string _value;

        private AutoscaleType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AutoscaleType ScheduleBased { get; } = new AutoscaleType("ScheduleBased");
        public static AutoscaleType LoadBased { get; } = new AutoscaleType("LoadBased");

        public static bool operator ==(AutoscaleType left, AutoscaleType right) => left.Equals(right);
        public static bool operator !=(AutoscaleType left, AutoscaleType right) => !left.Equals(right);

        public static explicit operator string(AutoscaleType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AutoscaleType other && Equals(other);
        public bool Equals(AutoscaleType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The comparison operator.
    /// </summary>
    [EnumType]
    public readonly struct ComparisonOperator : IEquatable<ComparisonOperator>
    {
        private readonly string _value;

        private ComparisonOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComparisonOperator GreaterThan { get; } = new ComparisonOperator("greaterThan");
        public static ComparisonOperator GreaterThanOrEqual { get; } = new ComparisonOperator("greaterThanOrEqual");
        public static ComparisonOperator LessThan { get; } = new ComparisonOperator("lessThan");
        public static ComparisonOperator LessThanOrEqual { get; } = new ComparisonOperator("lessThanOrEqual");

        public static bool operator ==(ComparisonOperator left, ComparisonOperator right) => left.Equals(right);
        public static bool operator !=(ComparisonOperator left, ComparisonOperator right) => !left.Equals(right);

        public static explicit operator string(ComparisonOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComparisonOperator other && Equals(other);
        public bool Equals(ComparisonOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This property indicates if the content is encoded and is case-insensitive. Please set the value to base64 if the content is base64 encoded. Set it to none or skip it if the content is plain text.
    /// </summary>
    [EnumType]
    public readonly struct ContentEncoding : IEquatable<ContentEncoding>
    {
        private readonly string _value;

        private ContentEncoding(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ContentEncoding Base64 { get; } = new ContentEncoding("Base64");
        public static ContentEncoding None { get; } = new ContentEncoding("None");

        public static bool operator ==(ContentEncoding left, ContentEncoding right) => left.Equals(right);
        public static bool operator !=(ContentEncoding left, ContentEncoding right) => !left.Equals(right);

        public static explicit operator string(ContentEncoding value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContentEncoding other && Equals(other);
        public bool Equals(ContentEncoding other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Managed Disk Type.
    /// </summary>
    [EnumType]
    public readonly struct DataDiskType : IEquatable<DataDiskType>
    {
        private readonly string _value;

        private DataDiskType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataDiskType Standard_HDD_LRS { get; } = new DataDiskType("Standard_HDD_LRS");
        public static DataDiskType Standard_SSD_LRS { get; } = new DataDiskType("Standard_SSD_LRS");
        public static DataDiskType Standard_SSD_ZRS { get; } = new DataDiskType("Standard_SSD_ZRS");
        public static DataDiskType Premium_SSD_LRS { get; } = new DataDiskType("Premium_SSD_LRS");
        public static DataDiskType Premium_SSD_ZRS { get; } = new DataDiskType("Premium_SSD_ZRS");
        public static DataDiskType Premium_SSD_v2_LRS { get; } = new DataDiskType("Premium_SSD_v2_LRS");

        public static bool operator ==(DataDiskType left, DataDiskType right) => left.Equals(right);
        public static bool operator !=(DataDiskType left, DataDiskType right) => !left.Equals(right);

        public static explicit operator string(DataDiskType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataDiskType other && Equals(other);
        public bool Equals(DataDiskType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct DaysOfWeek : IEquatable<DaysOfWeek>
    {
        private readonly string _value;

        private DaysOfWeek(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DaysOfWeek Monday { get; } = new DaysOfWeek("Monday");
        public static DaysOfWeek Tuesday { get; } = new DaysOfWeek("Tuesday");
        public static DaysOfWeek Wednesday { get; } = new DaysOfWeek("Wednesday");
        public static DaysOfWeek Thursday { get; } = new DaysOfWeek("Thursday");
        public static DaysOfWeek Friday { get; } = new DaysOfWeek("Friday");
        public static DaysOfWeek Saturday { get; } = new DaysOfWeek("Saturday");
        public static DaysOfWeek Sunday { get; } = new DaysOfWeek("Sunday");

        public static bool operator ==(DaysOfWeek left, DaysOfWeek right) => left.Equals(right);
        public static bool operator !=(DaysOfWeek left, DaysOfWeek right) => !left.Equals(right);

        public static explicit operator string(DaysOfWeek value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DaysOfWeek other && Equals(other);
        public bool Equals(DaysOfWeek other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The authentication mode to connect to your Hive metastore database. More details: https://learn.microsoft.com/en-us/azure/azure-sql/database/logins-create-manage?view=azuresql#authentication-and-authorization
    /// </summary>
    [EnumType]
    public readonly struct DbConnectionAuthenticationMode : IEquatable<DbConnectionAuthenticationMode>
    {
        private readonly string _value;

        private DbConnectionAuthenticationMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The password-based authentication to connect to your Hive metastore database. 
        /// </summary>
        public static DbConnectionAuthenticationMode SqlAuth { get; } = new DbConnectionAuthenticationMode("SqlAuth");
        /// <summary>
        /// The managed-identity-based authentication to connect to your Hive metastore database. 
        /// </summary>
        public static DbConnectionAuthenticationMode IdentityAuth { get; } = new DbConnectionAuthenticationMode("IdentityAuth");

        public static bool operator ==(DbConnectionAuthenticationMode left, DbConnectionAuthenticationMode right) => left.Equals(right);
        public static bool operator !=(DbConnectionAuthenticationMode left, DbConnectionAuthenticationMode right) => !left.Equals(right);

        public static explicit operator string(DbConnectionAuthenticationMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DbConnectionAuthenticationMode other && Equals(other);
        public bool Equals(DbConnectionAuthenticationMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A string property that indicates the deployment mode of Flink cluster. It can have one of the following enum values =&gt; Application, Session. Default value is Session
    /// </summary>
    [EnumType]
    public readonly struct DeploymentMode : IEquatable<DeploymentMode>
    {
        private readonly string _value;

        private DeploymentMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentMode Application { get; } = new DeploymentMode("Application");
        public static DeploymentMode Session { get; } = new DeploymentMode("Session");

        public static bool operator ==(DeploymentMode left, DeploymentMode right) => left.Equals(right);
        public static bool operator !=(DeploymentMode left, DeploymentMode right) => !left.Equals(right);

        public static explicit operator string(DeploymentMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentMode other && Equals(other);
        public bool Equals(DeploymentMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The directory type.
    /// </summary>
    [EnumType]
    public readonly struct DirectoryType : IEquatable<DirectoryType>
    {
        private readonly string _value;

        private DirectoryType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DirectoryType ActiveDirectory { get; } = new DirectoryType("ActiveDirectory");

        public static bool operator ==(DirectoryType left, DirectoryType right) => left.Equals(right);
        public static bool operator !=(DirectoryType left, DirectoryType right) => !left.Equals(right);

        public static explicit operator string(DirectoryType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DirectoryType other && Equals(other);
        public bool Equals(DirectoryType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Algorithm identifier for encryption, default RSA-OAEP.
    /// </summary>
    [EnumType]
    public readonly struct JsonWebKeyEncryptionAlgorithm : IEquatable<JsonWebKeyEncryptionAlgorithm>
    {
        private readonly string _value;

        private JsonWebKeyEncryptionAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static JsonWebKeyEncryptionAlgorithm RSA_OAEP { get; } = new JsonWebKeyEncryptionAlgorithm("RSA-OAEP");
        public static JsonWebKeyEncryptionAlgorithm RSA_OAEP_256 { get; } = new JsonWebKeyEncryptionAlgorithm("RSA-OAEP-256");
        public static JsonWebKeyEncryptionAlgorithm RSA1_5 { get; } = new JsonWebKeyEncryptionAlgorithm("RSA1_5");

        public static bool operator ==(JsonWebKeyEncryptionAlgorithm left, JsonWebKeyEncryptionAlgorithm right) => left.Equals(right);
        public static bool operator !=(JsonWebKeyEncryptionAlgorithm left, JsonWebKeyEncryptionAlgorithm right) => !left.Equals(right);

        public static explicit operator string(JsonWebKeyEncryptionAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JsonWebKeyEncryptionAlgorithm other && Equals(other);
        public bool Equals(JsonWebKeyEncryptionAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of key vault object: secret, key or certificate.
    /// </summary>
    [EnumType]
    public readonly struct KeyVaultObjectType : IEquatable<KeyVaultObjectType>
    {
        private readonly string _value;

        private KeyVaultObjectType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static KeyVaultObjectType Key { get; } = new KeyVaultObjectType("Key");
        public static KeyVaultObjectType Secret { get; } = new KeyVaultObjectType("Secret");
        public static KeyVaultObjectType Certificate { get; } = new KeyVaultObjectType("Certificate");

        public static bool operator ==(KeyVaultObjectType left, KeyVaultObjectType right) => left.Equals(right);
        public static bool operator !=(KeyVaultObjectType left, KeyVaultObjectType right) => !left.Equals(right);

        public static explicit operator string(KeyVaultObjectType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeyVaultObjectType other && Equals(other);
        public bool Equals(KeyVaultObjectType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of managed identity.
    /// </summary>
    [EnumType]
    public readonly struct ManagedIdentityType : IEquatable<ManagedIdentityType>
    {
        private readonly string _value;

        private ManagedIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ManagedIdentityType Cluster { get; } = new ManagedIdentityType("cluster");
        public static ManagedIdentityType User { get; } = new ManagedIdentityType("user");
        public static ManagedIdentityType @Internal { get; } = new ManagedIdentityType("internal");

        public static bool operator ==(ManagedIdentityType left, ManagedIdentityType right) => left.Equals(right);
        public static bool operator !=(ManagedIdentityType left, ManagedIdentityType right) => !left.Equals(right);

        public static explicit operator string(ManagedIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedIdentityType other && Equals(other);
        public bool Equals(ManagedIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The authentication mode to connect to your Hive metastore database. More details: https://learn.microsoft.com/en-us/azure/azure-sql/database/logins-create-manage?view=azuresql#authentication-and-authorization
    /// </summary>
    [EnumType]
    public readonly struct MetastoreDbConnectionAuthenticationMode : IEquatable<MetastoreDbConnectionAuthenticationMode>
    {
        private readonly string _value;

        private MetastoreDbConnectionAuthenticationMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The password-based authentication to connect to your Hive metastore database. 
        /// </summary>
        public static MetastoreDbConnectionAuthenticationMode SqlAuth { get; } = new MetastoreDbConnectionAuthenticationMode("SqlAuth");
        /// <summary>
        /// The managed-identity-based authentication to connect to your Hive metastore database. 
        /// </summary>
        public static MetastoreDbConnectionAuthenticationMode IdentityAuth { get; } = new MetastoreDbConnectionAuthenticationMode("IdentityAuth");

        public static bool operator ==(MetastoreDbConnectionAuthenticationMode left, MetastoreDbConnectionAuthenticationMode right) => left.Equals(right);
        public static bool operator !=(MetastoreDbConnectionAuthenticationMode left, MetastoreDbConnectionAuthenticationMode right) => !left.Equals(right);

        public static explicit operator string(MetastoreDbConnectionAuthenticationMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MetastoreDbConnectionAuthenticationMode other && Equals(other);
        public bool Equals(MetastoreDbConnectionAuthenticationMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of operating system.
    /// </summary>
    [EnumType]
    public readonly struct OSType : IEquatable<OSType>
    {
        private readonly string _value;

        private OSType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OSType Windows { get; } = new OSType("Windows");
        public static OSType Linux { get; } = new OSType("Linux");

        public static bool operator ==(OSType left, OSType right) => left.Equals(right);
        public static bool operator !=(OSType left, OSType right) => !left.Equals(right);

        public static explicit operator string(OSType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OSType other && Equals(other);
        public bool Equals(OSType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A value to describe how the outbound dependencies of a HDInsight cluster are managed. 'Managed' means that the outbound dependencies are managed by the HDInsight service. 'External' means that the outbound dependencies are managed by a customer specific solution.
    /// </summary>
    [EnumType]
    public readonly struct OutboundDependenciesManagedType : IEquatable<OutboundDependenciesManagedType>
    {
        private readonly string _value;

        private OutboundDependenciesManagedType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OutboundDependenciesManagedType Managed { get; } = new OutboundDependenciesManagedType("Managed");
        public static OutboundDependenciesManagedType External { get; } = new OutboundDependenciesManagedType("External");

        public static bool operator ==(OutboundDependenciesManagedType left, OutboundDependenciesManagedType right) => left.Equals(right);
        public static bool operator !=(OutboundDependenciesManagedType left, OutboundDependenciesManagedType right) => !left.Equals(right);

        public static explicit operator string(OutboundDependenciesManagedType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OutboundDependenciesManagedType other && Equals(other);
        public bool Equals(OutboundDependenciesManagedType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This can only be set at cluster pool creation time and cannot be changed later. 
    /// </summary>
    [EnumType]
    public readonly struct OutboundType : IEquatable<OutboundType>
    {
        private readonly string _value;

        private OutboundType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The load balancer is used for egress through an AKS assigned public IP. This supports Kubernetes services of type 'loadBalancer'. 
        /// </summary>
        public static OutboundType LoadBalancer { get; } = new OutboundType("loadBalancer");
        /// <summary>
        /// Egress paths must be defined by the user. This is an advanced scenario and requires proper network configuration. 
        /// </summary>
        public static OutboundType UserDefinedRouting { get; } = new OutboundType("userDefinedRouting");

        public static bool operator ==(OutboundType left, OutboundType right) => left.Equals(right);
        public static bool operator !=(OutboundType left, OutboundType right) => !left.Equals(right);

        public static explicit operator string(OutboundType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OutboundType other && Equals(other);
        public bool Equals(OutboundType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The method that private IP address is allocated.
    /// </summary>
    [EnumType]
    public readonly struct PrivateIPAllocationMethod : IEquatable<PrivateIPAllocationMethod>
    {
        private readonly string _value;

        private PrivateIPAllocationMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateIPAllocationMethod @Dynamic { get; } = new PrivateIPAllocationMethod("dynamic");
        public static PrivateIPAllocationMethod @Static { get; } = new PrivateIPAllocationMethod("static");

        public static bool operator ==(PrivateIPAllocationMethod left, PrivateIPAllocationMethod right) => left.Equals(right);
        public static bool operator !=(PrivateIPAllocationMethod left, PrivateIPAllocationMethod right) => !left.Equals(right);

        public static explicit operator string(PrivateIPAllocationMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateIPAllocationMethod other && Equals(other);
        public bool Equals(PrivateIPAllocationMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether or not private link is enabled.
    /// </summary>
    [EnumType]
    public readonly struct PrivateLink : IEquatable<PrivateLink>
    {
        private readonly string _value;

        private PrivateLink(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateLink Disabled { get; } = new PrivateLink("Disabled");
        public static PrivateLink Enabled { get; } = new PrivateLink("Enabled");

        public static bool operator ==(PrivateLink left, PrivateLink right) => left.Equals(right);
        public static bool operator !=(PrivateLink left, PrivateLink right) => !left.Equals(right);

        public static explicit operator string(PrivateLink value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateLink other && Equals(other);
        public bool Equals(PrivateLink other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The concrete private link service connection.
    /// </summary>
    [EnumType]
    public readonly struct PrivateLinkServiceConnectionStatus : IEquatable<PrivateLinkServiceConnectionStatus>
    {
        private readonly string _value;

        private PrivateLinkServiceConnectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateLinkServiceConnectionStatus Approved { get; } = new PrivateLinkServiceConnectionStatus("Approved");
        public static PrivateLinkServiceConnectionStatus Rejected { get; } = new PrivateLinkServiceConnectionStatus("Rejected");
        public static PrivateLinkServiceConnectionStatus Pending { get; } = new PrivateLinkServiceConnectionStatus("Pending");
        public static PrivateLinkServiceConnectionStatus Removed { get; } = new PrivateLinkServiceConnectionStatus("Removed");

        public static bool operator ==(PrivateLinkServiceConnectionStatus left, PrivateLinkServiceConnectionStatus right) => left.Equals(right);
        public static bool operator !=(PrivateLinkServiceConnectionStatus left, PrivateLinkServiceConnectionStatus right) => !left.Equals(right);

        public static explicit operator string(PrivateLinkServiceConnectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateLinkServiceConnectionStatus other && Equals(other);
        public bool Equals(PrivateLinkServiceConnectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// User &amp; groups can be synced automatically or via a static list that's refreshed.
    /// </summary>
    [EnumType]
    public readonly struct RangerUsersyncMode : IEquatable<RangerUsersyncMode>
    {
        private readonly string _value;

        private RangerUsersyncMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RangerUsersyncMode @Static { get; } = new RangerUsersyncMode("static");
        public static RangerUsersyncMode Automatic { get; } = new RangerUsersyncMode("automatic");

        public static bool operator ==(RangerUsersyncMode left, RangerUsersyncMode right) => left.Equals(right);
        public static bool operator !=(RangerUsersyncMode left, RangerUsersyncMode right) => !left.Equals(right);

        public static explicit operator string(RangerUsersyncMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RangerUsersyncMode other && Equals(other);
        public bool Equals(RangerUsersyncMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of identity used for the cluster. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities.
    /// </summary>
    [EnumType]
    public readonly struct ResourceIdentityType : IEquatable<ResourceIdentityType>
    {
        private readonly string _value;

        private ResourceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceIdentityType SystemAssigned { get; } = new ResourceIdentityType("SystemAssigned");
        public static ResourceIdentityType UserAssigned { get; } = new ResourceIdentityType("UserAssigned");
        public static ResourceIdentityType SystemAssigned_UserAssigned { get; } = new ResourceIdentityType("SystemAssigned, UserAssigned");
        public static ResourceIdentityType None { get; } = new ResourceIdentityType("None");

        public static bool operator ==(ResourceIdentityType left, ResourceIdentityType right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityType left, ResourceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityType other && Equals(other);
        public bool Equals(ResourceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The direction for the resource provider connection.
    /// </summary>
    [EnumType]
    public readonly struct ResourceProviderConnection : IEquatable<ResourceProviderConnection>
    {
        private readonly string _value;

        private ResourceProviderConnection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceProviderConnection Inbound { get; } = new ResourceProviderConnection("Inbound");
        public static ResourceProviderConnection Outbound { get; } = new ResourceProviderConnection("Outbound");

        public static bool operator ==(ResourceProviderConnection left, ResourceProviderConnection right) => left.Equals(right);
        public static bool operator !=(ResourceProviderConnection left, ResourceProviderConnection right) => !left.Equals(right);

        public static explicit operator string(ResourceProviderConnection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceProviderConnection other && Equals(other);
        public bool Equals(ResourceProviderConnection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The action type.
    /// </summary>
    [EnumType]
    public readonly struct ScaleActionType : IEquatable<ScaleActionType>
    {
        private readonly string _value;

        private ScaleActionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScaleActionType Scaleup { get; } = new ScaleActionType("scaleup");
        public static ScaleActionType Scaledown { get; } = new ScaleActionType("scaledown");

        public static bool operator ==(ScaleActionType left, ScaleActionType right) => left.Equals(right);
        public static bool operator !=(ScaleActionType left, ScaleActionType right) => !left.Equals(right);

        public static explicit operator string(ScaleActionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScaleActionType other && Equals(other);
        public bool Equals(ScaleActionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScheduleDay : IEquatable<ScheduleDay>
    {
        private readonly string _value;

        private ScheduleDay(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScheduleDay Sunday { get; } = new ScheduleDay("Sunday");
        public static ScheduleDay Monday { get; } = new ScheduleDay("Monday");
        public static ScheduleDay Tuesday { get; } = new ScheduleDay("Tuesday");
        public static ScheduleDay Wednesday { get; } = new ScheduleDay("Wednesday");
        public static ScheduleDay Thursday { get; } = new ScheduleDay("Thursday");
        public static ScheduleDay Friday { get; } = new ScheduleDay("Friday");
        public static ScheduleDay Saturday { get; } = new ScheduleDay("Saturday");

        public static bool operator ==(ScheduleDay left, ScheduleDay right) => left.Equals(right);
        public static bool operator !=(ScheduleDay left, ScheduleDay right) => !left.Equals(right);

        public static explicit operator string(ScheduleDay value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScheduleDay other && Equals(other);
        public bool Equals(ScheduleDay other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The cluster tier.
    /// </summary>
    [EnumType]
    public readonly struct Tier : IEquatable<Tier>
    {
        private readonly string _value;

        private Tier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Tier Standard { get; } = new Tier("Standard");
        public static Tier Premium { get; } = new Tier("Premium");

        public static bool operator ==(Tier left, Tier right) => left.Equals(right);
        public static bool operator !=(Tier left, Tier right) => !left.Equals(right);

        public static explicit operator string(Tier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Tier other && Equals(other);
        public bool Equals(Tier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A string property that indicates the upgrade mode to be performed on the Flink job. It can have one of the following enum values =&gt; STATELESS_UPDATE, UPDATE, LAST_STATE_UPDATE.
    /// </summary>
    [EnumType]
    public readonly struct UpgradeMode : IEquatable<UpgradeMode>
    {
        private readonly string _value;

        private UpgradeMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UpgradeMode STATELESS_UPDATE { get; } = new UpgradeMode("STATELESS_UPDATE");
        public static UpgradeMode UPDATE { get; } = new UpgradeMode("UPDATE");
        public static UpgradeMode LAST_STATE_UPDATE { get; } = new UpgradeMode("LAST_STATE_UPDATE");

        public static bool operator ==(UpgradeMode left, UpgradeMode right) => left.Equals(right);
        public static bool operator !=(UpgradeMode left, UpgradeMode right) => !left.Equals(right);

        public static explicit operator string(UpgradeMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UpgradeMode other && Equals(other);
        public bool Equals(UpgradeMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
