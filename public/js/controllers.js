var app = angular.module('expense-tracker.controllers', []);

app.controller('headerCtrl', function($scope, $location) {
  $scope.isActive = function(viewLocation) {
    return viewLocation === $location.path();
  };
});

app.controller('mainCtrl', function($scope) {

});

// Controller shared between login and sign up pages
app.controller('entryCtrl', function($scope, $http, Entry) {
  Entry.formFunction();

  $scope.form = {};
  $scope.loginSuccess;
  $scope.signupSuccess;
  $scope.err;

  $scope.login = function() {
    Entry.postLoginInfo($scope.form).then(function(res) {
      $scope.loginSuccess = res.status;
      $scope.err = res.err;
    });
  }

  $scope.signup = function() {
    Entry.postSignupInfo($scope.form).then(function(res) {
      $scope.signupSuccess = res.status;
      $scope.err = res.err;
    });
  }
});

app.controller('addCtrl', function($scope) {

});

app.controller('searchCtrl', function($scope) {

});