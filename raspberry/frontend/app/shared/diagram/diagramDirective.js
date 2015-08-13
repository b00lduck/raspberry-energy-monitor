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

            function getData() {

                var ret = [];

                $http.get(API_BASEURL + "counter/" + $scope.counter + "/events").then(function(payload) {

                    var data = payload.data,
                        i,
                        len = data.length;

                    for (i = 0; i < len; i++) {
                        ret.push({
                            "x": data[i].Timestamp,
                            "y": data[i].Reading
                        });
                    }

                    if (i > 0) {
                        ret.push({
                            "x": new Date().getTime(),
                            "y": data[i-1].Reading
                        });
                    }

                    $scope.data = [{
                        values: ret,
                        key: 'm³',
                        color: '#ff7f0e'
                    }];

                }, function(error) {
                    console.log(error);
                });

            }

            getData();

            $scope.options = {
                chart: {
                    type: 'lineChart',
                    height: 200,
                    width: 500,
                    margin : {
                        top: 20,
                        right: 20,
                        bottom: 40,
                        left: 70
                    },
                    x: function(d) {
                        return d.x;
                    },
                    y: function(d) {
                        return d.y;
                    },
                    useInteractiveGuideline: true,
                    dispatch: {
                        stateChange: function(e){ console.log("stateChange"); },
                        changeState: function(e){ console.log("changeState"); },
                        tooltipShow: function(e){ console.log("tooltipShow"); },
                        tooltipHide: function(e){ console.log("tooltipHide"); }
                    },
                    xAxis: {
                        axisLabel: 'Time (UTC)',
                        tickFormat: function(d){
                            var c = new Date(d),
                                time = c.getHours() + ":" + c.getMinutes(),
                                date = c.getDate() + "." + (c.getMonth() + 1) + "." + c.getFullYear();

                            return  time + " " + date;
                        },
                        axisLabelDistance: 0
                    },
                    yAxis: {
                        axisLabel: 'Counter (m³)',
                        tickFormat: function(d){
                            return d3.format('.03f')(d / 1000);
                        },
                        axisLabelDistance: 0
                    }
                },
                title: {
                    enable: true,
                    text: $scope.title
                }
            };
        }
    ]);