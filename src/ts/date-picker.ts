/// <reference path="../../typings/tsd.d.ts" />

class DatePickerApp {

    $selectDate = $('#select-date');
    $datePicker = $('#date-picker');

    constructor() {
        this.events();
    }

    events() {

        var datepickerOptions = {
            language: "ja",
            todayHighlight: true
        }

        this.$datePicker
            .datepicker(datepickerOptions)
            .on("changeDate", e => this.updateDate(e));
    }

    updateDate(e) {
        if(_.isUndefined(e.date)) {
            this.$selectDate.html('')
        } else {
            this.$selectDate.html(e.date.toString())
        }
    }

}
