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

app.factory('Home', function($http) {
  return {
    getAccountName: function() {
      return $http({
        method: 'GET',
        url: '/accounts',
        params: {email: sessionStorage.email}
      });
    },

    getRecentExpenses: function() {
      return $http({
        method: 'GET',
        url: '/expenses/recent',
        params: {
          email: sessionStorage.email,
        }
      });
    }
  }
});

app.factory('Entry', function($http) {
  var form = {
    fname: '',
    lname: '',
    email: '',
    password: ''
  };

  return {
    form: function() {
      return form;
    },

    clearForm: function() {
      form.fname = '';
      form.lname = '',
      form.email = '',
      form.password = ''
    },

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

    login: function() {
      // Source: http://goo.gl/wPHJrE
      // Send login form data to the server
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
        data: form
      });
    },

    signup: function() {
      // Source: http://goo.gl/wPHJrE
      // Send sign up form data to the server
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
      });
    }
  }
});

app.factory('Search', function($http) {
  var form = {
    email: sessionStorage.email,
    name: '',
    date: ''
  };

  return {
    form: function() {
      return form;
    },

    clearForm: function() {
      form.name = '';
      form.date = '';
    },

    search: function() {
      return $http({
        method: 'POST',
        url: '/search',
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
      });
    }
  }
});

app.factory('Add', function($http) {
  var expenses = [{
    name: '',
    amount: '',
    date: '',
    index: 0
  }];

  return {
    expenses: function() {
      return expenses;
    },

    // Simply add another empty expense form with the index equaling the number
    // of existing expenses
    addExpense: function() {
      expenses.push({
        name: '',
        amount: '',
        date: '',
        index: expenses.length
      });
    },

    // Remove the expense based on its index and assign the original array to
    // the new array of expenses
    removeExpense: function(index) {
      var newExpenses = [];

      _.each(expenses, function(expense) {
        if (expense.index != index) {
          newExpenses.push(expense);
        }
      });

      expenses = newExpenses;
    },

    // Source: http://goo.gl/wPHJrE
    // Send the expenses form data to the server
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