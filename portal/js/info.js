angular.module('info', ['ngResource']).
    factory('Info', function ($resource) {
        return $resource('v1/appinfo/');
    });

function InfoController($scope, Info) {
    $scope.info = Info.get();
}
