CREATE DATABASE capstone5;

USE capstone5;

CREATE TABLE Customers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    token VARCHAR(255),
    reset_token VARCHAR(255),
    reset_token_expire DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE Admins (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    create_at DATETIME,
    update_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE Fundraising_categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE Organizations (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    admin_id BIGINT,
    name VARCHAR(255) NOT NULL,
    is_valid ENUM('yes', 'no') DEFAULT 'no',
    description TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (admin_id) REFERENCES Admins(id)
);

CREATE TABLE Volunteer_vacancies (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    organization_id BIGINT,
    title VARCHAR(255) NOT NULL,
    content_activity TEXT,
    location VARCHAR(255),
    date VARCHAR(255),
    target_volunteer INT,
    registered_volunteer INT,
    registration_deadline INT,
    image_url VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (organization_id) REFERENCES Organizations(id)
);

CREATE TABLE Articles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    admin_id BIGINT,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    image_url VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (admin_id) REFERENCES Admins(id)
);

CREATE TABLE Fundraising (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    organization_id BIGINT,
    fundraising_category_id BIGINT,
    goal_amount FLOAT,
    current_amount FLOAT,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    status ENUM('active', 'inactive') DEFAULT 'active',
    is_bookmark ENUM('yes', 'no') DEFAULT 'no',
    image VARCHAR(255),
    start_date DATETIME,
    end_date DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (organization_id) REFERENCES Organizations(id),
    FOREIGN KEY (fundraising_category_id) REFERENCES Fundraising_categories(id)
);

CREATE TABLE Donations (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    fundraising_id BIGINT,
    customer_id BIGINT,
    nominal FLOAT,
    status ENUM('pending', 'success', 'failed') DEFAULT 'pending',
    payment_url VARCHAR(255),
    create_at DATETIME,
    update_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (fundraising_id) REFERENCES Fundraising(id),
    FOREIGN KEY (customer_id) REFERENCES Customers(id)
);

CREATE TABLE Testimoni_Fundraising (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    donation_id BIGINT,
    total_like INT,
    testimoni TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (donation_id) REFERENCES Donations(id)
);

CREATE TABLE Customer_Bookmark_Fundraising (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT,
    fundraising_id BIGINT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (customer_id) REFERENCES Customers(id),
    FOREIGN KEY (fundraising_id) REFERENCES Fundraising(id)
);

CREATE TABLE Customer_Bookmark_Volunteer_Vacancies (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT,
    volunteer_vacancies_id BIGINT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (customer_id) REFERENCES Customers(id),
    FOREIGN KEY (volunteer_vacancies_id) REFERENCES Volunteer_vacancies(id)
);

CREATE TABLE Customer_Bookmark_Article (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT,
    article_id BIGINT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (customer_id) REFERENCES Customers(id),
    FOREIGN KEY (article_id) REFERENCES Articles(id)
);

CREATE TABLE Comment (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT,
    article_id BIGINT,
    comment VARCHAR(255),
    total_like INT,
    is_like ENUM('yes', 'no') DEFAULT 'no',
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (customer_id) REFERENCES Customers(id),
    FOREIGN KEY (article_id) REFERENCES Articles(id)
);

CREATE TABLE Application (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT,
    vacancy_id BIGINT,
    image_url VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (customer_id) REFERENCES Customers(id),
    FOREIGN KEY (vacancy_id) REFERENCES Volunteer_vacancies(id)
);