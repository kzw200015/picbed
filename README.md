# picbed

一个简单的图床

## 特点
- 无数据库
- 无外部依赖，仅一个二进制，简单易部署
- 支持删除上传的图片（上传后三分钟内）
- 一键复制各类链接
- 低资源占用
- 简单的防盗链

## 使用方法
### 配置文件
- 参照config.example.yml编写配置文件
- `./picbed -c /path/to/config`
### 命令行参数
```
./picbed -p port -s /path/to/images
```
**不支持指定Referer，默认为"*"，即允许所有**
### 环境变量
- `PB_PORT` 端口
- `PB_PORT_IMG_SRC` 图片目录
- `PB_PORT_DEBUG` 调试模式

**不支持指定Referer，默认为"*"，即允许所有**

三者可结合使用 优先级：环境变量 > 配置文件 > 命令行参数