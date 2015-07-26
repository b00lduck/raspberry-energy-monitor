/*jslint node: true */
'use strict';

angular.module('welcome', ['nvd3ChartDirectives'])

    .controller('WelcomeController', ['$scope', '$http',
        function ($scope, $http) {

            function getData() {

                var ret = [];

                $http.get("counterEvents.json").then(function(payload) {

                    var data = payload.data,
                        i,
                        len = data.length;

                    for (i = 0; i < len; i++) {
                        ret[i] = [ data[i].timestamp.millis, data[i].reading ];
                    }

                    $scope.exampleData = [{
                        "values" : ret
                    }];

                }, function(error) {
                    console.log(error);
                });

            }

            $scope.hello = "Hello, Controller!";

            getData();

        }]);

