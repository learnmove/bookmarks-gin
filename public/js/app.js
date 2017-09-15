var myApp = angular.module('ngclient', ['ngRoute']);

myApp.config(function($routeProvider, $httpProvider, $locationProvider) {

  $httpProvider.interceptors.push('TokenInterceptor');

  $routeProvider
    .when('/login', {
      templateUrl: 'partials/login.html',
      controller: 'LoginCtrl',
      access: {
        requiredLogin: false
      }
    }).when('/', {
      templateUrl: 'partials/home.html',
      controller: 'HomeCtrl',
      access: {
        requiredLogin: false
      }
    }).when('/bookmarks', {
      templateUrl: 'partials/bookmarks.html',
      controller: 'BookmarkCtrl',
      access: {
        requiredLogin: true
      }
    }).when('/bookmark/:id', {
      templateUrl: 'partials/edit.html',
      controller: 'EditCtrl',
      access: {
        requiredLogin: true
      }
    }).when('/new', {
      templateUrl: 'partials/new.html',
      controller: 'NewCtrl',
      access: {
        requiredLogin: true
      }
    }).when('/admin', {
      templateUrl: 'partials/admin.html',
      controller: 'AdminCtrl',
      access: {
        requiredLogin: true
      }
    }).otherwise({
      redirectTo: '/login'
    });
});

myApp.run(function($rootScope, $window, $location, AuthenticationFactory) {
  // when the page refreshes, check if the user is already logged in
  AuthenticationFactory.check();

  $rootScope.$on("$routeChangeStart", function(event, nextRoute, currentRoute) {
    if ((nextRoute.access && nextRoute.access.requiredLogin) && !AuthenticationFactory.isLogged) {
      $location.path("/login");
    } else {
      // check if user object exists else fetch it. This is incase of a page refresh
      if (!AuthenticationFactory.user) AuthenticationFactory.user = $window.sessionStorage.user;
      if (!AuthenticationFactory.userRole) AuthenticationFactory.userRole = $window.sessionStorage.userRole;
    }
  });

  $rootScope.$on('$routeChangeSuccess', function(event, nextRoute, currentRoute) {
    $rootScope.showLogin = AuthenticationFactory.isLogged;
    $rootScope.showMenu = AuthenticationFactory.isLogged;
    $rootScope.role = AuthenticationFactory.userRole;
    // if the user is already logged in, take him to the home page
    if (AuthenticationFactory.isLogged === true && $location.path() === '/login') {
      $location.path('/');
    }
  });
});