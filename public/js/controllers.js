myApp.controller("HeaderCtrl", ['$scope', '$location',
  function($scope, $location) {

    $scope.isActive = function(route) {
      return route === $location.path();
    };
  }
]);

myApp.controller("HomeCtrl", ['$scope', 'bookmarkFactory',
  function($scope, bookmarkFactory) {
    $scope.name = "Bookmarks";

    $scope.bookmarks = [];

    // Access the factory and get all the bookmarks
    bookmarkFactory.getAllBookmarks().then(function (data) {
      $scope.bookmarks = data.data;
    });
  }

]);

myApp.controller("NewCtrl", ['$scope', '$location', 'bookmarkFactory',
  function($scope, $location, bookmarkFactory) {
    $scope.name = "New Bookmark";

    $scope.new = function (bookmark) {
      bookmarkFactory.new(bookmark).then(function (data){
        $location.path("/");
      });
    };
  }
]);

myApp.controller("EditCtrl", ['$scope', '$window', '$location', 'bookmarkFactory', '$routeParams',
  function($scope, $window, $location, bookmarkFactory, $routeParams) {
    $scope.name = "Edit " + $routeParams.id;

    $scope.bookmark = {};

    bookmarkFactory.getBookmark($routeParams.id).then(function (data){
      $scope.bookmark = data.data;
    });



    $scope.update = function (bookmark) {
      bookmarkFactory.update(bookmark).then(function (data){
        $location.path("/");
      });
    };

    $scope.delete = function (bookmark) {
      var deleteBookmark = $window.confirm('Are you absolutely sure you want to delete?');

      if(deleteBookmark){
        bookmarkFactory.delete(bookmark).then(function (data){
          $location.path("/");
        });
      }else{
        return;
      }
    };
  }
]);