#新建数据库
create  database  `gin-vue-nezha` default character  set utf8 collate utf8_general_ci;
#新建用户
create user 'gin-vue-nezha'@'%' identified by ' gin-vue-nezha@123 ';
#用户赋权
grant all privileges on `gin-vue-nezha`.* to 'gin-vue-nezha'@'%' ;
#赋权生效
flush privileges;
