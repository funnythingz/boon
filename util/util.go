package util

import(
    "fmt"
    "reflect"
)

func CreateTimeOptions() []string {

    var timeOptions []string

    for i := 0; i < 24; i++ {

        count := fmt.Sprintf("%d", i)

        if reflect.DeepEqual(len(count), 2) {
            timeOptions = append(timeOptions, count + ":00")
            timeOptions = append(timeOptions, count + ":30")
        } else {
            timeOptions = append(timeOptions, "0" + count + ":00")
            timeOptions = append(timeOptions, "0" + count + ":30")
        }
    }

    return timeOptions

}
