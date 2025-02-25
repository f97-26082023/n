package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/f97/gofire/pkg/errs"
)

const (
	longDateTimeFormat              = "2006-01-02 15:04:05"
	longDateTimeWithoutSecondFormat = "2006-01-02 15:04"
	shortDateTimeFormat             = "2006-1-2 15:4:5"
	yearMonthDateTimeFormat         = "2006-01"
	westernmostTimezoneUtcOffset    = -720 // Etc/GMT+12 (UTC-12:00)
	easternmostTimezoneUtcOffset    = 840  // Pacific/Kiritimati (UTC+14:00)
)

// FormatUnixTimeToLongDateTimeInServerTimezone returns a textual representation of the unix time formatted by long date time format
func FormatUnixTimeToLongDateTimeInServerTimezone(unixTime int64) string {
	return parseFromUnixTime(unixTime).Format(longDateTimeFormat)
}

// FormatUnixTimeToLongDateTimeWithoutSecond returns a textual representation of the unix time formatted by long date time format (no second)
func FormatUnixTimeToLongDateTimeWithoutSecond(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(longDateTimeWithoutSecondFormat)
}

// FormatUnixTimeToYearMonth returns year and month of specified unix time
func FormatUnixTimeToYearMonth(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(yearMonthDateTimeFormat)
}

// ParseFromLongDateTimeToMinUnixTime parses a formatted string in long date time format to minimal unix time (the westernmost timezone)
func ParseFromLongDateTimeToMinUnixTime(t string) (time.Time, error) {
	timezone := time.FixedZone("Timezone", easternmostTimezoneUtcOffset*60)
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}

// ParseFromLongDateTimeToMaxUnixTime parses a formatted string in long date time format to maximal unix time (the easternmost timezone)
func ParseFromLongDateTimeToMaxUnixTime(t string) (time.Time, error) {
	timezone := time.FixedZone("Timezone", westernmostTimezoneUtcOffset*60)
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}

// ParseFromLongDateTime parses a formatted string in long date time format
func ParseFromLongDateTime(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}

// ParseFromLongDateTimeWithoutSecond parses a formatted string in long date time format (no second)
func ParseFromLongDateTimeWithoutSecond(t string, timezone *time.Location) (time.Time, error) {
	return time.ParseInLocation(longDateTimeWithoutSecondFormat, t, timezone)
}

// ParseFromShortDateTime parses a formatted string in short date time format
func ParseFromShortDateTime(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	return time.ParseInLocation(shortDateTimeFormat, t, timezone)
}

func ParseFromElapsedSeconds(elapsedSeconds int) (string, error) {
	if elapsedSeconds < 0 || elapsedSeconds >= 86400 {
		return "", errs.ErrFormatInvalid
	}

	second := elapsedSeconds % 60
	elapsedSeconds = elapsedSeconds - second
	elapsedSeconds = elapsedSeconds / 60

	minute := elapsedSeconds % 60
	elapsedSeconds = elapsedSeconds - minute
	elapsedSeconds = elapsedSeconds / 60

	hour := elapsedSeconds

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), nil
}

// IsUnixTimeEqualsYearAndMonth returns whether year and month of the unix time are equals to the specified year and month
func IsUnixTimeEqualsYearAndMonth(unixTime int64, timezone *time.Location, year int32, month int32) bool {
	date := parseFromUnixTime(unixTime).In(timezone)
	return date.Year() == int(year) && int(date.Month()) == int(month)
}

// GetTimezoneOffsetMinutes returns offset minutes according specified timezone
func GetTimezoneOffsetMinutes(timezone *time.Location) int16 {
	_, tzOffset := time.Now().In(timezone).Zone()
	tzMinuteOffset := int16(tzOffset / 60)

	return tzMinuteOffset
}

// FormatTimezoneOffset returns "+/-HH:MM" format of timezone
func FormatTimezoneOffset(timezone *time.Location) string {
	tzMinutesOffset := GetTimezoneOffsetMinutes(timezone)

	sign := "+"
	hourAbsOffset := tzMinutesOffset / 60
	minuteAbsOffset := tzMinutesOffset % 60

	if hourAbsOffset < 0 {
		sign = "-"
		hourAbsOffset = -hourAbsOffset
		minuteAbsOffset = -minuteAbsOffset
	}

	return fmt.Sprintf("%s%02d:%02d", sign, hourAbsOffset, minuteAbsOffset)
}

// ParseFromTimezoneOffset parses a formatted string in timezone offset format
func ParseFromTimezoneOffset(tzOffset string) (*time.Location, error) {
	if len(tzOffset) != 6 { // +/-HH:MM
		return nil, errs.ErrFormatInvalid
	}

	sign := tzOffset[0]

	if sign != '-' && sign != '+' {
		return nil, errs.ErrFormatInvalid
	}

	offsets := strings.Split(tzOffset[1:], ":")

	if len(offsets) != 2 {
		return nil, errs.ErrFormatInvalid
	}

	hourAbsOffset, err := StringToInt(offsets[0])

	if err != nil {
		return nil, err
	}

	minuteAbsOffset, err := StringToInt(offsets[1])

	if err != nil {
		return nil, err
	}

	totalMinuteOffset := hourAbsOffset*60 + minuteAbsOffset

	if sign == '-' {
		totalMinuteOffset = -totalMinuteOffset
	}

	totalOffset := totalMinuteOffset * 60
	return time.FixedZone("Timezone", totalOffset), nil
}

// GetMinTransactionTimeFromUnixTime returns the minimum transaction time from unix time
func GetMinTransactionTimeFromUnixTime(unixTime int64) int64 {
	return unixTime * 1000
}

// GetMaxTransactionTimeFromUnixTime returns the maximum transaction time from unix time
func GetMaxTransactionTimeFromUnixTime(unixTime int64) int64 {
	return unixTime*1000 + 999
}

// GetUnixTimeFromTransactionTime returns unix time from the transaction time
func GetUnixTimeFromTransactionTime(transactionTime int64) int64 {
	return transactionTime / 1000
}

// parseFromUnixTime parses a unix time and returns a golang time struct
func parseFromUnixTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}
