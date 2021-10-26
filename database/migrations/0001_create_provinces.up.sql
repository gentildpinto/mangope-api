CREATE TABLE IF NOT EXISTS `provinces` (
    id              VARCHAR (256) PRIMARY KEY,
    name            VARCHAR (50) UNIQUE NOT NULL,
    foundation      VARCHAR (50) NULL,
    capital         VARCHAR (10) NULL,
    population      INTEGER NULL,
    area            INTEGER NULL,
    phone_prefix    VARCHAR (4) NULL,
    government_site VARCHAR (256) NULL,
    created_at      DATETIME (3) NULL,
    updated_at      DATETIME (3) NULL,
    deleted_at      DATETIME (3) NULL
);
