/*jslint node: true */
'use strict';

angular.module('display', [])

    .directive('myDisplay', [
        function() {
            return {
                restrict: 'E',
                templateUrl: 'shared/display/displayView.html',
                controller: 'MyDisplayController',
                scope: {}
            };
        }
    ])

    .controller('MyDisplayController', ['$scope', '$interval', '$http', 'DISPLAY_URL', 'CLICK_URL',

        function($scope, $interval, $http, DISPLAY_URL, CLICK_URL) {

            console.log(DISPLAY_URL);

            $interval(function() {
                $scope.imageUrl = DISPLAY_URL + '?' + new Date().getTime();
            }, 500);

            $scope.doClick = function(event) {
                var x = event.offsetX,
                    y = event.offsetY;
                $http.get(CLICK_URL + "?x=" + x + "&y=" + y);
            };

        }

    ]);