create database users;
use  users;
CREATE TABLE users (
    user_id BIGINT AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    city VARCHAR(50),
    introduction TEXT,
    avatar VARCHAR(255),
    create_time DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_time DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (user_id),
    UNIQUE KEY idx_email (email),
    UNIQUE KEY idx_username (username),
    UNIQUE KEY idx_phone (phone)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;