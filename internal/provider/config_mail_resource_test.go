/*
 * Copyright (c) 2019-present Sonatype, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccConfigMailResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccConfigMailResource(),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify mail configuration
					resource.TestCheckResourceAttrSet("sonatypeiq_config_mail.mail_config", "id"),
					resource.TestCheckResourceAttr("sonatypeiq_config_mail.mail_config", "port", "25"),
					resource.TestCheckResourceAttr("sonatypeiq_config_mail.mail_config", "ssl_enabled", "false"),
					resource.TestCheckResourceAttr("sonatypeiq_config_mail.mail_config", "start_tls_enabled", "false"),
					resource.TestCheckResourceAttr("sonatypeiq_config_mail.mail_config", "system_email", "no-reply@my-domain.tld"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccConfigMailResource() string {
	return fmt.Sprintf(providerConfig + `
resource "sonatypeiq_config_mail" "mail_config" {
  hostname          = "smtp.my-domain.tld"
  port              = 25
  ssl_enabled       = false
  start_tls_enabled = false
  system_email      = "no-reply@my-domain.tld"
}`)
}
