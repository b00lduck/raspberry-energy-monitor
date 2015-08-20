/*jslint node: true */
'use strict';

angular.module('dateTools', [])

    .service('DateToolsService', function () {

        return {

            /**
             * Format the given Date object to the time format which is common in germany.
             * i.e.: 13:37 or 05:42
             **/
            timeFormat: function (c) {
                return sprintf("%02d:%02d", c.getHours(), c.getMinutes());
            },

            /**
             * Format the given Date object to a date format which is common in germany.
             * i.e.: 31.03.2015
             */
            dateFormat: function (c) {
                return sprintf("%02d.%d.%4d", c.getDate(), c.getMonth() + 1, c.getFullYear());
            },

            /**
             * Format the given millisecond timestamp to a date and time format which is common in germany.
             * Uses timeFormat() and dateFormat() from above.
             */
            timeDateFormatMilli: function (d) {
                var date = new Date(d);
                return this.timeFormat(date) + " " + this.dateFormat(date);
            },

            /**
             * Get the millisecond timestamp of the start of the first interval of given type
             * before the given timestamp.
             */
            getPreviousFullInterval: function (date, intervalType) {
                var interval = this.getMillisByIntervalType(intervalType),
                    offset = this.getOffsetByIntervalType(intervalType);
                return Math.floor(date / interval) * interval + offset;
            },

            /**
             * Get length of interval type in milliseconds.
             * Unimplemented:: switching years and months
             */
            getMillisByIntervalType: function (intervalType) {
                switch (intervalType) {
                case "hour":
                    return 3600000;
                case "day":
                    return 3600000 * 24;
                case "week":
                    return 3600000 * 24 * 7;
                case "year":
                    return 3600000 * 24 * 365;
                default:
                    console.log("invalid interval type " + intervalType);
                    return;
                }
            },

            /**
             * Get adjust offset for start of intervals. The 1.1.1970 was a Thursday so we have to substract 3
             * days to come out on Monday.
             * Unimplemented: switching years and months
             */
            getOffsetByIntervalType: function (intervalType) {
                if ("week" === intervalType) {
                    return -3600000 * 24 * 3;
                }
                return 0;
            }
        };
    });
