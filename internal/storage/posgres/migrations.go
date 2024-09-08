package posgres

//todo migrations
/*-- auto-generated definition
create table owner
(
name             varchar(255)                                                                not null,
chat_id          bigint                                                                      not null,
id               bigint                   default nextval('outline_admins_id_seq'::regclass) not null
constraint outline_admins_pkey
primary key,
created_at       timestamp with time zone default now()                                      not null,
is_outline_admin boolean                  default false                                      not null
);

comment on table owner is 'клиенты бота';

comment on column owner.name is 'тг юзернейм';

comment on column owner.id is 'администраторы серверов outline';

comment on column owner.is_outline_admin is 'проставляется если у пользователя есть добавленые серверы';

alter table owner
owner to postgres;

create index owner_name_index
on owner (name);

comment on index owner_name_index is 'телеграм юзер нейм';

create index owner_is_outline_admin_index
on owner (is_outline_admin);

create unique index owner_chat_id_uindex
on owner (chat_id);

comment on index owner_chat_id_uindex is 'id чата с пользователем, основной  его идентификатор';*/
