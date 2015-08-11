/*jslint node: true */
'use strict';

angular.module('welcome', ['nvd3'])

    .controller('WelcomeController', ['$scope', '$http', 'API_BASEURL', 'DISPLAY_URL', 'CLICK_URL', '$interval',
        function ($scope, $http, API_BASEURL, DISPLAY_URL, CLICK_URL, $interval) {

            function getData() {

                var ret = [];

                $http.get(API_BASEURL + "counter/1/events").then(function(payload) {

                    var data = payload.data,
                        i,
                        len = data.length;

                    for (i = 0; i < len; i++) {
                        ret.push({
                            "x": data[i].Timestamp,
                            "y": data[i].Reading
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

            $scope.options = {
                chart: {
                    type: 'lineChart',
                    height: 450,
                    margin : {
                        top: 20,
                        right: 20,
                        bottom: 40,
                        left: 55
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
                        }
                    },
                    yAxis: {
                        axisLabel: 'Counter (m³)',
                        tickFormat: function(d){
                            return d3.format('.03f')(d);
                        },
                        axisLabelDistance: 30
                    }
                },
                title: {
                    enable: true,
                    text: 'Title'
                },
                subtitle: {
                    enable: true,
                    text: 'Subtitle',
                    css: {
                        'text-align': 'center',
                        'margin': '10px 13px 0px 7px'
                    }
                },
                caption: {
                    enable: true,
                    html: '<b>Description</b>',
                    css: {
                        'text-align': 'justify',
                        'margin': '10px 13px 0px 7px'
                    }
                }
            };

            $scope.hello = "Hello, Controller!";

            //$interval(function() {
            //   $scope.imageUrl = DISPLAY_URL + '?' + new Date().getTime();
            //}, 500);

            $scope.doClick = function(event) {
                var x = event.offsetX,
                    y = event.offsetY;
                $http.get(CLICK_URL + "?x=" + x + "&y=" + y);
            };

            getData();

        }]);

