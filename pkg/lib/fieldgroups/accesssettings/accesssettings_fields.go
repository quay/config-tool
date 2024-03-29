package accesssettings

// Fields returns a list of strings representing the fields in this field group
func (fg *AccessSettingsFieldGroup) Fields() []string {
	return []string{"AUTHENTICATION_TYPE", "FEATURE_ANONYMOUS_ACCESS", "FEATURE_DIRECT_LOGIN", "FEATURE_GITHUB_LOGIN", "FEATURE_GOOGLE_LOGIN", "FEATURE_INVITE_ONLY_USER_CREATION", "FEATURE_PARTIAL_USER_AUTOCOMPLETE", "FEATURE_USERNAME_CONFIRMATION", "FEATURE_USER_CREATION", "FEATURE_USER_LAST_ACCESSED", "FEATURE_USER_LOG_ACCESS", "FEATURE_USER_METADATA", "FEATURE_USER_RENAME", "FRESH_LOGIN_TIMEOUT", "USER_RECOVERY_TOKEN_LIFETIME", "FEATURE_EXTENDED_REPOSITORY_NAMES", "CREATE_REPOSITORY_ON_PUSH_PUBLIC", "FEATURE_USER_INITIALIZE"}
}
