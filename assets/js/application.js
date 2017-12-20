require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

// require("jquery-ui/build/release.js");
require("blueimp-file-upload/js/jquery.fileupload.js");

$(() => {
    $('#fileupload').fileupload({
        dataType: 'json',
        autoUpload: true,
        //acceptFileTypes: /(\.|\/)(gif|jpe?g|png)$/i,
        //maxFileSize: 999000,
        // change: function (e, data) {
        //     console.log(data.files[0].name)
        //     $('#progress label').text(data.files[0].name)
        // },
        done: function (e, data) {
            console.log("DONE")
            $.each(data.files, function (index, file) {
                console.log("done: "+file.name)
                $('<p/>').text(file.name).appendTo('#uploaded-file');
                $('#medium-FileID').attr("value", data.result.fileID);
                $('#medium-FileName').attr("value", data.result.fileName);
            });
        },
        progressall: function (e, data) {
            var progress = parseInt(data.loaded / data.total * 100, 10);
            $('#progress .progress-bar').css(
                'width',
                progress + '%'
            );
        }
    });
});
