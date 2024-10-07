terraform-provider-azureipam/
│
├── main.go                         # Entry point for the Terraform provider
├── go.mod                          # Go module dependencies
├── go.sum                          # Go module checksums
├── provider.go                     # Main provider definition
│
├── resources/                      # Directory for all Terraform resources
│   ├── resource_space.go           # CRUD operations for IPAM spaces
│   ├── resource_block.go           # CRUD operations for IPAM blocks
│   ├── resource_reservation.go     # CRUD operations for IP reservations
│   ├── resource_subnet.go          # CRUD operations for subnets inside blocks
│   ├── resource_next_subnet.go     # Automatic allocation of the next available subnet
│   ├── resource_external_network.go # CRUD operations for external networks
│   ├── resource_user.go            # Future: CRUD operations for managing users
│   ├── resource_admin.go           # Future: CRUD operations for managing admins
│   └── resource_exclusion.go       # Optional: CRUD operations for managing excluded subscriptions
│
├── data_sources/                   # Directory for all Terraform data sources
│   ├── data_source_space.go        # Fetch single space
│   ├── data_source_spaces.go       # Fetch all spaces
│   ├── data_source_block.go        # Fetch single block
│   ├── data_source_blocks.go       # Fetch all blocks in a space
│   ├── data_source_reservation.go  # Fetch single reservation
│   ├── data_source_reservations.go # Fetch all reservations in a block
│   ├── data_source_next_subnet.go  # Fetch next available subnet
│   ├── data_source_next_vnet.go    # Fetch next available VNet CIDR
│   ├── data_source_cidr_check.go   # Perform CIDR overlap check
│   ├── data_source_available_networks.go # Fetch available networks in a block
│   ├── data_source_utilization.go  # Fetch utilization data for spaces or blocks
│   ├── data_source_user.go         # Future: Fetch user data
│   ├── data_source_admin.go        # Future: Fetch admin data
│   ├── data_source_exclusions.go   # Optional: Fetch excluded subscriptions
│   ├── data_source_logs.go         # Optional: Fetch event logs (if available)
│
├── client/                         # Directory for the Go client used to interact with the Azure IPAM API
│   ├── client.go                   # Core client setup and configuration
│   ├── spaces.go                   # API client functions for managing spaces
│   ├── blocks.go                   # API client functions for managing blocks
│   ├── reservations.go             # API client functions for managing reservations
│   ├── subnets.go                  # API client functions for managing subnets in blocks
│   ├── next_subnet.go              # API client functions for next available subnet
│   ├── external_networks.go        # API client functions for external networks
│   ├── users.go                    # Future: API client functions for managing users
│   ├── admins.go                   # Future: API client functions for managing admins
│   ├── exclusions.go               # Optional: API client functions for managing excluded subscriptions
│   └── utils.go                    # Helper functions (e.g., request creation, error handling)
│
├── tests/                          # Directory for unit and acceptance tests
│   ├── unit/                       # Unit tests for all Go client functions
│   │   ├── client_test.go          # Unit tests for core client functions
│   │   ├── spaces_test.go          # Unit tests for space management
│   │   ├── blocks_test.go          # Unit tests for block management
│   │   ├── reservations_test.go    # Unit tests for reservation management
│   │   ├── subnets_test.go         # Unit tests for subnet management
│   ├── acceptance/                 # Acceptance tests for Terraform resources and data sources
│   │   ├── test_azureipam_space.go # Acceptance tests for IPAM space resource
│   │   ├── test_azureipam_block.go # Acceptance tests for IPAM block resource
│   │   ├── test_azureipam_reservation.go # Acceptance tests for IPAM reservation resource
│   │   ├── test_azureipam_subnet.go # Acceptance tests for IPAM subnet resource
│   │   ├── test_azureipam_next_subnet.go # Acceptance tests for next available subnet resource
│   │   ├── test_azureipam_external_network.go # Acceptance tests for external networks
│   │   ├── test_azureipam_user.go   # Future: Acceptance tests for user resource
│   │   ├── test_azureipam_admin.go  # Future: Acceptance tests for admin resource
│   │   └── test_azureipam_exclusion.go # Optional: Acceptance tests for excluded subscriptions
│
├── examples/                       # Directory containing example Terraform configurations
│   ├── example_spaces.tf           # Example usage for creating spaces
│   ├── example_blocks.tf           # Example usage for managing blocks
│   ├── example_reservations.tf     # Example for IP reservations
│   ├── example_subnets.tf          # Example for automatic subnet assignment
│   ├── example_next_subnet.tf      # Example for reserving the next available subnet
│   ├── example_external_network.tf # Example for managing external networks
│   ├── example_cidr_check.tf       # Example for using the CIDR check tool
│   ├── example_utilization.tf      # Example for fetching utilization data
│   ├── example_user.tf             # Future: Example for user management
│   ├── example_admin.tf            # Future: Example for admin management
│   └── example_exclusion.tf        # Optional: Example for managing excluded subscriptions
│
└── README.md                       # Documentation for the Terraform provider
