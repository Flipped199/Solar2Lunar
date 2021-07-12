package calendar

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	baseYear   = 1950 //基准年限
	maxYear    = 2100
	baseDay    = 17
	baseYMonth = 2
)

var (
	heavenlyStems   = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}           //天干
	earthlyBranches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"} //地支
	//zodiac          = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}                                                                                     //对应地支十二生肖
	//solarTerm       = []string{"小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至"} //二十四节气
	monthCn = []string{"正", "二", "三", "四", "五", "六", "七", "八", "九", "十", "冬", "腊"}
	dateCn  = []string{"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}

	lunarInfo = [][]int{{0, 2, 17, 27808}, {0, 2, 6, 46416}, {5, 1, 27, 21928}, {0, 2, 14, 19872}, {0, 2, 3, 42416}, {3, 1, 24, 21176}, {0, 2, 12, 21168}, {8, 1, 31, 43344}, {0, 2, 18, 59728}, {0, 2, 8, 27296}, {6, 1, 28, 44368}, {0, 2, 15, 43856}, {0, 2, 5, 19296}, {4, 1, 25, 42352}, {0, 2, 13, 42352}, {0, 2, 2, 21088}, {3, 1, 21, 59696}, {0, 2, 9, 55632}, {7, 1, 30, 23208}, {0, 2, 17, 22176}, {0, 2, 6, 38608}, {5, 1, 27, 19176}, {0, 2, 15, 19152}, {0, 2, 3, 42192}, {4, 1, 23, 53864}, {0, 2, 11, 53840}, {8, 1, 31, 54568}, {0, 2, 18, 46400}, {0, 2, 7, 46752}, {6, 1, 28, 38608}, {0, 2, 16, 38320}, {0, 2, 5, 18864}, {4, 1, 25, 42168}, {0, 2, 13, 42160}, {10, 2, 2, 45656}, {0, 2, 20, 27216}, {0, 2, 9, 27968}, {6, 1, 29, 44448}, {0, 2, 17, 43872}, {0, 2, 6, 38256}, {5, 1, 27, 18808}, {0, 2, 15, 18800}, {0, 2, 4, 25776}, {3, 1, 23, 27216}, {0, 2, 10, 59984}, {8, 1, 31, 27432}, {0, 2, 19, 23232}, {0, 2, 7, 43872}, {5, 1, 28, 37736}, {0, 2, 16, 37600}, {0, 2, 5, 51552}, {4, 1, 24, 54440}, {0, 2, 12, 54432}, {0, 2, 1, 55888}, {2, 1, 22, 23208}, {0, 2, 9, 22176}, {7, 1, 29, 43736}, {0, 2, 18, 9680}, {0, 2, 7, 37584}, {5, 1, 26, 51544}, {0, 2, 14, 43344}, {0, 2, 3, 46240}, {4, 1, 23, 46416}, {0, 2, 10, 44368}, {9, 1, 31, 21928}, {0, 2, 19, 19360}, {0, 2, 8, 42416}, {6, 1, 28, 21176}, {0, 2, 16, 21168}, {0, 2, 5, 43312}, {4, 1, 25, 29864}, {0, 2, 12, 27296}, {0, 2, 1, 44368}, {2, 1, 22, 19880}, {0, 2, 10, 19296}, {6, 1, 29, 42352}, {0, 2, 17, 42208}, {0, 2, 6, 53856}, {5, 1, 26, 59696}, {0, 2, 13, 54576}, {0, 2, 3, 23200}, {3, 1, 23, 27472}, {0, 2, 11, 38608}, {11, 1, 31, 19176}, {0, 2, 19, 19152}, {0, 2, 8, 42192}, {6, 1, 28, 53848}, {0, 2, 15, 53840}, {0, 2, 4, 54560}, {5, 1, 24, 55968}, {0, 2, 12, 46496}, {0, 2, 1, 22224}, {2, 1, 22, 19160}, {0, 2, 10, 18864}, {7, 1, 30, 42168}, {0, 2, 17, 42160}, {0, 2, 6, 43600}, {5, 1, 26, 46376}, {0, 2, 14, 27936}, {0, 2, 2, 44448}, {3, 1, 23, 21936}, {0, 2, 11, 37744}, {8, 2, 1, 18808}, {0, 2, 19, 18800}, {0, 2, 8, 25776}, {6, 1, 28, 27216}, {0, 2, 15, 59984}, {0, 2, 4, 27424}, {4, 1, 24, 43872}, {0, 2, 12, 43744}, {0, 2, 2, 37600}, {3, 1, 21, 51568}, {0, 2, 9, 51552}, {7, 1, 29, 54440}, {0, 2, 17, 54432}, {0, 2, 5, 55888}, {5, 1, 26, 23208}, {0, 2, 14, 22176}, {0, 2, 3, 42704}, {4, 1, 23, 21224}, {0, 2, 11, 21200}, {8, 1, 31, 43352}, {0, 2, 19, 43344}, {0, 2, 7, 46240}, {6, 1, 27, 46416}, {0, 2, 15, 44368}, {0, 2, 5, 21920}, {4, 1, 24, 42448}, {0, 2, 12, 42416}, {0, 2, 2, 21168}, {3, 1, 22, 43320}, {0, 2, 9, 26928}, {7, 1, 29, 29336}, {0, 2, 17, 27296}, {0, 2, 6, 44368}, {5, 1, 26, 19880}, {0, 2, 14, 19296}, {0, 2, 3, 42352}, {4, 1, 24, 21104}, {0, 2, 10, 53856}, {8, 1, 30, 59696}, {0, 2, 18, 54560}, {0, 2, 7, 55968}, {6, 1, 27, 27472}, {0, 2, 15, 22224}, {0, 2, 5, 19168}, {4, 1, 25, 42216}, {0, 2, 12, 42192}, {0, 2, 1, 53584}, {2, 1, 21, 55592}, {0, 2, 9, 54560}}
)

