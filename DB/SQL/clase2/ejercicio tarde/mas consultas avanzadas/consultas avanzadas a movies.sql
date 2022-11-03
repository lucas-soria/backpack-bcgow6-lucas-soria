use movies_db;

-- 1. Agregar una película a la tabla movies.
insert into movies(created_at, updated_at, title, rating, awards, release_date, length, genre_id) values
	("2022-11-01", "2022-11-02", "Peli de mentira", 9.5, 3, "2022-11-03", 125, 3)
;
select * from movies where length = 125;
-- 2. Agregar un género a la tabla genres.
insert into genres(created_at, updated_at, `name`, ranking, `active`) values
	("2022-11-01", "2022-11-02", "Genero de mentira", 101, 1)
;
select * from genres where ranking = "101";
-- 3. Asociar a la película del punto 1. con el género creado en el punto 2.
update movies set genre_id = 16 where id = 22;
select * from movies where id = 22;
-- 4. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
update actors set favorite_movie_id = 22 where id = 1;
select * from actors where id = 1;
-- 5. Crear una tabla temporal copia de la tabla movies.
create temporary table temp_movies (select * from movies);
select * from temp_movies;
-- 6. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
delete from temp_movies where awards < 5;
select * from temp_movies;
SET SQL_SAFE_UPDATES = 1;
-- 7. Obtener la lista de todos los géneros que tengan al menos una película.
select genres.* from genres join movies on movies.genre_id = genres.id group by id;
-- 8. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
select actors.* from actors join movies on actors.favorite_movie_id = movies.id where movies.awards > 3;
-- 9. Crear un índice sobre el nombre en la tabla movies.
create index nombre on movies(title);
-- 10. Chequee que el índice fue creado correctamente.
show index from movies;
-- 11. En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
/*
Se suelen hacer consultas basadas en los nombres de las peliculas, por lo que se veria beneficiada.
No se hacen tantas inserciones, por lo que el timpo que se agrega a los queries de escritura no deberian importar.
*/
-- 12. ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
/*
Al igual que las peliculas, las series tambien se pueden beneficiar de un indice en titulo, mas de eso no veo ventajas
*/
