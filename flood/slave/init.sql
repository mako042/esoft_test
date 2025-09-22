CREATE DATABASE IF NOT EXISTS test_db;
USE test_db;

CREATE TABLE user_cart (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL DEFAULT 1,
  item_price_cents INT NOT NULL,
  total_price_cents INT NOT NULL,
  original_product_id INT,
  category_id INT,
  warehouse_id INT,
  promotion_id INT DEFAULT NULL,

  product_name VARCHAR(255) NOT NULL,
  product_attributes VARCHAR(255),
  session_id VARCHAR(255),

  PRIMARY KEY (id),
  INDEX idx_user_id (user_id)
);


CREATE USER 'exporter'@'%' IDENTIFIED BY 'admin' WITH MAX_USER_CONNECTIONS 3;
GRANT PROCESS, REPLICATION CLIENT, SELECT ON *.* TO 'exporter'@'%';

