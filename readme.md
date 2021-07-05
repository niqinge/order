要求：

0. 数据模型： demo_order

id
order_no string
user_name string
amount float64
status string
file_url string



1. 项目结构
   your_project/
   db/ (数据库配置， GORM配置)
   model/ (model)
   service/ (service)
   handler/ (Gin handler)
   router/ (Gin router)


2. 根据数据模型，使用 Gin、Gorm、数据库使用mysql

3. 需要实现 4个需求：
   1） 创建 demo_order
   2） 更新 demo_order （amount、status、file_url）
   3） 获取 demo_order 详情
   4） 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）

4. 根据需求编写 Restful Api接口

5. service 下所有代码需要编写测试

6. 需要实现一个 Gorm 事物 SQL

7. 需要实现一个Gin 文件上传的功能， 并且更新 demo_order ： file_url

8. 同时实现Gin 文件下载

9. 使用 https://github.com/tealeg/xlsx 这个库 将demo_order 所有数据以excel形式导出来(可以下载)
