create table emojis (
    id integer primary key autoincrement,
    emoji varchar(32) not null,
    created_at datetime not null,
    updated_at datetime not null
);
