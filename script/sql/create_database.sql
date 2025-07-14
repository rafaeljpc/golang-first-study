-- Create database go_test
CREATE DATABASE go_test;

-- Create user go_test with password (replace 'password' with a secure password)
CREATE USER go_test WITH PASSWORD 'go';

-- Grant privileges to the user
GRANT ALL ON DATABASE go_test TO go_test;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
