/*jslint unparam: true, node: true */
'use strict';

var app = angular.module('app', ['ngRoute', 'thermometer', 'counter', 'templates', 'navigation']);

app.config(['$routeProvider', function ($routeProvider) {

    $routeProvider.when('/counter', {
        controller: 'CounterController',
        templateUrl: 'components/counter/counterView.html'
    })
    .when('/thermometer', {
        controller: 'ThermometerController',
        templateUrl: 'components/thermometer/thermometerView.html'
    })
    .otherwise({redirectTo: '/thermometer'});

}]);

