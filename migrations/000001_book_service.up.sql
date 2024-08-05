CREATE TABLE IF NOT EXISTS book_table(
    id UUID PRIMARY KEY ,
    author_name VARCHAR(255),
    publisher_year TIMESTAMP ,
    title VARCHAR(255)
);