type Lunar struct {
	HeavenlyStem    string // 天干
	EarthlyBranches string // 地支
	YearNum         int
	YearStr         string
	MonthNum        int
	MonthStr        string
	DayNum          int
	DayStr          string
	IsLeap          bool // 是否闰月
}

type Solar struct {
	Year  int
	Month int
	Day   int
}

// SolarToLunar 公历转农历
// 闰年数字采用负数表示
func SolarToLunar(solar *Solar) (*Lunar, error) {
	y, m, d := solar.Year, solar.Month, solar.Day

	if y < baseYear || y > maxYear || m < 0 || m > 12 || d < 0 || d > 31 {
		return nil, errors.New(fmt.Sprintf("日期超出范围[%d-01-01 - %d-01-01]", baseYear, maxYear))
	}
	// 1950 年，农历1950正月初一为公历1950 2月17
	// 以它为基准进行计算
	baseTime := time.Date(baseYear, baseYMonth, baseDay, 0, 0, 0, 0, time.Local)
	// 目标日期
	targetTime := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	// 间隔天数
	betweenDays := int(math.Abs(baseTime.Sub(targetTime).Hours() / 24))

	lunarY := 1950
	lunarM := 1
	lunarD := 1
	leapMonth := false

	for i := 0; i < len(lunarInfo); i++ {
		days, months := getLunarYearDays(lunarY)
		betweenDays -= days

		if betweenDays <= 0 {
			betweenDays += days
			for j := 0; j < len(months); j++ {
				betweenDays -= months[j]
				if betweenDays < 0 {
					lunarM = j + 1
					lunarD = months[j] + betweenDays + 1
					break
				} else if betweenDays == 0 {
					lunarM = j + 2
					lunarD = 1
					break
				}
			}
			break
		}
		lunarY++
	}

	// 判断是否闰年，是的话月份如果大于等于闰月要减1，小于闰月不管
	lunarIndex := lunarY - baseYear
	if lunarInfo[lunarIndex][0] != 0 {
		if lunarM >= lunarInfo[lunarIndex][0]+1 {
			lunarM -= 1
			if lunarM == lunarInfo[lunarIndex][0] {
				leapMonth = true
			}
		}
	}

	mStr := monthCn[lunarM-1] + "月"

	if leapMonth {
		mStr = "闰" + mStr
	}

	h := getHeavenlyStems(lunarY)
	e := getEarthlyBranches(lunarY)
	lunar := &Lunar{
		HeavenlyStem:    h,
		EarthlyBranches: e,
		YearNum:         lunarY,
		YearStr:         h + e + "年",
		MonthNum:        lunarM,
		MonthStr:        mStr,
		DayNum:          lunarD,
		DayStr:          dateCn[lunarD-1],
		IsLeap:          leapMonth,
	}

	return lunar, nil
}

