/*jslint node: true */
'use strict';

angular.module('welcome', ['nvd3'])

    .controller('WelcomeController', ['$scope', '$http',
        function ($scope, $http) {

            function getData() {

                var ret = [];

                $http.get("http://localhost:8080/counterEvents").then(function(payload) {

                    console.log(payload);

                    var data = payload.data._embedded.counterEvents,
                        i,
                        len = data.length;

                    for (i = 0; i < len; i++) {
                        ret.push({
                            "x": data[i].timestamp,
                            "y": data[i].reading
                        });
                    }

                    $scope.data = [{
                        values: ret,
                        key: 'Sine Wave',
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
                        axisLabel: 'Time (ms)'
                    },
                    yAxis: {
                        axisLabel: 'Voltage (v)',
                        tickFormat: function(d){
                            return d3.format('.02f')(d);
                        },
                        axisLabelDistance: 30
                    },
                    callback: function(chart){
                        console.log("!!! lineChart callback !!!");
                    }
                },
                title: {
                    enable: true,
                    text: 'Title for Line Chart'
                },
                subtitle: {
                    enable: true,
                    text: 'Subtitle for simple line chart. Lorem ipsum dolor sit amet, at eam blandit sadipscing, vim adhuc sanctus disputando ex, cu usu affert alienum urbanitas.',
                    css: {
                        'text-align': 'center',
                        'margin': '10px 13px 0px 7px'
                    }
                },
                caption: {
                    enable: true,
                    html: '<b>Figure 1.</b> Lorem ipsum dolor sit amet, at eam blandit sadipscing, <span style="text-decoration: underline;">vim adhuc sanctus disputando ex</span>, cu usu affert alienum urbanitas. <i>Cum in purto erat, mea ne nominavi persecuti reformidans.</i> Docendi blandit abhorreant ea has, minim tantas alterum pro eu. <span style="color: darkred;">Exerci graeci ad vix, elit tacimates ea duo</span>. Id mel eruditi fuisset. Stet vidit patrioque in pro, eum ex veri verterem abhorreant, id unum oportere intellegam nec<sup>[1, <a href="https://github.com/krispo/angular-nvd3" target="_blank">2</a>, 3]</sup>.',
                    css: {
                        'text-align': 'justify',
                        'margin': '10px 13px 0px 7px'
                    }
                }
            };

            $scope.hello = "Hello, Controller!";

            getData();

        }]);

