## Project setup
```
npm cache clean --force
npm install
npm install vue-loader vue-template-compiler
```

###安装mysql：
https://www.runoob.com/mysql/mysql-install.html
set password for root@localhost = password('root'); 
GRANT ALL PRIVILEGES ON *.* TO root@"%" IDENTIFIED BY "root";
flush privileges;

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
