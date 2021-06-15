
-- ----------------------------
-- 新建数据库和用户名
-- ----------------------------
create database `ubook` default character set utf8mb4 collate utf8mb4_unicode_ci;
use `ubook`;
create user `ubook`@`%` identified by 'ubook';
grant all privileges on `ubook`.* to `ubook`@`%`;
flush privileges;