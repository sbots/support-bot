drop table if exists customers;
drop table if exists messages;
drop table if exists chats;
alter table bots
    alter column type type text;
drop type messaging_platform;
drop type message_type;