///<reference path="../../../tools/typings/tsd.d.ts" />
///<reference path="../../../tools/typings/typescriptApp.d.ts" />

module app {

    class WelcomeController {

        counterEvents: ICounterEvent[] = null;

        static $inject = ['$routeParams', 'app.dataService'];
        constructor($routeParams, dataService: DataService) {

            dataService.getCounterEvents()
              .then((counterEvents: ICounterEvent[]) => {
                 this.counterEvents = counterEvents;
              });
        }
    }

    angular.module('app').controller('app.WelcomeController', WelcomeController);

}
