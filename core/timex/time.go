package timex

import (
	"errors"
	"regexp"
	"time"
)

var WeekStartDay = time.Sunday

// TimeFormats default time formats will be parsed as
var TimeFormats = []string{
	"2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2",
	"15:4:5", "15:4", "15",
	"15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05Z0700", "2006-01-02T15:04:05Z07",
	"2006.1.2", "2006.1.2 15:04:05", "2006.01.02", "2006.01.02 15:04:05", "2006.01.02 15:04:05.999999999",
	"1/2/2006", "1/2/2006 15:4:5", "2006/01/02", "20060102", "2006/01/02 15:04:05",
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
	time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}

// Config configuration for t package
type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
}

type Time struct {
	time.Time
	*Config
}

// DefaultConfig default config
var DefaultConfig *Config

func (config *Config) With(t time.Time) *Time {
	return &Time{Time: t, Config: config}
}

func formatTimeToList(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, min, hour, day, int(month), year}
}

func With(t time.Time) *Time {
	config := DefaultConfig
	if config == nil {
		config = &Config{
			WeekStartDay: WeekStartDay,
			TimeFormats:  TimeFormats,
		}
	}

	return &Time{Time: t, Config: config}
}

func New(t time.Time) *Time {
	return With(t)
}

// BeginningOfMinute beginning of minute
func (t *Time) BeginningOfMinute() time.Time {
	return t.Truncate(time.Minute)
}

// BeginningOfHour beginning of hour
func (t *Time) BeginningOfHour() time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Time.Hour(), 0, 0, 0, t.Time.Location())
}

// BeginningOfDay beginning of day
func (t *Time) BeginningOfDay() time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Time.Location())
}

