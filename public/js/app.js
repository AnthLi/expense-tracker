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

app.config(function($routeProvider, $locationProvider) {
  $routeProvider

    // Home page
    .when('/', {
      templateUrl: '/public/templates/home.html',
      controller: 'HomeCtrl'
    })

    // Login page
    .when('/login', {
      templateUrl: '/public/templates/login.html',
      controller: 'EntryCtrl'
    })

    // Sign up page
    .when('/signup', {
      templateUrl: '/public/templates/signup.html',
      controller: 'EntryCtrl'
    })

    // Search expenses page
    .when('/search', {
      templateUrl: '/public/templates/search.html',
      controller: 'SearchCtrl'
    })

    // Add expenses page
    .when('/add', {
      templateUrl: '/public/templates/add.html',
      controller: 'AddCtrl'
    });

  $locationProvider.html5Mode(true);
});