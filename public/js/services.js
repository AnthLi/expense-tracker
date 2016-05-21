var app = angular.module('expense-tracker.services', []);

app.factory('Nav', function($http) {
  return {
    logout: function() {
      return $http({
        method: 'POST',
        url: '/logout'
      });
    }
  }
});

app.factory('Entry', function($http) {
  return {
    // Source: http://goo.gl/GxKNXk
    // Shift the field labels when user input is detected
    formFieldAnimations: function() {
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
    },

    login: function(credentials) {
      // Source: http://goo.gl/wPHJrE
      // Send login POST values to the server
      return $http({
        method: 'POST',
        url: '/login',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        transformRequest: function(obj) {
          var str = [];
          for (var p in obj) {
            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
          }

          return str.join("&");
        },
        data: credentials
      }).then(function(res) {
        return {
          status: true,
          err: null
        };
      }, function(err) {
        return {
          status: false,
          err: err.data
        };
      });
    },

    signup: function(form) {
      // Source: http://goo.gl/wPHJrE
      // Send sign up POST values to the server
      return $http({
        method: 'POST',
        url: '/signup',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        transformRequest: function(obj) {
          var str = [];
          for (var p in obj) {
            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
          }

          return str.join("&");
        },
        data: form
      }).then(function(res) {
        return {
          status: true,
          err: null
        };
      }, function(err) {
        return {
          status: false,
          err: err.data
        };
      });
    }
  }
});

app.factory('Search', function($http) {
  return {

  }
});

app.factory('Add', function($http) {
  return {
    submitExpense: function(expense) {
      return $http({
        method: 'POST',
        url: '/add',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        transformRequest: function(obj) {
          var str = [];
          for (var p in obj) {
            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
          }

          return str.join("&");
        },
        data: expense
      });
    }
  }
});