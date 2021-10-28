CREATE TABLE IF NOT EXISTS provinces(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR (50) UNIQUE NOT NULL,
    foundation      VARCHAR (50) NULL,
    capital         VARCHAR (50) NULL,
    population      INTEGER NULL,
    area            DECIMAL NULL,
    phone_prefix    VARCHAR (4) NULL,
    government_site VARCHAR (100) NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NULL
);

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Bengo', '26 de Abril de 1980', 'Caxito', 0, 31.371, '034', 'https://www.bengo.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Benguela', '1617', 'Benguela', 0, 39.827, '+244', 'https://www.benguela.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Bié', '1 de Maio de 1922', 'Cuíto', 0, 70.314, '+244', 'https://www.bie.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Cabinda', '28 de Fevereiro de 1919', 'Cabinda', 0, 7.283, '+244', 'https://www.cabinda.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Cuando Cubango', '21 de Outubro de 1961', 'Menongue', 0, 199.049, '049', 'https://www.cuandocubango.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Cuanza Norte', '15 de Agosto de 1914', 'Ndalatando', 0, 24.110, '035', 'https://www.cuanzanorte.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Cuanza Sul', '15 de Setembro de 1917', 'Sumbe', 0, 55.660, '+244', 'https://www.cuanzasul.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Cunene', '10 de Julho de 1970', 'Ondjiva', 0, 78.342, '035', 'https://www.cunene.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Huambo', '21 de Setembro de 1912', 'Huambo', 0, 35.771, '+244', 'https://www.huambo.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Huíla', '2 de Setembro de 1901', 'Lubango', 0, 79.022, '+244', 'https://www.huila.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Luanda', '1605', 'Luanda', 0, 18.826, '222', 'https://www.luanda.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Lunda Norte', '4 de Julho de 1978', 'Dundo', 0, 103.760, '+244', 'https://www.lundanorte.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Lunda Sul', '13 de Julho de 1895', 'Saurimo', 0, 77.636, '+244', 'https://www.lundasul.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Malanje', '1921', 'Malanje', 0, 98.302, '+244', 'https://www.malanje.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Moxico', '15 de Setembro de 1917', 'Luena', 0, 223.023, '+244', 'https://www.moxico.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Namibe', '19 de Abril de 1849', 'Moçâmedes', 0, 57.091, '+244', 'https://www.namibe.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Uíge', '31 de Maio de 1887', 'Uíge', 0, 58.698, '+244', 'https://www.uige.gov.ao');

INSERT INTO provinces(name, foundation, capital, population, area, phone_prefix, government_site)
VALUES ('Zaire', '1 de Abril de 1961', 'Mbanza Congo', 0, 40.138, '+232', 'https://www.zaire.gov.ao');
