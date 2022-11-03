drop database if exists biblioteca;
create database biblioteca;
use biblioteca;

create table libro (
	id int not null auto_increment,
    titulo varchar(255),
    editorial varchar(255),
    area varchar(255),
    primary key (id)
);

create table autor (
	id int not null auto_increment,
    nombre varchar(255),
    pais varchar(255),
    primary key (id)
);

create table libro_autor (
	autor_id int,
    libro_id int,
    foreign key (autor_id) references autor(id),
    foreign key (libro_id) references libro(id)
);

create table estudiante (
	id int not null auto_increment,
    nombre varchar(255),
    apellido varchar(255),
    direccion varchar(255),
    carrera varchar(255),
    edad int,
    primary key (id)
);

create table prestamo (
	estudiante_id int,
    libro_id int,
    fecha_prestamo timestamp,
    fecha_devolucion timestamp,
    devuelto bool,
    foreign key (estudiante_id) references estudiante(id),
    foreign key (libro_id) references libro(id)
);

insert into libro (titulo, editorial, area) values
	("fisica 1", "Salamandra", "fisica"),
    ("fisica 2", "libros de ciencia", "fisica"),
    ("Deep dive into design patterns", "Guru", "Internet"),
    ("El Universo: Guía de viaje", "Puerto de palos", "Ficcion"),
    ("H. P.", "Salamandra", "Ficcion")
;

insert into autor (nombre, pais) values
	("Le pier", "Francia"),
    ("Giorgio", "Italia"),
    ("J.K. Rowling", "Reino Unido"),
    ("Paco Perez", "Argentina"),
    ("Juan Gomez", "Francia")
;

insert into libro_autor (autor_id, libro_id) values
	(3, 5),
    (1, 1),
    (1, 2),
    (2, 2),
    (4, 5),
    (5, 4),
    (5, 3)
;

insert into estudiante (nombre, apellido, direccion, carrera, edad) values
	("Lucas", "Soria", "Punzon 273", "Informatica", 22),
    ("Filippo", "Galli", "Broken dreams", "Arquitectura", 24),
    ("Juan", "Cruz", "Espejo 321", "Informatica", 21),
    ("Paco", "Rodripguez", "Georgias del sur 213", "Electronica", 18),
    ("Martina", "Flores", "San Martin 983", "Derecho", 20)
;

insert into prestamo (estudiante_id, libro_id, fecha_prestamo, fecha_devolucion, devuelto) values
	(1, 1, "2022-01-10", "2022-01-24", true),
    (2, 2, "2022-01-10", "2022-01-24", true),
    (3, 5, "2022-05-21", "2022-06-5", true),
    (4, 3, "2022-11-02", null, false),
    (5, 4, "2022-10-25", null, false),
    (3, 3, "2022-03-01", "2022-03-15", true),
    (4, 1, "2022-05-10", "2022-05-24", true),
    (5, 3, "2022-08-07", "2022-08-21", true)
;

-- 1. Listar los datos de los autores.
select * from autor;
-- 2. Listar nombre y edad de los estudiantes.
select nombre, edad from estudiante;
-- 3. ¿Qué estudiantes pertenecen a la carrera informática?
select id, nombre, apellido, direccion, edad from estudiante where carrera = "Informatica";
-- 4. ¿Qué autores son de nacionalidad francesa o italiana?
select id, nombre from autor where pais = "Francia" or pais = "Italia";
-- 5. ¿Qué libros no son del área de internet?
select * from libro where area != "Internet";
-- 6. Listar los libros de la editorial Salamandra.
select id, titulo, area from libro where editorial = "Salamandra";
-- 7. Listar los datos de los estudiantes cuya edad es mayor al promedio.
select id, nombre, apellido, direccion, carrera, edad from estudiante join (select avg(edad) as a from estudiante) as e where edad > e.a;
-- 8. Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
select nombre from estudiante where apellido like "G%";
-- 9. Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
select autor.nombre from autor join (select libro_autor.autor_id from libro_autor join (select * from libro where titulo = "El Universo: Guía de viaje") l on l.id = libro_autor.libro_id) as ale on autor.id = ale.autor_id;
-- 10. ¿Qué libros se prestaron al lector “Filippo Galli”?
select titulo from libro join (select * from prestamo join (select * from estudiante where nombre = "Filippo" and apellido = "Galli") as idf on idf.id = prestamo.estudiante_id) as lpf on lpf.libro_id = libro.id;
-- 11. Listar el nombre del estudiante de menor edad.
select nombre from estudiante where edad in (select min(edad) from estudiante);
-- 12. Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos. -> mejor el "Deep dive into design patterns" porque ya cree la DB
select nombre from estudiante join (select * from prestamo join (select * from libro where titulo = "Deep dive into design patterns") as idf on idf.id = prestamo.libro_id) as lpf on lpf.estudiante_id = estudiante.id;
-- 13. Listar los libros que pertenecen a la autora J.K. Rowling.
select titulo from libro join (select * from libro_autor join (select * from autor where nombre = "J.K. Rowling") as idj on idj.id = libro_autor.autor_id) as laj on laj.libro_id = libro.id;
-- 14. Listar títulos de los libros que debían devolverse el 16/07/2021. -> mejor el 2022-01-24 porque ya cree la DB
select titulo from libro join prestamo on libro.id = prestamo.libro_id where prestamo.fecha_devolucion = "2022-01-24";
