CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS book_table(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    author_name VARCHAR(255),
    publisher_year TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255)
);