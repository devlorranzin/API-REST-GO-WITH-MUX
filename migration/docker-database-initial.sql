create table empresas(
    id serial primary key,
    nome varchar,
    cnpj varchar,
    razao_social varchar,
    endereco_empresa varchar,
    ramo_de_atuacao varchar
);

INSERT INTO empresas(nome, cnpj, razao_social, endereco_empresa, ramo_de_atuacao) VALUES
('Hartman e Pricing LTDA', '03228997542', 'Hartman Pneus', 'Avenida das Pedras', 'Comercio'),
('Jose de Souza Neto MEI', '02323397542', 'Papagaio Motos', 'Rua 32 de Dezembro', 'Oficina')