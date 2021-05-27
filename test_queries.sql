
use drouotDB;
show tables from drouotDB;

select description from articles
where id in (select articleID from bids where userID = 1);

select * from auctions;
select * from users;
select * from articles;
select * from bids;