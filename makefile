up_build: script_permissions export_variables
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose --env-file test_influxdb.env up --build -d
	@echo "Docker images built and started!"

up: script_permissions
	@echo "Starting Docker images..."
	docker-compose --env-file test_influxdb.env up -d
	@echo "Docker images started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

script_permissions:
	@echo "Setting execution permissions for init-influxdb.sh..."
	chmod +x db/init.sh
