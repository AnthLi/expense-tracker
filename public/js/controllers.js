var app = angular.module('expense-tracker.controllers', []);

app.controller('headerCtrl', function($scope, $location) {
  $scope.isActive = function(viewLocation) {
    return viewLocation === $location.path();
  };

  // Close the menu when clicking outside of it.
  // This only applies to when the dropdown button appears, which is when the
  // browser width is less than 768 pixels.
  $(document).on('click', function(event) {
    if ($(window).width() < 768) {
      // Get each significant section of the navbar/dropdown menu
      var navbar = $(event.target).closest('.navbar').length;
      var collapsed = $(event.target).closest('.navbar-collapse').length;
      var toggled = $(event.target).closest('.navbar-toggle').length;
      var expanded = $('#navbar-collapse[aria-expanded="true"]').length;
      var navList = $(event.target).closest('.navbar-nav li a').length;

      if (!navbar && !collapsed && !toggled && expanded || navList) {
        // Close the menu by triggering a click
        $('.navbar-toggle').click();
      }
    }
  });

  // $scope.toggle = function() {
  //   $scope.toggled = !$scope.toggled;
  //   $scope.state = !$scope.state;
  // };

  // $(document).on('click', function(event) {
  //   var length = $(event.target).closest('#menu').length;
  //   if (!length && event.target.id !== 'navButton' && $scope.toggled) {
  //     $scope.$apply(function() {
  //       $scope.toggle();
  //     });
  //   }
  // });
});

app.controller('mainCtrl', function($scope) {});

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