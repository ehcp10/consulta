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
    sexo int references sexo(id)  -- Adicionei uma vírgula aqui
);

CREATE TABLE tipo_contato (
    id int primary key,
    tipo_contato text
);

CREATE TABLE contato (
    id int primary key,
    client int references cliente(id),  -- Corrigi "client" para "cliente"
    tipo_contato int references tipo_contato(id),
    valor text
);

CREATE TABLE contato_emergencial (
    id int primary key,
    client int references cliente(id),  -- Corrigi "client" para "cliente"
    tipo_contato int references tipo_contato(id),
    valor text
);

CREATE TABLE custo_cliente(
    id int primary key,
    client int references cliente(id),  -- Corrigi "client" para "cliente"
    custo_sessao numeric(1000,2)  -- Adicionei uma vírgula aqui
);
