**Warning:** when `AzureADOnlyAuthentication` is enabled, the Azure SQL API rejects any `AdministratorLoginPassword`, even if it is the same as the current one.

According to the Azure team, this API design owes to the following reasons:
- Changing the password is not allowed when Entra-only authentication is enabled because it could lead to invalid templates.
- Any updates containing the same, unchanged password are also rejected because different behavior for same vs different passwords would be a vector for brute forcing the password.

To work around this, you can comment out `AdministratorLoginPassword` when enabling `AzureADOnlyAuthentication`. To update the password, you can disable `AzureADOnlyAuthentication` and re-enable it after the update.

For more details and discussion please see [this issue](https://github.com/pulumi/pulumi-azure-native/issues/2937).