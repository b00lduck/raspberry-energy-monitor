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
                return sprintf("%02d:%02d", c.getUTCHours(), c.getMinutes());
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
                return this.dateFormat(date) + " " + this.timeFormat(date);
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
             * Get the millisecond timestamp of the start of the first interval of given type
             * before the given timestamp.
             */
            getNextFullInterval: function (date, intervalType) {
                var interval = this.getMillisByIntervalType(intervalType),
                    offset = this.getOffsetByIntervalType(intervalType);
                return Math.ceil(date / interval) * interval + offset;
            },

            /**
             * Get length of interval type in milliseconds.
             * Unimplemented:: switching years and months
             */
            getMillisByIntervalType: function (intervalType) {
                const hour = 3600000,
                      day = hour * 24,
                      week = day * 7,
                      year = day * 365;
                switch (intervalType) {
                case "hour":
                    return hour;
                case "3hour":
                    return hour * 3;
                case "6hour":
                    return hour * 3;
                case "day":
                    return day;
                case "2day":
                    return day * 2;
                case "week":
                    return week;
                case "2week":
                    return week * 2;
                case "year":
                    return year;
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
