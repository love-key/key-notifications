# Variables
DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
MIGRATIONS_DIR=database/migrations  # Add this line to specify the migration directory

# Initialize database folder structure
init-db:
	mkdir -p database/migrations database/seeders database/models database/config
	touch database/config/config.go database/database.go
	echo "package config\n\nimport (\n\t\"os\"\n\t\"github.com/joho/godotenv\"\n)\n\ntype DBConfig struct {\n\tHost     string\n\tPort     string\n\tUser     string\n\tPassword string\n\tDBName   string\n}\n\nfunc LoadDBConfig() DBConfig {\n\tgodotenv.Load()\n\treturn DBConfig{\n\t\tHost:     os.Getenv(\"DB_HOST\"),\n\t\tPort:     os.Getenv(\"DB_PORT\"),\n\t\tUser:     os.Getenv(\"DB_USER\"),\n\t\tPassword: os.Getenv(\"DB_PASSWORD\"),\n\t\tDBName:   os.Getenv(\"DB_NAME\"),\n\t}\n}" > database/config/config.go
	echo "package database\n\nimport (\n\t\"database/sql\"\n\t\"log\"\n\t_ \"github.com/lib/pq\"\n\t\"notifications-system/database/config\"\n)\n\nvar DB *sql.DB\n\nfunc Connect() {\n\tcfg := config.LoadDBConfig()\n\tconnStr := \"postgres://\" + cfg.User + \":\" + cfg.Password + \"@\" + cfg.Host + \":\" + cfg.Port + \"/\" + cfg.DBName + \"?sslmode=disable\"\n\n\tvar err error\n\tDB, err = sql.Open(\"postgres\", connStr)\n\tif err != nil {\n\t\tlog.Fatalf(\"Unable to connect to database: %v\", err)\n\t}\n\n\tlog.Println(\"Connected to PostgreSQL!\")\n}" > database/database.go

# Create a new model
create-model:
	@read -p "Enter model name (e.g., user): " name; \
	touch database/models/$${name}.go; \
	echo "package models\n\ntype $$(echo $${name} | awk '{print toupper(substr($$0,1,1)) substr($$0,2)}') struct {\n\tID int \`json:\"id\"\`\n}" > database/models/$${name}.go

# Create a new migration
create-migration:
	@read -p "Enter migration name: " name; \
    migrate create -seq -ext sql -dir $(MIGRATIONS_DIR) $${name}

# Create a new seeder
create-seeder:
	@read -p "Enter seeder name (e.g., seed_users): " name; \
	touch database/seeders/$${name}.go; \
	echo "package main\n\nimport (\n\t\"database/sql\"\n\t\"log\"\n\t_ \"github.com/lib/pq\"\n)\n\nfunc main() {\n\tconnStr := \"postgres://user:password@localhost:5432/dbname?sslmode=disable\"\n\tdb, err := sql.Open(\"postgres\", connStr)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\tdefer db.Close()\n\n\t// Insert seed data\n\t_, err = db.Exec(\`INSERT INTO users (name, email) VALUES ($1, $2)\`, \"John Doe\", \"john@example.com\")\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\tlog.Println(\"Seed data inserted successfully!\")\n}" > database/seeders/$${name}.go

# Revert all migrations
reset-migrations:
	migrate -path ./migrations -database "$(DB_URL)" down -all
	@echo "All migrations have been reverted."

# Revert all seeders
reset-seeders:
	@echo "Reverting seeders..."
	@go run database/seeders/cleanup_seeders.go || echo "No cleanup script for seeders. Add custom logic if needed."
	@echo "Seeders cleanup is complete."
	
# Run migrations
# Run migrations using Go script
migrate-up:
	go run database/migrations/migrate.go
	@echo "Migrations have been run successfully."

# Rollback migration using Go script
migrate-down:
	go run database/migrations/rollback.go
	@echo "Rollback has been completed."

# Rollback all migrations
rollback-all:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down -all
	@echo "All migrations have been reverted."

# Run seeders
seed-users:
	go run seeders/seed_users.go

# Run tests
test:
	go test ./...

# Start the application with nodemon
start:
	nodemon --exec go run main.go --ext go

# Clean up
clean:
	rm -rf bin/
