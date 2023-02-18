package data

// TransactionType -
type TransactionType string

const (
	TransactionTypeInvoke         = "INVOKE"
	TransactionTypeInvokeFunction = "INVOKE_FUNCTION"
	TransactionTypeDeclare        = "DECLARE"
	TransactionTypeDeploy         = "DEPLOY"
	TransactionTypeDeployAccount  = "DEPLOY_ACCOUNT"
	TransactionTypeL1Handler      = "L1_HANDLER"
)
