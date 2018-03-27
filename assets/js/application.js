require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");
require("admin-lte/dist/js/adminlte.js");

$(() => {
    let $pushMenu = $('[data-toggle="push-menu"]').data('lte.pushmenu')
    $pushMenu.expandOnHover();

});
