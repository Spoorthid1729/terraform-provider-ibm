// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmDb2SaasAutoscaleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2SaasAutoscaleDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "x_db_profile"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "auto_scaling_allow_plan_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "auto_scaling_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "auto_scaling_max_storage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "auto_scaling_over_time_period"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "auto_scaling_pause_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "auto_scaling_threshold"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "storage_unit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "storage_utilization_percentage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.Db2-v0-test-public", "support_auto_scaling"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2SaasAutoscaleDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		 data "ibm_db2_saas_autoscale" "Db2-v0-test-public" {
    x_db_profile = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:f9455c22-07af-4a86-b9df-f02fd4774471::"
}
	`)
}
