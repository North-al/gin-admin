// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {"schemas":{"biz.Permission":{"properties":{"code":{"type":"string"},"created_at":{"type":"string"},"description":{"type":"string"},"id":{"type":"integer"},"name":{"type":"string"},"updated_at":{"type":"string"}},"type":"object"},"biz.Role":{"properties":{"created_at":{"type":"string"},"description":{"type":"string"},"id":{"type":"integer"},"name":{"type":"string"},"permissions":{"items":{"$ref":"#/components/schemas/biz.Permission"},"type":"array","uniqueItems":false},"updated_at":{"type":"string"}},"type":"object"},"biz.Users":{"description":"系统用户信息","properties":{"created_at":{"description":"创建时间","type":"string"},"email":{"description":"邮箱","type":"string"},"id":{"description":"用户id","type":"integer"},"phone":{"description":"手机号","type":"string"},"roles":{"description":"角色","items":{"$ref":"#/components/schemas/biz.Role"},"type":"array","uniqueItems":false},"status":{"description":"状态","type":"boolean"},"updated_at":{"description":"更新时间","type":"string"},"username":{"description":"用户名 唯一","type":"string"}},"type":"object"},"data":{"properties":{"data":{"$ref":"#/components/schemas/biz.Users"}},"type":"object"},"response.Response":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"},"services.LoginParams":{"properties":{"account":{"default":"zhangsan","description":"账号 邮箱/手机号/用户名","type":"string"},"password":{"default":"123456","description":"密码 6-16位","type":"string"}},"required":["account","password"],"type":"object"},"services.RegisterParams":{"properties":{"account":{"default":"zhangsan","description":"账号 邮箱/手机号/用户名","type":"string"},"password":{"default":"123456","description":"密码 6-16位","type":"string"}},"required":["account","password"],"type":"object"}},"securitySchemes":{"ApiKeyAuth":{"in":"header","name":"Authorization","type":"apiKey"}}},
    "info": {"contact":{"email":"1227379879@qq.com","name":"North"},"description":"{{escape .Description}}","termsOfService":"http://swagger.io/terms/","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/user/info":{"get":{"description":"获取用户信息","requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"获取用户信息","tags":["用户模块"]}},"/user/login":{"post":{"description":"用户登录，返回token","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/services.LoginParams"}}},"description":"用户信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"用户登录","tags":["用户模块"]}},"/user/register":{"post":{"description":"注册","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/services.RegisterParams"}}},"description":"用户信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"注册","tags":["用户模块"]}}},
    "openapi": "3.1.0",
    "servers": [
        {"url":"/api"}
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Title:            "权限管理系统",
	Description:      "通过在请求头中添加 `Authorization` 字段进行认证",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
