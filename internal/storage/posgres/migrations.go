package posgres

//todo migrations
/*-- auto-generated definition
CREATE SEQUENCE owner_id_seq;
create table owner
(
id               bigint                   default nextval('owner_id_seq'::regclass) not null
constraint outline_admins_pkey primary key,
name             varchar(255)                                                not null,
chat_id          bigint                                                      not null,
created_at       timestamp with time zone default now()                      not null,
is_outline_admin boolean                  default false                      not null
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

/*-- auto-generated definition
CREATE SEQUENCE outline_server_id_seq;
create table outline_server
(
id           bigint                                 not null
constraint outline_server_pk
primary key,
chat_id      integer                                not null,
api_url      varchar(255)                           not null,
cert_sha_256 varchar                                not null,
created_at   timestamp with time zone default now() not null
);

comment on table outline_server is 'ключи от серверов outline';

alter table outline_server
owner to postgres;

create unique index outline_server_id_uindex
on outline_server (id);

create unique index outline_server_chat_id_uindex
on outline_server (chat_id);*/
