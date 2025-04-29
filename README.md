# kubernetes-Management
# 获取所有用户
curl --location --request GET 'http://localhost:12833/api/users' \
--header 'Content-Type: application/json'

# 注册用户
curl --location --request POST 'http://localhost:12833/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "testuser0",
"password": "testuser0",
"email": "email0@email.com"
}'
# 登录用户
curl --location --request POST 'http://localhost:12833/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "testuser",
"password": "testpassword"
}'


# 通过自动的结构体创建deployment,但是参数太多，只设置了几个参数
# 参见httpServer/types.go下的结构体
curl --location --request POST 'http://localhost:12833/api/v1/k8s/deployment/create' \
--header 'Content-Type: application/json' \
--data-raw '{
"namespace": "default",
"name": "my-deployment",
"image": "nginx:latest",
"replicas": 2,
"labels": {
"app": "my-app"
},
"containers": [
{
"name": "nginx-container",
"image": "nginx:latest"
}
]
}'

# 获取k8s集群信息
curl --location --request GET 'http://localhost:12833/api/v1/k8s/information'

# 获取节点信息
curl --location --request GET 'http://localhost:12833/api/v1/k8s/nodes'

# 通过client-go原生的对象创建deployment，可以传任意所需的参数
curl --location --request POST 'http://localhost:12833/api/v1/k8s/deployment/create' \
--header 'Content-Type: application/json' \
--data-raw '{
"namespace": "default",
"name": "my-deployment",
"image": "nginx:latest",
"replicas": 2,
"labels": {
"app": "my-app"
},
"containers": [
{
"name": "nginx-container",
"image": "nginx:latest"
}
]
}'