package test

import (
  "fmt"
  "testing"
  "time"

  http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
  "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformAwsHelloWorldExample(t *testing.T) {
  t.Parallel()
  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../examples/terraform-hello-world-example",
  })

  // 最後にリソースを破棄する
  defer terraform.Destroy(t, terraformOptions)

  terraform.InitAndApply(t, terraformOptions)

  publicIp := terraform.Output(t, terraformOptions, "public_ip")

  url := fmt.Sprintf("http://%s:8080", publicIp)

  // テスト成功パターン
  http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello, World!", 30, 5*time.Second)

  // テスト失敗パターン
  //http_helper.HttpGetWithRetry(t, url, nil, 301, "Hello, World!", 30, 5*time.Second)
}
