# Solar2Lunar
一个简单的阳历-农历互转程序

> 时间区间 1950 - 2100

使用示例：
```go
// 阳历转农历 &Solar{Year, Month, Day}
lunar, err := SolarToLunar(&Solar{Year: 2020, Month: 6, Day: 19}) // 2020 闰 4 28
log.Println(err) // nil
log.Printf("%#v\n", lunar) // &Lunar{HeavenlyStem:"庚", EarthlyBranches:"子", YearNum:2020, YearStr:"庚子年", MonthNum:4, MonthStr:"闰四月", DayNum:28, DayStr:"廿八", IsLeap:true}
log.Println(lunar.String()) // 2020-04-28
  
// 农历转阳历 &calendar.Lunar{YearNum,	MonthNum,	IsLeap, DayNum}
solar, _ := LunarToSolar(&Lunar{
  YearNum:  2020,
  MonthNum: 4,
  IsLeap:   false,
  DayNum:   28,
}) // 2020 5 20
log.Printf("%v\n", solar) // 2020-05-20

solar, _ = LunarToSolar(&Lunar{
  YearNum:  2020,
  MonthNum: 4,
  IsLeap:   true,
  DayNum:   28,
}) // 2020 6 19
log.Printf("%v\n", solar) // 2020-06-19
```

### 如在使用过程中发现任何bug，欢迎issue，小白一枚，代码写的烂，希望不要嫌弃，如有可能，大佬可以贴一份优化后的代码共同学习
