// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.SecurityInsights.V20211001Preview
{
    /// <summary>
    /// Alert detail
    /// </summary>
    [EnumType]
    public readonly struct AlertDetail : IEquatable<AlertDetail>
    {
        private readonly string _value;

        private AlertDetail(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Alert display name
        /// </summary>
        public static AlertDetail DisplayName { get; } = new AlertDetail("DisplayName");
        /// <summary>
        /// Alert severity
        /// </summary>
        public static AlertDetail Severity { get; } = new AlertDetail("Severity");

        public static bool operator ==(AlertDetail left, AlertDetail right) => left.Equals(right);
        public static bool operator !=(AlertDetail left, AlertDetail right) => !left.Equals(right);

        public static explicit operator string(AlertDetail value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertDetail other && Equals(other);
        public bool Equals(AlertDetail other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of the alert rule
    /// </summary>
    [EnumType]
    public readonly struct AlertRuleKind : IEquatable<AlertRuleKind>
    {
        private readonly string _value;

        private AlertRuleKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertRuleKind Scheduled { get; } = new AlertRuleKind("Scheduled");
        public static AlertRuleKind MicrosoftSecurityIncidentCreation { get; } = new AlertRuleKind("MicrosoftSecurityIncidentCreation");
        public static AlertRuleKind Fusion { get; } = new AlertRuleKind("Fusion");
        public static AlertRuleKind MLBehaviorAnalytics { get; } = new AlertRuleKind("MLBehaviorAnalytics");
        public static AlertRuleKind ThreatIntelligence { get; } = new AlertRuleKind("ThreatIntelligence");
        public static AlertRuleKind NRT { get; } = new AlertRuleKind("NRT");

        public static bool operator ==(AlertRuleKind left, AlertRuleKind right) => left.Equals(right);
        public static bool operator !=(AlertRuleKind left, AlertRuleKind right) => !left.Equals(right);

        public static explicit operator string(AlertRuleKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertRuleKind other && Equals(other);
        public bool Equals(AlertRuleKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The severity for alerts created by this alert rule.
    /// </summary>
    [EnumType]
    public readonly struct AlertSeverity : IEquatable<AlertSeverity>
    {
        private readonly string _value;

        private AlertSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// High severity
        /// </summary>
        public static AlertSeverity High { get; } = new AlertSeverity("High");
        /// <summary>
        /// Medium severity
        /// </summary>
        public static AlertSeverity Medium { get; } = new AlertSeverity("Medium");
        /// <summary>
        /// Low severity
        /// </summary>
        public static AlertSeverity Low { get; } = new AlertSeverity("Low");
        /// <summary>
        /// Informational severity
        /// </summary>
        public static AlertSeverity Informational { get; } = new AlertSeverity("Informational");

        public static bool operator ==(AlertSeverity left, AlertSeverity right) => left.Equals(right);
        public static bool operator !=(AlertSeverity left, AlertSeverity right) => !left.Equals(right);

        public static explicit operator string(AlertSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertSeverity other && Equals(other);
        public bool Equals(AlertSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The severity for alerts created by this alert rule.
    /// </summary>
    [EnumType]
    public readonly struct AttackTactic : IEquatable<AttackTactic>
    {
        private readonly string _value;

        private AttackTactic(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AttackTactic Reconnaissance { get; } = new AttackTactic("Reconnaissance");
        public static AttackTactic ResourceDevelopment { get; } = new AttackTactic("ResourceDevelopment");
        public static AttackTactic InitialAccess { get; } = new AttackTactic("InitialAccess");
        public static AttackTactic Execution { get; } = new AttackTactic("Execution");
        public static AttackTactic Persistence { get; } = new AttackTactic("Persistence");
        public static AttackTactic PrivilegeEscalation { get; } = new AttackTactic("PrivilegeEscalation");
        public static AttackTactic DefenseEvasion { get; } = new AttackTactic("DefenseEvasion");
        public static AttackTactic CredentialAccess { get; } = new AttackTactic("CredentialAccess");
        public static AttackTactic Discovery { get; } = new AttackTactic("Discovery");
        public static AttackTactic LateralMovement { get; } = new AttackTactic("LateralMovement");
        public static AttackTactic Collection { get; } = new AttackTactic("Collection");
        public static AttackTactic Exfiltration { get; } = new AttackTactic("Exfiltration");
        public static AttackTactic CommandAndControl { get; } = new AttackTactic("CommandAndControl");
        public static AttackTactic Impact { get; } = new AttackTactic("Impact");
        public static AttackTactic PreAttack { get; } = new AttackTactic("PreAttack");
        public static AttackTactic ImpairProcessControl { get; } = new AttackTactic("ImpairProcessControl");
        public static AttackTactic InhibitResponseFunction { get; } = new AttackTactic("InhibitResponseFunction");

        public static bool operator ==(AttackTactic left, AttackTactic right) => left.Equals(right);
        public static bool operator !=(AttackTactic left, AttackTactic right) => !left.Equals(right);

        public static explicit operator string(AttackTactic value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AttackTactic other && Equals(other);
        public bool Equals(AttackTactic other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The V3 type of the mapped entity
    /// </summary>
    [EnumType]
    public readonly struct EntityMappingType : IEquatable<EntityMappingType>
    {
        private readonly string _value;

        private EntityMappingType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// User account entity type
        /// </summary>
        public static EntityMappingType Account { get; } = new EntityMappingType("Account");
        /// <summary>
        /// Host entity type
        /// </summary>
        public static EntityMappingType Host { get; } = new EntityMappingType("Host");
        /// <summary>
        /// IP address entity type
        /// </summary>
        public static EntityMappingType IP { get; } = new EntityMappingType("IP");
        /// <summary>
        /// Malware entity type
        /// </summary>
        public static EntityMappingType Malware { get; } = new EntityMappingType("Malware");
        /// <summary>
        /// System file entity type
        /// </summary>
        public static EntityMappingType File { get; } = new EntityMappingType("File");
        /// <summary>
        /// Process entity type
        /// </summary>
        public static EntityMappingType Process { get; } = new EntityMappingType("Process");
        /// <summary>
        /// Cloud app entity type
        /// </summary>
        public static EntityMappingType CloudApplication { get; } = new EntityMappingType("CloudApplication");
        /// <summary>
        /// DNS entity type
        /// </summary>
        public static EntityMappingType DNS { get; } = new EntityMappingType("DNS");
        /// <summary>
        /// Azure resource entity type
        /// </summary>
        public static EntityMappingType AzureResource { get; } = new EntityMappingType("AzureResource");
        /// <summary>
        /// File-hash entity type
        /// </summary>
        public static EntityMappingType FileHash { get; } = new EntityMappingType("FileHash");
        /// <summary>
        /// Registry key entity type
        /// </summary>
        public static EntityMappingType RegistryKey { get; } = new EntityMappingType("RegistryKey");
        /// <summary>
        /// Registry value entity type
        /// </summary>
        public static EntityMappingType RegistryValue { get; } = new EntityMappingType("RegistryValue");
        /// <summary>
        /// Security group entity type
        /// </summary>
        public static EntityMappingType SecurityGroup { get; } = new EntityMappingType("SecurityGroup");
        /// <summary>
        /// URL entity type
        /// </summary>
        public static EntityMappingType URL { get; } = new EntityMappingType("URL");
        /// <summary>
        /// Mailbox entity type
        /// </summary>
        public static EntityMappingType Mailbox { get; } = new EntityMappingType("Mailbox");
        /// <summary>
        /// Mail cluster entity type
        /// </summary>
        public static EntityMappingType MailCluster { get; } = new EntityMappingType("MailCluster");
        /// <summary>
        /// Mail message entity type
        /// </summary>
        public static EntityMappingType MailMessage { get; } = new EntityMappingType("MailMessage");
        /// <summary>
        /// Submission mail entity type
        /// </summary>
        public static EntityMappingType SubmissionMail { get; } = new EntityMappingType("SubmissionMail");

        public static bool operator ==(EntityMappingType left, EntityMappingType right) => left.Equals(right);
        public static bool operator !=(EntityMappingType left, EntityMappingType right) => !left.Equals(right);

        public static explicit operator string(EntityMappingType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityMappingType other && Equals(other);
        public bool Equals(EntityMappingType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The entity query kind
    /// </summary>
    [EnumType]
    public readonly struct EntityTimelineKind : IEquatable<EntityTimelineKind>
    {
        private readonly string _value;

        private EntityTimelineKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// activity
        /// </summary>
        public static EntityTimelineKind Activity { get; } = new EntityTimelineKind("Activity");
        /// <summary>
        /// bookmarks
        /// </summary>
        public static EntityTimelineKind Bookmark { get; } = new EntityTimelineKind("Bookmark");
        /// <summary>
        /// security alerts
        /// </summary>
        public static EntityTimelineKind SecurityAlert { get; } = new EntityTimelineKind("SecurityAlert");

        public static bool operator ==(EntityTimelineKind left, EntityTimelineKind right) => left.Equals(right);
        public static bool operator !=(EntityTimelineKind left, EntityTimelineKind right) => !left.Equals(right);

        public static explicit operator string(EntityTimelineKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityTimelineKind other && Equals(other);
        public bool Equals(EntityTimelineKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Grouping matching method. When method is Selected at least one of groupByEntities, groupByAlertDetails, groupByCustomDetails must be provided and not empty.
    /// </summary>
    [EnumType]
    public readonly struct MatchingMethod : IEquatable<MatchingMethod>
    {
        private readonly string _value;

        private MatchingMethod(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Grouping alerts into a single incident if all the entities match
        /// </summary>
        public static MatchingMethod AllEntities { get; } = new MatchingMethod("AllEntities");
        /// <summary>
        /// Grouping any alerts triggered by this rule into a single incident
        /// </summary>
        public static MatchingMethod AnyAlert { get; } = new MatchingMethod("AnyAlert");
        /// <summary>
        /// Grouping alerts into a single incident if the selected entities, custom details and alert details match
        /// </summary>
        public static MatchingMethod Selected { get; } = new MatchingMethod("Selected");

        public static bool operator ==(MatchingMethod left, MatchingMethod right) => left.Equals(right);
        public static bool operator !=(MatchingMethod left, MatchingMethod right) => !left.Equals(right);

        public static explicit operator string(MatchingMethod value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MatchingMethod other && Equals(other);
        public bool Equals(MatchingMethod other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The source of the watchlist
    /// </summary>
    [EnumType]
    public readonly struct Source : IEquatable<Source>
    {
        private readonly string _value;

        private Source(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Source Local_file { get; } = new Source("Local file");
        public static Source Remote_storage { get; } = new Source("Remote storage");

        public static bool operator ==(Source left, Source right) => left.Equals(right);
        public static bool operator !=(Source left, Source right) => !left.Equals(right);

        public static explicit operator string(Source value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Source other && Equals(other);
        public bool Equals(Source other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
