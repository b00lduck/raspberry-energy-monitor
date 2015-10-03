/*jslint node: true */
'use strict';

angular.module('thermometer', ['nvd3', 'ui.bootstrap', 'thermometerDiagram', 'flagDiagram', 'thermoFlagDiagram'])

    .controller('ThermometerController', ['$scope',
        function ($scope) {

            $scope.selectInterval = function(range) {
                console.log("Select interval " + range);
                $scope.interval = range * 1000;
            };

            $scope.selectInterval(3600 * 24);

        }]);
