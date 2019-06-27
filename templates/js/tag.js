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

function subForm_tag() {
    $.ajax({
        url: 'http://localhost:8080/v1/tag  ', //请求的url
        type: 'get', //请求的方式
        data: '', //form表单里要提交的内容，里面的input等写上name就会提交，这是序列化
        error:function (data) {
            alert('请求失败');
        },
        success:function (data) {
            var str1 = "";
            //清空table中的html
            $("#body_tag").html("");
            // alert('查询成功，共有数据'+data.data.productList.length+'条')
            // 第一个data是外面的总数据包，再往下细分
            for(var i = 0;i<data.data.tagList.length;i++){
                $id = data.data.tagList[i].id;
                $name = data.data.tagList[i].name;
                $source = data.data.tagList[i].source;
                $category = data.data.tagList[i].category;
                $property = data.data.tagList[i].property;
                $state = data.data.tagList[i].state;
                $createTime = data.data.tagList[i].create_time;
                $updateTime = data.data.tagList[i].update_time;
                $operator = data.data.tagList[i].operator;

                //这里有用\" 对引号进行转译，因为这段是html格式，字符串要用""包裹起来，但是"又用来分割字符串，所以包裹的"需要用\转译
                str1 = "<tr>" +
                    "<td>"+$id + "</td>" +
                    "<td>"+$name + "</td>" +
                    "<td>"+$source + "</td>" +
                    "<td>"+$category + "</td>" +
                    "<td>"+$property + "</td>" +
                    "<td>"+$state + "</td>" +
                    "<td>"+$createTime + "</td>" +
                    "<td>"+$updateTime + "</td>" +
                    "<td>"+$operator + "</td>" +
                    "<td><button class='layui-btn layui-btn-sm' id='btn_update' onclick='update()'>同步</button><button class='layui-btn layui-btn-sm' id='btn_up'>上架</button><button class='layui-btn layui-btn-sm' id='btn_update'>下架</button></td>" +
                    "</tr>";
                $("#body_tag").append(str1);
            }
        }
    });
}

$(document).ready(function(){
    $("#button1").click(function(){
        $("#div1").load("http://localhost:8080/v1/user");
    });
});

function call_create_tag() {
    layui.use('layer', function () {
        var layer = layui.layer;

        layer.open({
            title: "标签配置详情",
            type: 2,
            content: "http://localhost:8080/tag_config",
            area: ['500px', '600px']
        })

    });
}

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

