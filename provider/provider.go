package provider

   import (
       "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
   )

   func Provider() *schema.Provider {
       return &schema.Provider{
           Schema: map[string]*schema.Schema{
               // Add your provider configuration here
           },
           ResourcesMap: map[string]*schema.Resource{
               // Add your resources here
           },
           DataSourcesMap: map[string]*schema.Resource{
               // Add your data sources here
           },
       }
   }