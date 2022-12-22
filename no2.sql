-- 2 A
select to_char(created_at,'Month') as tx_month, count(id) as mon from "user" GROUP BY tx_month;

-- 2 B

-- 2 C

-- 2 D
select a.id, a.username, a.fullname, sum(b.total_item) as total_watchlist 
	from "user" as a
	LEFT JOIN (select user_id , count(id) as total_item from  watchlist group by user_id ) as b on a.id = b.user_id
	group by a.id, a.fullname
	order by a.fullname asc

-- 2 E
select 
  b.tx_month, 
  RANK () OVER (
    ORDER BY 
      b.total_item
  ) rank, 
  a.title, 
  b.total_item 
from 
  movie as a 
  LEFT JOIN (
    select 
      movie_id, 
      to_char(created_at, 'Month') as tx_month, 
      count(id) as total_item 
    from 
      watchlist_movie 
    group by 
      movie_id, 
      tx_month
  ) as b on a.id_external = b.movie_id 
group by 
  a.title, 
  b.total_item, 
  b.tx_month;

