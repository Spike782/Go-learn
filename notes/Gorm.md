# GORM

## 什么是gorm

**GORM**（全称：**Go Object-Relational Mapping**）是 **Go 语言中最流行的 ORM 框架**。
 它的主要作用是：**让你用 Go 代码来操作数据库，避免手写繁琐的 SQL**，并且具备以下特性：

 **ORM（对象关系映射）**
 它会自动将 Go 语言的结构体和数据库表进行映射。比如你写一个 `User` 结构体，GORM 会帮你把它对应到数据库的 `users` 表。

**简化数据库操作**
 你可以像写普通的 Go 代码一样写：`db.Create(&user)`、`db.Find(&user)`，而不用写 `INSERT`、`SELECT` 这些 SQL 语句。

**丰富的功能**

- 自动建表 / 自动迁移（`AutoMigrate`）
- 事务支持（`db.Transaction`）
- 预处理语句（`PrepareStmt`）
- 连接池管理
- 支持钩子（Hook）
- 支持多表关联（`has one`, `has many`, `belongs to`）
- 日志和慢查询分析

**跨数据库支持**
 支持 MySQL、Postgres、SQLite、SQL Server 等主流数据库。

## 环境搭建

go mod init 文件名

go get [gorm](https://so.csdn.net/so/search?q=gorm&spm=1001.2101.3001.7020).io/driver/mysql //mysql的驱动

go get "gorm.io/gorm"

## 链接数据库

```go

// 定义一个全局变量db，用于后面数据库的读写操作,通常就放在全局里面
var DB *gorm.DB
 
func init() {
	username := "root"       //账号
	password := "password"   //密码
	host := "IP"             //数据库地址
	port := "3306"           //端口
	Dnname := "dtbase"       //数据库名
	timeout := "10s"         //连接超时，10s
 
	//root:root@tcp(127.0.0.1:3306)/test？
    //mb4兼容emoji表情符号。
	// 想要正确的处理time.Time ，需要带上parseTime参数。
	// loc=Local采用机器本地的时区。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dnname, timeout)
	//连接mysql，获得DB类型实例，用于后面数据库的读写操作
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}
	DB = db
 
	//连接成功
	fmt.Println("连接数据库成功")
}
 
```

可以在连接时自定义配置

```go
	db, err := gorm.Open(mysql.Open(DataSourceName), &gorm.Config{
		PrepareStmt:            true, 
        //执行任何SQL时都会创建一个prepared statement并将其缓存，以提高后续的效率
		SkipDefaultTransaction: true, 
        // 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升。
		NamingStrategy: schema.NamingStrategy{ 
            //覆盖默认的NamingStrategy来更改命名约定
			// TablePrefix:   "t_", 把所有的命名加上前缀"t_", table for `User` would be `t_users`
			SingularTable: true, 
            //表名映射时不加复数，仅是驼峰-->蛇形
			// NoLowerCase:   true,                              // skip the snake_casing of names
			// NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger:                   newLogger, //日志控制
		DryRun:                   false,     
        //true代表生成SQL但不执行，可以用于准备或测试生成的 SQL
		DisableNestedTransaction: true,      
        //在一个事务中使用Transaction方法，GORM会使用 SavePoint(savedPointName)，RollbackTo(savedPointName) 为你提供嵌套事务支持。如果不需要嵌套事务，可以将其禁用
		DisableAutomaticPing:     false,     
        //在完成初始化后，GORM 会自动ping数据库以检查数据库的可用性
	})
```



### 创建表

gorm的命名规则如下：

```go
// gorm.Model:默认情况下，GORM使用ID作为主键，使用结构体名的蛇形复数作为表名，字段名的蛇形作为列名
type User struct {
	//ID	名为ID的字段为默认主键
	Id         int       `gorm:"primarykey;column:id"` //也可以显式指定主键，显式指定表里对应的列名`
	UserId     int       `gorm:"column:uid"`
	Degree     string    //如果不显式指定，表里对应的列名就是转为蛇形的degree
	Keywords   []string  `gorm:"json"`                         //切片转为json，可以对应DB里的char、varchar或text
	CreateTime time.Time `gorm:"column:create_time"`           //在creat时gorm会自动把当前时间赋给createtime
	updateTime time.Time `gorm:"column:update_time;type:date"` //在update时gorm会自动把当前时间赋给updatetime
	//`gorm:"type:date`可以显式指定精确度，此处对应的不是datetime，而是date（天）
	Gender   string
	City     string
	Province string `gorm:"-"` //表里没有这个字段，但结构体里需要有它
}
```

可以通过函数自动创建

```go
DB.AutoMigrate(&User{})
fmt.Println("创建表成功！")
```

### 日志显示

gorm默认日志是只打印错误和慢sql，我们可以设置日志的显示等级

```go
	//日志控制
	//先指定个log文件
	logFile, _ := os.OpenFile("路径名", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer，可以输出到文件，也可以输出到os.Stdout
		//日志配置
		logger.Config{
			SlowThreshold:             500 * time.Millisecond, //耗时超过此值认定为慢查询
			LogLevel:                  logger.Info,            // LogLevel的最低阈值，Silent为不输出日志
			IgnoreRecordNotFoundError: true,                   // 忽略RecordNotFound这种错误日志
			ParameterizedQueries:      false,                  // true代表SQL日志里不包含参数
			Colorful:                  false,                  // 禁用颜色
		},
	)
