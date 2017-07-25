/**
 * Created by Administrator on 2017/7/7.
 */
var theme = {
    color: [
        '#26B99A', '#34495E', '#BDC3C7', '#3498DB',
        '#9B59B6', '#8abb6f', '#759c6a', '#bfd3b7'
    ],

    title: {
        itemGap: 8,
        textStyle: {
            fontWeight: 'normal',
            color: '#408829'
        }
    },

    dataRange: {
        color: ['#1f610a', '#97b58d']
    },

    toolbox: {
        color: ['#408829', '#408829', '#408829', '#408829']
    },

    tooltip: {
        backgroundColor: 'rgba(0,0,0,0.5)',
        axisPointer: {
            type: 'line',
            lineStyle: {
                color: '#408829',
                type: 'dashed'
            },
            crossStyle: {
                color: '#408829'
            },
            shadowStyle: {
                color: 'rgba(200,200,200,0.3)'
            }
        }
    },

    dataZoom: {
        dataBackgroundColor: '#eee',
        fillerColor: 'rgba(64,136,41,0.2)',
        handleColor: '#408829'
    },
    grid: {
        borderWidth: 0
    },

    categoryAxis: {
        axisLine: {
            lineStyle: {
                color: '#408829'
            }
        },
        splitLine: {
            lineStyle: {
                color: ['#eee']
            }
        }
    },

    valueAxis: {
        axisLine: {
            lineStyle: {
                color: '#408829'
            }
        },
        splitArea: {
            show: true,
            areaStyle: {
                color: ['rgba(250,250,250,0.1)', 'rgba(200,200,200,0.1)']
            }
        },
        splitLine: {
            lineStyle: {
                color: ['#eee']
            }
        }
    },
    timeline: {
        lineStyle: {
            color: '#408829'
        },
        controlStyle: {
            normal: {color: '#408829'},
            emphasis: {color: '#408829'}
        }
    },

    k: {
        itemStyle: {
            normal: {
                color: '#68a54a',
                color0: '#a9cba2',
                lineStyle: {
                    width: 1,
                    color: '#408829',
                    color0: '#86b379'
                }
            }
        }
    },
    map: {
        itemStyle: {
            normal: {
                areaStyle: {
                    color: '#ddd'
                },
                label: {
                    textStyle: {
                        color: '#c12e34'
                    }
                }
            },
            emphasis: {
                areaStyle: {
                    color: '#99d2dd'
                },
                label: {
                    textStyle: {
                        color: '#c12e34'
                    }
                }
            }
        }
    },
    force: {
        itemStyle: {
            normal: {
                linkStyle: {
                    strokeColor: '#408829'
                }
            }
        }
    },
    chord: {
        padding: 4,
        itemStyle: {
            normal: {
                lineStyle: {
                    width: 1,
                    color: 'rgba(128, 128, 128, 0.5)'
                },
                chordStyle: {
                    lineStyle: {
                        width: 1,
                        color: 'rgba(128, 128, 128, 0.5)'
                    }
                }
            },
            emphasis: {
                lineStyle: {
                    width: 1,
                    color: 'rgba(128, 128, 128, 0.5)'
                },
                chordStyle: {
                    lineStyle: {
                        width: 1,
                        color: 'rgba(128, 128, 128, 0.5)'
                    }
                }
            }
        }
    },
    gauge: {
        startAngle: 225,
        endAngle: -45,
        axisLine: {
            show: true,
            lineStyle: {
                color: [[0.2, '#86b379'], [0.8, '#68a54a'], [1, '#408829']],
                width: 8
            }
        },
        axisTick: {
            splitNumber: 10,
            length: 12,
            lineStyle: {
                color: 'auto'
            }
        },
        axisLabel: {
            textStyle: {
                color: 'auto'
            }
        },
        splitLine: {
            length: 18,
            lineStyle: {
                color: 'auto'
            }
        },
        pointer: {
            length: '90%',
            color: 'auto'
        },
        title: {
            textStyle: {
                color: '#333'
            }
        },
        detail: {
            textStyle: {
                color: 'auto'
            }
        }
    },
    textStyle: {
        fontFamily: 'Arial, Verdana, sans-serif'
    }
};
var theme2 = {
    color: [
        '#d87c7c',
        '#919e8b',
        '#d7ab82',
        '#61a0a8',
        '#6e7074',
        '#efa18d',
        '#787464',
        '#cc7e63',
        '#724e58',
        '#4b565b'
    ],

    title: {
        itemGap: 8,
        textStyle: {
            fontWeight: 'normal',
            color: '#408829'
        }
    },

    dataRange: {
        color: ['#1f610a', '#97b58d']
    },

    toolbox: {
        color: ['#408829', '#408829', '#408829', '#408829']
    },

    tooltip: {
        backgroundColor: 'rgba(0,0,0,0.5)',
        axisPointer: {
            type: 'line',
            lineStyle: {
                color: '#408829',
                type: 'dashed'
            },
            crossStyle: {
                color: '#408829'
            },
            shadowStyle: {
                color: 'rgba(200,200,200,0.3)'
            }
        }
    },

    dataZoom: {
        dataBackgroundColor: '#eee',
        fillerColor: 'rgba(64,136,41,0.2)',
        handleColor: '#408829'
    },
    grid: {
        borderWidth: 0
    },

    categoryAxis: {
        axisLine: {
            lineStyle: {
                color: '#408829'
            }
        },
        splitLine: {
            lineStyle: {
                color: ['#eee']
            }
        }
    },

    valueAxis: {
        axisLine: {
            lineStyle: {
                color: '#408829'
            }
        },
        splitArea: {
            show: true,
            areaStyle: {
                color: ['rgba(250,250,250,0.1)', 'rgba(200,200,200,0.1)']
            }
        },
        splitLine: {
            lineStyle: {
                color: ['#eee']
            }
        }
    },
    timeline: {
        lineStyle: {
            color: '#408829'
        },
        controlStyle: {
            normal: {color: '#408829'},
            emphasis: {color: '#408829'}
        }
    },

    k: {
        itemStyle: {
            normal: {
                color: '#68a54a',
                color0: '#a9cba2',
                lineStyle: {
                    width: 1,
                    color: '#408829',
                    color0: '#86b379'
                }
            }
        }
    },
    map: {
        itemStyle: {
            normal: {
                areaStyle: {
                    color: '#ddd'
                },
                label: {
                    textStyle: {
                        color: '#c12e34'
                    }
                }
            },
            emphasis: {
                areaStyle: {
                    color: '#99d2dd'
                },
                label: {
                    textStyle: {
                        color: '#c12e34'
                    }
                }
            }
        }
    },
    force: {
        itemStyle: {
            normal: {
                linkStyle: {
                    strokeColor: '#408829'
                }
            }
        }
    },
    chord: {
        padding: 4,
        itemStyle: {
            normal: {
                lineStyle: {
                    width: 1,
                    color: 'rgba(128, 128, 128, 0.5)'
                },
                chordStyle: {
                    lineStyle: {
                        width: 1,
                        color: 'rgba(128, 128, 128, 0.5)'
                    }
                }
            },
            emphasis: {
                lineStyle: {
                    width: 1,
                    color: 'rgba(128, 128, 128, 0.5)'
                },
                chordStyle: {
                    lineStyle: {
                        width: 1,
                        color: 'rgba(128, 128, 128, 0.5)'
                    }
                }
            }
        }
    },
    gauge: {
        startAngle: 225,
        endAngle: -45,
        axisLine: {
            show: true,
            lineStyle: {
                color: [[0.2, '#86b379'], [0.8, '#68a54a'], [1, '#408829']],
                width: 8
            }
        },
        axisTick: {
            splitNumber: 10,
            length: 12,
            lineStyle: {
                color: 'auto'
            }
        },
        axisLabel: {
            textStyle: {
                color: 'auto'
            }
        },
        splitLine: {
            length: 18,
            lineStyle: {
                color: 'auto'
            }
        },
        pointer: {
            length: '90%',
            color: 'auto'
        },
        title: {
            textStyle: {
                color: '#333'
            }
        },
        detail: {
            textStyle: {
                color: 'auto'
            }
        }
    },
    textStyle: {
        fontFamily: 'Arial, Verdana, sans-serif'
    }
};
var cache = {
    index:[],
    page: 1,
    tasklist:[],
    taskCount:0,
    machineCount:0,
    clusterCount:0,
    poolCount:0,
    expandTime:0,
    stackTime:0,

}

