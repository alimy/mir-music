angular.module('SpringMusic', ['albums', 'errors', 'status', 'info', 'ngRoute', 'ui.directives']).
    config(function ($locationProvider, $routeProvider) {
        // $locationProvider.html5Mode(true);

        $routeProvider.when('/errors', {
            controller: 'ErrorsController',
            templateUrl: 'static/templates/errors.html'
        });
        $routeProvider.otherwise({
            controller: 'AlbumsController',
            templateUrl: 'static/templates/albums.html'
        });
    }
);
