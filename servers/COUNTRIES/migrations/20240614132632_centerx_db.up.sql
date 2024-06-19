CREATE TABLE IF NOT EXISTS countries (
    country_id SERIAL PRIMARY KEY,
    country_name VARCHAR(64),
    latitude FLOAT,
    longitude FLOAT
);
