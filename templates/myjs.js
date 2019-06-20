//添加导航栏
$(function(){
    $("#nav-placeholder").load("nav");
});
// //添加导航栏监听
// layui.use('element', function(){
//     var element = layui.element;
//     element.on('nav(chaoren-nav)', function(elem){
//         console.log(elem); //得到当前点击的DOM对象
//         alert(elem);
//     });
//     //…
// });

//刷新表格
function subForm() {
    $.ajax({
        url: 'http://localhost:8080/v1/user', //请求的url
        type: 'get', //请求的方式
        data: '', //form表单里要提交的内容，里面的input等写上name就会提交，这是序列化
        error:function (data) {
            alert('请求失败');
        },
        success:function (data) {
            var str1 = "";
            //清空table中的html
            $("#body_user").html("");
            // alert('查询成功，共有数据'+data.data.userList.length+'条')
            // 第一个data是外面的总数据包，再往下细分
            for(var i = 0;i<data.data.userList.length;i++){
                $id = data.data.userList[i].id;
                $username = data.data.userList[i].username;
                $sayHello = data.data.userList[i].sayHello;
                $password = data.data.userList[i].password;
                $createTime = data.data.userList[i].createdAt;
                $updateTime = data.data.userList[i].updatedAt;
                //这里有用\" 对引号进行转译，因为这段是html格式，字符串要用""包裹起来，但是"又用来分割字符串，所以包裹的"需要用\转译
                str1 = "<tr>" +
                    "<td>"+$id + "</td>" +
                    "<td>"+$username + "</td>" +
                    "<td>"+$sayHello + "</td>" +
                    "<td>"+$password + "</td>" +
                    "<td>"+$createTime + "</td>" +
                    "<td>"+$updateTime + "</td>" +
                    "<td><button class='layui-btn layui-btn-sm' id='btn_update' onclick='update("+$id+",\""+$username+"\",\""+$password+"\",\""+$updateTime+"\")'>更新</button></td>" +
                    "</tr>";
                $("#body_user").append(str1);
            }
        }
    });
}

function subForm2() {
    $.ajax({
        url: 'http://localhost:8080/v1/product  ', //请求的url
        type: 'get', //请求的方式
        data: '', //form表单里要提交的内容，里面的input等写上name就会提交，这是序列化
        error:function (data) {
            alert('请求失败');
        },
        success:function (data) {
            var str1 = "";
            //清空table中的html
            $("#body_product").html("");
            // alert('查询成功，共有数据'+data.data.productList.length+'条')
            // 第一个data是外面的总数据包，再往下细分
            for(var i = 0;i<data.data.productList.length;i++){
                $id = data.data.productList[i].id;
                $name = data.data.productList[i].name;
                $author = data.data.productList[i].author;
                $intro = data.data.productList[i].intro;
                $label = data.data.productList[i].label;
                $createTime = data.data.productList[i].create_time;
                $updateTime = data.data.productList[i].update_time;
                $state = data.data.productList[i].state;
                $pic = data.data.productList[i].pic;

                //这里有用\" 对引号进行转译，因为这段是html格式，字符串要用""包裹起来，但是"又用来分割字符串，所以包裹的"需要用\转译
                str1 = "<tr>" +
                    "<td>"+$id + "</td>" +
                    "<td>"+$name + "</td>" +
                    "<td>"+$author + "</td>" +
                    "<td>"+$intro + "</td>" +
                    "<td>"+$label + "</td>" +
                    "<td>"+$createTime + "</td>" +
                    "<td>"+$updateTime + "</td>" +
                    "<td>"+$state + "</td>" +
                    "<td>"+$pic + "</td>" +
                    "<td><button class='layui-btn layui-btn-sm' id='btn_update'>更新</button></td>" +
                    "</tr>";
                $("#body_product").append(str1);
            }
        }
    });
}

$(document).ready(function(){
    $("#button1").click(function(){
        $("#div1").load("http://localhost:8080/v1/user");
    });
});

//更新updateTime，更新完后刷新表格
function update(id,username,password,updateTime) {
    layui.use('layer', function(){
        var layer = layui.layer;

        layer.open({
            title:"配置详情",
            type:2,
            content:"http://localhost:8080/user_config",
            area: ['500px', '600px']
        })
    });

    // $.ajax({
    //     url: 'http://localhost:8080/v1/user/'+id, //请求的url
    //     type: 'put', //请求的方式
    //     dateType:'json',
    //     data: '{"username":"'+username+'","password":"'+password+'","updatedAt":"'+updateTime+'"}',
    //     // data: '?id='+id+'&username='+username+'&deletedAt='+'&password='+password+'&createTime='+createTime+'&updatedAt='+updateTime, //form表单里要提交的内容，里面的input等写上name就会提交，这是序列化
    //     error:function (data) {
    //         alert('请求失败');
    //     },
    //     success:function (data) {
    //         subForm()
    //     }
    // });
}

