var app = angular.module('expense-tracker.controllers', []);

app.controller('homeCtrl', function($scope, $http, $location) {
  $scope.fname;

  // Make the user log in
  if (!sessionStorage.loggedIn) {
    $location.path('/login');
  }

  $http({
    method: 'GET',
    url: '/accounts',
    params: {email: sessionStorage.userEmail}
  }).then(function(res) {
    $scope.fname = res.data.fname;
  }, function(err) {
    console.log(err.data);
  });
});

app.controller('navCtrl', function($scope, $location, Nav) {
  $scope.isActive = function(viewLocation) {
    return viewLocation === $location.path();
  };

  $scope.isLoggedIn = function() {
    return sessionStorage.loggedIn === 'true';
  }

  $scope.logout = function() {
    Nav.logout().then(function() {
      sessionStorage.removeItem("loggedIn");
      sessionStorage.removeItem("userEmail");
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
  if (sessionStorage.loggedIn === 'true') {
    $location.path('/');
  }

  $scope.login = function() {
    Entry.login($scope.form).then(function(res) {
      $scope.loggedIn = sessionStorage.loggedIn = res.status;
      $scope.err = res.err;

      // User logged in, now redirect to home
      if ($scope.loggedIn) {
        sessionStorage.userEmail = $scope.form.email;
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

app.controller('searchCtrl', function($scope, $location, Search) {

});

app.controller('addCtrl', function($scope, $location, Add) {
  $scope.expenses = [{
    name: '',
    amount: '',
    date: '',
    index: 0
  }];

  $scope.add = function() {
    $scope.expenses.push({
      name: '',
      amount: '',
      date: '',
      index: $scope.expenses.length
    });
  }

  $scope.remove = function(index) {
    var newExpenses = [];

    _.each($scope.expenses, function(expense) {
      if (expense.index != index) {
        newExpenses.push(expense);
      }
    });

    $scope.expenses = newExpenses;
  }

  $scope.submit = function() {
    _.each($scope.expenses, function(expense) {
      Add.submitExpense(expense);
    })
  }
});
