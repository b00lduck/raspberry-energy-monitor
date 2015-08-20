/*jslint node: true */
'use strict';

angular.module('diagram', ['nvd3', 'dateTools', 'data'])

    .directive('myDiagram', [
        function () {
            return {
                restrict: 'E',
                templateUrl: 'shared/diagram/diagramView.html',
                controller: 'MyDiagramController',
                scope: {
                    title: '@',
                    counter: '@',
                    intervalUnit: '@'
                }
            };
        }
    ])

    .controller('MyDiagramController', ['$scope', 'DateToolsService', 'CounterDataService',

        function ($scope, DateToolsService, CounterDataService) {

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
                            key: 'm³/h',
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
                    x: function (d) {
                        return d.x;
                    },
                    y: function (d) {
                        return d.y;
                    },
                    useInteractiveGuideline: true,
                    xAxis: {
                        axisLabel: 'Time (UTC)',
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
                        axisLabel: 'Counter (m³)',
                        tickFormat: function (d) {
                            return sprintf("%.1f", d / 1000);
                        },
                        axisLabelDistance: 0
                    },
                    y1Axis: {
                        axisLabel: 'Rate (m³/h)',
                        tickFormat: function (d) {
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

            refreshData();

        }]);
