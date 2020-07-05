-- ----------------------------
-- 删除数据库和用户名
-- ----------------------------
drop database if exists `gin-vue-nezha`;
use mysql;
delete from user where user='gin-vue-nezha' and host='localhost';
flush privileges;

-- ----------------------------
-- 新建数据库和用户名
-- ----------------------------
-- 支持emoji：需要mysql数据库参数： character_set_server=utf8mb4
create database `gin-vue-nezha` default character set utf8mb4 collate utf8mb4_unicode_ci;
use `gin-vue-nezha`;
create user `gin-vue-nezha`@`localhost` identified by 'gin-vue-nezha@123';
grant all privileges on `gin-vue-nezha`.* to `gin-vue-nezha`@`localhost`;
flush privileges;
