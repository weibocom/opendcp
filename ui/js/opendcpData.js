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
var cache = {
    index:[],
    test:{},
    page: 1,
    tasklist:[],
    check_step: {}, //Action详情
    task_step: [], //步骤
    task_tpl: {}, //模板列表
    cluster: [], //集群列表
    service: [], //服务列表
    pool: [], //服务池列表
    ip: [], //选中IP列表
}

$(document).ready(function() {
    getTask(1);
    getTheMachine();
});
var getModuleMachine = function(pie_data){
    var echartPie = echarts.init(document.getElementById('echart_pie'), theme);
    echartPie.setOption({
        tooltip: {
            trigger: 'item',
            formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        legend: {
            enabled:false
        },
        toolbox: {
            show: false,
        },
        calculable: true,
        series: [{
            name: '模板比例',
            type: 'pie',
            radius: '55%',
            center: ['50%', '48%'],
            data: pie_data
        }]
    });
}
var getTheMachine = function(){
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });
    var mychart = new Highcharts.Chart('container', {
        chart:{
            type: 'spline',
            animation: Highcharts.svg, // don't animate in IE < IE 10.
            marginRight: 5,
            events: {
                load: function () {
                    var series = this.series;
                    var loadData = function() {
                        $.ajax({
                            type : "post",
                            async : true,            //异步请求（同步请求将会锁住浏览器，用户其他操作必须等待请求完成才可以执行）
                            url : '/api/for_cloud/cluster.php?action=machine',    //请求发送到TestServlet处
                            data : {},
                            dataType : "json",        //返回数据形式为json
                            success : function(result) {
                                //请求成功时执行该函数内容，result即为服务器返回的json对象
                                if (result.code == 0) {
                                    var x = (new Date()).getTime();// current time
                                    var y = parseInt(result.content.all);
                                    series[0].addPoint([x, y], true, true);
                                    var json = eval(result.content); //数组
                                    var pie_data = [];
                                    var pie_index = 0;
                                    $.each(json, function (k, v) {
                                        var name = k + "";
                                        // if(name == "time") alert(true);
                                        if(name != "time" && name != "all"){
                                            var element = {};
                                            element["name"] = name;
                                            element["value"] = parseInt(v);
                                            pie_data[pie_index] = element;
                                            pie_index++;
                                        }
                                        //循环获取数据
                                    });
                                    getModuleMachine(pie_data);

                                }else{
                                    var x = (new Date()).getTime();// current time
                                    var y = 0;
                                    series[0].addPoint([x, y], true, true);
                                    var pie_data = [];
                                    getModuleMachine(pie_data);
                                }
                            },
                            error : function() {
                                pageNotify('error','加载失败！','错误信息：接口不可用');
                            }
                        });
                    }
                    loadData();
                    setInterval(loadData, 60000);
                }
            }
        },
        title: {
            text: '',
            x: -20
        },
        subtitle: {
            text: '',
            x: -20
        },
        xAxis: {
            type: 'datetime',
            tickPixelInterval: 100,
        },
        yAxis: {
            title: {
                text: ''
            },
            plotLines: [{
                value: 0,
                width: 1,
                color: '#808080'
            }]
        },
        legend: {
            enabled: false
        },
        credits: {
            enabled:false
        },
        exporting: {
            enabled:false
        },
        tooltip: {
            trigger: 'axis'
        },
        plotOptions: {
            series: {
                cursor: 'pointer',
                point: {
                    events: {
                        // 数据点点击事件
                        // 其中 e 变量为事件对象，this 为当前数据点对象
                        click: function (e) {
                            $.ajax({
                                type : "post",
                                async : true,            //异步请求（同步请求将会锁住浏览器，用户其他操作必须等待请求完成才可以执行）
                                url : '/api/for_cloud/cluster.php?action=machine',    //请求发送到TestServlet处
                                data : {},
                                dataType : "json",        //返回数据形式为json
                                success : function(result) {
                                    //请求成功时执行该函数内容，result即为服务器返回的json对象
                                    if (result.code == 0) {
                                        var json = eval(result.content); //数组
                                        var pie_data = [];
                                        var pie_index = 0;
                                        $.each(json, function (k, v) {
                                            var name = k + "";
                                            if(name != "time" && name != "all"){
                                                var element = {};
                                                element["name"] = name;
                                                element["value"] = parseInt(v);
                                                pie_data[pie_index] = element;
                                                pie_index++;
                                            }
                                            //循环获取数据
                                        });
                                        getModuleMachine(pie_data);

                                    }else{
                                        var pie_data = [];
                                        getModuleMachine(pie_data);
                                    }
                                },
                                error : function() {
                                    pageNotify('error','加载失败！','错误信息：接口不可用');
                                }
                            });
                        }
                    }
                },
                marker: {
                    lineWidth: 1
                }
            }
        },
        series: [{
            name: '机器数量',
            type:"spline",
            data: (function () {
                // generate an array of random data
                var data = [],time = (new Date()).getTime(),i;
                for (i = -19; i <= 0; i += 1) {
                    data.push({
                        x: time + i * 60000,
                        y: parseInt(Math.random()*10)
                    });
                }
                return data;
            }())
        }]
    });
}
//获取任务列表
var getTask=function(page){
    var url='/api/for_layout/task.php?action=list';
    if (!page) {
        page = cache.page;
    }else{
        cache.page = page;
    }
    var postData={"action":"list","page":page,"pagesize":20};
    $.ajax({
        type: "POST",
        url: url,
        data: postData,
        dataType: "json",
        success: function (listdata) {
            if(listdata.code==0) {
                //更新最新状态
                var pageinfo = $("#table-pageinfo");//分页信息
                var paginate = $("#table-paginate");//分页代码
                var head = $("#table-head");//数据表头
                var body = $("#table-body");//数据列表
                //清除当前页面数据
                pageinfo.html("");
                paginate.html("");
                head.html("");
                body.html("");
                //生成页面
                //生成分页
                //生成分页
                listdata.title = ["#","服务池名称","任务名称","执行中","成功","失败","总计","成功率","执行时间"];
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
                var count = v.Stat[0]+v.Stat[1]+v.Stat[2]+v.Stat[3];
                td = '<td><span class="label label-info">' + running + '</span></td>';
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