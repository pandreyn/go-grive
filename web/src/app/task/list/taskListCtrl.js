'use strict';

angular.module('asc.taskControllers')
    .controller('taskListCtrl',
    [
      '$scope',
      '$state',
      'Task',
      '$mdDialog',
      function ($scope,
                $state,
                Task,
                $mdDialog) {

        function refreshTasks() {
          Task.getTasks().then(function (result) {
            $scope.Tasks = result;
          });
        }

        function addTask($event) {

          $mdDialog.show({
            targetEvent: $event,
            templateUrl: 'app/task/list/newTask.tpl.html',
            locals: {name: 'Bob'},
            controller: ['$scope', 'name', function ($scope, name) {
              $scope.name = name;
              $scope.closeDialog = function () {
                $mdDialog.hide({title: 'new' + new Date()});
              }
            }]
          }).then(
              function (task) {
                //Task.addTask();
                //refreshTasks();
                console.log('modal succses result:', task);
              },
              function () {
                console.log('modal error');
              }
          );

        }

        $scope.addTask = addTask;


        function activate() {
          refreshTasks();
        }

        activate();
        console.log('Task list')
      }]);
