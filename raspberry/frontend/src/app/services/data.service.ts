///<reference path="../../../tools/typings/tsd.d.ts" />
///<reference path="../../../tools/typings/typescriptApp.d.ts" />

module app {

    export interface ICounterEvent {
        id: number;
        type: string;
        value: number;
    }

    export class DataService {

        static $inject = ['$http'];
        constructor(private $http: ng.IHttpService) {}

        getCounterEvents(): ng.IPromise<ICounterEvent[]> {
            return this.$http.get('counterEvents.json').then(response => {
                return response.data;
            });
        }

    }

    angular.module('app').service('app.dataService', DataService);

}
