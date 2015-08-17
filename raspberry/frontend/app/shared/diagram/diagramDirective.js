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

            function getNextFullHour(date) {
                var ret = new Date(date);
                ret.setHours(ret.getHours() + 1);
                ret.setMinutes(0);
                ret.setSeconds(0);
                ret.setMilliseconds(0);
                return ret.getTime();
            }

            function getPerHourConsumption(start, data) {

                var i = 0,
                    len = data.length,
                    end = start + 3599999,
                    startValue,
                    endValue;

                while(data[i].Timestamp < start && i < len) {
                    i++;
                }

                if (data[i].Timestamp > end) {
                    // no data in this interval
                    return 0;
                }

                startValue = data[i].Reading;

                while(data[i+1].Timestamp < end && i+1 < len) {
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

                if (0 < i) {
                    ret.push({
                        "x": new Date().getTime(),
                        "y": data[i-1].Reading
                    });
                }
                return ret;

            }

            function createDeltaValues(data) {

                var startDate = findStartDate(data),
                    endDate = findEndDate(data),
                    firstHour = getNextFullHour(startDate),
                    interval = endDate - firstHour,
                    numHours = Math.floor(interval / 3600000),
                    i,
                    ret = [],
                    x;

                for (i = 0; i < numHours; i++ ) {
                    console.log(i);
                    x = firstHour + i * 3600000;
                    ret.push({
                        x: x,
                        y: getPerHourConsumption(x, data)
                    });
                }

                return ret;

            }

            function getData() {

                $http.get(API_BASEURL + "counter/" + $scope.counter + "/events").then(function(payload) {

                    var counterValues = createCounterValues(payload.data),
                        deltaValues = createDeltaValues(payload.data);

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
                            return sprintf("%d", d);
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
