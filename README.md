# 解决Docker容器时区不一致的问题
1、利用Dockerfile创建镜像时，在Dockerfile中加入
```
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
```
2、容器创建时，加入时区挂载选项：-v /etc/localtime:/etc/localtime
```
docker run -d -p 6379:6379 -v /etc/localtime:/etc/localtime --name test-redis redis
```
3、容器已启动时
```
docker exec -it container /bin/sh // 进入交互模式，container为容器ID或名称，下同
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
docker restart container // 重启容器
docker exec container date -R // 查看时区
```