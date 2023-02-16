package api

// TransactionType -
type TransactionType string

const (
	TransactionTypeInvoke        = "INVOKE"
	TransactionTypeDeclare       = "DECLARE"
	TransactionTypeDeploy        = "DEPLOY"
	TransactionTypeDeployAccount = "DEPLOY_ACCOUNT"
)