var intervalminute = 1;//表示1分钟刷新一次整体数据

$(document).ready(function() {
    getTask(1);
    changeTime(0);
    changeOpenTime(0);
    window.setTimeout('getInstanceCount()',200);
    window.setTimeout('getClusterCount()',300);
    window.setTimeout('getPoolCount()',400);
    setInterval('getTask(1)',intervalminute*60*1000);
    setInterval('loadAllData()',intervalminute*60*1000);
    setInterval('loadOpenStackData()',intervalminute*60*1000);
    setInterval('getInstanceCount()',intervalminute*60*1000 + 200);
    setInterval('getClusterCount()',intervalminute*60*1000 + 400);
    setInterval('getPoolCount()',intervalminute*60*1000 + 600);
});

var refreshTaskCountView = function(){
    $("#taskCount").html(cache.taskCount);
}
var refreshMachineCountView = function(){
    $("#machineCount").html(cache.machineCount);
}
var refreshClusterCountView = function(){
    $("#clusterCount").html(cache.clusterCount);
}
var refreshPoolCountView = function () {
    $("#poolCount").html(cache.poolCount);
}

var getInstanceCount = function () {
    var url='/api/for_cloud/ecs.php?action=list';
    var postData={'pagesize':1000};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (data) {
            NProgress.done();
            if(data.code==0){
                if(typeof data.content != 'undefined') {
                    cache.machineCount = data.content.length;
                    refreshMachineCountView();
                }
            }
        },
        error: function (){
        }
    });
}

