'use strict';

var taskControllers = angular.module('asc.taskControllers', []);

require('angular-resource/angular-resource');

var taskServices = angular.module('asc.taskServices', ['ngResource']);

require('app/task/list/taskListCtrl');
require('app/task/details/taskDetailsCtrl');


taskServices.config(function ($locationProvider, $stateProvider, $urlRouterProvider) {
  $stateProvider
      .state('tasks', {
        url: '/',
        templateUrl: '/app/task/list/list.html',
        controller: 'taskListCtrl',
        ncyBreadcrumb: {
          label: 'Tasks'
        }
      })
      .state('tasks.details', {
        url: '/:Id',
        templateUrl: '/app/task/details/details.html',
        controller: 'taskDetailsCtrl',
        ncyBreadcrumb: {
          label: 'Details {{currentTask.Id}}'
        }
      })
});

taskServices.service('Task', function ($q) {
  var _self = this;
  this.tasks =  [{
    id: '1',
    title: "Get er done",
    description: "Where is that from?",
    complete: false
  }, {
    id: '2',
    title: "Add some more tasks",
    description: "We need alot of them",
    complete: true
  }, {
    id: '3',
    title: "Make the add task btn work",
    description: "That'd be nice wouldn't it? I think a dialog would work perfectly for this.",
    complete: false
  }, {
    id: '4',
    title: "Finish implementing the task details",
    description: "Also find out what to put in there? May make this description have an ellipse when it's really long",
    complete: false
  }, {
    id: '5',
    title: "Persist tasks to local storage/backend w/ mongodb",
    description: "Which will we choose? Easy and simple local storage or mongodb?",
    complete: false
  }];

  this.getTaskDetails = function (id) {
    var deferred = $q.defer();

    deferred.resolve(_self.tasks[id]);

    return deferred.promise
  };

  this.getTasks = function () {
    var deferred = $q.defer();

    deferred.resolve(_self.tasks);

    return deferred.promise
  };

  this.addTask = function (task) {
    var deferred = $q.defer();

    _self.tasks.push(task);
    deferred.resolve(_self.tasks);

    return deferred.promise
  }

});
