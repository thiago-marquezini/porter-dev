package repository

type Repository interface {
	User() UserRepository
	Project() ProjectRepository
	Release() ReleaseRepository
	Environment() EnvironmentRepository
	Session() SessionRepository
	GitRepo() GitRepoRepository
	Cluster() ClusterRepository
	Database() DatabaseRepository
	HelmRepo() HelmRepoRepository
	Registry() RegistryRepository
	Infra() InfraRepository
	GitActionConfig() GitActionConfigRepository
	Invite() InviteRepository
	AuthCode() AuthCodeRepository
	DNSRecord() DNSRecordRepository
	PWResetToken() PWResetTokenRepository
	KubeIntegration() KubeIntegrationRepository
	BasicIntegration() BasicIntegrationRepository
	OIDCIntegration() OIDCIntegrationRepository
	OAuthIntegration() OAuthIntegrationRepository
	GCPIntegration() GCPIntegrationRepository
	AWSIntegration() AWSIntegrationRepository
	AzureIntegration() AzureIntegrationRepository
	GithubAppInstallation() GithubAppInstallationRepository
	GithubAppOAuthIntegration() GithubAppOAuthIntegrationRepository
	SlackIntegration() SlackIntegrationRepository
	AppEventWebhook() AppEventWebhookRepository
	GitlabIntegration() GitlabIntegrationRepository
	GitlabAppOAuthIntegration() GitlabAppOAuthIntegrationRepository
	NotificationConfig() NotificationConfigRepository
	JobNotificationConfig() JobNotificationConfigRepository
	BuildEvent() BuildEventRepository
	KubeEvent() KubeEventRepository
	ProjectUsage() ProjectUsageRepository
	Onboarding() ProjectOnboardingRepository
	CredentialsExchangeToken() CredentialsExchangeTokenRepository
	BuildConfig() BuildConfigRepository
	Allowlist() AllowlistRepository
	APIToken() APITokenRepository
	Policy() PolicyRepository
	Tag() TagRepository
	Stack() StackRepository
	MonitorTestResult() MonitorTestResultRepository
	APIContractRevisioner() APIContractRevisioner
	AWSAssumeRoleChainer() AWSAssumeRoleChainer
	PorterApp() PorterAppRepository
	PorterAppEvent() PorterAppEventRepository
	SystemServiceStatus() SystemServiceStatusRepository
	DeploymentTarget() DeploymentTargetRepository
	AppRevision() AppRevisionRepository
	AppTemplate() AppTemplateRepository
	GithubWebhook() GithubWebhookRepository
	Datastore() DatastoreRepository
	AppInstance() AppInstanceRepository
	Referral() ReferralRepository
}