var getClusterCount = function () {
    var url='/api/for_layout/cluster.php?page=1';
    var postData={"action":"list","fIdx":""};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (listdata) {
            if(listdata.code==0){
                cache.clusterCount = listdata.count;
                refreshClusterCountView();
            }
        },
        error: function (){
        }
    });
}

var getPoolCount = function () {
    var postData={"action":"poolList","pagesize":1000};
    $.ajax({
        type: "POST",
        url: '/api/for_layout/pool.php',
        data: postData,
        dataType: "json",
        success: function (data) {
            if(data.code==0){
                if(data.content.length>0){
                    cache.poolCount = data.count;
                    refreshPoolCountView();
                }
            }
        },
        error: function (){
        }
    });
}

var changeTime = function(timeUnit){
    var timeHour = 0;
    if(timeUnit == 0){
        $("#time_unit").html('<a onclick="changeTime(0)">小时 </a><span class="caret"></span>');
        timeHour = 1;
    }
    if(timeUnit == 1){
        $("#time_unit").html('<a onclick="changeTime(1)">天 </a><span class="caret"></span>');
        timeHour = 1*24;
    }
    if(timeUnit == 2){
        $("#time_unit").html('<a onclick="changeTime(2)">周 </a><span class="caret"></span>');
        timeHour = 1*24*7;
    }
    if(timeUnit == 3){
        $("#time_unit").html('<a onclick="changeTime(3)">月 </a><span class="caret"></span>');
        timeHour = 1*24*31;
    }
    if(timeUnit == 4){
        $("#time_unit").html('<a onclick="changeTime(4)">年 </a><span class="caret"></span>');
        timeHour = 1*24*31*365;

    }
    var time_data = $("#time_nume").val();
    if($.isNumeric(time_data)){
        var time = parseInt(time_data)*timeHour;
        cache.expandTime = time;
        loadAllData();
    }
}

var changeOpenTime = function(timeUnit){
    var timeHour = 0;
    if(timeUnit == 0){
        $("#time_open_unit").html('<a onclick="changeTime(0)">小时 </a><span class="caret"></span>');
        timeHour = 1;
    }
    if(timeUnit == 1){
        $("#time_open_unit").html('<a onclick="changeTime(1)">天 </a><span class="caret"></span>');
        timeHour = 1*24;
    }
    if(timeUnit == 2){
        $("#time_open_unit").html('<a onclick="changeTime(2)">周 </a><span class="caret"></span>');
        timeHour = 1*24*7;
    }
    if(timeUnit == 3){
        $("#time_open_unit").html('<a onclick="changeTime(3)">月 </a><span class="caret"></span>');
        timeHour = 1*24*31;
    }
    if(timeUnit == 4){
        $("#time_open_unit").html('<a onclick="changeTime(4)">年 </a><span class="caret"></span>');
        timeHour = 1*24*31*365;

    }
    var time_data = $("#time_open_nume").val();
    if($.isNumeric(time_data)){
        var time = parseInt(time_data)*timeHour;
        cache.stackTime = time;
        loadOpenStackData();
    }
}

