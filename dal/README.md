# dal 目录

## 此处为 数据访问层，根据链接的不同数据库划分为多个文件夹
```text
一、注意！通常在 Golang 项目，我们都会使用 GORM 作为 ORM 框架来对数据库进行操作，在此项目中，我们为您提供了两种选择：
    1、您可能较熟悉的 GORM，地址：https://gorm.io/zh_CN/docs/
    2、您可能不太熟悉的 GORM Gen，地址：https://gorm.io/zh_CN/gen/

二、在此模板项目中，我们会把两种使用示例都提供给您，请您根据自己的使用习惯去选择
    1、如您使用 GORM：
        直接使用 dal 包下以 DB 开头的 *gorm.DB 对象即可
    2、如您使用 GORM/Gen：
        1）完善 db.go 中的 InitGormGen 函数
        2）在入口文件中进行初始化（dal.InitGormGen()）
        
三、使用建议：
    按照链接的不同数据库，分成不同的目录，在目录下创建以下文件目录
```

### cache
```text
缓存相关
```
### dao
```text
封装的数据库查询
```
### model
```text
gorm/gen 生成的数据模型对象（DO NOT EDIT）
```
### query
```text
gorm/gen 生成的CRUD代码（DO NOT EDIT）
```
