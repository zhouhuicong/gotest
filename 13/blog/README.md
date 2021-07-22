#练习kratos的 example的 blog 


# How to run this blog example server
1. You should ensure that your mysql server is running.
2. Ensure that the database named `testdb` has been created, 
   otherwise you should execute the following database script:
```mysql
create database testdb;
```
3. Modify the `configs/config.yaml` file and add your mysql information in the data source:
```yaml
data:
  database:
    driver: mysql
    source: root:password@tcp(127.0.0.1:3306)/testdb?parseTime=True
```
4. Run your blog server:
```yaml
make all

insert into articles(title,content,created_at,updated_at) values('test','123456',now(),now());
commit;
# 请求测试
http://118.25.11.63:8001/v1/article 
```
