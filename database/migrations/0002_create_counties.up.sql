CREATE TABLE IF NOT EXISTS counties(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR (50) UNIQUE NOT NULL,
    province_id     INTEGER NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NULL,
    CONSTRAINT fk_province
        FOREIGN KEY (province_id)
            REFERENCES provinces (id)
);
