# Gin - layout

一个开箱即用的gin脚手架

## 特性
- 基于 Gin web 框架
- 基于 Gorm 的数据库存储
- 基于 Viper 配置解析，支持动态载入配置文件
- 基于 Zap 的日志实现
- 提供常用的 Jwt ，哈希加密，随机字符串生成等工具

## 库
- Gin
- Gorm
- Zap
- Viper
- go-jwt

## Overview
### Project structure
```json
├─config
├─controller
├─dao
│  └─user
├─errorx
├─etc
├─logger
├─middleware
├─response
├─router
├─service
└─utils
```