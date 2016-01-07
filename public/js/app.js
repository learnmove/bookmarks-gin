var myApp = angular.module('ngclient', ['ngRoute']);

myApp.config(function($routeProvider, $httpProvider, $locationProvider) {
  $routeProvider
    .when('/', {
      templateUrl: 'partials/home.html',
      controller: 'HomeCtrl'
    }).when('/bookmark/:id', {
      templateUrl: 'partials/edit.html',
      controller: 'EditCtrl'
    }).when('/new', {
      templateUrl: 'partials/new.html',
      controller: 'NewCtrl'
    });
});