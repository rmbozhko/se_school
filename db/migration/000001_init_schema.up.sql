create domain email as text
    check ( length(value) = 0 or value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );
comment on domain email is 'Регулярний вираз специфікації HTML5 для type=email';

create table if not exists subscriptions (
    id serial primary key,
    email email unique not null
);
