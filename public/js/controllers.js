var app = angular.module('expense-tracker.controllers', []);

app.controller('mainCtrl', function($scope, $location) {
  // Make the user log in
  if (!localStorage.loggedIn) {
    $location.path('/login');
  }
});

app.controller('navCtrl', function($scope, $location, Nav) {
  $scope.isActive = function(viewLocation) {
    return viewLocation === $location.path();
  };

  $scope.isLoggedIn = function() {
    return localStorage.loggedIn === 'true';
  }

  $scope.logout = function() {
    Nav.logout().then(function(res) {
      localStorage.removeItem("loggedIn");
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
      var expanded = $('.navbar-collapse[aria-expanded="true"]').length;
      var navList = $(event.target).closest('.navbar-nav li a').length;
      var logout = $(event.target).closest('#nav-logout form button').length;

      if (!navbar && !collapsed && !toggled && expanded || navList || logout) {
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

  // Redirect to home since the user is already logged in
  if ($scope.loggedIn) {
    $location.path('/');
  }

  $scope.login = function() {
    Entry.login($scope.form).then(function(res) {
      localStorage.loggedIn = res.status;
      $scope.loggedIn = res.status;
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

app.controller('searchCtrl', function($scope) {

});

app.controller('addCtrl', function($scope) {

});
