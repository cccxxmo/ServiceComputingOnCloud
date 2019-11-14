$(document).ready(function() {
    $.ajax({
        url: "/js"
    }).then(function(data) {
       $('.title').append(data.title);
    });
});