var loadOpenStackData = function () {
    var time = cache.stackTime;
    var url = '/api/for_openstack/machine.php?action=getcomputepowerbytime&time='+ time;
    $.ajax({
        type : "post",
        async : true,
        url : url,
        data : {},
        dataType : "json",
        success : function(result) {
            //请求成功时执行该函数内容，result即为服务器返回的json对象
            if (result.code == 0) {
                var line_open_data = [];
                var line_open_time = [];
                for(var i = 0; i < result.data.length; i++){
                    var cpusCout = result.data[i].data.vcpus;
                    var memory = result.data[i].data.memory_gb;
                    var machine_count = result.data[i].data.machine_count;
                    line_open_time.push(result.data[i].create_time);
                    if(i == 0){
                        var element = {
                            "name":"CPU(个)",
                            "data":[parseInt(cpusCout)]
                        }
                        var element2 = {
                            "name":"Memory(G)",
                            "data":[parseFloat(memory)]
                        }
                        line_open_data.push(element);
                        line_open_data.push(element2);
                    }else{
                        line_open_data[0].data.push(parseInt(cpusCout));
                        line_open_data[1].data.push(parseFloat(memory));
                    }
                }
                testMachineStackChart(line_open_data,line_open_time);
            }
        },
        error : function() {
            pageNotify('error','加载失败！','错误信息：接口不可用');
        }
    });
}

var loadAllData = function (){
    var time = cache.expandTime;
    var postData = {'action':'number','hour':time};
    $.ajax({
        type : "post",
        async : true,
        url : '/api/for_cloud/cluster.php?action=machine',
        data : {"data":JSON.stringify(postData)},
        dataType : "json",
        success : function(result) {
            if (result.code == 0) {
                var line_data = [];
                var line_time = [];
                var phydevCount = 0;
                var lineName = [];

                for(var i = 0; i < result.content.length; i++){
                    var map = eval(result.content[i]);
                    $.each(map, function (k, v) {
                        var name = k + "";
                        var flag = false;
                        for(var k = 0; k < lineName.length; k++){
                            if(name == lineName[k]){
                                flag = true;
                                break;
                            }
                        }
                        if(!flag){
                            lineName.push(name);
                        }
                    });
                }
                for(var i = 0; i < lineName.length; i++){
                    if(lineName[i] == "time" || lineName[i] == "phydev"){
                        continue;
                    }
                    var theLine = {
                        'name':lineName[i] ,
                        'data':[]
                    }
                    line_data.push(theLine);
                }

                for(var i = 0; i < result.content.length; i++){
                    var map = eval(result.content[i]);
                    $.each(map, function (k, v) {
                        var name = k + "";
                        if(name=="time") {
                            line_time.push(v);
                        }
                        if(name=="phydev") {
                            phydevCount = parseInt(v);
                        }
                    });
                    var current_data_length = 0;
                    $.each(map, function (k, v) {
                        var name = k + "";
                        for(var p = 0; p < line_data.length; p++){
                            if(line_data[p].name == name){
                                if(name == "total"){
                                    line_data[p].data.push(parseInt(v)-phydevCount);
                                    current_data_length = line_data[p].data.length;
                                }else{
                                    line_data[p].data.push(parseInt(v));
                                }
                            }
                        }
                    });
                    for(var p = 0; p < line_data.length; p++){
                        if(line_data[p].data.length < current_data_length){
                            line_data[p].data.push(0);
                        }
                    }
                }
                testMachineChart(line_data,line_time);
            }
        },
        error : function() {
            pageNotify('error','加载失败！','错误信息：接口不可用');
        }
    });
}

// var iniMachineLineChart = function (macheineData, time) {
//     var echartPie = echarts.init(document.getElementById('container'), theme2);
//     echartPie.setOption({
//         tooltip : {
//             trigger: 'axis',
//         },
//         calculable: true,
//         legend: {
//             enabled:false,
//         },
//         toolbox: {
//             show: false,
//         },
//         xAxis : [
//             {
//                 type : 'category',
//                 boundaryGap : false,
//                 data : time,
//             }
//         ],
//         yAxis : [
//             {
//                 type : 'value'
//             }
//         ],
//         series : macheineData
//     });
// }

