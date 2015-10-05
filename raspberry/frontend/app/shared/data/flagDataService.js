/*jslint node: true */
'use strict';

angular.module('data')

    .service('FlagDataService', ['$q', '$http', 'API_BASEURL', 'DateToolsService',
        function ($q, $http, API_BASEURL, DateToolsService) {

            return {

                getData: function (flagId, interval) {
                    return $q(function (resolve, reject) {

                        var start = Date.now() - interval,
                            end = Date.now();

                        $http.get(API_BASEURL + "flag/" + flagId + "/states?start=" + start + "&end=" + end).then(function (payload) {
                            var i,
                                len = payload.data.length,
                                out = [];

                            console.log("Got payload from REST service with length " + payload.data.length);

                            for (i = 0; i < len; i++) {

                                if ((0 < i) && (i < len - 1)) {
                                    out.push({
                                        Timestamp: payload.data[i].Timestamp - 1,
                                        State: payload.data[i-1].State
                                    });
                                }

                                out.push(payload.data[i]);

                            }

                            console.log(out);

                            resolve(out);
                        }, function (error) {
                            reject(error);
                        });

                    });
                }
            };

        }]);
