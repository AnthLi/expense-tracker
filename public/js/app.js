var app = angular.module('expense-tracker', [
  'ngAnimate',
  'ngRoute',
  'ngTouch',
  'expense-tracker.controllers',
  'expense-tracker.services'
]);

// lodash
app.constant('_',
  window._
);

app.config(['$routeProvider', function($routeProvider) {
  $routeProvider

    // Home page
    .when('/', {
      templateUrl: '/public/templates/home.html',
      controller: 'homeCtrl'
    })

    // Login page
    .when('/login', {
      templateUrl: '/public/templates/login.html',
      controller: 'entryCtrl'
    })

    // Sign up page
    .when('/signup', {
      templateUrl: '/public/templates/signup.html',
      controller: 'entryCtrl'
    })

    // Search expenses page
    .when('/search', {
      templateUrl: '/public/templates/search.html',
      controller: 'searchCtrl'
    })

    // Add expenses page
    .when('/add', {
      templateUrl: '/public/templates/add.html',
      controller: 'addCtrl'
    });
}]);