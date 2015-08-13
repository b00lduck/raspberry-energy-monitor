/*jslint node: true */
'use strict';

var baseUrl = document.location.protocol + "//" + document.location.hostname + "/";

console.log(baseUrl);

angular.module('app')

    .constant("API_BASEURL", baseUrl + "dataservice/")
    .constant("DISPLAY_URL", baseUrl + "display/display")
    .constant("CLICK_URL", baseUrl + "display/click");
