需要在mysql容器中执行
```bash
# Execute the following commands in the MySQL container:
docker exec -it mysql mysql -uroot -p

# Enter the password when prompted:
# Password: PXDN93VRKUm8TeE7

# Then execute the following SQL commands:
USE mysql;
UPDATE user SET host='%' WHERE user='root';
FLUSH PRIVILEGES;
```
# End of Selection
```