var testMachineChart = function(macheineData, xaixs_time){
    // alert(JSON.stringify(macheineData));
    var echartPie = echarts.init(document.getElementById('container'), theme2);
    echartPie.setOption({
        title: {
            text: null
        },
        tooltip : {
            trigger: 'axis',
        },
        legend: {
            x: 'center',
            y: 20,
            data:(function () {
                // generate an array of random data
                var data = [];
                var totalIndex = -1;
                for(var i = 0; i < macheineData.length; i++){
                    if(macheineData[i].name == "phydev"){
                        continue;
                    }
                    if(macheineData[i].name == "total"){
                        totalIndex = i;
                    }else{
                        data.push(macheineData[i].name);
                    }
                }
                if(totalIndex != -1){
                    data.push(macheineData[totalIndex].name);
                }
                return data;
            }())
        },
        toolbox: {
            show: false,
        },
        xAxis : [
            {
                type : 'category',
                boundaryGap : false,
                data : xaixs_time
            }
        ],
        yAxis : [
            {
                type : 'value'
            }
        ],
        series :  (function () {
            // generate an array of random data
            var data = [];
            var totalIndex = -1;
            for(var i = 0; i < macheineData.length; i++){
                if(macheineData[i].name == "phydev"){
                    continue;
                }
                if(macheineData[i].name == "total"){
                    totalIndex = i;
                    continue;
                }
                var line = {
                    name:macheineData[i].name,
                    type:'line',
                    smooth: true,
                    stack: '总量',
                    areaStyle: {normal: {}},
                    data:macheineData[i].data
                }
                data.push(line);
            }
            if(totalIndex != -1){
                var line = {
                    name:macheineData[totalIndex].name,
                    type:'line',
                    smooth: true,
                    stack: '总量',
                    label: {
                        normal: {
                            show: true,
                            position: 'top'
                        }
                    },
                    areaStyle: {normal: {}},
                    data:macheineData[totalIndex].data
                }
                data.push(line);
            }
            return data;
        }())
    });
}

