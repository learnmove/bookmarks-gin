myApp.factory('bookmarkFactory', function($http) {
  var urlBase = '/api/v1/bookmarks';
  var _bookmarkFactory = {};

  _bookmarkFactory.getAllBookmarks = function() {
    return $http.get(urlBase);
  };

  _bookmarkFactory.getBookmark = function(id) {
    return $http.get(urlBase + "/" + id);
  };

  _bookmarkFactory.new = function(bookmark) {
    return $http.post(urlBase, bookmark);
  };

  _bookmarkFactory.update = function(bookmark) {
    return $http.put(urlBase + "/" + bookmark.id, bookmark);
  };

  _bookmarkFactory.delete = function(bookmark) {
    return $http.delete(urlBase + "/" + bookmark.id);
  };

  return _bookmarkFactory;
});