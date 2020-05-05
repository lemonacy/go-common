create database go_common default character set utf8mb4 default collate utf8mb4_general_ci;

use go_common;

create table t_user (
  id int not null primary key auto_increment,
  name varchar(64) not null,
  age int not null default 0,
  created_time timestamp default current_timestamp,
  created_by varchar(64) not null default 'system',
  modified_time timestamp default current_timestamp on update current_timestamp,
  modified_by varchar(64) not null default 'system'
) engine = InnoDB;
