package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := migrate.NewBusinessCaseOperation(ctx, "businessCaseOperation", &migrate.BusinessCaseOperationArgs{
BusinessCaseName: pulumi.String("sample-business-case"),
ProjectName: pulumi.String("multipleto8617project"),
ResourceGroupName: pulumi.String("rgopenapi"),
Settings: migrate.SettingsResponse{
AzureSettings: &migrate.AzureSettingsArgs{
AvsLaborCostPercentage: pulumi.Float64(0),
BusinessCaseType: pulumi.String("OptimizeForCost"),
ComfortFactor: pulumi.Float64(29),
Currency: pulumi.String("USD"),
DiscountPercentage: pulumi.Float64(83),
IaasLaborCostPercentage: pulumi.Float64(94),
InfrastructureGrowthRate: pulumi.Float64(83),
NetworkCostPercentage: pulumi.Float64(40),
PaasLaborCostPercentage: pulumi.Float64(47),
PerYearMigrationCompletionPercentage: pulumi.Float64Map{
"Year0": pulumi.Float64(20),
"Year1": pulumi.Float64(30),
"Year2": pulumi.Float64(60),
"Year3": pulumi.Float64(90),
},
PerformanceDataEndTime: pulumi.String("2023-11-08T07:10:07.764Z"),
PerformanceDataStartTime: pulumi.String("2023-11-08T07:10:07.764Z"),
PerformanceUtilizationPercentile: pulumi.Float64(4),
SavingsOption: pulumi.String("RI3Year"),
TargetLocation: pulumi.String("WestUs2"),
Wacc: pulumi.Float64(79),
WorkloadDiscoverySource: pulumi.String("Appliance"),
},
OnPremiseSettings: interface{}{
ComputeSettings: interface{}{
HyperthreadCoreToMemoryRatio: pulumi.Float64(12),
Price: pulumi.Float64(16),
SqlServerLicensing: migrate.SqlServerLicensingSettingsArray{
&migrate.SqlServerLicensingSettingsArgs{
LicenseCost: pulumi.Float64(27),
SoftwareAssuranceCost: pulumi.Float64(16),
Version: pulumi.String("Enterprise"),
},
},
VirtualizationSoftwareSettings: interface{}{
LicenseAndSupportList: migrate.VsphereLicenseArray{
&migrate.VsphereLicenseArgs{
BasicSupportCost: pulumi.Float64(22),
LicenseCost: pulumi.Float64(8),
LicenseType: pulumi.String("VSphereStandard"),
ProductionSupportCost: pulumi.Float64(22),
},
},
NumberOfPhysicalCoresPerLicense: pulumi.Int(17),
SoftwareAssuranceCost: pulumi.Float64(14),
},
WindowsServerLicensing: &migrate.WindowsServerLicensingSettingsArgs{
LicenseCost: pulumi.Float64(9),
LicensesPerCore: pulumi.Int(11),
SoftwareAssuranceCost: pulumi.Float64(1),
},
},
FacilitySettings: &migrate.FacilitySettingsArgs{
FacilitiesCost: pulumi.Float64(7),
},
LaborSettings: &migrate.LaborSettingsArgs{
HourlyAdminCost: pulumi.Float64(25),
PhysicalServersPerAdmin: pulumi.Int(6),
VirtualMachinesPerAdmin: pulumi.Int(24),
},
ManagementSettings: interface{}{
HypervVirtualizationManagementSettings: interface{}{
LicenseAndSupportList: migrate.HypervLicenseArray{
&migrate.HypervLicenseArgs{
LicenseCost: pulumi.Float64(4),
LicenseType: pulumi.String("Standard"),
},
},
NumberOfPhysicalCoresPerLicense: pulumi.Int(2),
SoftwareAssuranceCost: pulumi.Float64(11),
},
OtherManagementCostsSettings: &migrate.OtherManagementCostsSettingsArgs{
DataProtectionCostPerServerPerYear: pulumi.Float64(18),
MonitoringCostPerServerPerYear: pulumi.Float64(10),
PatchingCostPerServerPerYear: pulumi.Float64(18),
},
ThirdPartyManagementSettings: &migrate.ThirdPartyManagementSettingsArgs{
LicenseCost: pulumi.Float64(23),
SupportCost: pulumi.Float64(9),
},
VsphereManagementSettings: interface{}{
LicenseAndSupportList: migrate.VsphereManagementLicenseArray{
&migrate.VsphereManagementLicenseArgs{
BasicSupportCost: pulumi.Float64(1),
LicenseCost: pulumi.Float64(6),
LicenseType: pulumi.String("VSphereServerStandard"),
ProductionSupportCost: pulumi.Float64(18),
},
},
},
},
NetworkSettings: &migrate.NetworkSettingsArgs{
HardwareSoftwareCostPercentage: pulumi.Float64(50),
MaintenanceCostPercentage: pulumi.Float64(48),
},
SecuritySettings: &migrate.SecuritySettingsArgs{
ServerSecurityCostPerServerPerYear: pulumi.Float64(14),
SqlServerSecurityCostPerServerPerYear: pulumi.Float64(7),
},
StorageSettings: &migrate.StorageSettingsArgs{
CostPerGbPerMonth: pulumi.Float64(22),
MaintainanceCostPercentageToAcquisitionCost: pulumi.Float64(1),
},
},
},
})
if err != nil {
return err
}
return nil
})
}
