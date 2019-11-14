create table posts (
    id serial primary key,
    content text,
    author varchar(255)
);

 -- 在终端执行psql -U gwp -f setup.sql -d gwp
 -- 每次执行后续代码之前都要执行这条命令，以便刷新和设置数据库