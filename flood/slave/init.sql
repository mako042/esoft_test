CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;

CREATE TABLE IF NOT EXISTS test_data (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	int_field1 INT, int_field2 INT, int_field3 INT, int_field4 INT, int_field5 INT,
	int_field6 INT, int_field7 INT, int_field8 INT, int_field9 INT, int_field10 INT,
	varchar_field1 VARCHAR(255), varchar_field2 VARCHAR(255), varchar_field3 VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE USER 'exporter'@'%' IDENTIFIED BY 'admin' WITH MAX_USER_CONNECTIONS 3;
GRANT PROCESS, REPLICATION CLIENT, SELECT ON *.* TO 'exporter'@'%';

