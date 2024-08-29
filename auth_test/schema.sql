CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL -- In real applications, use hashed passwords
);

CREATE TABLE IF NOT EXISTS income (
    id SERIAL PRIMARY KEY,
    amount NUMERIC NOT NULL,
    description TEXT,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);
