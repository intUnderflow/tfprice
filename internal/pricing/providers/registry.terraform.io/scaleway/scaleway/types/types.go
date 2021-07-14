package types

type ProviderConfig struct {
	Expressions ProviderConfigExpressions `json:"expressions"`
}

type ProviderConfigExpressions struct {
	Region ProviderConfigExpressionsValue `json:"region"`
	Zone   ProviderConfigExpressionsValue `json:"zone"`
}

type ProviderConfigExpressionsValue struct {
	ConstantValue *string `json:"constant_value"`
}
