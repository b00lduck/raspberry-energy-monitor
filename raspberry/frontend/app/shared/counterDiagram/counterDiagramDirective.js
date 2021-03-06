/*jslint node: true */
'use strict';

angular.module('counterDiagram', ['nvd3', 'dateTools', 'data'])

    .directive('myCounterDiagram', [
        function () {
            return {
                restrict: 'E',
                templateUrl: 'shared/counterDiagram/counterDiagramView.html',
                controller: 'MyCounterDiagramController',
                scope: {
                    title: '@',
                    counter: '@',
                    intervalUnit: '@'
                }
            };
        }
    ])

    .controller('MyCounterDiagramController', ['$scope', 'DateToolsService', 'CounterDataService',

        function ($scope, DateToolsService, CounterDataService) {

            var counterUnit = "m³";

            function getRateUnit() {
                return counterUnit + '/' + $scope.intervalType;
            }

            function formatReading(x) {
                return sprintf("%0.1f", x / 1000);
            }

            function formatRate(x) {
                return sprintf("%0.2f", x / 1000);
            }

            function getLineTooltip(data) {
                return "<h1>Counter reading</h1>" +
                    "Instant: " + DateToolsService.timeDateFormatMilli(data.point.x) + "<br>" +
                    formatReading(data.point.y) + counterUnit;
            }

            function getBarTooltip(data) {
                var begin = data.data.x,
                    end = data.data.x + DateToolsService.getMillisByIntervalType($scope.intervalType);
                return "<h1>Consumption period</h1>" + "<br>" +
                    "Begin: " + DateToolsService.timeDateFormatMilli(begin) + "<br>" +
                    "End: " + DateToolsService.timeDateFormatMilli(end) + "<br>" +
                    "Average: " + formatRate(data.data.y) + getRateUnit();
            }

            function getTooltip(data) {
                var content,
                    cssClass;
                if (undefined === data.index) {
                    content = getLineTooltip(data);
                    cssClass = "lineTip";
                } else {
                    content = getBarTooltip(data);
                    cssClass = "barTip";
                }
                return '<div class="' + cssClass + '">' + content + '</div>';
            }

            function createTickArray(value, lastValue, intervalType) {
                var currentValue = value,
                    ret = [];
                while (currentValue < lastValue) {
                    currentValue = DateToolsService.getNextFullInterval(currentValue, intervalType);
                    ret.push(currentValue);
                    currentValue += 1;
                }
                return ret;
            }

            function getTickValues(d) {
                var maxTickValues = 7,
                    hour = 3600000,
                    day = hour * 24,
                    week = day * 7,
                    values = d[0].values,
                    len = values.length,
                    firstValue = values[0].x,
                    lastValue = values[len - 1].x,
                    range = lastValue - firstValue,
                    currentValue = firstValue;

                // tick mode    range max
                // ---------------------------------
                // hours        maxTickValues * 3600000
                // 3 hours      maxTickValues * 3600000 * 3
                // 6 hours      maxTickValues * 3600000 * 6
                // day          maxTickValues * 3600000 * 24
                // 2 days       maxTickValues * 3600000 * 48
                // week         maxTickValues * 3600000 * 24 * 7
                // 2 weeks      maxTickValues * 3600000 * 24 * 14

                if (range < maxTickValues * hour) {
                    console.log("hour tick");
                    return createTickArray(currentValue, lastValue, "hour");
                }
                if (range < maxTickValues * hour * 3) {
                    console.log("3 hour tick");
                    return createTickArray(currentValue, lastValue, "3hour");
                }
                if (range < maxTickValues * hour * 6) {
                    console.log("6 hour tick");
                    return createTickArray(currentValue, lastValue, "6hour");
                }
                if (range < maxTickValues * day) {
                    console.log("day tick");
                    return createTickArray(currentValue, lastValue, "day");
                }
                if (range < maxTickValues * day * 2) {
                    console.log("2 day tick");
                    return createTickArray(currentValue, lastValue, "2day");
                }
                if (range < maxTickValues * week) {
                    console.log("week tick");
                    return createTickArray(currentValue, lastValue, "week");
                }
                if (range < maxTickValues * week * 2) {
                    console.log("2 week tick");
                    return createTickArray(currentValue, lastValue, "2week");
                }
            }

            function setOptions() {
                $scope.options = {
                    chart: {
                        type: 'linePlusBarChart',
                        height: 400,
                        width: 1000,
                        margin: {
                            top: 20,
                            right: 20,
                            bottom: 40,
                            left: 140
                        },
                        line1: {},
                        x: function (d) {
                            return d.x;
                        },
                        y: function (d) {
                            return d.y;
                        },
                        tooltip: {
                            contentGenerator: getTooltip
                        },
                        useInteractiveGuideline: false,
                        showBarLabels: true,
                        xAxis: {
                            axisLabel: 'Time (UTC)',
                            tickValues: getTickValues,
                            tickFormat: function (d) {
                                return DateToolsService.timeDateFormatMilli(d);
                            },
                            axisLabelDistance: 0
                        },
                        x2Axis: {
                            axisLabel: 'Time (UTC)',
                            tickFormat: function (d) {
                                return DateToolsService.timeDateFormatMilli(d);
                            },
                            axisLabelDistance: 0
                        },
                        y2Axis: {
                            axisLabel: 'Counter (' + counterUnit + ')',
                            tickFormat: formatReading,
                            axisLabelDistance: 0
                        },
                        y1Axis: {
                            axisLabel: 'Rate (' + getRateUnit() + ')',
                            tickFormat: formatRate,
                            axisLabelDistance: 0
                        },
                        focusEnable: false
                    },
                    title: {
                        enable: true,
                        text: $scope.title
                    }
                };
            }

            function refreshData() {

                console.log("refreshData() for intervalType " + $scope.intervalType);

                CounterDataService.getData($scope.counter, $scope.intervalType)

                    .then(function (payload) {

                        $scope.data = [{
                            values: payload.counterValues,
                            key: 'm³',
                            color: '#ff7f0e'
                        }, {
                            values: payload.deltaValues,
                            key: 'm³/' + $scope.intervalType,
                            bar: true,
                            color: '#0eff7f'
                        }];

                    }, function (error) {
                        console.log(error);
                    });
            }

            $scope.intervalType = "day";

            $scope.selectIntervalType = function (newIntervalType) {
                $scope.intervalType = newIntervalType;
                refreshData();
            };

            setOptions();
            refreshData();

        }]);
