本期来学一学viper的使用

```Go
go get github.com/spf13/viper
```

## 设置与读取值

```Go
viper.SetDefault("name1", "默认值") // 设置默认值，最低优先级
viper.Set("name2", "显式设置")       // 显式设置的值，最高优先级
fmt.Println(viper.Get("name1"))
fmt.Println(viper.Get("name2"))
fmt.Println(viper.Get("Name2")) // 大小写不敏感
// 优先级高的会覆盖优先级低的
viper.Set("name1", "被覆盖的默认值")
fmt.Println(viper.Get("name1"))
// 同级被后面的覆盖
viper.Set("name1", "被后面的值覆盖")
fmt.Println(viper.Get("name1"))
```

两个知识点

1. key值是大小写不敏感的
2. 设置的值有优先级，优先级高的会覆盖优先级低的，同级的会被后面的覆盖，和css层叠样式表有点像

## 读写配置文件

写配置文件

```Go
viper.Set("name", "枫枫1")
viper.Set("age", 21)
viper.Set("likes", []string{"唱", "跳", "rap"})

type Info struct {
  Addr string `json:"addr"`
  No   string `json:"no"`
}
viper.Set("info", Info{Addr: "长沙", No: "001"})

viper.SetConfigType("yaml")
viper.SafeWriteConfigAs("config.yaml") // 文件存在的时候就不写入了

viper.SetConfigType("json")
viper.WriteConfigAs("config.json")

viper.SetConfigType("toml")
viper.WriteConfigAs("config.toml")
```

读配置文件

```Go
viper.SetConfigType("yaml")
viper.SetConfigFile("config.yaml")

// 读配置
err := viper.ReadInConfig()
if err != nil {
  fmt.Println(err)
  return
}

// 读配置
fmt.Println(viper.Get("name"))

// 读嵌套的配置
fmt.Println(viper.Get("info.addr"))

// 访问数组的配置
fmt.Println(viper.Get("likes.0"))

type Config struct {
  Name string `json:"mapstructure:name"`
  Age  int    `json:"mapstructure:age"`
}
// 映射结构体
var config Config
err = viper.Unmarshal(&config)
if err != nil {
  fmt.Println(err)
  return
}
fmt.Println(config)
```

## 监控配置文件

```Go
viper.SetConfigType("yaml")
viper.SetConfigFile("config.yaml")

viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  // name是文件名，op是变化的类型
  fmt.Println("配置文件发生变化", e.Name, e.Op)
})
select {}
```

可以通过这种方式，实现服务改配置文件实时生效，不用重启服务

## 从环境变量读取配置

```Go
viper.AutomaticEnv()

// 读全部的环境变量
viper.SetEnvPrefix("")

// 读系统的一些环境变量
fmt.Println(viper.Get("USERNAME"))                  // 系统用户名
fmt.Println(viper.Get("USERDOMAIN_ROAMINGPROFILE")) // 系统电脑名称
fmt.Println(viper.Get("ALLUSERSPROFILE"))           // 用户数据目录

// 指定前缀
viper.SetEnvPrefix("ENV")
fmt.Println(viper.Get("version"))
viper.WriteConfig()
```

补充知识

windows环境变量

```Bash
# 对于cmd
# 查看全部环境变量
set
# 设置环境变量
set Key=Value
# 查看某个环境变量
set Key
```

linux环境变量

```Bash
export name=xx
echo $name
```

## 命令行读取配置

我们有时候会带参数运行项目

例如 ./main -db

./main -export data

```Go
var serverPort = pflag.Int("server.port", 8080, "server listen port")
var serverIp = pflag.String("server.ip", "", "server listen ip")

var name string
var age int
pflag.StringVar(&name, "name", "", "name")
pflag.IntVar(&age, "age", 0, "age")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)
fmt.Println(*serverIp, *serverPort)
fmt.Println(name, age)
```

通过双横杠表示参数

```Go
go run main.go --server.port 80 --server.ip 127.0.0.1 --name 枫枫
```

## 远程读取配置

viper还支持从远程读取配置

例如 etcd和consul

```Go
package main

import (
  "fmt"
  "github.com/spf13/viper"
  _ "github.com/spf13/viper/remote"
)

func main() {

  viper.AddRemoteProvider("etcd", "http://192.168.100.107:2379", "config/test")
  viper.SetConfigType("json") // 配置文件的类型

  // 读取远程配置
  err := viper.ReadRemoteConfig()
  if err != nil {
    fmt.Println("Failed to read remote config:", err)
    return
  }

  fmt.Println(viper.Get("name"))
}
docker run --name etcd -d -p 2379:2379 -p 2380:2380 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.3.11 etcd

docker exec -it etcd bash
etcdctl set config/test '{"name":"fengfeng"}'
```