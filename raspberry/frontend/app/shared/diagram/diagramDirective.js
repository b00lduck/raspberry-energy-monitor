/*jslint node: true */
'use strict';

angular.module('diagram', ['nvd3'])

    .directive('myDiagram', [
        function() {
            return {
                restrict: 'E',
                templateUrl: 'shared/diagram/diagramView.html',
                controller: 'MyDiagramController',
                scope: {
                    title: '@',
                    counter: '@'
                }
            };
        }
    ])

    .controller('MyDiagramController', ['$scope', '$http', 'API_BASEURL',

        function($scope, $http, API_BASEURL) {

            function timeFormat(c) {
                return sprintf("%02d:%02d", c.getHours(), c.getMinutes());
            }

            function dateFormat(c) {
                return sprintf("%02d.%d.%4d", c.getDate(), c.getMonth() + 1, c.getFullYear());
            }

            function timeDateFormatMilli(d) {
                var date = new Date(d);
                return timeFormat(date) + " " + dateFormat(date);
            }

            function findStartDate(data) {
                return data[0].Timestamp;
            }

            function findEndDate(data) {
                return data[data.length - 1].Timestamp;
            }

            function getPreviousFullInterval(date, interval) {
                var x = Math.floor(date /interval) * interval;
                return x;
            }

            function getPerIntervalConsumption(start, data, interval) {

                var i = 0,
                    len = data.length,
                    end = start + interval - 1,
                    startValue,
                    endValue;

                while(i < len && data[i].Timestamp < start) {
                    i++;
                }

                if (i >= len || data[i].Timestamp > end) {
                    // no data in this interval
                    return 0;
                }

                startValue = data[i].Reading;

                while(i+1 < len && data[i+1].Timestamp < end) {
                    i++;
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

                for (i = 0; i < len; i++) {
                    ret.push({
                        "x": data[i].Timestamp,
                        "y": data[i].Reading
                    });
                }

                return ret;

            }

            function createDeltaValues(data, singleInterval) {

                var startDate = findStartDate(data),
                    endDate = findEndDate(data),
                    firstIntervalStart = getPreviousFullInterval(startDate, singleInterval),
                    displayInterval = endDate - firstIntervalStart,
                    numIntervals = Math.floor(displayInterval / singleInterval),
                    i,
                    ret = [],
                    x;

                for (i = 0; i < numIntervals; i++ ) {
                    x = firstIntervalStart + i * singleInterval;
                    ret.push({
                        x: x,
                        y: getPerIntervalConsumption(x, data, singleInterval)
                    });
                }

                return ret;

            }

            function getData() {

                $http.get(API_BASEURL + "counter/" + $scope.counter + "/events").then(function(payload) {

                    var counterValues = createCounterValues(payload.data),
                        deltaValues = createDeltaValues(payload.data, 3600000 * 24);

                    $scope.data = [{
                        values: counterValues,
                        key: 'm³',
                        color: '#ff7f0e'
                    },{
                        values: deltaValues,
                        key: 'm³/h',
                        bar: true,
                        color: '#0eff7f'
                    }];

                }, function(error) {
                    console.log(error);
                });

            }

            getData();

            $scope.options = {
                chart: {
                    type: 'linePlusBarChart',
                    height: 400,
                    width: 1000,
                    margin : {
                        top: 20,
                        right: 20,
                        bottom: 40,
                        left: 140
                    },
                    x: function(d) {
                        return d.x;
                    },
                    y: function(d) {
                        return d.y;
                    },
                    useInteractiveGuideline: true,
                    xAxis: {
                        axisLabel: 'Time (UTC)',
                        tickFormat: function(d){
                            return timeDateFormatMilli(d);
                        },
                        axisLabelDistance: 0
                    },
                    x2Axis: {
                        axisLabel: 'Time (UTC)',
                        tickFormat: function(d){
                            return timeDateFormatMilli(d);
                        },
                        axisLabelDistance: 0
                    },
                    y2Axis: {
                        axisLabel: 'Counter (m³)',
                        tickFormat: function(d){
                            return sprintf("%.1f", d / 1000);
                        },
                        axisLabelDistance: 0
                    },
                    y1Axis: {
                        axisLabel: 'Rate (m³/h)',
                        tickFormat: function(d){
                            return sprintf("%0.1f", d / 1000);
                        },
                        axisLabelDistance: 0
                    },
                    bars: {
                        forceY: [0]
                    }
                },
                title: {
                    enable: true,
                    text: $scope.title
                }
            };
        }
    ]);
