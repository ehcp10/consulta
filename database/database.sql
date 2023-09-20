CREATE DATABASE mental WITH ENCODING 'UTF8' TEMPLATE template1;

CREATE TABLE profissao (
    id int primary key,
    profissao text
);

CREATE TABLE estado_civil (
    id int primary key,
    estado_civil text
);

CREATE table sexo (
    id int primary key,
    sexo text
);

CREATE TABLE cliente (
    id int primary key,
    nome text,
    data_nascimento date,
    profissao int references profissao(id),
    idade int,
    estado_civil int references estado_civil(id),
    sexo int references sexo(id)  
);

CREATE TABLE tipo_contato (
    id int primary key,
    tipo_contato text
);

CREATE TABLE contato (
    id int primary key,
    client int references cliente(id),  
    tipo_contato int references tipo_contato(id),
    valor text
);

CREATE TABLE contato_emergencial (
    id int primary key,
    client int references cliente(id),  
    tipo_contato int references tipo_contato(id),
    valor text
);

CREATE TABLE custo_cliente(
    id int primary key,
    client int references cliente(id),  
    custo_sessao numeric(1000,2)  
);

CREATE TABLE status(
    id int primary key,
    status text
);

CREATE TABLE pagamento(
    id int primary key,
    situacao text
);

CREATE TABLE sessao (
    id int primary key,
    data_sessao date,
    hora int,
    cliente int references cliente(id),
    presenca boolean,
    status int references status(id),
    pagamento int references pagamento(id),
    atualizado_em  date
);

CREATE TABLE credito (
    id int primary key,
    cliente int references cliente(id),
    quantidade_sessoes int
)
