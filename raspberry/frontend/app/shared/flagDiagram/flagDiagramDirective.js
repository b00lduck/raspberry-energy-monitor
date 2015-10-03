/*jslint node: true */
'use strict';

angular.module('flagDiagram', ['nvd3', 'dateTools', 'data'])

    .directive('myFlagDiagram', [
        function () {
            return {
                restrict: 'E',
                templateUrl: 'shared/flagDiagram/flagDiagramView.html',
                controller: 'MyFlagDiagramController',
                scope: {
                    title: '@',
                    flag: '@',
                    interval: '='
                }
            };
        }
    ])

    .controller('MyFlagDiagramController', ['$scope', '$rootScope', 'DateToolsService', 'FlagDataService',

        function ($scope, $rootScope, DateToolsService, FlagDataService) {

            function formatFlag(x) {
                return x;
            }

            function getLineTooltip(data) {
                return "<h1>Flag status</h1>" +
                    "Instant: " + DateToolsService.timeDateFormatMilli(data.point.x) + "<br>" +
                    formatFlag(data.point.y);
            }

            function setOptions() {
                $scope.options = {
                    chart: {
                        type: 'lineChart',
                        height: 80,
                        margin: {
                            top: 0,
                            right: 100,
                            bottom: 0,
                            left: 140
                        },
                        line1: {},
                            x: function (d) {
                                return d.Timestamp;
                            },
                            y: function (d) {
                                return d.State;
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
                            axisLabel: 'Flag',
                            tickFormat: formatFlag,
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

                FlagDataService.getData($scope.flag, $scope.interval)
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
