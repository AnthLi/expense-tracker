var app = angular.module('expense-tracker.controllers', []);

var loggedIn;

app.controller('mainCtrl', function($scope, $location) {
  // Make the user log in
  if (!loggedIn) {
    $location.path('/login');
  }
});

app.controller('navCtrl', function($scope, $location, Nav) {
  $scope.isActive = function(viewLocation) {
    return viewLocation === $location.path();
  };

  $scope.logout = function() {
    Nav.logout().then(function(res) {
      loggedIn = false;
      $location.path('/login');
    });
  }

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
});

// Controller shared between login and sign up pages
app.controller('entryCtrl', function($scope, $http, $location, Entry) {
  Entry.formFieldAnimations();

  $scope.form = {};
  $scope.loggedIn;
  $scope.signedUp;
  $scope.err;

  // Redirect to home since the user is already logged int
  if (loggedIn) {
    $location.path('/');
  }

  $scope.login = function() {
    Entry.login($scope.form).then(function(res) {
      $scope.loggedIn = loggedIn = res.status;
      $scope.err = res.err;

      // User logged in, now redirect to home
      if ($scope.loggedIn) {
        $location.path('/');
      }
    });
  }

  $scope.signup = function() {
    Entry.signup($scope.form).then(function(res) {
      $scope.signedUp = res.status;
      $scope.err = res.err;

      // Allow the user to login after signing up
      if ($scope.signedUp) {
        $location.path('/login');
      }
    });
  }
});

app.controller('addCtrl', function($scope) {

});

app.controller('searchCtrl', function($scope) {

});