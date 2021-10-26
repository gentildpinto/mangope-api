CREATE TABLE IF NOT EXISTS counties(
    id              VARCHAR (256) NOT NULL,
    name            VARCHAR (50) UNIQUE NOT NULL,
    province_id     VARCHAR (256) NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_province
        FOREIGN KEY (province_id)
            REFERENCES provinces (id)
);
