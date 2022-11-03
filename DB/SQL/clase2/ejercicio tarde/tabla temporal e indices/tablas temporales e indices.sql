use movies_db;

-- Ejercicio 1
select episodes.* from episodes join seasons on episodes.season_id = seasons.id where seasons.serie_id = (select id from series where title = "The Walking Dead");

create temporary table TWD (select episodes.* from episodes join seasons on episodes.season_id = seasons.id where seasons.serie_id = (select id from series where title = "The Walking Dead"));

-- 20 es la primer temposrada
select * from TWD where season_id = 20;

-- Ejercicio 2
create index movie_title on movies(title);
/* ¿Por qué?
Si bien puede que se repitan los nombres (y mas ahora que se hacen remakes),
es raro y las busquedas en esta práctica son por lo general son por nombre,
no por id */
show index from movies;

explain select * from movies where title = "La Guerra de las galaxias: Episodio VI";
explain select * from movies where title like "%Guerra%";

select * from movies;
