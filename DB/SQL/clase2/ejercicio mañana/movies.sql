use movies_db;

-- Mostrar el título y el nombre del género de todas las series.
select movies.title, genres.name as genre from movies join genres on movies.genre_id = genres.id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
select m.f as 'First name', m.l as 'Last name', episodes.title as 'Episode name' from (select actors.first_name as f, actors.last_name as l, actor_episode.episode_id as e from actors join actor_episode on actors.id = actor_episode.actor_id) as m join episodes on episodes.id = m.e order by m.f, m.l;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select series.title, count(seasons.serie_id) as number_of_seasons from series join seasons on series.id = seasons.serie_id group by seasons.serie_id;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select genres.name as genre, count(movies.genre_id) as count from genres join movies on movies.genre_id = genres.id group by genres.id having count >= 3;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select actors.first_name, actors.last_name from actors join (select actor_movie.actor_id as actor_id, actor_movie.movie_id, msw.title from actor_movie join (select id, title from movies where title like "%La Guerra de las galaxias%") as msw on msw.id = actor_movie.movie_id) as amsw on actors.id = amsw.actor_id;
