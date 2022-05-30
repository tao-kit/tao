# 配置项管理

| 名称              | 是否可选 | 说明                                          |
|-------------------|----------|-----------------------------------------------|
| namingFormat      | YES      | 文件名称格式化符                      |

# naming-format
`namingFormat`可以用于对生成代码的文件名称进行格式化，和日期格式化符（yyyy-MM-dd）类似，在代码生成时可以根据这些配置项的格式化符进行格式化。

## 格式化符(gotao)
格式化符有`go`,`tao`组成，如常见的三种格式化风格你可以这样编写：
* lower: `gotao`
* camel: `goTao`
* snake: `go_tao`

常见格式化符生成示例
源字符：welcome_to_go_tao

| 格式化符   | 格式化结果            | 说明                      |
|------------|-----------------------|---------------------------|
| gotao     | welcometogotao       | 小写                      |
| goTao     | welcomeToGoZero       | 驼峰                      |
| go_tao    | welcome_to_go_tao    | snake                     |
| Go#tao    | Welcome#to#go#tao    | #号分割Title类型          |
| GOTAO     | WELCOMETOGOTAO       | 大写                      |
| \_go#tao_ | \_welcome#to#go#tao_ | 下划线做前后缀，并且#分割 |

错误格式化符示例
* go
* gOZero
* tao
* goZEro
* goZERo
* goZeRo
* tal

# 使用方法
目前可通过在生成api、rpc、model时通过`--style`参数指定format格式，如：
```shell script
taoctl api go test.api -dir . -style gotao
```
```shell script
 taoctl rpc proto -src test.proto -dir . -style go_tao
```
```shell script
taoctl model mysql datasource -url="" -table="*" -dir ./snake -style GoZero
```

# 默认值
当不指定-style时默认值为`gotao`