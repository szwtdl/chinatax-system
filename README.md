### 财务系统


### 用户(企业端)


### 财务(财务公司)


### 平台(平台运营)


### 创建数据库

```mysql
    CREATE DATABASE tax_szwtdl_cn
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_general_ci;
```

### 安装redis

```bash
sudo apt install redis-server -y
```



### 上传配置

```bash
scp /Users/pengjian/work/go/platform_szwtdl_com/build/cmd/config.toml root@39.108.50.141:/var/www/tax.szwtdl.cn/

```

### 上传项目
```bash
scp /Users/pengjian/work/go/platform_szwtdl_com/build/cmd/chinatax-linux-prod root@39.108.50.141:/var/www/tax.szwtdl.cn/
```

### 启动项目

```bash
nohup ./chinatax-linux-prod >>logs.log 2>&1 &
```