create table if not exists infoRSS (
    id serial primary key, 
    title text,
    description text,
    link text,
    published text
);