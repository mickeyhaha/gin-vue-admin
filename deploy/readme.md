1. 安装并初始化msyql 5.7.* Community Server: https://www.cnblogs.com/ericli-ericli/p/6916285.html

2. 打开cmd，运行：
   mysql -uroot -p
   CREATE DATABASE IF NOT EXISTS gva DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

3. 打开cmd，到deploy目录，运行：
   mysql -uroot -proot gva < db/dump.sql

4. 解压nginx-1.20.0_configed.zip

4. 将web/dist.zip 解压，放到上面nginx-1.20.0/html目录

5. 修改nginx-1.20.0/conf/nginx.conf

server_name  localhost 127.0.0.1 121.36.196.22;
第三个IP改为实际服务器IP地址

6. 点击运行nginx.exe

7. 修改server/config.yaml的mysql和mssql部分，修改数据库名，用户名，密码

8. 点击运行server/main.exe

9. 打开 http://127.0.0.1:9090

