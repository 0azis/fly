CREATE TABLE chat_users(
    userid serial not null,
    email varchar(50) not null,
    password varchar(60) not null,
    primary key (userid)
);


CREATE TABLE chats(
    chatid varchar(50) not null,
    user_one integer not null,
    user_two integer not null
);

CREATE TABLE messages(
    messageid serial not null,
    chatid varchar(50) not null,
    from integer not null,
    message varchar(50) not null,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    primary key (messageid)
);