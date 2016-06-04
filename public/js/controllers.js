var app = angular.module('expense-tracker.controllers', []);

app.controller('NavCtrl', function($scope, $location, Nav) {
  // Labels which nav item is selected
  $scope.isActive = function(viewLocation) {
    return viewLocation === $location.path();
  };

  $scope.isLoggedIn = function() {
    return sessionStorage.loggedIn === 'true';
  }

  $scope.logout = function() {
    Nav.logout().then(function() {
      sessionStorage.removeItem('loggedIn');
      sessionStorage.removeItem('name');
      sessionStorage.removeItem('email');
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

app.controller('HomeCtrl', function($scope, $http, $location, Home) {
  // Make the user log in
  if (!sessionStorage.loggedIn) {
    $location.path('/login');
  }

  $scope.recentExpenses;
  $scope.name = sessionStorage.name;

  // Get the user's full name to display on the home page
  if (!sessionStorage.name) {
    Home.getAccountName().then(function(res) {
      $scope.name = sessionStorage.name = res.data.fname + ' ' + res.data.lname;
    });
  }

  Home.getRecentExpenses().then(function(res) {
    $scope.recentExpenses = res.data;
  });
});

// Controller shared between login and sign up pages
app.controller('EntryCtrl', function($scope, $http, $location, Entry) {
  Entry.formFieldAnimations();
  Entry.clearForm();

  $scope.loggedIn;
  $scope.signedUp;
  $scope.err;
  $scope.form = Entry.form();

  // Redirect to home since the user is already logged in
  if (sessionStorage.loggedIn === 'true') {
    $location.path('/');
  }

  $scope.login = function() {
    Entry.login().then(function(res) {
      $scope.loggedIn = sessionStorage.loggedIn = true;

      // User logged in, now redirect to the home page
      if ($scope.loggedIn) {
        sessionStorage.email = Entry.form().email;
        $location.path('/');
      }
    }, function(err) {
      $scope.loggedIn = sessionStorage.loggedIn = false;
      $scope.err = err.data;
    });
  }

  $scope.signup = function() {
    Entry.signup().then(function(res) {
      $scope.signedUp = true;

      // Allow the user to login directly after signing up
      if ($scope.signedUp) {
        $location.path('/login');
      }
    }, function(err) {
      $scope.signedUp = false;
      $scope.err = err.data;
    });
  }
});

app.controller('SearchCtrl', function($scope, $location, Search) {
  // Make the user log in
  if (!sessionStorage.loggedIn) {
    $location.path('/login');
  }

  Search.clearForm()

  $scope.searched;
  $scope.err;
  $scope.form = Search.form();

  $scope.search = function() {
    Search.search().then(function(res) {
      $scope.searched = true;
    }, function(err) {
      $scope.searched = false;
      $scope.err = err.data;
    });
  }
});

app.controller('AddCtrl', function($scope, $location, Add) {
  // Make the user log in
  if (!sessionStorage.loggedIn) {
    $location.path('/login');
  }

  $scope.expenses = Add.expenses();

  $scope.addExpense = function() {
    Add.addExpense();
  };

  $scope.removeExpense = function(index) {
    Add.removeExpense(index);
    $scope.expenses = Add.expenses();
  };

  $scope.submit = function() {
    _.each($scope.expenses, function(expense) {
      Add.submitExpense(expense);
    });
  };
});
