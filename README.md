# Web Demo

# 依赖管理
<pre>
cd generic-exporter
govendor init
govendor add +e
</pre>

- Build
<pre>
CGO_ENABLED=0 go build -o=bin/web-demo -ldflags "-s -w" -i  main.go
</pre>

- Run
<pre>
docker run --name web-demo --restart unless-stopped -v /etc/localtime:/etc/localtime --net=host -d web-demo:v1.0.0
</pre>