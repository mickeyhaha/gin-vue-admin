1. 安装并初始化msyql 5.7.* Community Server: https://www.cnblogs.com/ericli-ericli/p/6916285.html

2. CREATE DATABASE IF NOT EXISTS gva DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

3. mysql -uroot -p gva < db/dump20120510.sql

4. 将web/dist.zip 解压，放到http服务器根目录

5. 修改server/config.yaml的mysql部分，修改用户名密码

6. 启动server/main.exe

