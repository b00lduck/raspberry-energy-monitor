/*jslint node: true */
'use strict';

angular.module('data')

    .service('ThermometerDataService', ['$q', '$http', 'API_BASEURL', 'DateToolsService',
        function ($q, $http, API_BASEURL, DateToolsService) {

            return {

                getData: function (thermometerId, interval) {
                    return $q(function (resolve, reject) {

                        var start =  Date.now() - interval,
                            end =  Date.now();

                        $http.get(API_BASEURL + "thermometer/" + thermometerId + "/readings?start=" + start + "&end=" + end).then(function (payload) {
                            console.log("Got payload from REST service with length " + payload.data.length);
                            resolve(payload.data);
                        }, function (error) {
                            reject(error);
                        });
                    });
                }
            };

        }]);
