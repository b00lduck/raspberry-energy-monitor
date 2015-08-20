/*jslint node: true */
'use strict';

angular.module('data', ['dateTools'])

    .service('CounterDataService', ['$q', '$http', 'API_BASEURL', 'DateToolsService',
        function ($q, $http, API_BASEURL, DateToolsService) {

            /**
             * Return the first timestamp of the given data array.
             */
            function getFirstDataTimestamp(data) {
                return data[0].Timestamp;
            }

            /**
             * Return the last timestamp of the given data array.
             */
            function getLastDataTimestamp(data) {
                return data[data.length - 1].Timestamp;
            }

            function getPerIntervalConsumption(start, data, interval) {

                var i = 0,
                    len = data.length,
                    end = start + interval - 1,
                    startValue,
                    endValue;

                while (i < len && data[i].Timestamp < start) {
                    i += 1;
                }

                if (i >= len || data[i].Timestamp > end) {
                    // no data in this interval
                    return 0;
                }

                startValue = data[i].Reading;

                while (i + 1 < len && data[i + 1].Timestamp < end) {
                    i += 1;
                }

                if (data[i].Timestamp > end) {
                    return 0;
                }

                endValue = data[i].Reading;

                return endValue - startValue;
            }

            function createCounterValues(data) {
                var ret = [],
                    i,
                    len = data.length;

                for (i = 0; i < len; i += 1) {
                    ret.push({
                        "x": data[i].Timestamp,
                        "y": data[i].Reading
                    });
                }

                return ret;

            }

            function createDeltaValues(data, intervalType) {

                var singleInterval = DateToolsService.getMillisByIntervalType(intervalType),
                    firstIntervalStartsAt = DateToolsService.getPreviousFullInterval(getFirstDataTimestamp(data),
                        intervalType),
                    displayInterval = getLastDataTimestamp(data) - firstIntervalStartsAt,
                    numIntervals = Math.floor(displayInterval / singleInterval),
                    i,
                    ret = [],
                    x;

                for (i = 0; i < numIntervals; i += 1) {
                    x = firstIntervalStartsAt + i * singleInterval;
                    ret.push({
                        x: x,
                        y: getPerIntervalConsumption(x, data, singleInterval)
                    });
                }

                return ret;

            }

            return {

                getData: function (counterId, intervalType) {
                    return $q(function (resolve, reject) {
                        $http.get(API_BASEURL + "counter/" + counterId + "/events").then(function (payload) {
                            var counterValues = createCounterValues(payload.data),
                                deltaValues = createDeltaValues(payload.data, intervalType);
                            console.log("Got " + counterValues.length + " counter values and " + deltaValues.length +
                                " delta values for interval type " + intervalType + ".");
                            resolve({
                                counterValues: counterValues,
                                deltaValues: deltaValues
                            });
                        }, function (error) {
                            reject(error);
                        });

                    });
                }
            };

        }]);
