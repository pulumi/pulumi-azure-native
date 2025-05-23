// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.AzureArcData
{
    /// <summary>
    /// The service account provisioning mode for this Active Directory connector.
    /// </summary>
    [EnumType]
    public readonly struct AccountProvisioningMode : IEquatable<AccountProvisioningMode>
    {
        private readonly string _value;

        private AccountProvisioningMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccountProvisioningMode Automatic { get; } = new AccountProvisioningMode("automatic");
        public static AccountProvisioningMode Manual { get; } = new AccountProvisioningMode("manual");

        public static bool operator ==(AccountProvisioningMode left, AccountProvisioningMode right) => left.Equals(right);
        public static bool operator !=(AccountProvisioningMode left, AccountProvisioningMode right) => !left.Equals(right);

        public static explicit operator string(AccountProvisioningMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccountProvisioningMode other && Equals(other);
        public bool Equals(AccountProvisioningMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The activation state of the license.
    /// </summary>
    [EnumType]
    public readonly struct ActivationState : IEquatable<ActivationState>
    {
        private readonly string _value;

        private ActivationState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ActivationState Activated { get; } = new ActivationState("Activated");
        public static ActivationState Deactivated { get; } = new ActivationState("Deactivated");

        public static bool operator ==(ActivationState left, ActivationState right) => left.Equals(right);
        public static bool operator !=(ActivationState left, ActivationState right) => !left.Equals(right);

        public static explicit operator string(ActivationState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ActivationState other && Equals(other);
        public bool Equals(ActivationState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The aggregation type to use for the numerical columns in the dataset.
    /// </summary>
    [EnumType]
    public readonly struct AggregationType : IEquatable<AggregationType>
    {
        private readonly string _value;

        private AggregationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AggregationType Average { get; } = new AggregationType("Average");
        public static AggregationType Minimum { get; } = new AggregationType("Minimum");
        public static AggregationType Maximum { get; } = new AggregationType("Maximum");
        public static AggregationType Sum { get; } = new AggregationType("Sum");
        public static AggregationType Count { get; } = new AggregationType("Count");

        public static bool operator ==(AggregationType left, AggregationType right) => left.Equals(right);
        public static bool operator !=(AggregationType left, AggregationType right) => !left.Equals(right);

        public static explicit operator string(AggregationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AggregationType other && Equals(other);
        public bool Equals(AggregationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The license type to apply for this managed instance.
    /// </summary>
    [EnumType]
    public readonly struct ArcSqlManagedInstanceLicenseType : IEquatable<ArcSqlManagedInstanceLicenseType>
    {
        private readonly string _value;

        private ArcSqlManagedInstanceLicenseType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ArcSqlManagedInstanceLicenseType BasePrice { get; } = new ArcSqlManagedInstanceLicenseType("BasePrice");
        public static ArcSqlManagedInstanceLicenseType LicenseIncluded { get; } = new ArcSqlManagedInstanceLicenseType("LicenseIncluded");
        public static ArcSqlManagedInstanceLicenseType DisasterRecovery { get; } = new ArcSqlManagedInstanceLicenseType("DisasterRecovery");

        public static bool operator ==(ArcSqlManagedInstanceLicenseType left, ArcSqlManagedInstanceLicenseType right) => left.Equals(right);
        public static bool operator !=(ArcSqlManagedInstanceLicenseType left, ArcSqlManagedInstanceLicenseType right) => !left.Equals(right);

        public static explicit operator string(ArcSqlManagedInstanceLicenseType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ArcSqlManagedInstanceLicenseType other && Equals(other);
        public bool Equals(ArcSqlManagedInstanceLicenseType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server license type.
    /// </summary>
    [EnumType]
    public readonly struct BillingPlan : IEquatable<BillingPlan>
    {
        private readonly string _value;

        private BillingPlan(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BillingPlan PAYG { get; } = new BillingPlan("PAYG");
        public static BillingPlan Paid { get; } = new BillingPlan("Paid");

        public static bool operator ==(BillingPlan left, BillingPlan right) => left.Equals(right);
        public static bool operator !=(BillingPlan left, BillingPlan right) => !left.Equals(right);

        public static explicit operator string(BillingPlan value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BillingPlan other && Equals(other);
        public bool Equals(BillingPlan other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. sourceDatabaseId and restorePointInTime must be specified.
    /// </summary>
    [EnumType]
    public readonly struct DatabaseCreateMode : IEquatable<DatabaseCreateMode>
    {
        private readonly string _value;

        private DatabaseCreateMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatabaseCreateMode Default { get; } = new DatabaseCreateMode("Default");
        public static DatabaseCreateMode PointInTimeRestore { get; } = new DatabaseCreateMode("PointInTimeRestore");

        public static bool operator ==(DatabaseCreateMode left, DatabaseCreateMode right) => left.Equals(right);
        public static bool operator !=(DatabaseCreateMode left, DatabaseCreateMode right) => !left.Equals(right);

        public static explicit operator string(DatabaseCreateMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatabaseCreateMode other && Equals(other);
        public bool Equals(DatabaseCreateMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// State of the database.
    /// </summary>
    [EnumType]
    public readonly struct DatabaseState : IEquatable<DatabaseState>
    {
        private readonly string _value;

        private DatabaseState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatabaseState Online { get; } = new DatabaseState("Online");
        public static DatabaseState Restoring { get; } = new DatabaseState("Restoring");
        public static DatabaseState Recovering { get; } = new DatabaseState("Recovering");
        public static DatabaseState RecoveryPending { get; } = new DatabaseState("RecoveryPending");
        public static DatabaseState Suspect { get; } = new DatabaseState("Suspect");
        public static DatabaseState Emergency { get; } = new DatabaseState("Emergency");
        public static DatabaseState Offline { get; } = new DatabaseState("Offline");
        public static DatabaseState Copying { get; } = new DatabaseState("Copying");
        public static DatabaseState OfflineSecondary { get; } = new DatabaseState("OfflineSecondary");

        public static bool operator ==(DatabaseState left, DatabaseState right) => left.Equals(right);
        public static bool operator !=(DatabaseState left, DatabaseState right) => !left.Equals(right);

        public static explicit operator string(DatabaseState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatabaseState other && Equals(other);
        public bool Equals(DatabaseState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server edition.
    /// </summary>
    [EnumType]
    public readonly struct EditionType : IEquatable<EditionType>
    {
        private readonly string _value;

        private EditionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EditionType Evaluation { get; } = new EditionType("Evaluation");
        public static EditionType Enterprise { get; } = new EditionType("Enterprise");
        public static EditionType Standard { get; } = new EditionType("Standard");
        public static EditionType Web { get; } = new EditionType("Web");
        public static EditionType Developer { get; } = new EditionType("Developer");
        public static EditionType Express { get; } = new EditionType("Express");
        public static EditionType Business_Intelligence { get; } = new EditionType("Business Intelligence");

        public static bool operator ==(EditionType left, EditionType right) => left.Equals(right);
        public static bool operator !=(EditionType left, EditionType right) => !left.Equals(right);

        public static explicit operator string(EditionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EditionType other && Equals(other);
        public bool Equals(EditionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the extended location.
    /// </summary>
    [EnumType]
    public readonly struct ExtendedLocationTypes : IEquatable<ExtendedLocationTypes>
    {
        private readonly string _value;

        private ExtendedLocationTypes(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ExtendedLocationTypes CustomLocation { get; } = new ExtendedLocationTypes("CustomLocation");

        public static bool operator ==(ExtendedLocationTypes left, ExtendedLocationTypes right) => left.Equals(right);
        public static bool operator !=(ExtendedLocationTypes left, ExtendedLocationTypes right) => !left.Equals(right);

        public static explicit operator string(ExtendedLocationTypes value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ExtendedLocationTypes other && Equals(other);
        public bool Equals(ExtendedLocationTypes other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The partner sync mode of the SQL managed instance.
    /// </summary>
    [EnumType]
    public readonly struct FailoverGroupPartnerSyncMode : IEquatable<FailoverGroupPartnerSyncMode>
    {
        private readonly string _value;

        private FailoverGroupPartnerSyncMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FailoverGroupPartnerSyncMode @Async { get; } = new FailoverGroupPartnerSyncMode("async");
        public static FailoverGroupPartnerSyncMode Sync { get; } = new FailoverGroupPartnerSyncMode("sync");

        public static bool operator ==(FailoverGroupPartnerSyncMode left, FailoverGroupPartnerSyncMode right) => left.Equals(right);
        public static bool operator !=(FailoverGroupPartnerSyncMode left, FailoverGroupPartnerSyncMode right) => !left.Equals(right);

        public static explicit operator string(FailoverGroupPartnerSyncMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FailoverGroupPartnerSyncMode other && Equals(other);
        public bool Equals(FailoverGroupPartnerSyncMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of host for Azure Arc SQL Server
    /// </summary>
    [EnumType]
    public readonly struct HostType : IEquatable<HostType>
    {
        private readonly string _value;

        private HostType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HostType Azure_Virtual_Machine { get; } = new HostType("Azure Virtual Machine");
        public static HostType Azure_VMWare_Virtual_Machine { get; } = new HostType("Azure VMWare Virtual Machine");
        public static HostType Azure_Kubernetes_Service { get; } = new HostType("Azure Kubernetes Service");
        public static HostType AWS_VMWare_Virtual_Machine { get; } = new HostType("AWS VMWare Virtual Machine");
        public static HostType AWS_Kubernetes_Service { get; } = new HostType("AWS Kubernetes Service");
        public static HostType GCP_VMWare_Virtual_Machine { get; } = new HostType("GCP VMWare Virtual Machine");
        public static HostType GCP_Kubernetes_Service { get; } = new HostType("GCP Kubernetes Service");
        public static HostType Container { get; } = new HostType("Container");
        public static HostType Virtual_Machine { get; } = new HostType("Virtual Machine");
        public static HostType Physical_Server { get; } = new HostType("Physical Server");
        public static HostType AWS_Virtual_Machine { get; } = new HostType("AWS Virtual Machine");
        public static HostType GCP_Virtual_Machine { get; } = new HostType("GCP Virtual Machine");
        public static HostType Other { get; } = new HostType("Other");

        public static bool operator ==(HostType left, HostType right) => left.Equals(right);
        public static bool operator !=(HostType left, HostType right) => !left.Equals(right);

        public static explicit operator string(HostType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HostType other && Equals(other);
        public bool Equals(HostType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The infrastructure the data controller is running on.
    /// </summary>
    [EnumType]
    public readonly struct Infrastructure : IEquatable<Infrastructure>
    {
        private readonly string _value;

        private Infrastructure(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Infrastructure Azure { get; } = new Infrastructure("azure");
        public static Infrastructure Gcp { get; } = new Infrastructure("gcp");
        public static Infrastructure Aws { get; } = new Infrastructure("aws");
        public static Infrastructure Alibaba { get; } = new Infrastructure("alibaba");
        public static Infrastructure Onpremises { get; } = new Infrastructure("onpremises");
        public static Infrastructure Other { get; } = new Infrastructure("other");

        public static bool operator ==(Infrastructure left, Infrastructure right) => left.Equals(right);
        public static bool operator !=(Infrastructure left, Infrastructure right) => !left.Equals(right);

        public static explicit operator string(Infrastructure value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Infrastructure other && Equals(other);
        public bool Equals(Infrastructure other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The role of the SQL managed instance in this failover group.
    /// </summary>
    [EnumType]
    public readonly struct InstanceFailoverGroupRole : IEquatable<InstanceFailoverGroupRole>
    {
        private readonly string _value;

        private InstanceFailoverGroupRole(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InstanceFailoverGroupRole Primary { get; } = new InstanceFailoverGroupRole("primary");
        public static InstanceFailoverGroupRole Secondary { get; } = new InstanceFailoverGroupRole("secondary");
        public static InstanceFailoverGroupRole Force_primary_allow_data_loss { get; } = new InstanceFailoverGroupRole("force-primary-allow-data-loss");
        public static InstanceFailoverGroupRole Force_secondary { get; } = new InstanceFailoverGroupRole("force-secondary");

        public static bool operator ==(InstanceFailoverGroupRole left, InstanceFailoverGroupRole right) => left.Equals(right);
        public static bool operator !=(InstanceFailoverGroupRole left, InstanceFailoverGroupRole right) => !left.Equals(right);

        public static explicit operator string(InstanceFailoverGroupRole value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceFailoverGroupRole other && Equals(other);
        public bool Equals(InstanceFailoverGroupRole other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This property represents the choice between SQL Server Core and ESU licenses.
    /// </summary>
    [EnumType]
    public readonly struct LicenseCategory : IEquatable<LicenseCategory>
    {
        private readonly string _value;

        private LicenseCategory(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LicenseCategory Core { get; } = new LicenseCategory("Core");

        public static bool operator ==(LicenseCategory left, LicenseCategory right) => left.Equals(right);
        public static bool operator !=(LicenseCategory left, LicenseCategory right) => !left.Equals(right);

        public static explicit operator string(LicenseCategory value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LicenseCategory other && Equals(other);
        public bool Equals(LicenseCategory other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This field is required to be implemented by the Resource Provider if the service has more than one tier.
    /// </summary>
    [EnumType]
    public readonly struct PostgresInstanceSkuTier : IEquatable<PostgresInstanceSkuTier>
    {
        private readonly string _value;

        private PostgresInstanceSkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PostgresInstanceSkuTier Hyperscale { get; } = new PostgresInstanceSkuTier("Hyperscale");

        public static bool operator ==(PostgresInstanceSkuTier left, PostgresInstanceSkuTier right) => left.Equals(right);
        public static bool operator !=(PostgresInstanceSkuTier left, PostgresInstanceSkuTier right) => !left.Equals(right);

        public static explicit operator string(PostgresInstanceSkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PostgresInstanceSkuTier other && Equals(other);
        public bool Equals(PostgresInstanceSkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Status of the database.
    /// </summary>
    [EnumType]
    public readonly struct RecoveryMode : IEquatable<RecoveryMode>
    {
        private readonly string _value;

        private RecoveryMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RecoveryMode Full { get; } = new RecoveryMode("Full");
        public static RecoveryMode Bulk_logged { get; } = new RecoveryMode("Bulk-logged");
        public static RecoveryMode Simple { get; } = new RecoveryMode("Simple");

        public static bool operator ==(RecoveryMode left, RecoveryMode right) => left.Equals(right);
        public static bool operator !=(RecoveryMode left, RecoveryMode right) => !left.Equals(right);

        public static explicit operator string(RecoveryMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RecoveryMode other && Equals(other);
        public bool Equals(RecoveryMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The Azure scope to which the license will apply.
    /// </summary>
    [EnumType]
    public readonly struct ScopeType : IEquatable<ScopeType>
    {
        private readonly string _value;

        private ScopeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScopeType Tenant { get; } = new ScopeType("Tenant");
        public static ScopeType Subscription { get; } = new ScopeType("Subscription");
        public static ScopeType ResourceGroup { get; } = new ScopeType("ResourceGroup");

        public static bool operator ==(ScopeType left, ScopeType right) => left.Equals(right);
        public static bool operator !=(ScopeType left, ScopeType right) => !left.Equals(right);

        public static explicit operator string(ScopeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScopeType other && Equals(other);
        public bool Equals(ScopeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the SKU.
    /// </summary>
    [EnumType]
    public readonly struct SqlManagedInstanceSkuName : IEquatable<SqlManagedInstanceSkuName>
    {
        private readonly string _value;

        private SqlManagedInstanceSkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlManagedInstanceSkuName VCore { get; } = new SqlManagedInstanceSkuName("vCore");

        public static bool operator ==(SqlManagedInstanceSkuName left, SqlManagedInstanceSkuName right) => left.Equals(right);
        public static bool operator !=(SqlManagedInstanceSkuName left, SqlManagedInstanceSkuName right) => !left.Equals(right);

        public static explicit operator string(SqlManagedInstanceSkuName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlManagedInstanceSkuName other && Equals(other);
        public bool Equals(SqlManagedInstanceSkuName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The pricing tier for the instance.
    /// </summary>
    [EnumType]
    public readonly struct SqlManagedInstanceSkuTier : IEquatable<SqlManagedInstanceSkuTier>
    {
        private readonly string _value;

        private SqlManagedInstanceSkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlManagedInstanceSkuTier GeneralPurpose { get; } = new SqlManagedInstanceSkuTier("GeneralPurpose");
        public static SqlManagedInstanceSkuTier BusinessCritical { get; } = new SqlManagedInstanceSkuTier("BusinessCritical");

        public static bool operator ==(SqlManagedInstanceSkuTier left, SqlManagedInstanceSkuTier right) => left.Equals(right);
        public static bool operator !=(SqlManagedInstanceSkuTier left, SqlManagedInstanceSkuTier right) => !left.Equals(right);

        public static explicit operator string(SqlManagedInstanceSkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlManagedInstanceSkuTier other && Equals(other);
        public bool Equals(SqlManagedInstanceSkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server version.
    /// </summary>
    [EnumType]
    public readonly struct SqlVersion : IEquatable<SqlVersion>
    {
        private readonly string _value;

        private SqlVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlVersion SQL_Server_2012 { get; } = new SqlVersion("SQL Server 2012");
        public static SqlVersion SQL_Server_2014 { get; } = new SqlVersion("SQL Server 2014");
        public static SqlVersion SQL_Server_2016 { get; } = new SqlVersion("SQL Server 2016");
        public static SqlVersion SQL_Server_2017 { get; } = new SqlVersion("SQL Server 2017");
        public static SqlVersion SQL_Server_2019 { get; } = new SqlVersion("SQL Server 2019");
        public static SqlVersion SQL_Server_2022 { get; } = new SqlVersion("SQL Server 2022");
        public static SqlVersion Unknown { get; } = new SqlVersion("Unknown");

        public static bool operator ==(SqlVersion left, SqlVersion right) => left.Equals(right);
        public static bool operator !=(SqlVersion left, SqlVersion right) => !left.Equals(right);

        public static explicit operator string(SqlVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlVersion other && Equals(other);
        public bool Equals(SqlVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The activation state of the license.
    /// </summary>
    [EnumType]
    public readonly struct State : IEquatable<State>
    {
        private readonly string _value;

        private State(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static State Inactive { get; } = new State("Inactive");
        public static State Active { get; } = new State("Active");
        public static State Terminated { get; } = new State("Terminated");

        public static bool operator ==(State left, State right) => left.Equals(right);
        public static bool operator !=(State left, State right) => !left.Equals(right);

        public static explicit operator string(State value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is State other && Equals(other);
        public bool Equals(State other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The SQL Server version the license covers.
    /// </summary>
    [EnumType]
    public readonly struct Version : IEquatable<Version>
    {
        private readonly string _value;

        private Version(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Version SQL_Server_2012 { get; } = new Version("SQL Server 2012");
        public static Version SQL_Server_2014 { get; } = new Version("SQL Server 2014");

        public static bool operator ==(Version left, Version right) => left.Equals(right);
        public static bool operator !=(Version left, Version right) => !left.Equals(right);

        public static explicit operator string(Version value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Version other && Equals(other);
        public bool Equals(Version other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
