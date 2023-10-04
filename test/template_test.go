
package test

import(
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformTemplate(t *testing.T) {
	t.Parallel()

	fixtureFoler := "./fixture"

	defer test_structure.RunTestStage(t, "teardown", func {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFoler)
		terraformOptions.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "setup", func {
		terraformOptions := configureModuleOptions(t, fixtureFoler)
		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFoler)
		result := terraform.Output(t, terraformOptions, "var_module_output")
		// if err != nil {
		// 	t.Fatal("Wrong output")
		// }

		if len(result) <= 0 {
			t.Fatal("Wrong output")
		}
	})
}

func configureModuleOptions(t *testing.T, folder string) *terraform.Options {
	terraformOptions := &terraform.Options {
		TerraformDir: folder,
		Vars :map[string]interface {

		}
	}
}