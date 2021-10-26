CREATE TABLE IF NOT EXISTS `provinces` (
    id              VARCHAR (256) PRIMARY KEY,
    name            VARCHAR (50) UNIQUE NOT NULL,
    province_id     VARCHAR (256) NOT NULL REFERENCES provinces(id),
    created_at      DATETIME (3) NULL,
    updated_at      DATETIME (3) NULL,
    deleted_at      DATETIME (3) NULL
);
