package gen

// Some resources which require empty containers to be present in the request body
// aren't listed as required in the specification, so we add them here manually.
func additionalRequiredContainers(resourceTok string) RequiredContainers {
	switch resourceTok {
	case "azure-native:app:ManagedEnvironment":
		return [][]string{{"properties"}}
	}
	return nil
}
