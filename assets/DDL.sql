CREATE DATABASE server_pulsa_db;

CREATE TABLE member (
    id INT AUTO_INCREMENT PRIMARY KEY,
    member_id VARCHAR(10) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    address VARCHAR(255) NOT NULL,
    balance DECIMAL (15, 2),
    pin INT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE supplier (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    jabber VARCHAR(255) UNIQUE NOT NULL,
    address VARCHAR(255) NOT NULL,
    balance DECIMAL (15, 2),
    pin INT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product (
    id INT AUTO_INCREMENT PRIMARY KEY,
    category VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    code VARCHAR(10) UNIQUE NOT NULL,
    nominal INT NOT NULL,
    price DECIMAL (15, 2) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transaction (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_member INT NOT NULL,
    FOREIGN KEY (id_member) REFERENCES member(id),
    id_product INT NOT NULL,
    FOREIGN KEY (id_product) REFERENCES product(id),
    id_center INT NOT NULL,
    FOREIGN KEY (id_center) REFERENCES center(id),
    quantity INT NOT NULL,
    status ENUM('pending', 'success', 'failed') NOT NULL DEFAULT 'pending',
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE topup (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_member INT NOT NULL,
    FOREIGN KEY (id_member) REFERENCES member(id),
    id_center INT NOT NULL,
    FOREIGN KEY (id_center) REFERENCES center(id),
    amount DECIMAL (15, 2) NOT NULL,
    payment_method VARCHAR(255),
    status ENUM('pending', 'success', 'failed') NOT NULL DEFAULT 'pending',
    topup_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE center (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_supplier INT NOT NULL,
    FOREIGN KEY (id_supplier) REFERENCES supplier(id),
    name VARCHAR(255) UNIQUE NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);