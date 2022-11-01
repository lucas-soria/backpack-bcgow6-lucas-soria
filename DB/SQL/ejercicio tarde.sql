/* Create DataBase */
drop database if exists empresa_internet;
create database empresa_internet;
use empresa_internet;

/* Create Tables */
create table cliente (
	dni int not null,
	nombre varchar(50),
	apellido varchar(50),
	fecha_nacimiento datetime,
	provincia varchar(50),
	ciudad varchar(50),
    primary key (dni)
);

create table plan (
	id int not null auto_increment,
    velocidad int,
    precio double,
    descuento double,
    primary key (id)
);

create table cliente_plan (
	cliente_dni int,
    plan_id int,
    foreign key (cliente_dni) references cliente(dni),
    foreign key (plan_id) references plan(id)
);

/* Insert Values */
insert into cliente (dni, nombre, apellido, fecha_nacimiento, provincia, ciudad) values 
	(42938297, "Lucas", "Soria", "2000-05-09", "Mendoza", "Lujan"),
    (93826946, "Juan", "Perez", "1982-02-23", "Buenos Aires", "Olavarria"),
    (27625725, "Florencia", "Rios", "1984-02-23", "Buenos Aires", "Tandil"),
    (82589463, "Lautaro", "Ruiz", "1970-02-23", "Mendoza", "Guaymallen"),
    (98723827, "Juan", "Ramos", "2004-02-23", "Mendoza", "Tunuyan"),
    (09812876, "Raton", "Perez", "2006-02-23", "Mendoza", "Godoy Cruz"),
    (23434277, "Claudio", "Fuentes", "1999-02-23", "Mendoza", "Lujan"),
    (90876345, "Maximiliano", "Torres", "1985-02-23", "Buenos Aires", "Olavarria"),
    (98423783, "Alejandro", "Guinness", "1993-02-23", "Buenos Aires", "Tandil"),
    (09324877, "Maria", "Perez", "1997-02-23", "Buenos Aires", "Tandil")
;

insert into plan (velocidad, precio, descuento) values
	(10, 1500.30, 0),
    (20, 3300.34, 2),
    (50, 7000.54, 5),
    (100, 8947.94, 10),
    (500, 9999.99, 15)
;

insert into cliente_plan(cliente_dni, plan_id) values
	(42938297, 1),
    (93826946, 2),
    (27625725, 3),
    (82589463, 4),
    (98723827, 5),
    (09812876, 3),
    (23434277, 2),
    (90876345, 3),
    (98423783, 5),
    (23434277, 5),
    (90876345, 2),
    (98423783, 3),
    (09324877, 3)
;

/* Selects */
select * from plan;
select * from cliente;
select * from cliente_plan;

/* 10 queries */
-- 1. cliente Perez

-- 2. group cliente apellido

-- 3. group cliente_plan plan_id

-- 4. group cliente provincia y ciudad

-- 5. select rango etareo

--