```

## 插入

```go
// 根据struct创建
func Create(db *gorm.DB) {
	//插入一条记录
	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海", Keywords: []string{"编程", "golang"}}
	result := db.Create(&user) //必须传指针，因为要给User的主键赋值。主键为0值时Create会自动给主键赋值
	if result.Error != nil {
		slog.Error("插入记录失败", "error", result.Error)
	}
	fmt.Printf("record id is %d\n", user.Id)
	fmt.Printf("影响%d行\n", result.RowsAffected)

	//会话模式
	tx := db.Session(&gorm.Session{SkipHooks: true}) //不执行钩子Hook
	// db := db.Session(&gorm.Session{DryRun: true}) //生成SQL，但不执行
	//一次性插入多条
	user1 := user                    //发生拷贝
	user1.Id = 0                     //把主键置为0
	user1.UserId = rand.IntN(100000) //UserId上有唯一性约束
	user2 := user                    //发生拷贝
	user2.Id = 0                     //把主键置为0
	user2.UserId = rand.IntN(100000)
	users := []*User{&user1, &user2} //切片里的元素也可以不是指针
	result = tx.Create(users)        //一条SQL插入所有数据
	fmt.Printf("影响%d行\n", result.RowsAffected)

	//量太大时分批插入（SQL语句的长度是有上限的，同时避免长时间阻塞）
	batchSize := 1 //通常为几百
	user3 := user
	user3.Id = 0
	user3.UserId = rand.IntN(100000)
	user4 := user3
	user4.Id = 0
	db.CreateInBatches([]*User{&user3, &user4}, batchSize) //一个批次一条SQL。且所有批次被放到一个事务中来执行。由于user4插不进去，所以user3也会回滚。但如果设置了SkipDefaultTransaction就没有事务
}

// 根据map创建
func CreateByMap(db *gorm.DB) {
	//插入一条记录
	db.Model(User{}).Create(map[string]any{
		"uid": rand.IntN(100000), "degree": "本科", "gender": "男", "city": "上海",
	})

	//一次性插入多条
	db.Model(User{}).Create([]map[string]any{
		{"uid": rand.IntN(100000), "degree": "本科", "gender": "男", "city": "北京"},
		{"uid": rand.IntN(100000), "degree": "本科", "gender": "男", "city": "深圳"},
	})
}
```

## 钩子函数

用来在插入等操作之前/之后来做一些操作（如写日志等）

```go
/*
如果任何钩子回调返回错误，GORM将停止后续的操作并回滚事务。

Create时钩子的执行时机：
// 开始事务
BeforeSave
BeforeCreate
// 关联前的 save
// 插入记录至 db
// 关联后的 save
AfterCreate
AfterSave
// 提交或回滚事务
*/

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook BeforeSave")
	return nil
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook BeforeCreate")
	return nil
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook AfterCreate")
	return nil
}

func (u *User) AfterSave(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook AfterSave")
	return nil
}

```

## 删除

```go
func Delete(db *gorm.DB) {
	tx := db.Where("degree=?", "本科").Delete(User{})
	fmt.Printf("删除%d行\n", tx.RowsAffected)

	var user User = User{Id: 10}
	db.Delete(user) //暗含的Where条件是id=10

	db.Delete(User{}, 1)              //暗含的Where条件是id=1
	db.Delete(User{}, []int{1, 2, 3}) //暗含的Where条件是id IN (1,2,3)
}

/*
如果任何钩子回调返回错误，GORM将停止后续的操作并回滚事务。

Delete时钩子的执行时机：
// 开始事务
BeforeDelete
// 删除 db 中的数据
AfterDelete
// 提交或回滚事务
*/

