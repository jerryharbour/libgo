package utils

package utils

import (
	goflag "flag"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"bkcheck-tools/pkg/version"

	"github.com/spf13/pflag"
)

func wordSeparatorNormalize(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}

	return pflag.NormalizedName(name)
}

func addVersionFlag(cmdline *pflag.FlagSet) *bool {
	return cmdline.BoolP("version", "v", false, "show version")
}

func addHelpFlag(cmdline *pflag.FlagSet) *bool {
	return cmdline.BoolP("help", "h", false, "show help")
}

func wrapFlags(cmdline *pflag.FlagSet, optType reflect.Type, optValue reflect.Value) {
	count := optType.NumField()

	for i := 0; i < count; i++ {
		fieldType := optType.Field(i)
		fieldValue := optValue.Field(i)
		if !fieldValue.IsValid() || !fieldValue.CanSet() {
			continue
		}

		flagName, nameOk := fieldType.Tag.Lookup("opt")
		if !nameOk && !fieldType.Anonymous {
			continue
		}

		flagValue, valueOk := fieldType.Tag.Lookup("defvalue")
		flagUsage, usageOk := fieldType.Tag.Lookup("usage")
		flagShortName := fieldType.Tag.Get("short")

		switch fieldType.Type.Kind() {
		case reflect.Struct:
			wrapFlags(cmdline, fieldType.Type, fieldValue)
			continue
		case reflect.Ptr:
			wrapFlags(cmdline, fieldType.Type.Elem(), fieldValue.Elem())
			continue
		}

		if !nameOk || !valueOk || !usageOk || flagName == "" || flagUsage == "" {
			continue
		}

		ptrFieldValue := unsafe.Pointer(fieldValue.UnsafeAddr())
		switch fieldType.Type.Kind() {
		case reflect.String:
			cmdline.StringVarP((*string)(ptrFieldValue), flagName, flagShortName, flagValue, flagUsage)
		case reflect.Bool:
			v, _ := strconv.ParseBool(flagValue)
			cmdline.BoolVarP((*bool)(ptrFieldValue), flagName, flagShortName, v, flagUsage)
		case reflect.Uint:
			v, _ := strconv.ParseUint(flagValue, 10, 0)
			cmdline.UintVarP((*uint)(ptrFieldValue), flagName, flagShortName, uint(v), flagUsage)
		case reflect.Uint8:
			v, _ := strconv.ParseUint(flagValue, 10, 8)
			cmdline.Uint8VarP((*uint8)(ptrFieldValue), flagName, flagShortName, uint8(v), flagUsage)
		case reflect.Uint16:
			v, _ := strconv.ParseUint(flagValue, 10, 16)
			cmdline.Uint16VarP((*uint16)(ptrFieldValue), flagName, flagShortName, uint16(v), flagUsage)
		case reflect.Uint32:
			v, _ := strconv.ParseUint(flagValue, 10, 32)
			cmdline.Uint32VarP((*uint32)(ptrFieldValue), flagName, flagShortName, uint32(v), flagUsage)
		case reflect.Uint64:
			v, _ := strconv.ParseUint(flagValue, 10, 64)
			cmdline.Uint64VarP((*uint64)(ptrFieldValue), flagName, flagShortName, uint64(v), flagUsage)
		case reflect.Int:
			v, _ := strconv.ParseInt(flagValue, 10, 0)
			cmdline.IntVarP((*int)(ptrFieldValue), flagName, flagShortName, int(v), flagUsage)
		case reflect.Int8:
			v, _ := strconv.ParseInt(flagValue, 10, 8)
			cmdline.Int8VarP((*int8)(ptrFieldValue), flagName, flagShortName, int8(v), flagUsage)
		case reflect.Int16:
			v, _ := strconv.ParseInt(flagValue, 10, 16)
			cmdline.Int16VarP((*int16)(ptrFieldValue), flagName, flagShortName, int16(v), flagUsage)
		case reflect.Int32:
			v, _ := strconv.ParseInt(flagValue, 10, 32)
			cmdline.Int32VarP((*int32)(ptrFieldValue), flagName, flagShortName, int32(v), flagUsage)
		case reflect.Int64:
			v, _ := strconv.ParseInt(flagValue, 10, 64)
			cmdline.Int64VarP((*int64)(ptrFieldValue), flagName, flagShortName, int64(v), flagUsage)
		case reflect.Float32:
			v, _ := strconv.ParseFloat(flagValue, 32)
			cmdline.Float32VarP((*float32)(ptrFieldValue), flagName, flagShortName, float32(v), flagUsage)
		case reflect.Float64:
			v, _ := strconv.ParseFloat(flagValue, 64)
			cmdline.Float64VarP((*float64)(ptrFieldValue), flagName, flagShortName, float64(v), flagUsage)
		case reflect.Slice:
			arrStr := make([]string, 0)
			if flagValue != "" {
				arrStr = strings.Split(flagValue, ",")
			}
			switch fieldType.Type.Elem().Kind() {
			case reflect.String:
				cmdline.StringSliceVarP((*[]string)(ptrFieldValue), flagName, flagShortName, arrStr, flagUsage)
			case reflect.Int:
				arrInt := make([]int, 0, len(arrStr))
				for _, str := range arrStr {
					i, _ := strconv.ParseInt(str, 10, 0)
					arrInt = append(arrInt, int(i))
				}
				cmdline.IntSliceVarP((*[]int)(ptrFieldValue), flagName, flagShortName, arrInt, flagUsage)
			}
		default:
			continue
		}
	}
}

func IsFlagSet(name string) bool {
	return pflag.CommandLine.Changed(name)
}

func AddFlags(opt interface{}) {
	wrapFlags(pflag.CommandLine, reflect.TypeOf(opt).Elem(), reflect.ValueOf(opt).Elem())
}

func InitFlags(opt interface{}) {
	pflag.CommandLine.SetNormalizeFunc(wordSeparatorNormalize)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	ver := addVersionFlag(pflag.CommandLine)
	help := addHelpFlag(pflag.CommandLine)
	AddFlags(opt)

	pflag.Parse()

	if *ver {
		version.ShowVersion()
		os.Exit(0)
	}

	if *help {
		pflag.PrintDefaults()
		os.Exit(0)
	}
}

type IFlagSet interface {
	IsFlagSet(offset uintptr) bool
	InitOffsetOptionName(flag interface{})
}

type FlagSet struct {
	offsetOptName map[uintptr]string
}

func (f *FlagSet) IsFlagSet(offset uintptr) bool {
	if f.offsetOptName == nil {
		return false
	}

	optName, ok := f.offsetOptName[offset]
	if !ok {
		return false
	}

	return IsFlagSet(optName)
}

func (f *FlagSet) InitOffsetOptionName(flag interface{}) {
	f.offsetOptName = make(map[uintptr]string)
	optType := reflect.TypeOf(flag).Elem()
	count := optType.NumField()

	for i := 0; i < count; i++ {
		fieldType := optType.Field(i)
		flagName, nameOk := fieldType.Tag.Lookup("opt")
		if !nameOk && !fieldType.Anonymous {
			continue
		}

		f.offsetOptName[fieldType.Offset] = flagName
	}
}
