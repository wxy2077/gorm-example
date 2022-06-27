## 数据库迁移

### 方法一(推荐): 使用go-migrate工具管理MySQL表
> 不用关心`model` gorm tag的书写规则，更直观精确的管理表，降低学习成本。

- 1 安装使用
```
go get github.com/rubenv/sql-migrate
```
项目根目录下配置 `dbconfig.yaml`
(注意名字固定写法不能更改)
一般*.yaml等配置文件都被gitignore忽略，只传一个yaml.example文件,防止配置文件泄漏。

- 2 迁移

项目根目录下运行
```shell
sql-migrate up -env=dev-mysql

# 解释
up 指执行 +migrate Up 标注的SQL
-env 参数为环境名字可以自行在dbconfig.yaml配置

```
会在对应数据库多生成一个名为`migrations`的表，来记录管理迁移操作。

- 3 其他基本命令
```shell
# 回退 执行+migrate Down标注的SQL 每次使用只回退一步操作
sql-migrate down -env=dev-mysql


# 查看迁移状态
sql-migrate status -env=dev-mysql
```
- 4 迁移SQL文件命名规则(建议)
> 这一条只是建议，不强制规定。
```shell
# 为了保证文件目录直观性，迁移SQL文件遵循这样命名
年月日时分_操作简单描述.sql
# 如下 给user表加一个索引
202206261030_add_index2user.sql
```


### 方法二: gorm自带迁移功能迁移

项目根目录下运行
```
 go run 02_migrate/migrate.go
```
