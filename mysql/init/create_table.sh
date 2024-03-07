#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

$CMD_MYSQL -e "create table user (
        id char(28) NOT NULL primary key,
        user_name varchar(127) DEFAULT '',
        introduction varchar(1023) DEFAULT '',
        git_hub varchar(255) DEFAULT '',
        register_time timestamp,
        last_online_time timestamp
    );"

$CMD_MYSQL -e "create table blog (
        id char(26) NOT NULL primary key,
        user_id char(28),
        title varchar(127) DEFAULT '',
        content text,
        birth_time timestamp,
        update_time timestamp,
        public tinyint(1) DEFAULT 0,
        likes int DEFAULT 0,
        FOREIGN KEY (user_id) REFERENCES user(id)
    );"

$CMD_MYSQL -e "create table book (
        id char(26) NOT NULL primary key,
        user_id char(28),
        title varchar(127) DEFAULT '',
        introduction varchar(255) DEFAULT '',
        birth_time timestamp,
        update_time timestamp,
        public tinyint(1) DEFAULT 0,
        FOREIGN KEY (user_id) REFERENCES user(id)
    );"

$CMD_MYSQL -e "create table chapter (
        id char(26) NOT NULL primary key,
        book_id char(26),
        title varchar(127) DEFAULT '',
        content text,
        birth_time timestamp,
        update_time timestamp,
        FOREIGN KEY (book_id) REFERENCES book(id)
    );"

$CMD_MYSQL -e "create table video (
        id char(26) NOT NULL primary key,
        user_id char(28),
        title varchar(127) DEFAULT '',
        introduction varchar(255) DEFAULT '',
        url varchar(255) DEFAULT '',
        birth_time timestamp,
        update_time timestamp,
        public tinyint(1) DEFAULT 0,
        FOREIGN KEY (user_id) REFERENCES user(id)
    );"

$CMD_MYSQL -e "create table work (
        id char(26) NOT NULL primary key,
        user_id char(28),
        title varchar(127) DEFAULT '',
        introduction varchar(255) DEFAULT '',
        url varchar(255) DEFAULT '',
        birth_time timestamp,
        update_time timestamp,
        public tinyint(1) DEFAULT 0,
        FOREIGN KEY (user_id) REFERENCES user(id)
    );"

$CMD_MYSQL -e "create table comment_blog (
        id char(26) NOT NULL primary key,
        user_id char(28),
        blog_id char(26),
        content text,
        birth_time timestamp,
        update_time timestamp,
        user_name varchar(127) DEFAULT '',
        FOREIGN KEY (user_id) REFERENCES user(id),
        FOREIGN KEY (blog_id) REFERENCES blog(id)
    );"
