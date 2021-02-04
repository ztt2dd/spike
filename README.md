# spike
> go语言秒杀系统demo

## 本地安装步骤
1. 运行go mod tidy命令安装第三方包（goland报红：go modules配置https://goproxy.io）
2. 数据库：sql/spike.sql
3. 配置conf/app.ini文件

## 服务器部署
> 通过docker进行部署

## 项目架构

## 程序逻辑

## 测试流程及工具

## 待解决的问题
1. err还是尽量每层都需要处理
2. interface是否需要使用(直接定义var变量而不是每个方法都去new行不行)
3. 日志系统还需要改造下
4. 还需要加入kafka
6. 订单号
7. 分布式的redis怎么搞