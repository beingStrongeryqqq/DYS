类Rediit论坛项目，go实现
该项目是一个类似 Reddit 的社区论坛，采用 Vue.js 作为前端框架，后端基于 Go 和
Gin 框架开发，数据管理使用 MySQL 和 Redis。项目中的关键技术包括：JWT 进行身份认证，Zap 用
于高效的日志记录，Nginx 负责网站的部署和反向代理，Viper 用于配置管理。
论坛允许用户发布帖子、评论和点赞，支持实时的社交互动。Redis 主要用于存储和管理实时数
据，例如帖子热度排行和点赞数量，确保数据快速读取和更新。而 MySQL 则用于存储更为静态的信
息，比如用户资料、帖子内容及相关的元数据。
通过充分利用 Go 的并发优势，以及 Gin 框架的轻量级设计，应用程序具有高性能和良好的扩展
性，支持实时交互。项目还使用 Nginx 进行优化部署，提供负载均衡和反向代理功能，确保在高并发
情况下依然能够提供流畅的用户体验。

该项目只包含后端内容，可用postman测试接口

拉取镜像
docker pull yqim/myapp:latest


docker run -d --name myapp_container yqim/myapp:latest
