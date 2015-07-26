/*jslint node: true */
'use strict';

angular.module('welcome', [])

    .controller('WelcomeController', ['$scope',
        function ($scope) {
            $scope.hello = "Hello, Controller!";
        }]);

