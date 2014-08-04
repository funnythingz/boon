/// <reference path="../../typings/tsd.d.ts" />

/// <reference path="helper.ts" />

class DatePickerApp {

    $selectDate = $('#select-date');
    $datePicker = $('#date-picker');

    constructor() {
        this.ready();
    }

    ready() {

        var datepickerOptions = {
            language: "ja",
            todayHighlight: true,
            format: "dd/mm/yyyy"
        }

        this.$datePicker
            .datepicker(datepickerOptions)
            .on("changeDate", e => this.updateDate(e));
    }

    updateDate(e) {
        if(_.isUndefined(e.date)) {
            this.$selectDate.val('')
        } else {
            var date = Helper.createDate(new Date(e.date));
            this.$selectDate.val(date);
        }
    }

}
