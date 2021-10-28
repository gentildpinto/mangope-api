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

INSERT INTO counties(name, province_id)
VALUES ('Ambriz', 1),
       ('Bula Atumba', 1),
       ('Dande', 1),
       ('Dembos', 1),
       ('Nambuangongo', 1),
       ('Pango Aluquém', 1);

INSERT INTO counties(name, province_id)
VALUES  ('Balombo', 2),
        ('Baía Farta', 2),
        ('Benguela', 2),
        ('Bocoio', 2),
        ('Caimbambo', 2),
        ('Catumbela', 2),
        ('Chongorói', 2),
        ('Cubal', 2),
        ('Ganda', 2),
        ('Lobito', 2);

INSERT INTO counties(name, province_id)
VALUES  ('Andulo', 3),
        ('Camacupa', 3),
        ('Catabola', 3),
        ('Chinguar', 3),
        ('Chitembo', 3),
        ('Cuemba', 3),
        ('Cunhinga', 3),
        ('Cuíto', 3),
        ('Nharea', 3);

INSERT INTO counties(name, province_id)
VALUES  ('Belize', 4),
        ('Buco-Zau', 4),
        ('Cabinda', 4),
        ('Cacongo', 4);

INSERT INTO counties(name, province_id)
VALUES  ('Calai', 5),
        ('Cuangar', 5),
        ('Cuchi', 5),
        ('Cuito Cuanavale', 5),
        ('Dirico', 5),
        ('Mavinga', 5),
        ('Menongue', 5),
        ('Nancova', 5),
        ('Rivungo', 5);

INSERT INTO counties(name, province_id)
VALUES  ('Ambaca', 6),
        ('Banga', 6),
        ('Bolongongo', 6),
        ('Cambambe', 6),
        ('Cazengo', 6),
        ('Golungo Alto', 6),
        ('Gonguembo', 6),
        ('Lucala', 6),
        ('Quiculungo', 6),
        ('Samba Caju', 6);

INSERT INTO counties(name, province_id)
VALUES  ('Amboim', 7),
        ('Cassongue', 7),
        ('Cela', 7),
        ('Conda', 7),
        ('Ebo', 7),
        ('Libolo', 7),
        ('Mussende', 7),
        ('Porto Amboim', 7),
        ('Quibala', 7),
        ('Quilenda', 7),
        ('Seles', 7),
        ('Sumbe', 7);

INSERT INTO counties(name, province_id)
VALUES  ('Cahama', 8),
        ('Cuanhama', 8),
        ('Curoca', 8),
        ('Cuvelai', 8),
        ('Namacunde', 8),
        ('Ombadja', 8);

INSERT INTO counties(name, province_id)
VALUES  ('Bailundo', 9),
        ('Cachiungo', 9),
        ('Caála', 9),
        ('Ecunha', 9),
        ('Huambo', 9),
        ('Londuimbali', 9),
        ('Longonjo', 9),
        ('Mungo', 9),
        ('Chicala-Choloanga', 9),
        ('Chinjenje', 9),
        ('Ucuma', 9);

INSERT INTO counties(name, province_id)
VALUES  ('Caconda', 10),
        ('Cacula', 10),
        ('Caluquembe', 10),
        ('Chiange', 10),
        ('Chibia', 10),
        ('Chicomba', 10),
        ('Chipindo', 10),
        ('Cuvango', 10),
        ('Humpata', 10),
        ('Jamba', 10),
        ('Lubango', 10),
        ('Matala', 10),
        ('Quilengues', 10),
        ('Quipungo', 10);

INSERT INTO counties(name, province_id)
VALUES  ('Belas', 11),
        ('Cacuaco', 11),
        ('Cazenga', 11),
        ('Ícolo e Bengo', 11),
        ('Luanda', 11),
        ('Quilamba Quiaxi', 11),
        ('Quissama', 11),
        ('Talatona', 11),
        ('Viana', 11);

INSERT INTO counties(name, province_id)
VALUES  ('Cambulo', 12),
        ('Capenda-Camulemba', 12),
        ('Caungula', 12),
        ('Chitato', 12),
        ('Cuango', 12),
        ('Cuílo', 12),
        ('Lóvua', 12),
        ('Lubalo', 12),
        ('Lucapa', 12),
        ('Xá-Muteba', 12);

INSERT INTO counties(name, province_id)
VALUES  ('Cacolo', 13),
        ('Dala', 13),
        ('Muconda', 13),
        ('Saurimo', 13);

INSERT INTO counties(name, province_id)
VALUES  ('Cacuso', 14),
        ('Calandula', 14),
        ('Cambundi-Catembo', 14),
        ('Cangandala', 14),
        ('Caombo', 14),
        ('Cuaba Nzoji', 14),
        ('Cunda-Dia-Baze', 14),
        ('Luquembo', 14),
        ('Malanje', 14),
        ('Marimba', 14),
        ('Massango', 14),
        ('Mucari', 14),
        ('Quela', 14),
        ('Quirima', 14);

INSERT INTO counties(name, province_id)
VALUES  ('Alto Zambeze', 15),
        ('Bundas', 15),
        ('Camanongue', 15),
        ('Léua', 15),
        ('Luau', 15),
        ('Luacano', 15),
        ('Luchazes', 15),
        ('Cameia', 15),
        ('Moxico', 15);

INSERT INTO counties(name, province_id)
VALUES  ('Bibala', 16),
        ('Camucuio', 16),
        ('Moçâmedes', 16),
        ('Tômbua', 16),
        ('Virei', 16);

INSERT INTO counties(name, province_id)
VALUES  ('Alto Cauale', 17),
        ('Ambuíla', 17),
        ('Bembe', 17),
        ('Buengas', 17),
        ('Bungo', 17),
        ('Damba', 17),
        ('Milunga', 17),
        ('Mucaba', 17),
        ('Negage', 17),
        ('Puri', 17),
        ('Quimbele', 17),
        ('Quitexe', 17),
        ('Sanza Pombo', 17),
        ('Songo', 17),
        ('Uíge', 17),
        ('Zombo', 17);

INSERT INTO counties(name, province_id)
VALUES  ('Cuimba', 18),
        ('Mabanza Congo', 18),
        ('Nóqui', 18),
        ('Nezeto', 18),
        ('Soio', 18),
        ('Tomboco', 18);
