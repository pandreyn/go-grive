'use strict';
/*

app.js

 */

//console.info('Welcome to go-grive', VERSION);

// Required for angular-material
require('angular-aria');
require('angular-animate');
require('material');

// angular imports
require('angular-ui-router/release/angular-ui-router');
require('angular-breadcrumb/dist/angular-breadcrumb');

require('app/task/taskServices');


var asc = angular.module('gogrive', [
  'ngAnimate',
  'ui.router',
  'ngMaterial',
  'ncy-angular-breadcrumb',
  'asc.taskControllers',
  'asc.taskServices'
]);

asc.config(function ($locationProvider, $urlRouterProvider, $stateProvider, $mdThemingProvider, $breadcrumbProvider) {
  //$mdThemingProvider.theme('default')
  //    .primaryPallete('indigo')
  //    .accentPallete('light-blue');

  $mdThemingProvider.theme('default')
      .primaryPalette('indigo')
      .accentPalette('pink')
      .warnPalette('red')
      .backgroundPalette('grey');

  $locationProvider.html5Mode(true);
  $urlRouterProvider.otherwise("/");

  //$stateProvider
  //    .state('app', {
  //      url: '/',
  //      abstract: true,
  //      ncyBreadcrumb: {
  //        label: 'Gomado'
  //      }
  //    });

  $breadcrumbProvider.setOptions({
    templateUrl: '/app/material-breadcrumbs.html'
  })
});

asc.run(function ($rootScope) {
  // Set this so we can get it from html easily
  $rootScope.BUILDINFO = {
    version: VERSION,
    devMode: DEVMODE
  }
});