// BeginningOfWeek beginning of week
func (t *Time) BeginningOfWeek() time.Time {
	b := t.BeginningOfDay()
	weekday := int(b.Weekday())

	if t.WeekStartDay != time.Sunday {
		weekStartDayInt := int(t.WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

// BeginningOfMonth beginning of month
func (t *Time) BeginningOfMonth() time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// BeginningOfQuarter beginning of quarter
func (t *Time) BeginningOfQuarter() time.Time {
	month := t.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// BeginningOfHalf beginning of half year
func (t *Time) BeginningOfHalf() time.Time {
	month := t.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// BeginningOfYear BeginningOfYear beginning of year
func (t *Time) BeginningOfYear() time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMinute end of minute
func (t *Time) EndOfMinute() time.Time {
	return t.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func (t *Time) EndOfHour() time.Time {
	return t.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func (t *Time) EndOfDay() time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// EndOfWeek end of week
func (t *Time) EndOfWeek() time.Time {
	return t.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func (t *Time) EndOfMonth() time.Time {
	return t.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfQuarter end of quarter
func (t *Time) EndOfQuarter() time.Time {
	return t.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// EndOfHalf end of half year
func (t *Time) EndOfHalf() time.Time {
	return t.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (t *Time) EndOfYear() time.Time {
	return t.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Monday monday
func (t *Time) Monday() time.Time {
	b := t.BeginningOfDay()
	weekday := int(b.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}

// Sunday sunday
func (t *Time) Sunday() time.Time {
	b := t.BeginningOfDay()
	weekday := int(b.Weekday())
	if weekday == 0 {
		return b
	}
	return t.AddDate(0, 0, (7 - weekday))
}

// EndOfSunday end of sunday
func (t *Time) EndOfSunday() time.Time {
	return New(t.Sunday()).EndOfDay()
}

// Quarter returns the yearly quarter
func (t *Time) Quarter() uint {
	return (uint(t.Month())-1)/3 + 1
}

func (t *Time) parseWithFormat(str string, location *time.Location) (result time.Time, err error) {
	for _, format := range t.TimeFormats {
		result, err = time.ParseInLocation(format, str, location)

		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*|T)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))(\s*$|[Z+-])`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, 2021-07-20T00:59:10Z, 2021-07-20T00:59:10+08:00, 2021-07-20T00:00:10-07:00 etc
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)                // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc

// Parse parse string to time
func (t *Time) Parse(strs ...string) (result time.Time, err error) {
	var (
		setCurrentTime  bool
		parseTime       []int
		currentLocation = t.Location()
		onlyTimeInStr   = true
		currentTime     = formatTimeToList(t.Time)
	)

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if result, err = t.parseWithFormat(str, currentLocation); err == nil {
			location := t.Location()
			parseTime = formatTimeToList(result)

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			result = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = formatTimeToList(result)
		}
	}
	return
}

// MustParse must parse string to time or it will panic
func (t *Time) MustParse(strs ...string) (result time.Time) {
	result, err := t.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return result
}

// Between check time between the begin, end time or not
func (t *Time) Between(begin, end string) bool {
	beginTime := t.MustParse(begin)
	endTime := t.MustParse(end)
	return t.After(beginTime) && t.Before(endTime)
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute() time.Time {
	return With(time.Now()).BeginningOfMinute()
}

// BeginningOfHour beginning of hour
func BeginningOfHour() time.Time {
	return With(time.Now()).BeginningOfHour()
}

// BeginningOfDay beginning of day
func BeginningOfDay() time.Time {
	return With(time.Now()).BeginningOfDay()
}

// BeginningOfWeek beginning of week
func BeginningOfWeek() time.Time {
	return With(time.Now()).BeginningOfWeek()
}

// BeginningOfMonth beginning of month
func BeginningOfMonth() time.Time {
	return With(time.Now()).BeginningOfMonth()
}

// BeginningOfQuarter beginning of quarter
func BeginningOfQuarter() time.Time {
	return With(time.Now()).BeginningOfQuarter()
}

// BeginningOfYear beginning of year
func BeginningOfYear() time.Time {
	return With(time.Now()).BeginningOfYear()
}

// EndOfMinute end of minute
func EndOfMinute() time.Time {
	return With(time.Now()).EndOfMinute()
}

// EndOfHour end of hour
func EndOfHour() time.Time {
	return With(time.Now()).EndOfHour()
}

// EndOfDay end of day
func EndOfDay() time.Time {
	return With(time.Now()).EndOfDay()
}

// EndOfWeek end of week
func EndOfWeek() time.Time {
	return With(time.Now()).EndOfWeek()
}

// EndOfMonth end of month
func EndOfMonth() time.Time {
	return With(time.Now()).EndOfMonth()
}

// EndOfQuarter end of quarter
func EndOfQuarter() time.Time {
	return With(time.Now()).EndOfQuarter()
}

// EndOfYear end of year
func EndOfYear() time.Time {
	return With(time.Now()).EndOfYear()
}

// Monday monday
func Monday() time.Time {
	return With(time.Now()).Monday()
}

// Sunday sunday
func Sunday() time.Time {
	return With(time.Now()).Sunday()
}

// EndOfSunday end of sunday
func EndOfSunday() time.Time {
	return With(time.Now()).EndOfSunday()
}

// Quarter returns the yearly quarter
func Quarter() uint {
	return With(time.Now()).Quarter()
}

// Parse parse string to time
func Parse(strs ...string) (time.Time, error) {
	return With(time.Now()).Parse(strs...)
}

// ParseInLocation parse string to time in location
func ParseInLocation(loc *time.Location, strs ...string) (time.Time, error) {
	return With(time.Now().In(loc)).Parse(strs...)
}

// MustParse must parse string to time or will panic
func MustParse(strs ...string) time.Time {
	return With(time.Now()).MustParse(strs...)
}

// MustParseInLocation must parse string to time in location or will panic
func MustParseInLocation(loc *time.Location, strs ...string) time.Time {
	return With(time.Now().In(loc)).MustParse(strs...)
}

// Between check now between the begin, end time or not
func Between(time1, time2 string) bool {
	return With(time.Now()).Between(time1, time2)
}

// Timestamp returns the timestamp in seconds.
func (t *Time) Timestamp() int64 {
	return t.UnixNano() / 1e9
}

// TimestampMilli returns the timestamp in milliseconds.
func (t *Time) TimestampMilli() int64 {
	return t.UnixNano() / 1e6
}

// TimestampMicro returns the timestamp in microseconds.
func (t *Time) TimestampMicro() int64 {
	return t.UnixNano() / 1e3
}

// TimestampNano returns the timestamp in nanoseconds.
func (t *Time) TimestampNano() int64 {
	return t.UnixNano()
}
