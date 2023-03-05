CREATE TABLE IF NOT EXISTS Accounts(
    user_id SERIAL PRIMARY KEY,
    uuid uuid  UNIQUE NOT NULL,
    dob DATE NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    first_name VARCHAR(25) NULL DEFAULT NULL,
    last_name VARCHAR(25) NULL DEFAULT NULL,
    PASSWORD VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    rating INTEGER DEFAULT 0,
    active BOOLEAN,
    bio TEXT NULL DEFAULT NULL,
    link VARCHAR(50) NULL DEFAULT NULL,
    verified BOOLEAN,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL,
    post bigint[]
)