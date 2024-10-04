package main

   import (
       "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
       "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
       "github.com/gitopsiq/terraform-provider-azureipam/provider"
   )

   func main() {
       plugin.Serve(&plugin.ServeOpts{
           ProviderFunc: func() *schema.Provider {
               return provider.Provider()
           },
       })
   }