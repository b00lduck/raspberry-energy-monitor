/*jslint node: true */
'use strict';

angular.module('navigation', [])

    .directive('myNavigation', [
        function () {
            return {
                restrict: 'E',
                templateUrl: 'shared/navigation/navigationView.html'
            };
        }
    ]);

