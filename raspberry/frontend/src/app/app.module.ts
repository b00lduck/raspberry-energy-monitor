///<reference path="../../tools/typings/tsd.d.ts" />
///<reference path="../../tools/typings/typescriptApp.d.ts" />

((): void => {

    var app = angular.module('app', ['ngRoute']);

    app.config(['$routeProvider', ($routeProvider) => {
        $routeProvider.when('/', {
            controller: 'app.WelcomeController',
            templateUrl: 'app/views/welcome.html'
        });
    }]);

})();
