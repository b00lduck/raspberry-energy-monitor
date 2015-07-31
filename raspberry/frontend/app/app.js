/*jslint unparam: true, node: true */
'use strict';

var app = angular.module('app', ['ngRoute', 'welcome', 'templates']);

app.config(['$routeProvider', function ($routeProvider) {

    $routeProvider.when('/', {
        controller: 'WelcomeController',
        templateUrl: 'components/welcome/welcomeView.html'
    })

        .otherwise({redirectTo: '/'});

}]);

