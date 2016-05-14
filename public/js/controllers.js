var app = angular.module('expense-tracker.controllers', []);

app.controller('headerCtrl', function($scope) {
  $scope.isActive = function (viewLocation) {
    return viewLocation === $location.path();
  };
});

app.controller('mainCtrl', function($scope) {

});

// Controller shared between login and sign up pages
app.controller('entryCtrl', function($scope, $http) {
  // Source: http://goo.gl/GxKNXk
  // Shift the field labels when user input is detected
  $('.form').find('input').on('keyup blur focus', function(e) {
    var $this = $(this);
    var label = $this.prev('label');

    if (e.type === 'keyup') {
      if ($this.val() === '') {
        label.removeClass('active highlight');
      } else {
        label.addClass('active highlight');
      }
    } else if (e.type === 'blur') {
      if ($this.val() === '') {
        label.removeClass('active highlight');
      } else {
        label.removeClass('highlight');
      }
    } else if (e.type === 'focus') {
      if ($this.val() === '') {
        label.removeClass('highlight');
      } else if ($this.val() !== '') {
        label.addClass('highlight');
      }
    }
  });

  $scope.form = {};

  $scope.login = function() {
    $http({
      method: 'POST',
      url: '/login',
      data: $scope.form
    });
  };

  $scope.signup = function() {
    $http({
      method: 'POST',
      url: '/signup',
      data: $scope.form
    });
  };
});

app.controller('addCtrl', function($scope) {

});

app.controller('searchCtrl', function($scope) {

});