function doLogin() {
    var frm = $('#lForm')
    var params = JSON.stringify(frm.serializeJSON());
    console.log(3)
    console.log(params)
    console.log(123)
    $.ajax({
        url: "/user/login",
        type: "POST",
        dataType: 'json',
        data: params,
        contentType: "application/json",  //缺失会出现URL编码，无法转成json对象
        success: function (data,textStatus,request) {
            var file = new File(["123123"],  "ddd.json", { type: "text/plain;charset=utf-8" })
            saveAs(file)

/*            var FileSaver = require('file-saver');
            var blob = new Blob(["Hello, world!"], {type: "text/plain;charset=utf-8"});
            FileSaver.saveAs(blob, "hello world.txt");*/

            console.log(data);
            console.log(request.getResponseHeader('Token'))
            alert(request.getResponseHeader('Token'));
        },
        error: function (data) {
            console.log(data);
        }
    });

}

$('#lForm').submit(function () {
    var frm = $('#lForm')
    var params = JSON.stringify(frm.serializeJSON());
    console.log(3)
    console.log(params)
    console.log(123)
    $.ajax({
        url: "/user/login",
        type: "POST",
        dataType: 'json',
        data: params,
        contentType: "application/json",  //缺失会出现URL编码，无法转成json对象
        success: function (data, headers) {
            console.log(data);
            console.log(headers.getAllResponseHeaders());//
        },
        error: function (data) {
            console.log(data);
        }
    });
})