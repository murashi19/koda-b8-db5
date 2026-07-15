CREATE TABLE contacts(
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    email VARCHAR(30) NOT NULL UNIQUE,
    phone_number VARCHAR(13) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

SELECT id, name, email, phone_number, created_at, updated_at FROM contacts;

INSERT INTO contacts (name, email, phone_number) VALUES
('Bang AL', 'al@mail.com', '082327163521');