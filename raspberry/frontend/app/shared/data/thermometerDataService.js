/*jslint node: true */
'use strict';

angular.module('data')

    .service('ThermometerDataService', ['$q', '$http', 'API_BASEURL', 'DateToolsService',
        function ($q, $http, API_BASEURL, DateToolsService) {

            function getStartByIntervalType(intervalType) {

                switch(intervalType) {
                    case 'day':
                        return getEnd() - 24 * 3600 * 1000;
                    case '2day':
                        return getEnd() - 48 * 3600 * 1000;
                    default:
                        return getEnd() - 24 * 3600 * 1000;
                }
            }

            function getEnd() {
                return Date.now();
            }

            return {

                getData: function (thermometerId, interval) {
                    return $q(function (resolve, reject) {

                        var start = getEnd() - interval,
                            end = getEnd();

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
