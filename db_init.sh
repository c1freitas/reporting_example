# Run this in the Postgresql container

# Create the DB if it does not exist
docker exec -it reporting_db_1 sh -c  "psql -U postgres -tc \"SELECT 1 FROM pg_database WHERE datname = 'reporting'\" | grep -q 1 | psql -U postgres -c \"CREATE DATABASE reporting\""

# Load the seed data
docker exec -it reporting_db_1 sh -c "psql -h localhost -d reporting -U postgres -f  /var/data/db_init.sql"