```

## 更新数据

有save和update两种

```go
// Save会保存所有的字段，即使字段是零值。主键为0时Save相当于Create
func Save(db *gorm.DB) {
	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海"}
	db.Save(&user) //主键为0值，Save相当于Create

	var user2 User
	db.Last(&user2)
	user2.Degree = "硕士"
	db.Save(&user2) //必须传指针
}
```

```go
// Update指定需要更新的列
func Update(db *gorm.DB) {
    // 根据map更新
    tx := db.Model(&User{}). //必须传指针
                Where("city=?", "北京").Updates(
       map[string]any{"degree": "硕士", "gender": "男"},
    )
    fmt.Printf("更新了%d行\n", tx.RowsAffected)

    //根据结构体更新，只会更新非0值
    db.Model(&User{}). //必须传指针
             Where("city=?", "北京").Updates(
       User{Degree: "本科", Gender: "男", Id: 1},
    )
    fmt.Printf("更新了%d行\n", tx.RowsAffected)
}

/*
如果任何钩子回调返回错误，GORM将停止后续的操作并回滚事务。

Update时钩子的执行时机：
// 开始事务
BeforeSave
BeforeUpdate
// 关联前的 save
// 更新 db
// 关联后的 save
AfterUpdate
AfterSave
// 提交或回滚事务
*/
```

## 查询

```go
// 查询
func Read(db *gorm.DB) {
	user := User{City: "HongKong", Id: 3} //Id会自动放到Where条件里，其他非0字段不会
	tx := db.
		Select("uid,city,gender,keywords").       //参数也可以这样传"uid","city","gender"或者[]string{"uid","city","gender"}。没有Select时默认为select *
		Where("uid>100 and degree='大专'").         //容易发生SQL注入攻击
		Where("city in ?", []string{"北京", "上海"}). //多个Where之间是and关系
		Where("degree like ?", "%科").
		Or("gender=?", "女"). //用?占位，避免发生SQL注入攻击
		Order("id desc, uid").
		Order("city").
		Offset(3).
		Limit(1).
		First(&user) //Find可以传一个结构体，也可以传结构体切片。Take、First、Last查不到结果时会返回gorm.ErrRecordNotFound，但Find不会，Find查无结果时就不去修改结构体
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Printf("read结果：%+v\n", user)
		} else {
			slog.Info("查无结果", "user", user)
		}
	}

	var user2 *User //不同于var user2 User，还没申请内存空间
	// 通过反射给user2赋值时发现还没给user2申请好内存空间
	tx = db.Find(user2) //error: invalid value, should be pointer to struct or slice
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	}

	var user3 *User = new(User)
	tx = db.Find(user3)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Printf("read结果：%+v\n", user3)
		} else {
			slog.Info("查无结果", "user", user3)
		}
	}

	var users []User
	tx = db.Limit(3).Find(&users) //要修改切片的长度，所以要传切片的指针
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Println("多个read结果")
			for _, u := range users {
				fmt.Printf("%+v\n", u)
			}
		} else {
			slog.Info("查无结果")
		}
	}

	user4 := User{Id: 23212} //给主键赋值
	tx = db.Find(&user4)     //主键不为0值时暗含了一个where条件：id=47
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Printf("read结果：%+v\n", user4)
		} else {
			slog.Info("查无结果")
		}
	}

	// SELECT * FROM `user` USE INDEX (`id`,`idx_uid`) WHERE uid>0
	db.Where("uid>0").
		Clauses(hints.UseIndex("id", "idx_uid")). //给mysql一个建议的索引范围，这个范围之外的索引mysql就不再考虑了
		Find(&users)
	// SELECT * FROM `user` FORCE INDEX (`idx_uid`) WHERE uid>0
	db.Where("uid>0").
		Clauses(hints.ForceIndex("idx_uid")). //强制mysql使用某个索引
		Find(&users)
}

// 基于统计的查询
func ReadWithStatistics(db *gorm.DB) {
	type Result struct {
		City string
		Mid  float64
	}

	var results []Result
	db.Model(User{}).Select("city,avg(id) as mid").Group("city").Having("mid>0").Find(&results)
	fmt.Println("group by having查询结果：")
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	db.Table("user").Distinct("city").Find(&results)
	fmt.Println("distinct查询结果：")
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	var count int64
	db.Table("user").Where("city=?", "北京").Count(&count)
	fmt.Printf("count=%d\n", count)
}

/*
如果任何钩子回调返回错误，GORM将停止后续的操作并回滚事务。

查询时钩子的执行时机：
// 从 db 中加载数据
// Preloading (eager loading)
AfterFind
*/

```

