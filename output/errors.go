package output

import "github.com/fatih/color"

func PrintError(value any) {

	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")

	}

}

func sum[T int | int16 | int64](a, b T) T {

	return a + b
}
