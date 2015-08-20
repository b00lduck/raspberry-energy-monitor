/*jslint node: true */
'use strict';

angular.module('welcome', ['nvd3', 'ui.bootstrap', 'diagram', 'display'])

    .controller('WelcomeController', ['$scope',
        function ($scope) {

            $scope.status = {
                opened: false
            };

            $scope.dateOptions = {
                formatYear: 'yy',
                startingDay: 1
            };

            /*
            $scope.open = function($event) {
                $scope.status.opened = true;
            };
            */

        }]);

