### 1、在logger包最外层新增日志输出方法，使用日志包更方便
#### 新增方法前日志用法
```
import "wayne/pkg/logger"

var Log *logger.Entry
# step1: 根据配置文件初始化日志
logger.UseLogrusWithConfig(logConfig)
# step2: 创建Entry对象
Log = logger.CreateLogger()
# 使用日志
Log.Debug(“this is debug.”)
```
项目的其他包中使用需要导入Log对象所在的包，打印日志需要多写一层包名繁琐不够简洁
```buildoutcfg
import "localpkg/xxx"    # 导入本地初始化logger对象的包

xxx.Log.Debug(“this is debug”)
```

#### 新增方法后日志用法
```
import "wayne/pkg/logger"

# step1: 根据配置文件初始化日志
logger.UseLogrusWithConfig(logConfig)
# step2: 创建Entry对象
logger.CreateLogger()
# 使用日志
logger.Debug(“this is new debug.”)
```

项目的其他包中使用日志直接导入logger包即可使用，而不是导入Entry对象所在的包
```
import "wayne/pkg/logger"

logger.Debug(“this is new debug.”)
```

### 2、修复显示日志输出具体行号信息不准确bug
#### 修复前
```
# 在配置文件中设置是否禁用行号信息显示(WarnLevel以上才会显示)
disable_line_hook = false

import "wayne/pkg/logger"

logger.Error(“this is error.”)
```
输出的行号信息并非logger.Error(“this is error.”)代码所在的文件和行号
> type=error message="this is error." appno="" source="logger/exported.go:117:Error()" timestamp=1608287333

#### 修复后
```
import "wayne/pkg/logger"

logger.Error(“this is new error.”)
```
输出的行号信息即日志所在的文件和行号，方便我们定位bug
> type=error message="this is error." appno="" source="learn/main.go:90:main()" timestamp=1608287513


### 3、修复并非指定输出行号信息的日志级别也输出行号信息bug
#### 修复前
```
# 在配置文件中设置是否禁用行号信息显示(WarnLevel以上才会显示)
disable_line_hook = false

import "wayne/pkg/logger"

logger.Error(“this is error.”)    # Error级别会输出行号信息

logger.Debug(“this is debug”)    # Debug级别不会输出行号信息，可是输出了前面error行号信息
```
Error级别会输出行号信息
> type=error message="this is error." appno="" source="learn/main.go:90:main()" timestamp=1608287513

Debug级别不会输出行号信息，可是输出了Error行号的信息
> type=debug message="this is debug." appno="" source="learn/main.go:90:main()" timestamp=1608287534


#### 修复后
```
logger.Error(“this is new error.”)    # Error级别会输出行号信息

logger.Debug(“this is new debug”)    # Debug级别不再输出行号信息，因为没有达到输出行号信息的日志级别
```
Error级别会输出行号信息
> type=error message="this is new error." appno="" source="learn/main.go:90:main()" timestamp=1608287513

Debug级别不再输出行号信息，因为没有达到输出行号信息的日志级别
> type=debug message="this is new debug."  appno=""  timestamp=1608287534


### 4、日志中不输出appno和时间戳
- appno: 作为一个共用库，应该有足够的自由度交给使用者决定是否每行日志都显示appno信息
- 时间戳: 作为日志中输出的时间，应该让看日志的人可以一目了然知道日志的输出时间，而不是时间戳
```
# 日志时间显示格式（可选值timestamp|time，默认为timestamp）
time_format = "timestamp"

import "wayne/pkg/logger"

logger.UseLogrusWithConfig(logConfig)
# 设置日志中不输出appno(注：需要在创建Entry对象之前设置该属性, 即在CreateLogger方法前设置)
logger.DisableAppNo = true
logger.CreateLogger()

logger.Debug(“this is debug”)
```
不再输出appno和时间戳
> time="2020-12-18 10:44:57.392" type=debug message="this is debug."
