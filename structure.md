├── main.go                 # Main entry point for the Notifications service
├── internal                        # Private application code (service-specific)
│   ├── email_preferences           # Feature module: Email Preferences
│   │   ├── models                  # Database models
│   │   │   └── email_preference.go
│   │   ├── validations             # Validation logic
│   │   │   └── email_preference_validations.go
│   │   ├── services                # Business logic
│   │       └── email_preference_service.go
│   └── notification_preferences    # Feature module: Notification Preferences
│       ├── models
│       │   └── notification_preference.go
│       ├── validations
│       │   └── notification_preference_validations.go
│       └── services
│           └── notification_preference_service.go
├── messaging                       # Messaging system integration
│   ├── kafka                       # Kafka integration
│   │   ├── producer.go
│   │   ├── consumer.go
│   │   └── index.go
│   └── websocket                   # WebSocket server and client logic
│       ├── server.go
│       ├── client.go
│       └── index.go
├── routes                          # API routes
│   ├── v1                          # Version 1 API
│   │   ├── email_preference_routes.go
│   │   └── notification_preference_routes.go
│   └── v2                          # Version 2 API (future enhancements)
│       ├── email_preference_routes.go
│       └── notification_preference_routes.go
├── handlers                        # Request handlers
│   ├── v1                          # Version 1 handlers
│   │   ├── email_preference_handler.go
│   │   └── notification_preference_handler.go
│   └── v2                          # Version 2 handlers (future)
│       ├── email_preference_handler.go
│       └── notification_preference_handler.go
├── config                          # Configuration and environment management
│   └── config.go
├── database                        # Database-related code
│   ├── connection.go               # Database connection setup
│   ├── migrations                  # Database migrations
│   │   ├── 0001_create_email_preferences_table.up.sql
│   │   ├── 0001_create_email_preferences_table.down.sql
│   │   ├── 0002_create_notification_preferences_table.up.sql
│   │   └── 0002_create_notification_preferences_table.down.sql
│   └── seeders                     # Database seeders
│       ├── seed_email_preferences.go
│       └── seed_notification_preferences.go
├── utils                           # Shared utilities
│   ├── constants.go                # Application-wide constants
│   ├── errors.go                   # Centralized error definitions
│   └── logger.go                   # Logging utilities
├── tests                           # Test suites
│   ├── unit                        # Unit tests
│   │   ├── handlers
│   │   │   └── v1
│   │   │       ├── email_preference_handler_test.go
│   │   │       └── notification_preference_handler_test.go
│   │   ├── services
│   │   │   ├── email_preference_service_test.go
│   │   │   └── notification_preference_service_test.go
│   │   └── utils
│   │       └── helpers_test.go
│   ├── integration                 # Integration tests
│   │   ├── api
│   │   │   └── v1
│   │   │       ├── email_preference_handler_test.go
│   │   │       └── notification_preference_handler_test.go
│   │   └── database
│   │       ├── email_preference_test.go
│   │       └── notification_preference_test.go
│   └── e2e                         # End-to-end tests
│       ├── email_preference_e2e_test.go
│       └── notification_preference_e2e_test.go
├── scripts                         # Helper scripts
│   ├── migrate.sh                  # Script for database migrations
│   └── seed.sh                     # Script for database seeding
├── docs                            # API documentation
│   ├── swagger.yaml
│   └── index.html
├── go.mod                          # Go module definition
├── go.sum                          # Go module dependencies
└── README.md                       # Project documentation
