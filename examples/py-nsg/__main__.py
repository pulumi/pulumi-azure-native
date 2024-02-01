from pulumi import ResourceOptions
from pulumi_azure_native import resources, network

# Create a new resource group
resource_group = resources.ResourceGroup('resource_group')

# Create a new NetworkSecurityGroup
nsg = network.NetworkSecurityGroup('network_security_group',
                                   resource_group_name=resource_group.name,
                                   location=resource_group.location,
                                   # Can be removed after https://github.com/pulumi/pulumi-azure-native/issues/3049 is fixed.
                                   opts=ResourceOptions(ignore_changes=["securityRules"]))

# Create a default SecurityRule
rule = network.SecurityRule('security_rule',
                            resource_group_name=resource_group.name,
                            network_security_group_name=nsg.name,
                            access='Allow',
                            protocol='*',
                            direction='Inbound',
                            priority=100,
                            source_port_range='*',
                            source_address_prefix='*',
                            destination_port_range='80',
                            destination_address_prefix='*'
                            )
