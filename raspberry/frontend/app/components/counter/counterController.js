/*jslint node: true */
'use strict';

angular.module('counter', ['nvd3', 'ui.bootstrap', 'counterDiagram', 'display'])

    .controller('CounterController', ['$scope',
        function ($scope) {

            $scope.dateOptions = {
                formatYear: 'yy',
                startingDay: 1
            };

        }]);

