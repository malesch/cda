require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");
require("admin-lte/dist/js/adminlte.js");

$(() => {
    // Show/collapse sidebar on mouse hover
    $('.main-sidebar').hover(
        function() { $('body').removeClass('sidebar-collapse') },
        function() { $('body').addClass('sidebar-collapse')
    })
});