var testMachineStackChart = function(macheineData, xaixs_time){
    var echartPie = echarts.init(document.getElementById('container_stack'), theme2);
    echartPie.setOption({
        title: {
            text: null
        },
        tooltip : {
            trigger: 'axis',
        },
        legend: {
            x: 'center',
            y: 20,
            data:(function () {
                // generate an array of random data
                var data = [];
                for(var i = 0; i < macheineData.length; i++){
                    data.push(macheineData[i].name);
                }
                return data;
            }())
        },
        toolbox: {
            show: false,
        },
        xAxis : [
            {
                type : 'category',
                boundaryGap : false,
                data : xaixs_time
            }
        ],
        yAxis : [
            {
                type : 'value'
            }
        ],
        series :  (function () {
            // generate an array of random data
            var data = [];
            for(var i = 0; i < macheineData.length; i++){
                var line = {
                    name:macheineData[i].name,
                    type:'line',
                    smooth: true,
                    stack: '总量',
                    areaStyle: {normal: {}},
                    data:macheineData[i].data
                }
                data.push(line);
            }
            return data;
        }())
    });
}
// var initMachinePieChart = function(pie_data){
//     var echartPie = echarts.init(document.getElementById('echart_pie'), theme);
//     echartPie.setOption({
//         tooltip: {
//             trigger: 'item',
//             formatter: "{a} <br/>{b} : {c} ({d}%)"
//         },
//         legend: {
//             enabled:false
//         },
//         toolbox: {
//             show: false,
//         },
//         calculable: true,
//         series: [{
//             name: '模板比例',
//             type: 'pie',
//             radius: '55%',
//             center: ['50%', '48%'],
//             data: pie_data
//         }]
//     });
// }
//获取任务列表
var getTask=function(page){
    var url='/api/for_layout/task.php?action=list';
    if (!page) {
        page = cache.page;
    }else{
        cache.page = page;
    }
    var postData={"action":"list","page":page,"pagesize":8};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (listdata) {
            if(listdata.code==0) {
                //更新最新状态
                cache.taskCount = listdata.count;
                refreshTaskCountView();
                var pageinfo = $("#table-pageinfo");//分页信息
                var paginate = $("#table-paginate");//分页代码
                var head = $("#table-head");//数据表头
                var body = $("#table-body");//数据列表
                //清除当前页面数据
                pageinfo.html("");
                paginate.html("");
                head.html("");
                body.html("");
                listdata.title = ["#","服务池名称","任务名称","执行中","暂停","成功","失败","总计","成功率","执行时间"];
                processPage(listdata, pageinfo, paginate);
                //生成列表
                processBody(listdata, head, body);
                NProgress.done();
            }else{
                pageNotify('error','加载失败！','错误信息：'+listdata.msg);
                NProgress.done();
            }
        },
        error: function (){
            pageNotify('error','加载失败！','错误信息：接口不可用');
            NProgress.done();
        }
    });
}
//生成分页
var processPage = function(data,pageinfo,paginate){
    var begin = data.pageSize * ( data.page - 1 ) + 1;
    var end = ( data.count > begin + data.pageSize - 1 ) ? begin + data.pageSize - 1 : data.count;
    pageinfo.html('Showing '+begin+' to '+end+' of '+data.count+' records');
    var p1=(data.page-1>0)?data.page-1:1;
    var p2=data.page+1;
    prev='<li><a onclick="getTask('+p1+')"><i class="fa fa-angle-left"></i></a></li>';
    paginate.append(prev);
    for (var i = 1; i <= data.pageCount; i++) {
        var li='';
        if(i==data.page){
            li='<li class="active"><a onclick="getTask('+i+')">'+i+'</a></li>';
        }else{
            if(i==1||i==data.pageCount){
                li='<li><a onclick="getTask('+i+')">'+i+'</a></li>';
            }else{
                if(i==p1){
                    if(p1>2){
                        li='<li class="disabled"><a href="#">...</a></li>'+"\n"+'<li><a onclick="getTask('+i+')">'+i+'</a></li>';
                    }else{
                        li='<li><a onclick="getTask('+i+')">'+i+'</a></li>';
                    }
                }else{
                    if(i==p2){
                        if(p2<data.pageCount-1){
                            li='<li><a onclick="getTask('+i+')">'+i+'</a></li>'+"\n"+'<li class="disabled"><a href="#">...</a></li>';
                        }else{
                            li='<li><a onclick="getTask('+i+')">'+i+'</a></li>';
                        }
                    }
                }
            }
        }
        paginate.append(li);
    }
    if(p2>data.pageCount) p2=data.pageCount;
    next='<li class="next"><a title="Next" onclick="getTask('+p2+')"><i class="fa fa-angle-right"></i></a></li>';
    paginate.append(next);
}
//生成列表
var processBody = function(data,head,body){
    var td="";
    if(data.title){
        var tr = $('<tr></tr>');
        for (var i = 0; i < data.title.length; i++) {
            var v = data.title[i];
            td = '<th>' + v + '</th>';
            tr.append(td);
        }
        head.html(tr);
    }
    if(data.content){
        if(data.content.length>0){
            for (var i = 0; i < data.content.length; i++) {
                var v = data.content[i];
                var tr = $('<tr></tr>');
                //序号
                td = '<td>' + v.i + '</td>';
                tr.append(td);
                //服务池名称
                td = '<td>' + v.pool_name + '</td>';
                tr.append(td);
                //任务名称
                td = '<td>' + v.task_name + '</td>';
                tr.append(td);
                var stat = v.Stat;
                var running = v.Stat[1];
                var success = v.Stat[2];
                var failed = v.Stat[3];
                var stoped = v.Stat[4];
                var count = v.Stat[0]+v.Stat[1]+v.Stat[2]+v.Stat[3]+v.Stat[4];
                td = '<td><span class="label label-info">' + running + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-warning">' + stoped + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-success">' + success + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-danger">' + failed + '</span></td>';
                tr.append(td);
                td = '<td><span class="label label-default">' + count + '</span></td>';
                tr.append(td);
                var rate = 0.0;
                if(success + failed != 0){
                    rate = success * 100.0 / (success + failed)*1.0;
                }
                var tmp=rate.toString().substr(0,5);
                td = '<td><span class="label label-success">' + tmp + '%</span></td>';
                tr.append(td);

                var beginSec = (typeof(v.created)!='undefined') ? Date.parse(v.created) : 0;
                var endSec = (typeof(v.updated)!='undefined') ? ((v.updated!='0000-00-00 00:00:00') ? Date.parse(v.updated) : new Date().getTime()) : new Date().getTime();
                var timeLen=(beginSec>0)?Math.ceil((endSec-beginSec)/1000)+'秒':'-';
                td = '<td>' + timeLen + '</td>';
                tr.append(td);
                body.append(tr);
            }
        }else{
            pageNotify('info','Warning','数据为空！');
        }
    }else{
        pageNotify('warning','error','接口异常！');
    }
}