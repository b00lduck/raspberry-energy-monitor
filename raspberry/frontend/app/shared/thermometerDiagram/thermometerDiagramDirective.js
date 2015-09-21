/*jslint node: true */
'use strict';

angular.module('thermometerDiagram', ['nvd3', 'dateTools', 'data'])

    .directive('myThermometerDiagram', [
        function () {
            return {
                restrict: 'E',
                templateUrl: 'shared/thermometerDiagram/thermometerDiagramView.html',
                controller: 'MyThermometerDiagramController',
                scope: {
                    title: '@',
                    thermometer: '@',
                    interval: '='
                }
            };
        }
    ])

    .controller('MyThermometerDiagramController', ['$scope', '$rootScope', 'DateToolsService', 'ThermometerDataService',

        function ($scope, $rootScope, DateToolsService, ThermometerDataService) {

            function formatTemperature(x) {
                return sprintf("%0.1f", x / 1000);
            }

            function getLineTooltip(data) {
                return "<h1>Temperature reading</h1>" +
                    "Instant: " + DateToolsService.timeDateFormatMilli(data.point.x) + "<br>" +
                    formatTemperature(data.point.y) + "°C";
            }

            function getTooltip(data) {
                var content,
                    cssClass;
                content = getLineTooltip(data);
                cssClass = "lineTip";
                return '<div class="' + cssClass + '">' + content + '</div>';
            }

            function setOptions() {
                $scope.options = {
                    chart: {
                        type: 'lineChart',
                        height: 200,
                        margin: {
                            top: 20,
                            right: 100,
                            bottom: 20,
                            left: 140
                        },
                        line1: {},
                            x: function (d) {
                                return d.Timestamp;
                            },
                            y: function (d) {
                                return d.Reading;
                        },
                        useInteractiveGuideline: false,
                        showBarLabels: true,
                        xAxis: {
                            axisLabel: 'Time (UTC)',
                            tickFormat: function (d) {
                                return DateToolsService.timeDateFormatMilli(d);
                            },
                            axisLabelDistance: 450
                        },
                        yAxis: {
                            axisLabel: 'Temperature (°C)',
                            tickFormat: formatTemperature,
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

                console.log($scope.interval);

                ThermometerDataService.getData($scope.thermometer, $scope.interval)
                    .then(function (payload) {
                        $scope.data = [{
                            values: payload,
                            key: '°C',
                            color: '#ff0000'
                        }];
                    }, function (error) {
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