// LunarToSolar 农历转公历
// 闰年数字采用负数表示
func LunarToSolar(lunar *Lunar) (*Solar, error) {
	if lunar.YearNum < baseYear || lunar.YearNum > maxYear || lunar.MonthNum > 12 || lunar.DayNum < 0 || lunar.DayNum > 31 {
		return nil, errors.New(fmt.Sprintf("日期超出范围[%d-01-01 - %d-01-01]", baseYear, maxYear))
	}
	// 计算农历日距1950年正月初一（1950。2。17）间隔多少天
	lunarY := lunar.YearNum
	lunarM := lunar.MonthNum
	lunarD := lunar.DayNum

	betweenDays := -1

	days, months := getLunarYearDays(lunarY)

	// 判断月份是否是闰月
	if lunar.IsLeap {
		lunarM = lunarM + 1
	}
	// 加上months day
	for i := 1; i < lunarM; i++ {
		betweenDays += months[i-1]
	}

	// 加上days
	betweenDays += lunarD

	for lunarY > baseYear {
		days, _ = getLunarYearDays(lunarY - 1)
		betweenDays += days
		lunarY--
	}

	baseTime := time.Date(baseYear, baseYMonth, baseDay, 0, 0, 0, 0, time.Local)
	targetTime := baseTime.Add(time.Hour * 24 * time.Duration(betweenDays))

	return &Solar{
		Year:  targetTime.Year(),
		Month: int(targetTime.Month()),
		Day:   targetTime.Day(),
	}, nil
}

// getLunarYearDays 获取农历年的总天数
func getLunarYearDays(year int) (int, []int) {
	days := 0
	var monthSlice []int
	// 将年份转二进制，左补0到16位，从左往右取12｜13位,1为大月,0为小月
	lunar := lunarInfo[year-baseYear]
	yearMonth := 12
	lunarData := strconv.FormatInt(int64(lunar[3]), 2)

	leftZero := strings.Repeat("0", 16-len(lunarData))

	if lunar[0] != 0 {
		yearMonth = 13
	}
	// 循环取前12｜13位计算天数
	for _, i := range (leftZero + lunarData)[:yearMonth] {
		if i == 48 { // 48 是 0 的ascii编码
			days += 29
			monthSlice = append(monthSlice, 29)
		} else {
			days += 30
			monthSlice = append(monthSlice, 30)
		}
	}

	return days, monthSlice
}

// getHeavenlyStems 计算天干 (年份- 3）/10余数对天干：如1894-3=1891 ，1891除以10余数是1即为甲
// 简化后的天干地支：甲、乙、丙、丁、戊、己、庚、辛、壬、癸 称为十天干
func getHeavenlyStems(y int) string {
	return heavenlyStems[(y-3)%10-1]
}

// getEarthlyBranches 计算地支 （年份- 3）/12余数对地支：如1894-3=1891 ，1891除以12余数是7即为午
// 子、丑、寅、卯、辰、巳、午、未、申、酉、戌、亥 称为十二地支。
func getEarthlyBranches(y int) string {
	return earthlyBranches[(y-3)%12-1]
}

func (l *Lunar) String() string {
	return fmt.Sprintf("%d-%02d-%02d", l.YearNum, l.MonthNum, l.DayNum)
}

func (l *Lunar) Chinese() string {
	return fmt.Sprintf("%s %s%s", l.YearStr, l.MonthStr, l.DayStr)
}

func (s *Solar) String() string {
	return fmt.Sprintf("%d-%02d-%02d", s.Year, s.Month, s.Day)
}
