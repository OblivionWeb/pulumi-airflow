package shim

import (
	airflow "github.com/drfaust92/terraform-provider-airflow/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider(version string) *schema.Provider {
	return airflow.New(version)()
}
