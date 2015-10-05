/*jslint node: true */
'use strict';


angular.module('thermoFlagDiagram', ['nvd3', 'dateTools', 'data'])

    .directive('myThermoFlagDiagram', [
        function () {
            return {
                restrict: 'E',
                templateUrl: 'shared/thermoFlagDiagram/thermoFlagDiagramView.html',
                controller: 'MyThermoFlagDiagramController',
                scope: {
                    thermometer: '@',
                    title: '@',
                    flag1: '@',
                    flag1title: '@',
                    flag2: '@',
                    flag2title: '@',
                    interval: '='
                }
            };
        }
    ])

    .controller('MyThermoFlagDiagramController', ['$scope', '$q', '$rootScope', 'DateToolsService', 'ThermometerDataService',
        'FlagDataService',

        function ($scope, $q, $rootScope, DateToolsService, ThermometerDataService, FlagDataService) {

            function formatTemperature(x) {
                return sprintf("%0.1f", x / 1000);
            }

            function formatFlag(x) {
                return x;
            }

            function getLineTooltip(data) {
                return "<h1>Temperature reading</h1>" +
                    "Instant: " + DateToolsService.timeDateFormatMilli(data.point.x) + "<br>" +
                    formatTemperature(data.point.y) + "°C";
            }

            function setOptions() {
                $scope.options = {
                    chart: {
                        type: 'multiChart',
                        height: 200,
                        margin: {
                            top: 20,
                            right: 100,
                            bottom: 20,
                            left: 140
                        },
                        interpolate: "linear",
                        useInteractiveGuideline: false,
                        showBarLabels: true,
                        xAxis: {
                            axisLabel: 'Time (UTC)',
                            tickFormat: function (d) {
                                return DateToolsService.timeDateFormatMilli(d);
                            },
                            axisLabelDistance: 450
                        },
                        yAxis1: {
                            axisLabel: "State",
                            tickFormat: formatFlag,
                            axisLabelDistance: 0

                        },
                        yAxis2: {
                            axisLabel: 'Temperature (°C)',
                            tickFormat: formatTemperature,
                            axisLabelDistance: 0
                        }
                    },
                    title: {
                        enable: true,
                        text: $scope.title
                    }
                };
            }

            function refreshFlagPromise(code, title, color) {
                return $q(function (resolve, reject) {
                    if ('undefined' === typeof code) {
                        resolve(null);
                        return;
                    }
                    FlagDataService.getData(code, $scope.interval)
                        .then(function (payload) {
                            var i,
                                len = payload.length,
                                transformed = [];
                            for (i = 0; i < len; i++) {
                                transformed.push({
                                    x: payload[i].Timestamp,
                                    y: payload[i].State
                                });
                            }
                            resolve({
                                type: "area",
                                yAxis: 1,
                                values: transformed,
                                key: title,
                                color: color
                            });
                        }, function (error) {
                            console.log(error);
                            reject(error);
                        });
                });
            }

            function refreshThermoPromise(code) {
                return $q(function (resolve, reject) {
                    ThermometerDataService.getData(code, $scope.interval)
                        .then(function (payload) {
                            var i,
                                len = payload.length,
                                transformed = [];
                            for (i = 0; i < len; i++) {
                                transformed.push({
                                    x: payload[i].Timestamp,
                                    y: payload[i].Reading
                                });
                            }
                            resolve({
                                type: "line",
                                yAxis: 2,
                                values: transformed,
                                key: '°C',
                                color: '#000055'
                            });
                        }, function (error) {
                            console.log(error);
                            reject();
                        });
                });
            }

            function flagMerge(primary, secondary) {

                var primIndex = 0,
                    primLen = primary.values.length,
                    secIndex = 0,
                    secLen = secondary.values.length,
                    out = {
                        type: primary.type,
                        yAxis: primary.yAxis,
                        values: [],
                        key: primary.key,
                        color: primary.color
                    };

                while(primIndex < primLen) {

                    if (primary.values[primIndex].x === secondary.values[secIndex].x) {
                        out.values.push(primary.values[primIndex]);
                        primIndex++;
                        secIndex++;
                    } else if (primary.values[primIndex].x < secondary.values[secIndex].x) {
                        out.values.push(primary.values[primIndex]);
                        primIndex++;
                    } else {
                        out.values.push({
                            y: primary.values[primIndex - 1].y,
                            x: secondary.values[secIndex].x
                        });
                        secIndex++;
                    }
                }

                console.log(primary);
                console.log(secondary);
                console.log(out);

                return out;

            }

            function refreshData() {

                var promises = {
                    thermo: refreshThermoPromise($scope.thermometer),
                    flag1: refreshFlagPromise($scope.flag1, $scope.flag1title, '#ff5599'),
                    flag2: refreshFlagPromise($scope.flag2, $scope.flag2title, '#99ff55')
                };

                $q.all(promises).then(function(payload) {
                    var data = [payload.thermo];

                    if (null !== payload.flag1 && null !== payload.flag2) {
                        data.push(flagMerge(payload.flag1, payload.flag2));
                        data.push(flagMerge(payload.flag2, payload.flag1));
                    } else {
                        if (null !== payload.flag1) {
                            data.push(payload.flag1);
                        }
                        if (null !== payload.flag2) {
                            data.push(payload.flag2);
                        }
                    }

                    $scope.data = data;
                }, function(error) {
                    console.log(error);
                });

            }

            setOptions();

            $scope.$watch(function(scope) {
                return scope.interval;
            }, function() {
                refreshData();
            });

        }]);
